package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/AidanAlsaadoun1/market-settlement-engine/indexer/bindings"
)

type dbMarket struct {
	id      int64
	yesPool string // NUMERIC comes out as text; parsed with big.Int below
	noPool  string
}

func runReconciliation(ctx context.Context, caller *bindings.MarketCaller, pool *pgxpool.Pool) (failures int, err error) {
	rows, err := pool.Query(ctx,
		`SELECT id, yes_pool::text, no_pool::text FROM markets ORDER BY id`)
	if err != nil {
		return 0, fmt.Errorf("read derived markets: %w", err)
	}
	defer rows.Close()

	var markets []dbMarket
	for rows.Next() {
		var m dbMarket
		if err := rows.Scan(&m.id, &m.yesPool, &m.noPool); err != nil {
			return 0, fmt.Errorf("scan market: %w", err)
		}
		markets = append(markets, m)
	}
	if err := rows.Err(); err != nil {
		return 0, err
	}

	for _, m := range markets {
		passed, details, err := reconcileMarket(ctx, caller, m)
		if err != nil {
			// An RPC failure is a failed CHECK, not a failed SYSTEM:
			// record it and carry on. The next pass retries naturally.
			log.Printf("reconcile market %d: check errored: %v", m.id, err)
			details = map[string]any{"error": err.Error()}
			passed = false
		}
		if !passed {
			failures++
			log.Printf("RECONCILIATION FAILURE market %d: %v", m.id, details)
		}
		if err := recordRun(ctx, pool, m.id, passed, details); err != nil {
			return failures, fmt.Errorf("record run: %w", err)
		}
	}
	return failures, nil
}

// reconcileMarket compares one market's on-chain pools with the folds
func reconcileMarket(ctx context.Context, caller *bindings.MarketCaller, m dbMarket) (bool, map[string]any, error) {
	onChain, err := caller.Markets(nil, big.NewInt(m.id))
	if err != nil {
		return false, nil, fmt.Errorf("view call: %w", err)
	}

	ourYes, okY := new(big.Int).SetString(m.yesPool, 10)
	ourNo, okN := new(big.Int).SetString(m.noPool, 10)
	if !okY || !okN {
		return false, nil, fmt.Errorf("unparseable derived pools %q / %q", m.yesPool, m.noPool)
	}

	yesMatch := onChain.YesPool.Cmp(ourYes) == 0
	noMatch := onChain.NoPool.Cmp(ourNo) == 0

	details := map[string]any{
		"onchain_yes": onChain.YesPool.String(),
		"derived_yes": ourYes.String(),
		"onchain_no":  onChain.NoPool.String(),
		"derived_no":  ourNo.String(),
	}
	return yesMatch && noMatch, details, nil
}

// recordRun appends one check result to the audit table
func recordRun(ctx context.Context, pool *pgxpool.Pool, marketID int64, passed bool, details map[string]any) error {
	body, err := json.Marshal(details)
	if err != nil {
		return err
	}
	_, err = pool.Exec(ctx, `
		INSERT INTO reconciliation_runs (market_id, passed, details)
		VALUES ($1, $2, $3)`,
		marketID, passed, body)
	return err
}

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type foldEvent struct {
	eventType string
	payload   []byte
}

// foldEvents rebuilds markets and positions from the ledger
func foldEvents(ctx context.Context, pool *pgxpool.Pool) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin fold tx: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, `TRUNCATE markets, positions`); err != nil {
		return fmt.Errorf("truncate derived state: %w", err)
	}

	// Chain order is (block_number, log_index)
	rows, err := tx.Query(ctx, `
		SELECT event_type, payload
		FROM chain_events
		ORDER BY block_number, log_index`)
	if err != nil {
		return fmt.Errorf("read ledger: %w", err)
	}
	defer rows.Close()

	var events []foldEvent
	for rows.Next() {
		var ev foldEvent
		if err := rows.Scan(&ev.eventType, &ev.payload); err != nil {
			return fmt.Errorf("scan event: %w", err)
		}
		events = append(events, ev)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("ledger iteration: %w", err)
	}

	for _, ev := range events {
		if err := applyEvent(ctx, tx, ev); err != nil {
			return fmt.Errorf("apply %s: %w", ev.eventType, err)
		}
	}

	return tx.Commit(ctx)
}

func applyEvent(ctx context.Context, tx pgx.Tx, ev foldEvent) error {
	switch ev.eventType {
	case "MarketCreated":
		return applyMarketCreated(ctx, tx, ev.payload)
	case "SharesPurchased":
		return applySharesPurchased(ctx, tx, ev.payload)
	case "MarketResolved":
		return applyMarketResolved(ctx, tx, ev.payload)
	case "Claimed":
		return applyClaimed(ctx, tx, ev.payload)
	default:
		// An unknown type means the indexer wrote something the fold
		// doesn't understand: a bug, not a condition to skip silently.
		return fmt.Errorf("unknown event type %q", ev.eventType)
	}
}

func applyMarketCreated(ctx context.Context, tx pgx.Tx, payload []byte) error {
	var p struct {
		ID        string `json:"id"`
		Question  string `json:"question"`
		CloseTime string `json:"closeTime"`
	}
	if err := json.Unmarshal(payload, &p); err != nil {
		return err
	}
	_, err := tx.Exec(ctx, `
		INSERT INTO markets (id, question, close_time)
		VALUES ($1::bigint, $2, $3::bigint)`,
		p.ID, p.Question, p.CloseTime)
	return err
}

func applySharesPurchased(ctx context.Context, tx pgx.Tx, payload []byte) error {
	var p struct {
		MarketID string `json:"marketId"`
		Buyer    string `json:"buyer"`
		Yes      bool   `json:"yes"`
		Amount   string `json:"amount"`
	}
	if err := json.Unmarshal(payload, &p); err != nil {
		return err
	}

	pool := "no_pool"
	stake := "no_stake"
	if p.Yes {
		pool = "yes_pool"
		stake = "yes_stake"
	}

	if _, err := tx.Exec(ctx, fmt.Sprintf(`
		UPDATE markets SET %s = %s + $1::numeric WHERE id = $2::bigint`,
		pool, pool),
		p.Amount, p.MarketID); err != nil {
		return err
	}

	_, err := tx.Exec(ctx, fmt.Sprintf(`
		INSERT INTO positions (market_id, address, %s)
		VALUES ($1::bigint, $2, $3::numeric)
		ON CONFLICT (market_id, address)
		DO UPDATE SET %s = positions.%s + EXCLUDED.%s`,
		stake, stake, stake, stake),
		p.MarketID, p.Buyer, p.Amount)
	return err
}

func applyMarketResolved(ctx context.Context, tx pgx.Tx, payload []byte) error {
	var p struct {
		MarketID string `json:"marketId"`
		Outcome  bool   `json:"outcome"`
	}
	if err := json.Unmarshal(payload, &p); err != nil {
		return err
	}
	_, err := tx.Exec(ctx, `
		UPDATE markets SET resolved = true, outcome = $1
		WHERE id = $2::bigint`,
		p.Outcome, p.MarketID)
	return err
}

func applyClaimed(ctx context.Context, tx pgx.Tx, payload []byte) error {
	var p struct {
		MarketID string `json:"marketId"`
		Claimer  string `json:"claimer"`
	}
	if err := json.Unmarshal(payload, &p); err != nil {
		return err
	}
	_, err := tx.Exec(ctx, `
		UPDATE positions SET claimed = true
		WHERE market_id = $1::bigint AND address = $2`,
		p.MarketID, p.Claimer)
	return err
}

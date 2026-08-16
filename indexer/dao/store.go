package dao

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const deploymentBlock uint64 = 11488331 // your contract's creation block

func GetCheckpoint(ctx context.Context, pool *pgxpool.Pool) (uint64, error) {
	var lastBlock uint64
	err := pool.QueryRow(ctx,
		"SELECT last_block FROM checkpoints WHERE id = 'indexer'",
	).Scan(&lastBlock)
	if errors.Is(err, pgx.ErrNoRows) {
		return deploymentBlock, nil
	}
	if err != nil {
		return 0, err
	}
	return lastBlock, nil
}

func SetCheckpoint(ctx context.Context, pool *pgxpool.Pool, block uint64) error {
	_, err := pool.Exec(ctx, `
		INSERT INTO checkpoints (id, last_block, updated_at)
		VALUES ('indexer', $1, now())
		ON CONFLICT (id) DO UPDATE
		SET last_block = EXCLUDED.last_block, updated_at = now()`,
		block,
	)
	return err
}

func InsertEvent(ctx context.Context, pool *pgxpool.Pool,
	txHash string, logIndex uint, blockNumber uint64, blockHash string,
	eventType string, payload map[string]any) error {

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	_, err = pool.Exec(ctx, `
		INSERT INTO chain_events
			(tx_hash, log_index, block_number, block_hash, event_type, payload)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (tx_hash, log_index) DO NOTHING`,
		txHash, logIndex, blockNumber, blockHash, eventType, body,
	)
	return err
}

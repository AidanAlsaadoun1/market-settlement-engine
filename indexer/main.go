package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/AidanAlsaadoun1/market-settlement-engine/indexer/bindings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
)

const deploymentBlock uint64 = 11488340

func getEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s not set", key)
	}
	return v
}

func insertEvent(
	ctx context.Context,
	pool *pgxpool.Pool,
	txHash string,
	logIndex uint,
	blockNumber uint64,
	blockHash string,
	eventType string,
	payload map[string]any,
) error {

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

func main() {
	rpcURL := getEnv("RPC_URL")
	dbURL := getEnv("DATABASE_URL")
	contractAddr := common.HexToAddress(getEnv("CONTRACT"))

	ctx := context.Background()

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("dial failed: %v", err)
	}
	defer client.Close()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}
	defer pool.Close()

	filterer, err := bindings.NewMarketFilterer(contractAddr, client)
	if err != nil {
		log.Fatalf("filterer init failed: %v", err)
	}

	caller, err := bindings.NewMarketCaller(contractAddr, client)
	if err != nil {
		log.Fatalf("caller init failed: %v", err)
	}
	if err := runIndexer(ctx, client, filterer, caller, pool); err != nil {
		log.Fatalf("indexer stopped: %v", err)
	}
}

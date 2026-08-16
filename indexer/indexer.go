package main

import (
	"context"
	"log"
	"time"

	"github.com/AidanAlsaadoun1/market-settlement-engine/indexer/bindings"
	"github.com/AidanAlsaadoun1/market-settlement-engine/indexer/dao"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	confirmations uint64 = 6
	chunkSize     uint64 = 8
	pollInterval         = 15 * time.Second
)

func runIndexer(
	ctx context.Context,
	client *ethclient.Client,
	filterer *bindings.MarketFilterer,
	pool *pgxpool.Pool,
) error {

	for {
		latest, err := client.BlockNumber(ctx)
		if err != nil {
			log.Printf("block number fetch failed: %v", err)
			time.Sleep(pollInterval)
			continue
		}
		safe := latest - confirmations

		checkpoint, err := dao.GetCheckpoint(ctx, pool)
		if err != nil {
			return err
		}

		from := checkpoint + 1
		if from > safe {
			time.Sleep(pollInterval)
			continue
		}
		to := min(from+chunkSize, safe)

		opts := &bind.FilterOpts{Start: from, End: &to, Context: ctx}

		count := 0
		n, err := ingestMarketCreated(ctx, filterer, opts, pool)
		if err != nil {
			return err
		}
		count += n

		n, err = ingestSharesPurchased(ctx, filterer, opts, pool)
		if err != nil {
			return err
		}
		count += n

		n, err = ingestMarketResolved(ctx, filterer, opts, pool)
		if err != nil {
			return err
		}
		count += n

		n, err = ingestClaimed(ctx, filterer, opts, pool)
		if err != nil {
			return err
		}
		count += n

		if err := dao.SetCheckpoint(ctx, pool, to); err != nil {
			return err
		}
		log.Printf("indexed blocks %d-%d, %d events", from, to, count)

		if to >= safe {
			time.Sleep(pollInterval)
		} else {
			time.Sleep(300 * time.Millisecond) // gentle throttle
		}
	}
}

func ingestMarketCreated(ctx context.Context, f *bindings.MarketFilterer,
	opts *bind.FilterOpts, pool *pgxpool.Pool) (int, error) {

	iter, err := f.FilterMarketCreated(opts, nil)
	if err != nil {
		return 0, err
	}
	defer iter.Close()

	count := 0
	for iter.Next() {
		ev := iter.Event
		payload := map[string]any{
			"id":        ev.Id.String(),
			"question":  ev.Question,
			"closeTime": ev.CloseTime.String(),
		}
		if err := insertEvent(ctx, pool, ev.Raw.TxHash.Hex(), ev.Raw.Index,
			ev.Raw.BlockNumber, ev.Raw.BlockHash.Hex(), "MarketCreated", payload); err != nil {
			return count, err
		}
		count++
	}
	return count, iter.Error()
}

func ingestSharesPurchased(ctx context.Context, f *bindings.MarketFilterer,
	opts *bind.FilterOpts, pool *pgxpool.Pool) (int, error) {

	iter, err := f.FilterSharesPurchased(opts, nil, nil)
	if err != nil {
		return 0, err
	}
	defer iter.Close()

	count := 0
	for iter.Next() {
		ev := iter.Event
		payload := map[string]any{
			"marketId": ev.MarketId.String(),
			"buyer":    ev.Buyer.Hex(),
			"yes":      ev.Yes,
			"amount":   ev.Amount.String(),
		}
		if err := insertEvent(ctx, pool, ev.Raw.TxHash.Hex(), ev.Raw.Index,
			ev.Raw.BlockNumber, ev.Raw.BlockHash.Hex(), "SharesPurchased", payload); err != nil {
			return count, err
		}
		count++
	}
	return count, iter.Error()
}

func ingestMarketResolved(ctx context.Context, f *bindings.MarketFilterer,
	opts *bind.FilterOpts, pool *pgxpool.Pool) (int, error) {

	iter, err := f.FilterMarketResolved(opts, nil)
	if err != nil {
		return 0, err
	}
	defer iter.Close()

	count := 0
	for iter.Next() {
		ev := iter.Event
		payload := map[string]any{
			"marketId": ev.MarketId.String(),
			"outcome":  ev.Outcome,
		}
		if err := insertEvent(ctx, pool, ev.Raw.TxHash.Hex(), ev.Raw.Index,
			ev.Raw.BlockNumber, ev.Raw.BlockHash.Hex(), "MarketResolved", payload); err != nil {
			return count, err
		}
		count++
	}
	return count, iter.Error()
}

func ingestClaimed(ctx context.Context, f *bindings.MarketFilterer,
	opts *bind.FilterOpts, pool *pgxpool.Pool) (int, error) {

	iter, err := f.FilterClaimed(opts, nil, nil)
	if err != nil {
		return 0, err
	}
	defer iter.Close()

	count := 0
	for iter.Next() {
		ev := iter.Event
		payload := map[string]any{
			"marketId": ev.MarketId.String(),
			"claimer":  ev.Claimer.Hex(),
			"payout":   ev.Payout.String(),
		}
		if err := insertEvent(ctx, pool, ev.Raw.TxHash.Hex(), ev.Raw.Index,
			ev.Raw.BlockNumber, ev.Raw.BlockHash.Hex(), "Claimed", payload); err != nil {
			return count, err
		}
		count++
	}
	return count, iter.Error()
}

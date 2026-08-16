# Settlement Engine

A miniature prediction-market settlement system: a Solidity parimutuel market contract on Sepolia, with a Go service that ingests the contract's event logs into PostgreSQL and continuously reconciles on-chain state against the off-chain database, alerting on any divergence.

Built to explore the correctness problems real-money trading platforms face: idempotent event ingestion, auditable state, and detecting drift between two systems of record.

## Status

- [x] Parimutuel market contract (`createMarket` / `buy` / `resolve` / `claim`) with Foundry lifecycle tests
- [ ] Deployed and verified on Sepolia
- [ ] Go indexer: checkpointed backfill of event logs, idempotent upserts into an append-only ledger
- [ ] Derived state (markets, positions) folded from the ledger
- [ ] Reconciler: on-chain view calls vs off-chain SQL, divergence alerting
- [ ] Fault-injection demo

## How it works

```
Solidity contract (Sepolia)
        | event logs (JSON-RPC)
        v
   Go indexer  -- poll, checkpoint, decode
        v
 chain_events (Postgres, append-only)
        v  fold
 derived state: markets, positions
        ^
        |  compared every cycle by the
   Reconciler  <-- view calls to the contract
        |
        v
 reconciliation_runs + alert on divergence
```

The contract is the system of record. Events narrate every state change; the indexer writes the narration down; derived state is rebuilt purely by folding over the ledger. Because the contract's pools and the database's sums are two accounts of the same events, they must always agree — the reconciler checks that they do.

## The contract

A parimutuel (pool-betting) market: anyone creates a yes/no question with a close time, stakes ETH on a side before close, and after the owner resolves the outcome, winners claim their stake back plus a pro-rata share of the losing pool.

Design points:

- **Checks-effects-interactions** on `claim`: the claimed flag is set before ETH is sent, so re-entrant calls fail the `already claimed` check.
- **Multiply before divide** in the payout maths — integer division means `(stake / pool) * losing` truncates to zero.
- **`call{value:}` over `transfer`** for payouts, with the success flag checked.

## Running the tests

```bash
cd contracts
forge install
forge test -vv
```

The main test walks a full market lifecycle with three participants (buy on both sides, resolve, winning claims, losing claim reverts, double-claim reverts).

## Known limitations (deliberate scope decisions)

- **Centralised oracle**: `resolve` is owner-only. A production market would use a decentralised oracle or dispute mechanism.
- **Timestamp-based close**: market close compares `block.timestamp`. Validator timestamp drift (± seconds) is acceptable at the hour/day granularity these markets use.
- **Stuck funds edge case**: if a market resolves with an empty winning pool, the losing pool is unclaimable. Acceptable for v1; a production version would add a refund path.
- **Confirmation depth, not reorg rollback** (indexer, upcoming): the indexer will trail the chain head by N confirmations rather than storing block hashes and re-ingesting from fork points.
- **Polling, not WebSocket subscription** (indexer, upcoming): simpler and self-healing for this scale; the backfill path doubles as crash recovery.

## How this was built

Hand-written with AI assistance used in three modes: as a tutor (concept explanations before writing), as a reviewer (senior-engineer-style review of code after writing it), and as an unblocker (timeboxed help on toolchain errors). Notes on what that workflow caught — including a set of inverted `require` conditions the lifecycle test would also have flagged — are in [NOTES.md](NOTES.md).

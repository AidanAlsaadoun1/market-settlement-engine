# Settlement Engine

A miniature prediction-market settlement system: a Solidity parimutuel market contract on Sepolia, with a Go service that ingests the contract's event logs into PostgreSQL and continuously reconciles on-chain state against the off-chain database, alerting on any divergence.

Built to explore the correctness problems real-money trading platforms face: idempotent event ingestion, auditable state, and detecting drift between two systems of record.

## Architecture

```mermaid
flowchart TB
    subgraph chain["Sepolia (source of truth)"]
        C["SettlementMarket contract<br/><i>createMarket · buy · resolve · claim</i>"]
    end

    subgraph indexer["Go indexer"]
        P["Polling loop<br/><i>chunked ranges, latest − 6 confirmations</i>"]
        F["Fold<br/><i>replay ledger in chain order</i>"]
        R["Reconciler<br/><i>view calls vs derived state</i>"]
    end

    subgraph pg["PostgreSQL"]
        L[("chain_events<br/><i>append-only ledger</i>")]
        D[("markets · positions<br/><i>derived, disposable</i>")]
        A[("reconciliation_runs<br/><i>audit trail</i>")]
        K[("checkpoints")]
    end

    C -- "event logs (JSON-RPC)" --> P
    P -- "idempotent insert<br/>ON CONFLICT (tx_hash, log_index)" --> L
    P -- "advance after writes" --> K
    L -- "fold: truncate + replay" --> D
    D -- "derived pools" --> R
    C -- "view calls: on-chain pools" --> R
    R -- "record every check,<br/>alert on divergence" --> A
```

## Status

- [x] Parimutuel market contract (`createMarket` / `buy` / `resolve` / `claim`) with Foundry lifecycle tests
- [x] Deployed and verified on Sepolia: [`0x36A7cf4DCD4a722Afc3F0Fd3861729CeF5c8B516`](https://sepolia.etherscan.io/address/0x36A7cf4DCD4a722Afc3F0Fd3861729CeF5c8B516)
- [x] Go indexer: checkpointed backfill of event logs, idempotent upserts into an append-only ledger
- [x] Derived state (markets, positions) folded from the ledger
- [x] Reconciler: on-chain view calls vs off-chain SQL, divergence alerting
- [x] Fault-injection demo

## The fold and the reconciler

The contract is the system of record. Events narrate every state change; the indexer writes the narration down into an append-only ledger. Derived state (`markets`, `positions`) is **disposable by design**: on every pass the fold truncates it and rebuilds it by replaying the ledger in chain order, inside one transaction, so a corrupted projection heals itself on the next cycle. The reconciler then verifies the projection independently: view calls ask the contract for its pool totals, and every comparison is recorded in `reconciliation_runs`, pass or fail, with both sides' numbers. Because the contract's pools and the database's sums are two accounts of the same events, they must always agree. The reconciler is the proof that they do.

## Fault injection: catch and heal

Corrupting a derived pool by hand (`UPDATE markets SET yes_pool = yes_pool + 999999 ...`) demonstrates both halves of the design:

<img width="478" height="169" alt="Reconciliation failure log showing on-chain vs derived pool mismatch" src="https://github.com/user-attachments/assets/46e48742-f134-4c26-8203-2e50f273b97e" />

*The reconciler catches the divergence on its next pass, recording both sides' numbers in `reconciliation_runs`.*

<img width="277" height="88" alt="Subsequent pass showing consistency restored after the fold rebuilt derived state" src="https://github.com/user-attachments/assets/f9ede9a7-9fca-4b41-869c-797af2f7cc5c" />

*One pass later it has healed itself: the fold rebuilt the projection from the ledger. The corruption was never in the source of truth, so it could not survive.*

## The contract

A parimutuel (pool-betting) market: anyone creates a yes/no question with a close time, stakes ETH on a side before close, and after the owner resolves the outcome, winners claim their stake back plus a pro-rata share of the losing pool.

Design points:

- **Checks-effects-interactions** on `claim`: the claimed flag is set before ETH is sent, so re-entrant calls fail the `already claimed` check.
- **Multiply before divide** in the payout maths. Integer division means `(stake / pool) * losing` truncates to zero.
- **`call{value:}` over `transfer`** for payouts, with the success flag checked.

## The indexer

A Go service that copies the contract's event history into PostgreSQL and keeps following the chain:

- **Backfill and live tail are the same loop**: each pass closes the gap between the last checkpoint and the safe head, whether that gap is fifteen thousand blocks or five.
- **Confirmation depth**: the indexer only reads up to `latest - 6`, so events from blocks that could still be reorganised never enter the ledger.
- **Idempotent by design**: events are keyed on `(tx_hash, log_index)`, their identity on the chain itself, with `ON CONFLICT DO NOTHING`. Replaying any range is a no-op.
- **At-least-once delivery**: the checkpoint only advances after every insert in a range has succeeded, so a crash at any moment re-processes work rather than skipping it. Idempotent inserts make the re-processing harmless.
- **Transient vs fatal errors**: RPC failures (rate limits, timeouts) log and retry the same range on the next pass; database and checkpoint failures stop the process, because a broken store with an advancing loop risks inconsistency.

## Running it

```bash
# 1. Contract tests
cd contracts
forge install
forge test -vv

# 2. Database (Postgres 16, local or docker compose up -d)
migrate -path indexer/migrations \
  -database "postgres://postgres:dev@localhost:5432/postgres?sslmode=disable" up

# 3. Indexer
cp .env.example .env    # fill in your RPC URL, contract address, database URL
export $(grep -v '^#' .env | xargs)
cd indexer && go run .
# watch for: "reconciliation: all markets consistent"
```

The main contract test walks a full market lifecycle with three participants (buy on both sides, resolve, winning claims, losing claim reverts, double-claim reverts).

## Known limitations (deliberate scope decisions)

- **Centralised oracle**: `resolve` is owner-only. A production market would use a decentralised oracle or dispute mechanism.
- **Timestamp-based close**: market close compares `block.timestamp`. Validator timestamp drift (± seconds) is acceptable at the hour/day granularity these markets use.
- **Stuck funds edge case**: if a market resolves with an empty winning pool, the losing pool is unclaimable. Acceptable for v1; a production version would add a refund path.
- **Confirmation depth, not reorg rollback**: the indexer trails the chain head by 6 confirmations rather than storing block hashes and re-ingesting from fork points. Block hashes are stored in the ledger, so full rollback is a straightforward extension.
- **Polling, not WebSocket subscription**: simpler and self-healing at this scale; the backfill path doubles as crash recovery.
- **Full re-fold each pass, not incremental**: derived state is rebuilt from the entire ledger every cycle. Deliberate at this event volume, and it makes "disposable derived state" literally true. At production volume the fold would advance from its own checkpoint, mirroring the indexer.
- **Free-tier RPC constraints**: the block range per `eth_getLogs` call and the backfill pace are tuned to Alchemy's free-tier limits (10-block inclusive ranges, throttled request rate). With a paid or less restricted endpoint, `chunkSize` scales up with no code changes.
- **Fatal on repeated RPC failure is acceptable here**: the checkpoint makes restarts free. Production would add exponential backoff and alerting.

## How this was built

Hand-written with AI assistance used in three modes: as a tutor (concept explanations before writing), as a reviewer (senior-engineer-style review of code after writing it), and as an unblocker (timeboxed help on toolchain errors). Notes on what that workflow caught — including a set of inverted `require` conditions the lifecycle test would also have flagged, an off-by-one in inclusive block ranges, and the free-tier rate-limit handling above — are in [NOTES.md](NOTES.md).
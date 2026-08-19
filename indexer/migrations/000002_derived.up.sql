-- Derived state: everything in this migration is DISPOSABLE.
-- These tables are never the source of truth; they are a queryable
-- projection of chain_events, rebuilt by the fold at any time.
-- If a bug corrupts them, the fix is re-derivation, not restoration.

CREATE TABLE markets (
    id         BIGINT PRIMARY KEY,   -- the on-chain market id
    question   TEXT    NOT NULL,
    close_time BIGINT  NOT NULL,     -- unix seconds, as emitted
    resolved   BOOLEAN NOT NULL DEFAULT false,
    outcome    BOOLEAN,              -- NULL until resolved
    -- NUMERIC, not BIGINT: wei amounts reach 10^18 per ether and sums
    -- exceed int64. NUMERIC is exact and unbounded; floats never touch money.
    yes_pool   NUMERIC NOT NULL DEFAULT 0,
    no_pool    NUMERIC NOT NULL DEFAULT 0
);

CREATE TABLE positions (
    market_id BIGINT  NOT NULL,
    address   TEXT    NOT NULL,
    yes_stake NUMERIC NOT NULL DEFAULT 0,
    no_stake  NUMERIC NOT NULL DEFAULT 0,
    claimed   BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (market_id, address)  -- one row per participant per market
);
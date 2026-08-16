CREATE TABLE IF NOT EXISTS chain_events (
    tx_hash      TEXT        NOT NULL,
    log_index    BIGINT      NOT NULL,
    block_number BIGINT      NOT NULL,
    block_hash   TEXT        NOT NULL,
    event_type   TEXT        NOT NULL,
    payload      JSONB       NOT NULL,
    ingested_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chain_events_identity UNIQUE (tx_hash, log_index)
);

CREATE INDEX idx_chain_events_block ON chain_events (block_number, log_index);

-- I'm making a checkpoint table here so that even if for some reason this waas to creash we have some sort of progress market
CREATE TABLE checkpoints (
    id         TEXT        PRIMARY KEY,
    last_block BIGINT      NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- simple audit trailf ro now with a pass or fail, failures will have details in it, doing this as a JSON blob, which might be a pain but is fine for this
CREATE TABLE reconciliation_runs (
    id        BIGSERIAL   PRIMARY KEY,
    run_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    market_id BIGINT      NOT NULL,
    passed    BOOLEAN     NOT NULL,
    details   JSONB       NOT NULL DEFAULT '{}'::jsonb
);
-- chains --
CREATE TABLE IF NOT EXISTS chains (
    id SERIAL PRIMARY KEY,
    chain_id VARCHAR(32) UNIQUE NOT NULL,
    chain_name VARCHAR(128) NOT NULL,
    icon_url VARCHAR(256) NULL,
    website VARCHAR(256) NULL,
    inserted_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- blocks --
CREATE TABLE IF NOT EXISTS blocks (
    id SERIAL PRIMARY KEY,
    chain_id INT NOT NULL,
    height BIGINT NOT NULL,
    block_hash VARCHAR(256) NOT NULL,
    prev_hash VARCHAR(256) NOT NULL,
    proposer_address VARCHAR(256) NOT NULL,
    last_commit_hash VARCHAR(256) NOT NULL,
    data_hash VARCHAR(256) NOT NULL,
    validators_hash VARCHAR(256) NOT NULL,
    next_validators_hash VARCHAR(256) NOT NULL,
    consensus_hash VARCHAR(256) NOT NULL,
    app_hash VARCHAR(256) NOT NULL,
    last_result_hash VARCHAR(256) NOT NULL,
    evidence_hash VARCHAR(256) NOT NULL,
    block_time TIMESTAMP NOT NULL,
    inserted_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- transactions --
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    chain_id INT NOT NULL,
    transaction_hash VARCHAR(256) UNIQUE NOT NULL,
    height BIGINT NOT NULL,
    code int NOT NULL,
    code_space VARCHAR(256) NOT NULL,
    tx_data TEXT NOT NULL,
    raw_log TEXT NOT NULL,
    info TEXT NOT NULL,
    memo VARCHAR(1024),
    gas_wanted BIGINT NOT NULL,
    gas_used BIGINT NOT NULL,
    issued_at TIMESTAMP NOT NULL,
    inserted_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- events --
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    chain_id INT NOT NULL,
    tx_id int NOT NULL,
    block_height BIGINT NOT NULL,
    event_seq INT NOT NULL,
    event_type VARCHAR(256) NOT NULL,
    event_key VARCHAR(256) NOT NULL,
    event_value VARCHAR(256) NOT NULL,
    indexed BOOLEAN NOT NULL,
    inserted_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- account --
CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    chain_id INT NOT NULL,
    address VARCHAR(256) UNIQUE NOT NULL,
    inserted_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

-- account_balance --
CREATE TABLE IF NOT EXISTS account_balance (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    amount BIGINT NOT NULL,
    coin_name VARCHAR(20) NOT NULL,
    inserted_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES accounts(id)
);

-- messages --
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    transaction_id INT NOT NULL,
    seq INT NOT NULL,
    rawdata JSONB NOT NULL,
    inserted_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    CONSTRAINT fk_transaction_id FOREIGN KEY (transaction_id) REFERENCES transactions(id)
);

-- create indexes --
CREATE INDEX idx_blocks_chain_id ON blocks(chain_id);
CREATE INDEX idx_blocks_height ON blocks(height);
CREATE INDEX idx_blocks_block_hash ON blocks(block_hash);

CREATE INDEX idx_accounts_chain_id ON accounts(chain_id);

CREATE INDEX idx_events_chain_id ON events(chain_id);
CREATE INDEX idx_events_tx_id ON events(tx_id);

CREATE INDEX idx_transactions_chain_id ON events(chain_id);
CREATE INDEX idx_transactions_tx_hash ON transactions(transaction_hash);
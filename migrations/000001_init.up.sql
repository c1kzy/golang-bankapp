CREATE SCHEMA bankapp;

CREATE TABLE bankapp.users (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    full_name VARCHAR NOT NULL CHECK (char_length(full_name) BETWEEN 3 AND 100),
    balance BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE bankapp.transactions (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    amount BIGINT NOT NULL CHECK (amount > 0),

    transaction_type VARCHAR NOT NULL,
    transaction_status VARCHAR NOT NULL DEFAULT 'completed',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    sender_id INT REFERENCES bankapp.users(id),
    receiver_id INT REFERENCES bankapp.users(id)
);
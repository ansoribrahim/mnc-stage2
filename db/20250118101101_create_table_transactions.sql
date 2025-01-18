-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id),
    "type" VARCHAR(50) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL DEFAULT 0 CHECK (amount >= 0),
    balance_before DECIMAL(10, 2) NOT NULL DEFAULT 0 CHECK (balance_before >= 0),
    balance_after DECIMAL(10, 2) NOT NULL DEFAULT 0 CHECK (balance_after >= 0),
    remarks TEXT,
    status VARCHAR NOT NULL,
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd

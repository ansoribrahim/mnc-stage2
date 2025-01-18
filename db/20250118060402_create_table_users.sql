-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR (255) NOT NULL,
    phone_number VARCHAR (15) NOT NULL UNIQUE,
    pin VARCHAR(1000) NOT NULL,
    address TEXT,
    created_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp WITH TIME ZONE,
    CONSTRAINT idx_users_phone_number UNIQUE (phone_number)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

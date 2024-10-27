-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workflows (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    "name" VARCHAR,
    "type" VARCHAR,
    account_id INT,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    created_at TIMESTAMP WITHOUT TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workflows;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS flow_instances_accounts (
    id SERIAL PRIMARY KEY,
    instance_id VARCHAR,
    account_id INT REFERENCES accounts (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON flow_instances_accounts;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON flow_instances_accounts
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON flow_instances_accounts;

DROP TABLE IF EXISTS flow_instances_accounts;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS currencies (
    iso_code VARCHAR(10) PRIMARY KEY NOT NULL,
    name VARCHAR(128) NOT NULL,
    symbol VARCHAR(32),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON currencies;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON currencies
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON currencies;

DROP TABLE IF EXISTS currencies ;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    gst_no INT,
    primary_user_id UUID NOT NULL,
    default_workflow UUID DEFAULT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    CONSTRAINT fk_accounts_users_id FOREIGN KEY(primary_user_id) REFERENCES users(id)
);

DROP TRIGGER IF EXISTS set_timestamp ON accounts;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON accounts
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON accounts;

DROP TABLE IF EXISTS accounts ;
-- +goose StatementEnd

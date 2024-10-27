-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts_users_map (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    account_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    CONSTRAINT fk_users_accounts_map_user_id FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_users_accounts_map_account_id FOREIGN KEY(account_id) REFERENCES accounts(id)
);

DROP TRIGGER IF EXISTS set_timestamp ON accounts_users_map;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON accounts_users_map
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON accounts_users_map;

DROP TABLE IF EXISTS accounts_users_map ;
-- +goose StatementEnd

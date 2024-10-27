-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(256) NOT NULL,
    email VARCHAR(128) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    is_parent BOOLEAN DEFAULT FALSE,
    primary_location_id UUID NOT NULL,
    access_token UUID NOT NULL DEFAULT uuid_generate_v4(),
    role_id INT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    CONSTRAINT fk_users_locations_id FOREIGN KEY(primary_location_id) REFERENCES locations(id)
);

DROP TRIGGER IF EXISTS set_timestamp ON users;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON users;

DROP TABLE IF EXISTS users ;
-- +goose StatementEnd

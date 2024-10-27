-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_locations_map (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users (id),
    location_id UUID REFERENCES locations (id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    CONSTRAINT fk_users_locations_users_id FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_users_locations_locations_id FOREIGN KEY(location_id) REFERENCES locations(id)
);

DROP TRIGGER IF EXISTS set_timestamp ON users_locations_map;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON users_locations_map
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON users_locations_map;

DROP TABLE IF EXISTS users_locations_map ;
-- +goose StatementEnd

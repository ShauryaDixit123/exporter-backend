-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS locations (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(), 
    line1 VARCHAR NOT NULL,
    line2 VARCHAR,
    area VARCHAR,
    city VARCHAR(128),
    "state" VARCHAR(128) NOT NULL,
    pincode INT NOT NULL,
    country_id VARCHAR NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    CONSTRAINT fk_locations_country_id FOREIGN KEY(country_id) REFERENCES countries(iso_code)
);

DROP TRIGGER IF EXISTS set_timestamp ON locations;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON locations
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON locations;

DROP TABLE IF EXISTS locations ;
-- +goose StatementEnd

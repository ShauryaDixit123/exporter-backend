-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS countries (
    iso_code VARCHAR(10) PRIMARY KEY NOT NULL,
    name VARCHAR(128) NOT NULL,
    iso_3 VARCHAR(32) DEFAULT NULL,
    calling_code SMALLINT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    CHECK(calling_code > 0)
);

DROP TRIGGER IF EXISTS set_timestamp ON countries;
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON countries
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_timestamp ON countries;

DROP TABLE IF EXISTS countries;
-- +goose StatementEnd

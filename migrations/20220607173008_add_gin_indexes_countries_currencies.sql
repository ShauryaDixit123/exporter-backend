-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX IF NOT EXISTS countries_name_search_idx ON countries USING GIN (name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS currencies_name_search_idx ON currencies USING GIN (name gin_trgm_ops);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS countries_name_search_idx;
DROP INDEX IF EXISTS currencies_name_search_idx;

DROP EXTENSION IF EXISTS pg_trgm;
-- +goose StatementEnd

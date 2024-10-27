-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd

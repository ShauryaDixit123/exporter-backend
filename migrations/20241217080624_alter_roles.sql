-- +goose Up
-- +goose StatementBegin
ALTER TABLE roles
ADD COLUMN IF NOT EXISTS display_value VARCHAR(54);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

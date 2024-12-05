-- +goose Up
-- +goose StatementBegin
ALTER TABLE purchase_orders
ADD COLUMN IF NOT EXISTS quote_id UUID REFERENCES quotes (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd


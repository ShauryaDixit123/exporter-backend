-- +goose Up
-- +goose StatementBegin
ALTER TABLE accounts
ADD default_workflow_post_order UUID;
ALTER TABLE accounts RENAME COLUMN default_workflow TO default_workflow_pre_order;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

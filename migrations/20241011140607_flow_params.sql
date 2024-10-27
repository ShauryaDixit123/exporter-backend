-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS flow_params (
	id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    flow_id UUID REFERENCES flows (id),
    "name" VARCHAR,
    "type" VARCHAR,
    mandatory BOOLEAN,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    created_at TIMESTAMP WITHOUT TIME ZONE,
    updated_by VARCHAR,
    created_by VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

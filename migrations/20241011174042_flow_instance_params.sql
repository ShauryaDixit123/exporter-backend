-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS flow_instance_params (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    flow_instance_id UUID REFERENCES flow_instances (id),
    "name" VARCHAR,
    "type" VARCHAR,
    mandatory BOOLEAN,
    "value" VARCHAR,
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

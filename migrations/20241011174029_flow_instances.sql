-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS flow_instances (
	id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(), 
    workflow_id UUID REFERENCES workflows (id),
    "description" VARCHAR,
    "type" VARCHAR,
    "title" VARCHAR,
    "order" INTEGER,
    active BOOLEAN DEFAULT FALSE,
    tat INTEGER,
    instance_id VARCHAR,
    instance_type VARCHAR,
    "status"  VARCHAR,
    is_completed BOOLEAN DEFAULT FALSE,
    assigned_to VARCHAR,
    expires_at TIMESTAMP WITHOUT TIME ZONE,
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

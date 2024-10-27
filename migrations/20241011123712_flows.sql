-- +goose Up
-- +goose StatementBegin
 
CREATE TABLE IF NOT EXISTS flows (
	id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    workflow_id UUID REFERENCES workflows (id),
    "description" VARCHAR,
    "type" VARCHAR,
    "title" VARCHAR,
    "order" INTEGER,
    active BOOLEAN,
    tat INTEGER,
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

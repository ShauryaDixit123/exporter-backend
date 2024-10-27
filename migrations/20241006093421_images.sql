-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS images (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    s3_path VARCHAR NOT NULL,
    mime_type VARCHAR NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    uploaded_by UUID REFERENCES users (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON images;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON images
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON images;
DROP TABLE IF EXISTS images ;
-- +goose StatementEnd

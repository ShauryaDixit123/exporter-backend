-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS flow_instance_params_images (
    id SERIAL PRIMARY KEY,
    image_id UUID REFERENCES images (id),
    flow_instance_params_id UUID REFERENCES flow_instance_params(id),
    is_active BOOLEAN DEFAULT TRUE,
    uploaded_by UUID REFERENCES users (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON flow_instance_params_images;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON flow_instance_params_images
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON flow_instance_params_images;
DROP TABLE IF EXISTS flow_instance_params_images ;
-- +goose StatementEnd

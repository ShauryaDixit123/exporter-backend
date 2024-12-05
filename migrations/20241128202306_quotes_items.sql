-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quotes_items (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    quote_id UUID REFERENCES quotes (id),
    item_code VARCHAR(64),
    image_id UUID DEFAULT NULL,
    store_id VARCHAR(128) DEFAULT NULL,
    title VARCHAR(64),
    description VARCHAR(64) NOT NULL,
    quantity INT NOT NULL,
    quantity_unit VARCHAR(16) NOT NULL,
    rate INT,
    rate_unit VARCHAR(10) REFERENCES currencies (iso_code),
    delivery_date VARCHAR(16) NOT NULL,
    created_on TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON quotes_items;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON quotes_items
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON quotes_items;

DROP TABLE IF EXISTS quotes_items;
-- +goose StatementEnd


-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS po_so_lineitems (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    po_id VARCHAR(64) NOT NULL,
    so_id VARCHAR(64),
    li_ref_id VARCHAR(64),
    item_code VARCHAR(64) NOT NULL,
    description VARCHAR(64) NOT NULL,
    batch_count INT , 
    quantity INT NOT NULL,
    og BOOLEAN NOT NULL,
    delivery_date VARCHAR(16) NOT NULL,
    created_on TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON po_so_lineitems;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON po_so_lineitems
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON po_so_lineitems;

DROP TABLE IF EXISTS po_so_lineitems;
-- +goose StatementEnd


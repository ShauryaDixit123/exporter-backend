-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sales_orders (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    flow_instance_id UUID,
    flow_instance_params_id UUID,
    po_id UUID REFERENCES purchase_orders (id),
    so_number INT,
    supplier_id UUID REFERENCES users (id),
    due_date VARCHAR,
    status VARCHAR(64),
    rejection_reason VARCHAR(64),
    created_by UUID REFERENCES users (id),
    modified_by UUID,
    created_on TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON sales_orders;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON sales_orders
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON sales_orders;

DROP TABLE IF EXISTS sales_orders;
-- +goose StatementEnd


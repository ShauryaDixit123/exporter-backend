-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS purchase_orders (
    id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4() ,
    po_number INT,
    flow_instance_id UUID,
    flow_instance_params_id UUID,
    account_id INT REFERENCES accounts (id),
    buyer_id UUID REFERENCES users (id),
    supplier_id UUID REFERENCES users (id),
    pickup_location_id INT REFERENCES users_locations_map (id),
    drop_location_id INT REFERENCES users_locations_map (id),
    due_date VARCHAR,
    shipment_terms VARCHAR NOT NULL,
    terms_and_conditions VARCHAR,
    remarks VARCHAR,
    rejection_reason VARCHAR,
    status VARCHAR(16),
    shipment_mode VARCHAR,
    created_by UUID REFERENCES users (id),
    modified_by UUID REFERENCES users (id),
    created_on TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON purchase_orders;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON purchase_orders
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS set_timestamp ON purchase_orders;
DROP TABLE IF EXISTS purchase_orders;
-- +goose StatementEnd


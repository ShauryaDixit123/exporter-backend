-- +goose Up
-- +goose StatementBegin
 
CREATE TABLE IF NOT EXISTS request_for_quote (
	id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    account_id int REFERENCES accounts (id),
    buyer_id UUID REFERENCES users (id),
    "title" VARCHAR,
    "description" VARCHAR,
    inco_terms VARCHAR(18),
    pickup_location_id INT REFERENCES users_locations_map (id),
    drop_location_id INT REFERENCES users_locations_map (id),
    payment_terms VARCHAR(18),
    active BOOLEAN,
    tat INTEGER,
    due_date DATE,
    "status" VARCHAR(18),
    terms_and_conditions VARCHAR,
    created_by UUID REFERENCES users (id),
    remarks VARCHAR,
    rejection_reason VARCHAR,
    created_on TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON request_for_quote;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON request_for_quote
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

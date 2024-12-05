-- +goose Up
-- +goose StatementBegin
 
CREATE TABLE IF NOT EXISTS quotes (
	id UUID UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    rfq_id UUID REFERENCES request_for_quote (id),
    supplier_id UUID REFERENCES users (id),
    active BOOLEAN,
    due_date DATE,
    "status" VARCHAR(18),
    terms_and_conditions VARCHAR,
    remarks VARCHAR,
    rejection_reason VARCHAR,
    created_on TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    modified_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

DROP TRIGGER IF EXISTS set_timestamp ON quotes;
CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON quotes
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

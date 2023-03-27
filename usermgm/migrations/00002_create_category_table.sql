-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category (
	id BIGSERIAL,
    name VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS category ;
-- +goose StatementEnd

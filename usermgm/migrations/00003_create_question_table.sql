-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS question (
    id BIGSERIAL PRIMARY KEY,
    category_id INT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    published_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS question ;
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS answere (
	id BIGSERIAL,
    user_id INT NOT NULL,
    question_id INT NOT NULL,
    answere text NOT NULL,
    is_correct BOOLEAN NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ DEFAULT NULL,

	PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS answere ;
-- +goose StatementEnd

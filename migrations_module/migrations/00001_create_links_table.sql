-- +goose Up
CREATE TABLE IF NOT EXISTS links (
                                     id BIGSERIAL PRIMARY KEY,
                                     full_link TEXT NOT NULL,
                                     short_link TEXT NOT NULL
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS links;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
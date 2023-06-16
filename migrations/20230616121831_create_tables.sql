-- +goose Up
-- +goose StatementBegin
CREATE TABLE "runs"
(
    id   SERIAL UNIQUE PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE "url_state" AS ENUM (
    'OK',
    'NOT OK'
);

CREATE TABLE "urls"
(
    id   BIGSERIAL,
    run_id INTEGER NOT NULL REFERENCES runs(id),
    url VARCHAR(128) NOT NULL,
    state url_state NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "urls";
DROP TYPE "url_state";
DROP TABLE "runs";
-- +goose StatementEnd

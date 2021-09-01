-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "joke" (
    id serial primary key,
    text varchar,
    author_id bigint
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "joke";
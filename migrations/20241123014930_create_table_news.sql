-- +goose Up
-- +goose StatementBegin
CREATE TABLE news
(
    id serial not null primary key,
    title text not null,
    content text not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists news;
-- +goose StatementEnd

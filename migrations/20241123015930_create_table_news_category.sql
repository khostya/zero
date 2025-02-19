-- +goose Up
-- +goose StatementBegin
create table if not exists news_category
(
    news_id     int NOT NULL,
    category_id int NOT NULL,
    primary key (news_id, category_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table news_category;
-- +goose StatementEnd

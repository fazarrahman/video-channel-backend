-- +goose Up
-- +goose StatementBegin
create table users (
id serial primary key,
username varchar(50) unique not null,
password_hash varchar(100) not null,
email varchar(100) unique not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd

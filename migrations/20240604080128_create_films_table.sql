-- +goose Up
-- +goose StatementBegin
create table films (
id serial primary key,
title varchar(255) not null,
description text not null,
image_thumbnail varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table films;
-- +goose StatementEnd

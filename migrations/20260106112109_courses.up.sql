create table courses
(
    id           uuid primary key,
    title        varchar(100)               not null unique,
    description  text,
    access_type  varchar(100)               not null,
    published_at timestamp default null,
    owner_id     uuid references users (id) not null,
    deleted_at   timestamp default null
);
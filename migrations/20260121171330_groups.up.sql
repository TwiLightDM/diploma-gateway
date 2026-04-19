create table groups
(
    id          uuid primary key,
    title       varchar(100) not null unique,
    description text,
    owner_id    uuid references users (id) not null
);
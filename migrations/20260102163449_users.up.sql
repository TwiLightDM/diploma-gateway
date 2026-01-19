create table users
(
    id         uuid primary key,
    email      varchar(50)  not null unique,
    password   varchar(100) not null,
    full_name  varchar(100) not null,
    salt       varchar(20)  not null,
    role       varchar(20)  not null,
    deleted_at timestamp default null
);
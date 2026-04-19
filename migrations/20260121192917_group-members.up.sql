create table group_members
(
    id       uuid primary key,
    user_id  uuid references users (id)  not null,
    group_id uuid references groups (id) not null
);
create table completed_modules
(
    user_id   uuid references users (id)   not null,
    module_id uuid references modules (id) not null,

    primary key (user_id, module_id)
);
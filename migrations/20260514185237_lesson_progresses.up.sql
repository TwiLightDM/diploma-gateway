create table lesson_progresses(
    user_id   uuid references users (id)   not null,
    lesson_id uuid references lessons (id) not null,

    primary key (user_id, lesson_id)
);
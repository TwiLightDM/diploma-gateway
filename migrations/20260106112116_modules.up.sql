create table modules
(
    id           uuid primary key,
    title        varchar(100)                 not null unique,
    description  text,
    position     int                          not null,
    course_id    uuid references courses (id) not null,
    deleted_at   timestamp default null
);

create unique index uniq_modules_course_position
    on modules (course_id, position)
    where deleted_at is null;
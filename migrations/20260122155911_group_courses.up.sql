create table group_courses
(
    id        uuid primary key,
    group_id  uuid references groups (id)  not null,
    course_id uuid references courses (id) not null
);
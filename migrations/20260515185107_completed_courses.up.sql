create table completed_courses
(
    user_id   uuid references users (id)   not null,
    course_id uuid references courses (id) not null,

    primary key (user_id, course_id)
);
create index if not exists idx_modules_course_position
    on modules (course_id, position)
    where deleted_at is null;

create function reindex_modules_positions()
    returns trigger as
$$
begin
    update modules
    set position = position + 100000
    where course_id = old.course_id
      and deleted_at is null;

    update modules
    set position = sub.new_position
    from (
             select id,
                    row_number() over (order by position, id) as new_position
             from modules
             where course_id = old.course_id
               and deleted_at is null
         ) sub
    where modules.id = sub.id;

    return null;
end;
$$ language plpgsql;

create trigger trg_reindex_lessons_after_soft_delete
    after update of deleted_at on modules
    for each row
    when (old.deleted_at is null and new.deleted_at is not null)
execute function reindex_modules_positions();
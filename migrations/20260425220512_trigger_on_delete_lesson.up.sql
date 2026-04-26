create index if not exists idx_lessons_module_position
    on lessons (module_id, position)
    where deleted_at is null;

create function reindex_lessons_positions()
    returns trigger as
$$
begin
    update lessons
    set position = position + 100000
    where module_id = old.module_id
      and deleted_at is null;

    update lessons
    set position = sub.new_position
    from (
             select id,
                    row_number() over (order by position, id) as new_position
             from lessons
             where module_id = old.module_id
               and deleted_at is null
         ) sub
    where lessons.id = sub.id;

    return null;
end;
$$ language plpgsql;

create trigger trg_reindex_lessons_after_soft_delete
    after update of deleted_at on lessons
    for each row
    when (old.deleted_at is null and new.deleted_at is not null)
execute function reindex_lessons_positions();
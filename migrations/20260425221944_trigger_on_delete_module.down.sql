drop trigger if exists trg_reindex_modules_after_soft_delete on modules;

drop function if exists reindex_modules_positions();

drop index if exists idx_modules_course_position;
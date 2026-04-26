drop trigger if exists trg_reindex_lessons_after_soft_delete on lessons;

drop function if exists reindex_lessons_positions();

drop index if exists idx_lessons_module_position;
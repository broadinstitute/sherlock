alter table v2_environments
    drop column if exists prevent_deletion,
    drop column if exists delete_after;

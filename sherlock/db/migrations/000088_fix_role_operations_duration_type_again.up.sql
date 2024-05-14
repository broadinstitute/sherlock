alter table if exists role_operations
    alter column to_default_glass_break_duration type bigint using to_default_glass_break_duration::bigint;

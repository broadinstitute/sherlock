alter table if exists role_operations
    alter column from_default_glass_break_duration type bigint using from_default_glass_break_duration::bigint;

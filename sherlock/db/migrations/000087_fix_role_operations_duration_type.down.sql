alter table if exists role_operations
    alter column from_default_glass_break_duration type text using from_default_glass_break_duration::text;

alter table if exists role_operations
    alter column to_default_glass_break_duration type text using to_default_glass_break_duration::text;

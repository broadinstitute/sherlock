alter table users
    add column if not exists deactivated_at timestamp with time zone;

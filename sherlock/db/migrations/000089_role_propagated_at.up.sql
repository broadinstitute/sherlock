alter table roles
    add column if not exists propagated_at timestamp with time zone;

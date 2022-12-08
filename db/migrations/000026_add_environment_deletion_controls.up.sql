alter table v2_environments
    add column if not exists prevent_deletion boolean not null default false,
    add column if not exists delete_after timestamp;

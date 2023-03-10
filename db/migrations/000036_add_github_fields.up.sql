alter table v2_users
    add if not exists github_username text;

alter table v2_users
    add if not exists github_id text;

alter table v2_users
    add if not exists name text;

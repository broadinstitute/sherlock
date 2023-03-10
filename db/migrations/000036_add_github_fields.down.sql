alter table v2_users
    drop column if exists github_username;

alter table v2_users
    drop column if exists github_id;

alter table v2_users
    drop column if exists name;

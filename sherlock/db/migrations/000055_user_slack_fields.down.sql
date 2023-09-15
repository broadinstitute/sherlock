alter table v2_users
    drop constraint if exists name_from_valid;

alter table v2_users
    drop column if exists name_from;

alter table v2_users
    drop column if exists slack_username;

alter table v2_users
    drop column if exists slack_id;

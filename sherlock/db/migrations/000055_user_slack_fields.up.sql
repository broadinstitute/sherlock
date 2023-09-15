alter table v2_users
    add column if not exists slack_id text;

alter table v2_users
    add column if not exists slack_username text;

alter table v2_users
    add column if not exists name_from text;

alter table v2_users
    add constraint name_from_valid
    check (name_from is null
        or name_from = 'sherlock'
        or name_from = 'github'
        or name_from = 'slack');

update v2_users
    set name_from = 'github'
    where name_inferred_from_github is true;

update v2_users
    set name_from = 'sherlock'
    where name_inferred_from_github is false and name is not null;


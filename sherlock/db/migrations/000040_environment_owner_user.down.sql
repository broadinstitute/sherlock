alter table v2_environments
    drop constraint if exists fk_v2_environments_owner;

alter table v2_environments
    drop column if exists owner_id;

alter table v2_environments
    alter column legacy_owner set not null;

alter table v2_environments
    rename column legacy_owner to owner;

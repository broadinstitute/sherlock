alter table v2_environments
    rename column owner to legacy_owner;

alter table v2_environments
    alter column legacy_owner drop not null;

alter table v2_environments
    add if not exists owner_id bigint;

alter table v2_environments
    add constraint fk_v2_environments_owner
        foreign key (owner_id) references v2_users;

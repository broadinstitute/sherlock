alter table v2_changesets
    add if not exists planned_by_id bigint;

alter table v2_changesets
    add constraint fk_v2_changesets_planned_by
        foreign key (planned_by_id) references v2_users;

alter table v2_changesets
    add if not exists applied_by_id bigint;

alter table v2_changesets
    add constraint fk_v2_changesets_applied_by
        foreign key (applied_by_id) references v2_users;

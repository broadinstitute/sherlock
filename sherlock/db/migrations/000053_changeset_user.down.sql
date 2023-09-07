alter table v2_changesets
    drop constraint if exists fk_v2_changesets_applied_by;

alter table v2_changesets
    drop column if exists applied_by_id;

alter table v2_changesets
    drop constraint if exists fk_v2_changesets_planned_by;

alter table v2_changesets
    drop column if exists planned_by_id;

alter table clusters
    add column required_role_id bigint;

alter table clusters
    add constraint fk_clusters_required_role
        foreign key (required_role_id) references roles;

-- Now formally allow nulls in requires_suitability column to match with the semantics of required_role_id
alter table clusters
    alter column requires_suitability drop not null;

alter table environments
    add column required_role_id bigint;

alter table environments
    add constraint fk_environments_required_role
        foreign key (required_role_id) references roles;

-- The not null requirement on requires_suitability here is lifecycle-dependent, see 000062_environment_v3.up.sql,
-- so we need to drop the not null constraint here to match the semantics of required_role_id.
-- We do this by adding a new constraint as not valid, so that it commits instantly, and then validate it.
-- This requires much, much less locking than a normal add constraint.
-- https://www.postgresql.org/docs/current/sql-altertable.html#SQL-ALTERTABLE-NOTES
alter table environments
    add constraint lifecycle_valid_temp
        check ((lifecycle = 'template' and
                template_environment_id is null) or
               (lifecycle = 'dynamic' and
                template_environment_id is not null and
                base is not null and base != '' and
                default_cluster_id is not null) or
               (lifecycle = 'static' and
                base is not null and base != '' and
                default_cluster_id is not null)) not valid;

alter table environments
    validate constraint lifecycle_valid_temp;

alter table environments
    drop constraint lifecycle_valid;

alter table environments
    rename constraint lifecycle_valid_temp to lifecycle_valid;

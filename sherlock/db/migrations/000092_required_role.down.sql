alter table clusters
    alter column requires_suitability set not null;

alter table environments
    drop constraint if exists fk_environments_required_role;

alter table environments
    drop column if exists required_role_id;

alter table environments
    add constraint lifecycle_valid_temp
        check ((lifecycle = 'template' and
                template_environment_id is null) or
               (lifecycle = 'dynamic' and
                template_environment_id is not null and
                base is not null and base != '' and
                default_cluster_id is not null and
                requires_suitability is not null) or
               (lifecycle = 'static' and
                base is not null and base != '' and
                default_cluster_id is not null and
                requires_suitability is not null));

alter table environments
    validate constraint lifecycle_valid_temp;

alter table environments
    drop constraint lifecycle_valid;

alter table environments
    rename constraint lifecycle_valid_temp to lifecycle_valid;

alter table clusters
    drop constraint if exists fk_clusters_required_role;

alter table clusters
    drop column if exists required_role_id;

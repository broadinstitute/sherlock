alter table environments
    add constraint name_valid
        check (name != ''  and name similar to '[a-z0-9]([-a-z0-9]*[a-z0-9])?');

alter table environments
    add constraint owner_id_present
        check (owner_id is not null or
               (legacy_owner is not null or legacy_owner != ''));

alter table environments
    add constraint lifecycle_valid
        check ((lifecycle = 'template' and template_environment_id is null) or
               (lifecycle = 'dynamic' and template_environment_id is not null) or
               (lifecycle = 'static' and
                base != '' and
                default_cluster_id is not null and
                requires_suitability is not null));

alter table environments
    add constraint default_namespace_present
        check (default_namespace != '');

alter table environments
    add constraint helmfile_ref_present
        check (helmfile_ref is not null and helmfile_ref != '');

alter table environments
    add constraint unique_resource_prefix_present
        check (unique_resource_prefix != '');

alter table environments
    add constraint delete_after_valid
        check (delete_after is null or
               (lifecycle = 'dynamic' and
               (prevent_deletion is null or prevent_deletion is false)));

alter table environments
    add constraint offline_valid
        check (lifecycle = 'dynamic' or
               (offline is null or offline is false) and
               (offline_schedule_begin_enabled is null or offline_schedule_begin_enabled is false) and
               (offline_schedule_end_enabled is null or offline_schedule_end_enabled is false));

alter table environments
    add constraint offline_schedule_begin_time_present
        check (offline_schedule_begin_enabled is null or
               offline_schedule_begin_enabled is false or
               offline_schedule_begin_time is not null);

alter table environments
    add constraint offline_schedule_end_time_present
        check (offline_schedule_end_enabled is null or
               offline_schedule_end_enabled is false or
               offline_schedule_end_time is not null);

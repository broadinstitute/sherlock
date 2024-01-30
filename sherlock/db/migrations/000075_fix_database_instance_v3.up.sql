alter table database_instances
    add constraint google_info_present
    check (platform != 'google' or (google_project is not null and google_project != '' and instance_name is not null and instance_name != ''));

alter table database_instances
    add constraint azure_info_present
    check (platform != 'azure' or (instance_name is not null and instance_name != ''));

alter table database_instances
    add constraint platform_valid
    check (platform = 'google' or platform = 'azure' or platform = 'kubernetes');

alter table database_instances
    drop constraint if exists platform_present;

create unique index database_instances_chart_release_unique
    on database_instances (chart_release_id)
    where database_instances.deleted_at is null;

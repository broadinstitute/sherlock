alter table database_instances
    add constraint platform_present
        check (platform is not null and
               ((platform = 'google' and google_project is not null and google_project != '') or
                (platform = 'azure' and instance_name is not null and instance_name != '') or
                (platform = 'kubernetes')));

alter table database_instances
    drop constraint if exists google_info_present;

alter table database_instances
    drop constraint if exists azure_info_present;

alter table database_instances
    drop constraint if exists platform_valid;

drop index if exists database_instances_chart_release_unique;

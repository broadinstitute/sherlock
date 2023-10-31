alter table database_instances
    add constraint chart_release_id_present
        check (chart_release_id is not null and chart_release_id != 0);

alter table database_instances
    add constraint platform_present
    check (platform is not null and
            ((platform = 'google' and google_project is not null and google_project != '') or
             (platform = 'azure' and instance_name is not null and instance_name != '') or
             (platform = 'kubernetes')));

alter table database_instances
    add constraint default_database_present
        check (default_database is not null and default_database != '');

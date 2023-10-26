alter table database_instances
    drop constraint if exists chart_release_id_present;

alter table database_instances
    drop constraint if exists platform_present;

alter table database_instances
    drop constraint if exists default_database_present;

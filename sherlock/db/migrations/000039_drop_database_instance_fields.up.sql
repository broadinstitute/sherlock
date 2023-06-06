alter table v2_database_instances
    drop column if exists google_location;

alter table v2_database_instances
    drop column if exists azure_managed_resource_group;

alter table v2_database_instances
    drop column if exists azure_subscription;

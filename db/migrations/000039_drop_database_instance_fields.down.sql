alter table v2_database_instances
    add if not exists google_location text;

alter table v2_database_instances
    add if not exists azure_managed_resource_group text;

alter table v2_database_instances
    add if not exists azure_subscription text;

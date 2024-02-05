create unique index environments_name_unique
    on environments (name)
    where deleted_at is null;

create unique index environments_urp_unique
    on environments (unique_resource_prefix)
    where deleted_at is null;

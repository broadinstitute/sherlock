drop index if exists roles_grants_dev_firecloud_group_unique;

create unique index if not exists roles_grants_dev_firecloud_group_unique
    on roles (grants_dev_firecloud_group)
    where deleted_at is null and grants_dev_firecloud_group is not null and grants_dev_firecloud_group != '';

drop index if exists roles_grants_dev_azure_group_unique;

create unique index if not exists roles_grants_dev_azure_group_unique
    on roles (grants_dev_azure_group)
    where deleted_at is null and grants_dev_azure_group is not null and grants_dev_azure_group != '';

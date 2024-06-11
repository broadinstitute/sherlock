alter table roles
    add column if not exists grants_qa_firecloud_group text;

create unique index if not exists roles_grants_qa_firecloud_group_unique
    on roles (grants_qa_firecloud_group)
    where deleted_at is null and grants_qa_firecloud_group is not null and grants_qa_firecloud_group != '';

alter table role_operations
    add column if not exists from_grants_qa_firecloud_group text;

alter table role_operations
    add column if not exists to_grants_qa_firecloud_group text;

alter table roles
    add column if not exists grants_prod_firecloud_group text;

create unique index if not exists roles_grants_prod_firecloud_group_unique
    on roles (grants_prod_firecloud_group)
    where deleted_at is null and grants_prod_firecloud_group is not null and grants_prod_firecloud_group != '';

alter table role_operations
    add column if not exists from_grants_prod_firecloud_group text;

alter table role_operations
    add column if not exists to_grants_prod_firecloud_group text;

alter table roles
    add column if not exists grants_prod_azure_group text;

create unique index if not exists roles_grants_prod_azure_group_unique
    on roles (grants_prod_azure_group)
    where deleted_at is null and grants_prod_azure_group is not null and grants_prod_azure_group != '';

alter table role_operations
    add column if not exists from_grants_prod_azure_group text;

alter table role_operations
    add column if not exists to_grants_prod_azure_group text;

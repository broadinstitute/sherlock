alter table roles
    add column if not exists grants_dev_firecloud_folder_owner text;

create unique index if not exists roles_grants_dev_firecloud_folder_owner_unique
    on roles (grants_dev_firecloud_folder_owner)
    where deleted_at is null and grants_dev_firecloud_folder_owner is not null and grants_dev_firecloud_folder_owner != '';

alter table role_operations
    add column if not exists from_grants_dev_firecloud_folder_owner text;

alter table role_operations
    add column if not exists to_grants_dev_firecloud_folder_owner text;

alter table roles
    add column if not exists grants_qa_firecloud_folder_owner text;

create unique index if not exists roles_grants_qa_firecloud_folder_owner_unique
    on roles (grants_qa_firecloud_folder_owner)
    where deleted_at is null and grants_qa_firecloud_folder_owner is not null and grants_qa_firecloud_folder_owner != '';

alter table role_operations
    add column if not exists from_grants_qa_firecloud_folder_owner text;

alter table role_operations
    add column if not exists to_grants_qa_firecloud_folder_owner text;

alter table roles
    add column if not exists grants_prod_firecloud_folder_owner text;

create unique index if not exists roles_grants_prod_firecloud_folder_owner_unique
    on roles (grants_prod_firecloud_folder_owner)
    where deleted_at is null and grants_prod_firecloud_folder_owner is not null and grants_prod_firecloud_folder_owner != '';

alter table role_operations
    add column if not exists from_grants_prod_firecloud_folder_owner text;

alter table role_operations
    add column if not exists to_grants_prod_firecloud_folder_owner text;

drop index if exists roles_grants_dev_firecloud_folder_owner_unique;

alter table roles
    drop column if exists grants_dev_firecloud_folder_owner;

alter table role_operations
    drop column if exists from_grants_dev_firecloud_folder_owner;

alter table role_operations
    drop column if exists to_grants_dev_firecloud_folder_owner;

drop index if exists roles_grants_qa_firecloud_folder_owner_unique;

alter table roles
    drop column if exists grants_qa_firecloud_folder_owner;

alter table role_operations
    drop column if exists from_grants_qa_firecloud_folder_owner;

alter table role_operations
    drop column if exists to_grants_qa_firecloud_folder_owner;

drop index if exists roles_grants_prod_firecloud_folder_owner_unique;

alter table roles
    drop column if exists grants_prod_firecloud_folder_owner;

alter table role_operations
    drop column if exists from_grants_prod_firecloud_folder_owner;

alter table role_operations
    drop column if exists to_grants_prod_firecloud_folder_owner;

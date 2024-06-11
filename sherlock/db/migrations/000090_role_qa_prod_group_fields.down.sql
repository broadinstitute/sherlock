drop index if exists roles_grants_qa_firecloud_group_unique;

alter table roles
    drop column if exists grants_qa_firecloud_group;

alter table role_operations
    drop column if exists from_grants_qa_firecloud_group;

alter table role_operations
    drop column if exists to_grants_qa_firecloud_group;

drop index if exists roles_grants_prod_firecloud_group_unique;

alter table roles
    drop column if exists grants_prod_firecloud_group;

alter table role_operations
    drop column if exists from_grants_prod_firecloud_group;

alter table role_operations
    drop column if exists to_grants_prod_firecloud_group;

drop index if exists roles_grants_prod_azure_group_unique;

alter table roles
    drop column if exists grants_prod_azure_group;

alter table role_operations
    drop column if exists from_grants_prod_azure_group;

alter table role_operations
    drop column if exists to_grants_prod_azure_group;

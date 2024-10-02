drop index if exists roles_grants_dev_azure_account_unique;

alter table roles
    drop column if exists grants_dev_azure_account;

alter table role_operations
    drop column if exists from_grants_dev_azure_account;

alter table role_operations
    drop column if exists to_grants_dev_azure_account;

drop index if exists roles_grants_prod_azure_account_unique;

alter table roles
    drop column if exists grants_prod_azure_account;

alter table role_operations
    drop column if exists from_grants_prod_azure_account;

alter table role_operations
    drop column if exists to_grants_prod_azure_account;

drop index if exists roles_grants_dev_azure_directory_roles_unique;

alter table roles
    drop column if exists grants_dev_azure_directory_roles;

alter table role_operations
    drop column if exists from_grants_dev_azure_directory_roles;

alter table role_operations
    drop column if exists to_grants_dev_azure_directory_roles;

drop index if exists roles_grants_prod_azure_directory_roles_unique;

alter table roles
    drop column if exists grants_prod_azure_directory_roles;

alter table role_operations
    drop column if exists from_grants_prod_azure_directory_roles;

alter table role_operations
    drop column if exists to_grants_prod_azure_directory_roles;

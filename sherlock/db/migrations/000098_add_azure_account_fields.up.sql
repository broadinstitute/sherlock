alter table roles
    add column if not exists grants_dev_azure_account boolean;

create unique index if not exists roles_grants_dev_azure_account_unique
    on roles (grants_dev_azure_account)
    where deleted_at is null and grants_dev_azure_account is not null and grants_dev_azure_account is true;

alter table role_operations
    add column if not exists from_grants_dev_azure_account boolean;

alter table role_operations
    add column if not exists to_grants_dev_azure_account boolean;

alter table roles
    add column if not exists grants_prod_azure_account boolean;

create unique index if not exists roles_grants_prod_azure_account_unique
    on roles (grants_prod_azure_account)
    where deleted_at is null and grants_prod_azure_account is not null and grants_prod_azure_account is true;

alter table role_operations
    add column if not exists from_grants_prod_azure_account boolean;

alter table role_operations
    add column if not exists to_grants_prod_azure_account boolean;

alter table roles
    add column if not exists grants_dev_azure_directory_roles boolean;

create unique index if not exists roles_grants_dev_azure_directory_roles_unique
    on roles (grants_dev_azure_directory_roles)
    where deleted_at is null and grants_dev_azure_directory_roles is not null and grants_dev_azure_directory_roles is true;

alter table role_operations
    add column if not exists from_grants_dev_azure_directory_roles boolean;

alter table role_operations
    add column if not exists to_grants_dev_azure_directory_roles boolean;

alter table roles
    add column if not exists grants_prod_azure_directory_roles boolean;

create unique index if not exists roles_grants_prod_azure_directory_roles_unique
    on roles (grants_prod_azure_directory_roles)
    where deleted_at is null and grants_prod_azure_directory_roles is not null and grants_prod_azure_directory_roles is true;

alter table role_operations
    add column if not exists from_grants_prod_azure_directory_roles boolean;

alter table role_operations
    add column if not exists to_grants_prod_azure_directory_roles boolean;

drop index if exists roles_grants_broad_institute_group_unique;

alter table roles
    drop column if exists grants_broad_institute_group;

alter table role_operations
    drop column if exists from_grants_broad_institute_group;

alter table role_operations
    drop column if exists to_grants_broad_institute_group;

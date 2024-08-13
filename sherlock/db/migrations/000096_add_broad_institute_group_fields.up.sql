alter table roles
    add column if not exists grants_broad_institute_group text;

create unique index if not exists roles_grants_broad_institute_group_unique
    on roles (grants_broad_institute_group)
    where deleted_at is null and grants_broad_institute_group is not null and grants_broad_institute_group != '';

alter table role_operations
    add column if not exists from_grants_broad_institute_group text;

alter table role_operations
    add column if not exists to_grants_broad_institute_group text;

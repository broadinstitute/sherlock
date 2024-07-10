alter table roles
    drop column if exists auto_assign_all_users;

alter table role_operations
    drop column if exists from_auto_assign_all_users;

alter table role_operations
    drop column if exists to_auto_assign_all_users;

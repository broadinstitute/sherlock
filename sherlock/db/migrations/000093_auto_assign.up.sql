alter table roles
    add column if not exists auto_assign_all_users boolean;

alter table role_operations
    add column if not exists from_auto_assign_all_users boolean;

alter table role_operations
    add column if not exists to_auto_assign_all_users boolean;

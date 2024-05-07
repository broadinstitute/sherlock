create table if not exists roles
(
    id                             bigserial
        primary key,
    created_at                     timestamp with time zone,
    updated_at                     timestamp with time zone,
    deleted_at                     timestamp with time zone,

    name                           text unique not null,
    suspend_non_suitable_users     boolean,
    can_be_glass_broken_by_role_id bigint
        constraint fk_roles_can_be_glass_broken_by_role_id
            references roles,
    default_glass_break_duration   text,
    grants_sherlock_super_admin    boolean,
    grants_dev_firecloud_group     text,
    grants_dev_azure_group         text
);

create index if not exists idx_roles_deleted_at
    on roles (deleted_at);

create unique index if not exists roles_grants_sherlock_super_admin_unique
    on roles (grants_sherlock_super_admin)
    where deleted_at is null and grants_sherlock_super_admin is true;

create unique index if not exists roles_grants_dev_firecloud_group_unique
    on roles (grants_dev_firecloud_group)
    where deleted_at is null and grants_dev_firecloud_group is not null;

create unique index if not exists roles_grants_dev_azure_group_unique
    on roles (grants_dev_azure_group)
    where deleted_at is null and grants_dev_azure_group is not null;

create table if not exists role_operations
(
    id                                  bigserial
        primary key,
    created_at                          timestamp with time zone,
    updated_at                          timestamp with time zone,
    deleted_at                          timestamp with time zone,

    role_id                             bigint not null
        constraint fk_role_operations_role_id
            references roles,
    author_id                           bigint not null
        constraint fk_role_operations_author_id
            references users,

    operation                           text   not null
        constraint operation_valid check (operation in ('create', 'update', 'delete')),

    from_name                           text   not null,
    from_suspend_non_suitable_users     boolean,
    from_can_be_glass_broken_by_role_id bigint
        constraint fk_roles_from_can_be_glass_broken_by_role_id
            references roles,
    from_default_glass_break_duration   text,
    from_grants_sherlock_super_admin    boolean,
    from_grants_dev_firecloud_group     text,
    from_grants_dev_azure_group         text,

    to_name                             text   not null,
    to_suspend_non_suitable_users       boolean,
    to_can_be_glass_broken_by_role_id   bigint
        constraint fk_roles_to_can_be_glass_broken_by_role_id
            references roles,
    to_default_glass_break_duration     text,
    to_grants_sherlock_super_admin      boolean,
    to_grants_dev_firecloud_group       text,
    to_grants_dev_azure_group           text
);

create index if not exists idx_role_operations_deleted_at
    on role_operations (deleted_at);

create table if not exists role_assignments
(
    role_id    bigint not null
        constraint fk_role_assignments_role_id
            references roles
            on update cascade on delete cascade,
    user_id    bigint not null
        constraint fk_role_assignments_user_id
            references users
            on update cascade on delete cascade,
    primary key (role_id, user_id),

    suspended  boolean not null,
    expires_at timestamp with time zone
);

create table if not exists role_assignment_operations
(
    id              bigserial
        primary key,
    created_at      timestamp with time zone,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone,

    role_id         bigint not null
        constraint fk_roles_assignment_operations_role_id
            references roles,
    user_id         bigint not null
        constraint fk_roles_assignment_operations_user_id
            references users,
    author_id       bigint not null
        constraint fk_roles_assignment_operations_author_id
            references users,

    operation       text   not null
        constraint operation_valid check (operation in ('create', 'update', 'delete')),

    from_suspended  boolean,
    from_expires_at timestamp with time zone,

    to_suspended    boolean,
    to_expires_at   timestamp with time zone
);

create index if not exists idx_roles_assignment_operations_deleted_at
    on role_assignment_operations (deleted_at);

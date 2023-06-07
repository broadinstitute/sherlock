create table if not exists v2_users
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email      text not null
        unique,
    google_id  text not null
        unique
);

create index if not exists idx_v2_users_deleted_at
    on v2_users (deleted_at);

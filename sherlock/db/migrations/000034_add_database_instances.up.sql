create table if not exists v2_database_instances
(
    id                           bigserial
        primary key,
    created_at                   timestamp with time zone,
    updated_at                   timestamp with time zone,
    deleted_at                   timestamp with time zone,
    chart_release_id             bigint
        constraint fk_v2_database_instances_chart_release
            references v2_chart_releases,
    platform                     text,
    google_project               text,
    google_location              text,
    azure_subscription           text,
    azure_managed_resource_group text,
    instance_name                text,
    default_database             text
);

create index if not exists idx_v2_database_instances_deleted_at
    on v2_database_instances (deleted_at);
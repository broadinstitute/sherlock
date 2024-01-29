create table if not exists incidents
(
    id                  bigserial
        primary key,
    created_at          timestamp with time zone,
    updated_at          timestamp with time zone,
    deleted_at          timestamp with time zone,

    ticket              text,
    description         text,
    started_at          timestamp with time zone,
    remediated_at       timestamp with time zone,
    review_completed_at timestamp with time zone
);

create index if not exists idx_incidents_deleted_at
    on incidents (deleted_at);

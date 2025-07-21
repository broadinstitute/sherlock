create table if not exists service_alerts
(
    id                             bigserial primary key,
    created_at                     timestamp with time zone,
    updated_at                     timestamp with time zone,
    deleted_at                     timestamp with time zone,

    on_environment_id   bigint
        constraint fk_service_alerts_on_environment
            references environments,

    title                           text not null,
	alert_message                           text not null,
	link                           text,
    severity   text,
    uuid    uuid, 
    created_by bigint
        constraint fk_service_alerts_created_by
            references users, 
    updated_by bigint
        constraint fk_service_alerts_updated_by
            references users, 
    deleted_by bigint
        constraint fk_service_alerts_deleted_by
            references users

);

create index if not exists svc_alerts_deleted_at
    on service_alerts (deleted_at);

create table if not exists service_alerts
(
    id                             bigserial
        primary key,
    created_at                     timestamp with time zone,
    updated_at                     timestamp with time zone,
    deleted_at                     timestamp with time zone,

    title                           text not null,
	message                           text not null,
	link                           text,
    severity   text
);

create index if not exists svc_alerts_deleted_at
    on service_alerts (deleted_at);

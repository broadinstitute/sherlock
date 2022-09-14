create table if not exists v2_chart_deploy_records
(
    id                  bigint default nextval('v2_chart_deploy_records_id_seq'::regclass) not null
        primary key,
    created_at          timestamp with time zone,
    updated_at          timestamp with time zone,
    deleted_at          timestamp with time zone,
    chart_release_id    bigint                                                             not null
        constraint fk_v2_chart_deploy_records_chart_release
            references v2_chart_releases,
    exact_chart_version text                                                               not null,
    exact_app_version   text,
    helmfile_ref        text                                                               not null
);

create index if not exists idx_v2_chart_deploy_records_deleted_at
    on v2_chart_deploy_records (deleted_at);

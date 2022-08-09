create sequence if not exists v2_clusters_id_seq;

create sequence if not exists v2_environments_id_seq;

create sequence if not exists v2_charts_id_seq;

create sequence if not exists v2_chart_versions_id_seq;

create sequence if not exists v2_app_versions_id_seq;

create sequence if not exists v2_chart_releases_id_seq;

create sequence if not exists v2_chart_deploy_records_id_seq;

create table if not exists v2_clusters
(
    id                   bigint default nextval('v2_clusters_id_seq'::regclass) not null
        primary key,
    created_at           timestamp with time zone,
    updated_at           timestamp with time zone,
    deleted_at           timestamp with time zone,
    name                 text                                                   not null
        unique,
    provider             text                                                   not null,
    google_project       text,
    azure_subscription   text,
    base                 text                                                   not null,
    address              text                                                   not null,
    requires_suitability boolean                                                not null
);

create index if not exists idx_v2_clusters_deleted_at
    on v2_clusters (deleted_at);

create table if not exists v2_environments
(
    id                      bigint default nextval('v2_environments_id_seq'::regclass) not null
        primary key,
    created_at              timestamp with time zone,
    updated_at              timestamp with time zone,
    deleted_at              timestamp with time zone,
    base                    text,
    lifecycle               text                                                       not null,
    name                    text                                                       not null
        unique,
    template_environment_id bigint
        constraint fk_v2_environments_template_environment
            references v2_environments,
    values_name             text,
    default_cluster_id      bigint
        constraint fk_v2_environments_default_cluster
            references v2_clusters,
    default_namespace       text,
    owner                   text                                                       not null,
    requires_suitability    boolean
);

create index if not exists idx_v2_environments_deleted_at
    on v2_environments (deleted_at);

create table if not exists v2_charts
(
    id                        bigint default nextval('v2_charts_id_seq'::regclass) not null
        primary key,
    created_at                timestamp with time zone,
    updated_at                timestamp with time zone,
    deleted_at                timestamp with time zone,
    name                      text                                                 not null
        unique,
    chart_repo                text                                                 not null,
    app_image_git_repo        text,
    app_image_git_main_branch text
);

create index if not exists idx_v2_charts_deleted_at
    on v2_charts (deleted_at);

create table if not exists v2_chart_versions
(
    id            bigint default nextval('v2_chart_versions_id_seq'::regclass) not null
        primary key,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone,
    chart_id      bigint                                                       not null
        constraint fk_v2_chart_versions_chart
            references v2_charts,
    chart_version text                                                         not null
);

create index if not exists idx_v2_chart_versions_deleted_at
    on v2_chart_versions (deleted_at);

create table if not exists v2_app_versions
(
    id          bigint default nextval('v2_app_versions_id_seq'::regclass) not null
        primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    chart_id    bigint                                                     not null
        constraint fk_v2_app_versions_chart
            references v2_charts,
    app_version text                                                       not null,
    git_commit  text,
    git_branch  text
);

create index if not exists idx_v2_app_versions_deleted_at
    on v2_app_versions (deleted_at);

create table if not exists v2_chart_releases
(
    id                          bigint default nextval('v2_chart_releases_id_seq'::regclass) not null
        primary key,
    created_at                  timestamp with time zone,
    updated_at                  timestamp with time zone,
    deleted_at                  timestamp with time zone,
    chart_id                    bigint
        constraint fk_v2_chart_releases_chart
            references v2_charts,
    cluster_id                  bigint
        constraint fk_v2_chart_releases_cluster
            references v2_clusters,
    destination_type            text,
    environment_id              bigint
        constraint fk_v2_chart_releases_environment
            references v2_environments,
    name                        text                                                         not null
        unique,
    namespace                   text,
    current_app_version_exact   text,
    current_chart_version_exact text,
    helmfile_ref                text,
    target_app_version_branch   text,
    target_app_version_commit   text,
    target_app_version_exact    text,
    target_app_version_use      text,
    target_chart_version_exact  text,
    target_chart_version_use    text                                                         not null,
    thelma_mode                 text
);

create index if not exists idx_v2_chart_releases_deleted_at
    on v2_chart_releases (deleted_at);

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
    exact_app_version   text                                                               not null,
    helmfile_ref        text                                                               not null
);

create index if not exists idx_v2_chart_deploy_records_deleted_at
    on v2_chart_deploy_records (deleted_at);

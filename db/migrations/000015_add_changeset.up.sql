alter table v2_chart_releases
    add if not exists resolved_at timestamp with time zone;

alter table v2_chart_releases
    rename column target_app_version_use to app_version_resolver;

alter table v2_chart_releases
    rename column current_app_version_exact to app_version_exact;

alter table v2_chart_releases
    drop column if exists target_app_version_exact;

alter table v2_chart_releases
    rename column target_app_version_branch to app_version_branch;

alter table v2_chart_releases
    rename column target_app_version_commit to app_version_commit;

alter table v2_chart_releases
    add if not exists app_version_id bigint;

alter table v2_chart_releases
    add constraint fk_v2_chart_releases_app_version
        foreign key (app_version_id) references v2_app_versions;

alter table v2_chart_releases
    rename column target_chart_version_use to chart_version_resolver;

alter table v2_chart_releases
    alter column chart_version_resolver drop not null;

alter table v2_chart_releases
    rename column current_chart_version_exact to chart_version_exact;

alter table v2_chart_releases
    drop column if exists target_chart_version_exact;

alter table v2_chart_releases
    add if not exists chart_version_id bigint;

alter table v2_chart_releases
    add constraint fk_v2_chart_releases_chart_version
        foreign key (chart_version_id) references v2_chart_versions;

alter table v2_chart_releases
    drop column if exists thelma_mode;

create sequence if not exists v2_changesets_id_seq;

create table if not exists v2_changesets
(
    id                          bigint default nextval('v2_changesets_id_seq'::regclass) not null
        primary key,
    created_at                  timestamp with time zone,
    updated_at                  timestamp with time zone,
    deleted_at                  timestamp with time zone,
    chart_release_id            bigint
        constraint fk_v2_changesets_chart_release
            references v2_chart_releases,
    from_resolved_at            timestamp with time zone,
    from_app_version_resolver   text,
    from_app_version_exact      text,
    from_app_version_branch     text,
    from_app_version_commit     text,
    from_app_version_id         bigint,
    from_chart_version_resolver text,
    from_chart_version_exact    text,
    from_chart_version_id       bigint,
    from_helmfile_ref           text,
    to_resolved_at              timestamp with time zone,
    to_app_version_resolver     text,
    to_app_version_exact        text,
    to_app_version_branch       text,
    to_app_version_commit       text,
    to_app_version_id           bigint
        constraint fk_v2_changesets_app_version
            references v2_app_versions,
    to_chart_version_resolver   text,
    to_chart_version_exact      text,
    to_chart_version_id         bigint
        constraint fk_v2_changesets_chart_version
            references v2_chart_versions,
    to_helmfile_ref             text,
    applied_at                  timestamp with time zone,
    superseded_at               timestamp with time zone
);

create index if not exists idx_v2_changesets_deleted_at
    on v2_changesets (deleted_at);

create table if not exists v2_changeset_new_app_versions
(
    changeset_id   bigint not null
        constraint fk_v2_changeset_new_app_versions_changeset
            references v2_changesets on delete cascade on update cascade,
    app_version_id bigint not null
        constraint fk_v2_changeset_new_app_versions_app_version
            references v2_app_versions on delete cascade on update cascade,
    primary key (changeset_id, app_version_id)
);

create table if not exists v2_changeset_new_chart_versions
(
    changeset_id     bigint not null
        constraint fk_v2_changeset_new_chart_versions_changeset
            references v2_changesets on delete cascade on update cascade,
    chart_version_id bigint not null
        constraint fk_v2_changeset_new_chart_versions_chart_version
            references v2_chart_versions on delete cascade on update cascade,
    primary key (changeset_id, chart_version_id)
);

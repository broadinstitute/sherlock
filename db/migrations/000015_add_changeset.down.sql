alter table v2_chart_releases
    drop column if exists resolved_at;

alter table v2_chart_releases
    rename column app_version_resolver to target_app_version_use;

alter table v2_chart_releases
    rename column app_version_exact to current_app_version_exact;

alter table v2_chart_releases
    add if not exists target_app_version_exact text;

alter table v2_chart_releases
    rename column app_version_branch to target_app_version_branch;

alter table v2_chart_releases
    rename column app_version_commit to target_app_version_commit;

alter table v2_chart_releases
    drop constraint if exists fk_v2_chart_releases_app_version;

alter table v2_chart_releases
    drop column if exists app_version_id;

alter table v2_chart_releases
    rename column chart_version_resolver to target_chart_version_use;

alter table v2_chart_releases
    alter column target_chart_version_use set not null;

alter table v2_chart_releases
    rename column chart_version_exact to current_chart_version_exact;

alter table v2_chart_releases
    add if not exists target_chart_version_exact text;

alter table v2_chart_releases
    drop constraint if exists fk_v2_chart_releases_chart_version;

alter table v2_chart_releases
    drop column if exists chart_version_id;

alter table v2_chart_releases
    add if not exists thelma_mode text;

drop table if exists v2_changesets;

drop table if exists v2_changeset_new_app_versions;

drop table if exists v2_changeset_new_chart_versions;

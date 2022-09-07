alter table v2_chart_versions
    drop constraint if exists fk_v2_chart_versions_parent_chart_version;

alter table v2_chart_versions
    drop column if exists parent_chart_version_id;

alter table v2_app_versions
    drop constraint if exists fk_v2_app_versions_parent_app_version;

alter table v2_app_versions
    drop column if exists parent_app_version_id;

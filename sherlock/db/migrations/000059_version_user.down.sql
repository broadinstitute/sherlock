alter table v2_chart_versions
    drop constraint fk_v2_chart_versions_authored_by;

alter table v2_chart_versions
    drop if exists authored_by_id;

alter table v2_app_versions
    drop constraint fk_v2_app_versions_authored_by;

alter table v2_app_versions
    drop if exists authored_by_id;

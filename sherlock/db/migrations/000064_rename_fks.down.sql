alter table app_versions
    rename constraint fk_app_versions_authored_by to fk_v2_app_versions_authored_by;

alter table changesets
    rename constraint fk_changesets_applied_by to fk_v2_changesets_applied_by;

alter table changesets
    rename constraint fk_changesets_planned_by to fk_v2_changesets_planned_by;

alter table chart_versions
    rename constraint fk_chart_versions_authored_by to fk_v2_chart_versions_authored_by;

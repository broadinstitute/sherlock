alter table v2_chart_versions
    add column if not exists parent_chart_version_id bigint;

/* add constraint doesn't have `if not exists` syntax */
alter table v2_chart_versions
    add constraint fk_v2_chart_versions_parent_chart_version
        foreign key (parent_chart_version_id) references v2_chart_versions;

alter table v2_app_versions
    add column if not exists parent_app_version_id bigint;

/* add constraint doesn't have `if not exists` syntax */
alter table v2_app_versions
    add constraint fk_v2_app_versions_parent_app_version
        foreign key (parent_app_version_id) references v2_app_versions;

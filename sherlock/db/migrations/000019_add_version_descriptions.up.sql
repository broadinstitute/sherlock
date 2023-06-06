alter table v2_app_versions
    add if not exists description text;

alter table v2_chart_versions
    add if not exists description text;

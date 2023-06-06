alter table v2_app_versions
    drop column if exists description;

alter table v2_chart_versions
    drop column  if exists description;

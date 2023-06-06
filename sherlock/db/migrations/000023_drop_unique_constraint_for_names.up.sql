alter table v2_chart_releases
    drop constraint if exists v2_chart_releases_name_key;

alter table v2_environments
    drop constraint if exists v2_environments_name_key;

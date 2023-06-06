alter table v2_chart_releases
    add constraint v2_chart_releases_name_key unique (name);

alter table v2_environments
    add constraint v2_environments_name_key unique (name);

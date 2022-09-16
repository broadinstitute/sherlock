alter table v2_environments
    drop column if exists base_domain;

alter table v2_environments
    drop column if exists name_prefixes_domain;

alter table v2_charts
    drop column if exists default_subdomain;

alter table v2_charts
    drop column if exists default_protocol;

alter table v2_charts
    drop column if exists default_port;

alter table v2_chart_releases
    drop column if exists thelma_mode;

alter table v2_chart_releases
    drop column if exists subdomain;

alter table v2_chart_releases
    drop column if exists protocol;

alter table v2_chart_releases
    drop column if exists port;

alter table v2_changesets
    drop column if exists from_thelma_mode;

alter table v2_changesets
    drop column if exists to_thelma_mode;

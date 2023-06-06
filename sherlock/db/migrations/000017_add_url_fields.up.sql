alter table v2_environments
    add if not exists base_domain text;

alter table v2_environments
    add if not exists name_prefixes_domain boolean;

alter table v2_charts
    add if not exists default_subdomain text;

alter table v2_charts
    add if not exists default_protocol text;

alter table v2_charts
    add if not exists default_port bigint;

alter table v2_chart_releases
    add if not exists thelma_mode text;

alter table v2_chart_releases
    add if not exists subdomain text;

alter table v2_chart_releases
    add if not exists protocol text;

alter table v2_chart_releases
    add if not exists port bigint;

alter table v2_changesets
    add if not exists from_thelma_mode text;

alter table v2_changesets
    add if not exists to_thelma_mode text;



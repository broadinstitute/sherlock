alter table v2_chart_releases
    add if not exists pagerduty_integration_id bigint;

alter table v2_environments
    add if not exists pagerduty_integration_id bigint;

create table if not exists v2_pagerduty_integrations
(
    id           bigserial
        primary key,
    created_at   timestamp with time zone,
    updated_at   timestamp with time zone,
    deleted_at   timestamp with time zone,
    name         text,
    pagerduty_id text,
    key          text,
    type         text
);


alter table v2_chart_releases
    add constraint fk_v2_chart_releases_pagerduty_integration
        foreign key (pagerduty_integration_id) references v2_pagerduty_integrations;

alter table v2_environments
    add constraint fk_v2_environments_pagerduty_integration
        foreign key (pagerduty_integration_id) references v2_pagerduty_integrations;

create index if not exists idx_v2_pagerduty_integrations_deleted_at
    on v2_pagerduty_integrations (deleted_at);

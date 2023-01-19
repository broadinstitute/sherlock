alter table v2_chart_releases
    drop constraint fk_v2_chart_releases_pagerduty_integration;

alter table v2_chart_releases
    drop column pagerduty_integration_id;

alter table v2_environments
    drop constraint fk_v2_environments_pagerduty_integration;

alter table v2_environments
    drop column pagerduty_integration_id;

drop table v2_pagerduty_integrations;

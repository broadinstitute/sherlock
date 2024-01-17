alter table pagerduty_integrations
    drop constraint if exists pagerduty_id_present;

drop index if exists pagerduty_integrations_pagerduty_id_unique_constraint;

alter table pagerduty_integrations
    drop constraint if exists pagerduty_name_present;

alter table pagerduty_integrations
    drop constraint if exists pagerduty_key_present;

alter table pagerduty_integrations
    drop constraint if exists pagerduty_type_present;

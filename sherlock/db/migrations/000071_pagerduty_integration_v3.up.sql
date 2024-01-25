alter table pagerduty_integrations
    add constraint pagerduty_id_present
        check (pagerduty_id is not null and pagerduty_id != '');

create unique index pagerduty_integrations_pagerduty_id_unique_constraint
    on pagerduty_integrations (pagerduty_id)
    where deleted_at is null;

alter table pagerduty_integrations
    add constraint pagerduty_name_present
        check (name is not null and name != '');

alter table pagerduty_integrations
    add constraint pagerduty_key_present
        check (key is not null and key != '');

alter table pagerduty_integrations
    add constraint pagerduty_type_present
        check (type is not null and type != '');

alter table v2_environments
    drop column if exists pact_identifier;

alter table v2_charts
    drop column if exists pact_participant;
        
alter table v2_environments
    add column pact_identifier uuid;

alter table v2_charts
    add column pact_participant boolean
    DEFAULT false;


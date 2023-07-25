alter table v2_environments
    add column participates_in_pact boolean
    DEFAULT false;

alter table v2_charts
    add column pact_participant boolean
    DEFAULT false;

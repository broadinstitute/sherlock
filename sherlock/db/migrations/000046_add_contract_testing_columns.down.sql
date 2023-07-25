alter table v2_environments
drop column if exists participates_in_pact;

alter table v2_charts
drop column if exists pact_participant;


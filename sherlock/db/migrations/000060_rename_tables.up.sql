alter table v2_app_versions
    rename to app_versions;

alter table v2_changeset_new_app_versions
    rename to changeset_new_app_versions;

alter table v2_changeset_new_chart_versions
    rename to changeset_new_chart_versions;

alter table v2_changesets
    rename to changesets;

alter table v2_chart_releases
    rename to chart_releases;

alter table v2_chart_versions
    rename to chart_versions;

alter table v2_charts
    rename to charts;

alter table v2_ci_identifiers
    rename to ci_identifiers;

alter table v2_ci_runs
    rename to ci_runs;

alter table v2_ci_runs_for_identifiers
    rename to ci_runs_for_identifiers;

alter table v2_clusters
    rename to clusters;

alter table v2_database_instances
    rename to database_instances;

alter table v2_environments
    rename to environments;

alter table v2_pagerduty_integrations
    rename to pagerduty_integrations;

alter table v2_users
    rename to users;
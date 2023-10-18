alter table app_versions
    rename to v2_app_versions;

alter table changeset_new_app_versions
    rename to v2_changeset_new_app_versions;

alter table changeset_new_chart_versions
    rename to v2_changeset_new_chart_versions;

alter table changesets
    rename to v2_changesets;

alter table chart_releases
    rename to v2_chart_releases;

alter table chart_versions
    rename to v2_chart_versions;

alter table charts
    rename to v2_charts;

alter table ci_identifiers
    rename to v2_ci_identifiers;

alter table ci_runs
    rename to v2_ci_runs;

alter table ci_runs_for_identifiers
    rename to v2_ci_runs_for_identifiers;

alter table clusters
    rename to v2_clusters;

alter table database_instances
    rename to v2_database_instances;

alter table environments
    rename to v2_environments;

alter table pagerduty_integrations
    rename to v2_pagerduty_integrations;

alter table users
    rename to v2_users;
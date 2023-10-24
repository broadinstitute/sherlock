/*
 * app_version
 */
alter table v2_app_versions
    rename to app_versions;

-- keys
alter table app_versions
    rename constraint v2_app_versions_pkey to app_versions_pkey;

-- foreign keys
alter table app_versions
    rename constraint fk_v2_app_versions_chart to fk_app_versions_chart;

alter table app_versions
    rename constraint fk_v2_app_versions_parent_app_version to fk_app_versions_parent_app_version;

-- indexes
alter index idx_v2_app_versions_deleted_at rename to idx_app_versions_deleted_at;

/*
 * changeset_new_app_versions
 */
alter table v2_changeset_new_app_versions
    rename to changeset_new_app_versions;

-- keys
alter table changeset_new_app_versions
    rename constraint v2_changeset_new_app_versions_pkey to changeset_new_app_versions_pkey;

-- foreign keys
alter table changeset_new_app_versions
    rename constraint fk_v2_changeset_new_app_versions_app_version to fk_changeset_new_app_versions_app_version;

alter table changeset_new_app_versions
    rename constraint fk_v2_changeset_new_app_versions_changeset to fk_changeset_new_app_versions_changeset;

/*
 * changeset_new_chart_versions
 */
alter table v2_changeset_new_chart_versions
    rename to changeset_new_chart_versions;

-- keys
alter table changeset_new_chart_versions
    rename constraint v2_changeset_new_chart_versions_pkey to changeset_new_chart_versions_pkey;

-- foreign keys
alter table changeset_new_chart_versions
    rename constraint fk_v2_changeset_new_chart_versions_changeset to fk_changeset_new_chart_versions_changeset;

alter table changeset_new_chart_versions
    rename constraint fk_v2_changeset_new_chart_versions_chart_version to fk_changeset_new_chart_versions_chart_version;

/*
 * changesets
 */
alter table v2_changesets
    rename to changesets;

-- keys
alter table changesets
    rename constraint v2_changesets_pkey to changesets_pkey;

-- foreign keys
alter table changesets
    rename constraint fk_v2_changesets_chart_release to fk_changesets_chart_release;

alter table changesets
    rename constraint fk_v2_changesets_from_app_version to fk_changesets_from_app_version;

alter table changesets
    rename constraint fk_v2_changesets_from_chart_version to fk_changesets_from_chart_version;

alter table changesets
    rename constraint fk_v2_changesets_from_follow_chart_release_app_version to fk_changesets_from_follow_chart_release_app_version;

alter table changesets
    rename constraint fk_v2_changesets_from_follow_chart_release_chart_version to fk_changesets_from_follow_chart_release_chart_version;

alter table changesets
    rename constraint fk_v2_changesets_to_app_version to fk_changesets_to_app_version;

alter table changesets
    rename constraint fk_v2_changesets_to_chart_version to fk_changesets_to_chart_version;

alter table changesets
    rename constraint fk_v2_changesets_to_follow_chart_release_app_version to fk_changesets_to_follow_chart_release_app_version;

alter table changesets
    rename constraint fk_v2_changesets_to_follow_chart_release_chart_version to fk_changesets_to_follow_chart_release_chart_version;

-- indexes
alter index idx_v2_changesets_deleted_at rename to idx_changesets_deleted_at;

/*
 * chart_releases
 */
alter table v2_chart_releases
    rename to chart_releases;

-- key
alter table chart_releases
    rename constraint v2_chart_releases_pkey to chart_releases_pkey;

-- foreign keys
alter table chart_releases
    rename constraint fk_v2_chart_releases_app_version to fk_chart_releases_app_version;

alter table chart_releases
    rename constraint fk_v2_chart_releases_chart to fk_chart_releases_chart;

alter table chart_releases
    rename constraint fk_v2_chart_releases_chart_version to fk_chart_releases_chart_version;

alter table chart_releases
    rename constraint fk_v2_chart_releases_cluster to fk_chart_releases_cluster;

alter table chart_releases
    rename constraint fk_v2_chart_releases_environment to fk_chart_releases_environment;

alter table chart_releases
    rename constraint fk_v2_chart_releases_follow_chart_release_app_version to fk_chart_releases_follow_chart_release_app_version;

alter table chart_releases
    rename constraint fk_v2_chart_releases_follow_chart_release_chart_version to fk_chart_releases_follow_chart_release_chart_version;

alter table chart_releases
    rename constraint fk_v2_chart_releases_pagerduty_integration to fk_chart_releases_pagerduty_integration;

-- indexes
alter index idx_v2_chart_releases_deleted_at rename to idx_chart_releases_deleted_at;

/*
 * chart_versions
 */
alter table v2_chart_versions
    rename to chart_versions;

-- keys
alter table chart_versions
    rename constraint v2_chart_versions_pkey to chart_versions_pkey;

-- foreign keys
alter table chart_versions
    rename constraint fk_v2_chart_versions_chart to fk_chart_versions_chart;

alter table chart_versions
    rename constraint fk_v2_chart_versions_parent_chart_version to fk_chart_versions_parent_chart_version;

-- indexes
alter index idx_v2_chart_versions_deleted_at rename to idx_chart_versions_deleted_at;

/*
 * charts
 */
alter table v2_charts
    rename to charts;

-- keys
alter table charts
    rename constraint v2_charts_pkey to charts_pkey;

alter table charts
    rename constraint v2_charts_name_key to charts_name_key;

-- indexes
alter index idx_v2_charts_deleted_at rename to idx_charts_deleted_at;

/*
 * ci_identifiers
 */
alter table v2_ci_identifiers
    rename to ci_identifiers;

-- keys
alter table ci_identifiers
    rename constraint v2_ci_identifiers_pkey to ci_identifiers_pkey;

-- indexes
alter index idx_v2_ci_identifiers_deleted_at rename to idx_ci_identifiers_deleted_at;

alter index idx_v2_ci_identifiers_polymorphic_index rename to idx_ci_identifiers_polymorphic_index;

/*
 * ci_runs
 */
alter table v2_ci_runs
    rename to ci_runs;

-- keys
alter table ci_runs
    rename constraint v2_ci_runs_pkey to ci_runs_pkey;

-- indexes
alter index idx_v2_ci_runs_deleted_at rename to idx_ci_runs_deleted_at;

/*
 * ci_runs_for_identifiers
 */
alter table v2_ci_runs_for_identifiers
    rename to ci_runs_for_identifiers;

-- keys
alter table ci_runs_for_identifiers
    rename constraint v2_ci_runs_for_identifiers_pkey to ci_runs_for_identifiers_pkey;

-- foreign keys
alter table ci_runs_for_identifiers
    rename constraint fk_v2_ci_runs_for_identifiers_ci_identifier to fk_ci_runs_for_identifiers_ci_identifier;

alter table ci_runs_for_identifiers
    rename constraint fk_v2_ci_runs_for_identifiers_ci_run to fk_ci_runs_for_identifiers_ci_run;

/*
 * clusters
 */
alter table v2_clusters
    rename to clusters;

-- keys
alter table clusters
    rename constraint v2_clusters_pkey to clusters_pkey;

alter table clusters
    rename constraint v2_clusters_name_key to clusters_name_key;

-- indexes
alter index idx_v2_clusters_deleted_at rename to idx_clusters_deleted_at;

/*
 * database_instances
 */
alter table v2_database_instances
    rename to database_instances;

-- keys
alter table database_instances
    rename constraint v2_database_instances_pkey to database_instances_pkey;

-- foreign keys
alter table database_instances
    rename constraint fk_v2_database_instances_chart_release to fk_database_instances_chart_release;

-- indexes
alter index idx_v2_database_instances_deleted_at rename to idx_database_instances_deleted_at;

/*
 * deploy_hook_trigger_configs
 */
alter table v2_deploy_hook_trigger_configs
    rename to deploy_hook_trigger_configs;

-- keys
alter table deploy_hook_trigger_configs
    rename constraint v2_deploy_hook_trigger_configs_pkey to deploy_hook_trigger_configs_pkey;

-- foreign keys
alter table deploy_hook_trigger_configs
    rename constraint fk_v2_deploy_hook_trigger_configs_on_chart_release to fk_deploy_hook_trigger_configs_on_chart_release;

alter table deploy_hook_trigger_configs
    rename constraint fk_v2_deploy_hook_trigger_configs_on_environment to fk_deploy_hook_trigger_configs_on_environment;

--indexes
alter index idx_v2_deploy_hook_trigger_configs_deleted_at rename to idx_deploy_hook_trigger_configs_deleted_at;

/*
 * environments
 */
alter table v2_environments
    rename to environments;

-- keys
alter table environments
    rename constraint v2_environments_pkey to environments_pkey;

-- foreign keys
alter table environments
    rename constraint fk_v2_environments_default_cluster to fk_environments_default_cluster;

alter table environments
    rename constraint fk_v2_environments_owner to fk_environments_owner;

alter table environments
    rename constraint fk_v2_environments_pagerduty_integration to fk_environments_pagerduty_integration;

alter table environments
    rename constraint fk_v2_environments_template_environment to fk_environments_template_environment;

-- indexes
alter index idx_v2_environments_deleted_at rename to idx_environments_deleted_at;

/*
 * github_actions_deploy_hooks
 */
alter table v2_github_actions_deploy_hooks
    rename to github_actions_deploy_hooks;

-- keys
alter table github_actions_deploy_hooks
    rename constraint v2_github_actions_deploy_hooks_pkey to github_actions_deploy_hooks_pkey;

-- indexes
alter index idx_v2_github_actions_deploy_hooks_deleted_at rename to idx_github_actions_deploy_hooks_deleted_at;

/*
 * pagerduty_integration
 */
alter table v2_pagerduty_integrations
    rename to pagerduty_integrations;

-- keys
alter table pagerduty_integrations
    rename constraint v2_pagerduty_integrations_pkey to pagerduty_integrations_pkey;

-- indexes
alter index idx_v2_pagerduty_integrations_deleted_at rename to idx_pagerduty_integrations_deleted_at;

/*
 * slack_deploy_hooks
 */
alter table v2_slack_deploy_hooks
    rename to slack_deploy_hooks;

-- keys
alter table slack_deploy_hooks
    rename constraint v2_slack_deploy_hooks_pkey to slack_deploy_hooks_pkey;

-- indexes
alter index idx_v2_slack_deploy_hooks_deleted_at rename to idx_slack_deploy_hooks_deleted_at;

/*
 * users
 */
alter table v2_users
    rename to users;

-- keys
alter table users
    rename constraint v2_users_pkey to users_pkey;

alter table users
    rename constraint v2_users_email_key to users_email_key;

alter table users
    rename constraint v2_users_google_id_key to users_google_id_key;

-- indexes
alter index idx_v2_users_deleted_at rename to idx_users_deleted_at;

/*
 * sequences
 */
alter sequence v2_app_versions_id_seq rename to app_versions_id_seq;

alter sequence v2_changesets_id_seq rename to changesets_id_seq;

alter sequence v2_chart_releases_id_seq rename to chart_releases_id_seq;

alter sequence v2_chart_versions_id_seq rename to chart_versions_id_seq;

alter sequence v2_charts_id_seq rename to charts_id_seq;

alter sequence v2_ci_identifiers_id_seq rename to ci_identifiers_id_seq;

alter sequence v2_ci_runs_id_seq rename to ci_runs_id_seq;

alter sequence v2_clusters_id_seq rename to clusters_id_seq;

alter sequence v2_database_instances_id_seq rename to database_instances_id_seq;

alter sequence v2_environments_id_seq rename to environments_id_seq;

alter sequence v2_pagerduty_integrations_id_seq rename to pagerduty_integrations_id_seq;

alter sequence v2_users_id_seq rename to users_id_seq;

/*
 * app_version
 */
alter table app_versions
    rename to v2_app_versions;

-- keys
alter table v2_app_versions
    rename constraint app_versions_pkey to v2_app_versions_pkey;

-- foreign keys
alter table v2_app_versions
    rename constraint fk_app_versions_chart to fk_v2_app_versions_chart;

alter table v2_app_versions
    rename constraint fk_app_versions_parent_app_version to fk_v2_app_versions_parent_app_version;

-- indexes
alter index idx_app_versions_deleted_at rename to idx_v2_app_versions_deleted_at;

/*
 * changeset_new_app_versions
 */
alter table changeset_new_app_versions
    rename to v2_changeset_new_app_versions;

-- keys
alter table v2_changeset_new_app_versions
    rename constraint changeset_new_app_versions_pkey to v2_changeset_new_app_versions_pkey;

-- foreign keys
alter table v2_changeset_new_app_versions
    rename constraint fk_changeset_new_app_versions_app_version to fk_v2_changeset_new_app_versions_app_version;

alter table v2_changeset_new_app_versions
    rename constraint fk_changeset_new_app_versions_changeset to fk_v2_changeset_new_app_versions_changeset;

/*
 * changeset_new_chart_versions
 */
alter table changeset_new_chart_versions
    rename to v2_changeset_new_chart_versions;

-- keys
alter table v2_changeset_new_chart_versions
    rename constraint changeset_new_chart_versions_pkey to v2_changeset_new_chart_versions_pkey;

-- foreign keys
alter table v2_changeset_new_chart_versions
    rename constraint fk_changeset_new_chart_versions_changeset to fk_v2_changeset_new_chart_versions_changeset;

alter table v2_changeset_new_chart_versions
    rename constraint fk_changeset_new_chart_versions_chart_version to fk_v2_changeset_new_chart_versions_chart_version;

/*
 * changesets
 */
alter table changesets
    rename to v2_changesets;

-- keys
alter table v2_changesets
    rename constraint changesets_pkey to v2_changesets_pkey;

-- foreign keys
alter table v2_changesets
    rename constraint fk_changesets_chart_release to fk_v2_changesets_chart_release;

alter table v2_changesets
    rename constraint fk_changesets_from_app_version to fk_v2_changesets_from_app_version;

alter table v2_changesets
    rename constraint fk_changesets_from_chart_version to fk_v2_changesets_from_chart_version;

alter table v2_changesets
    rename constraint fk_changesets_from_follow_chart_release_app_version to fk_v2_changesets_from_follow_chart_release_app_version;

alter table v2_changesets
    rename constraint fk_changesets_from_follow_chart_release_chart_version to fk_v2_changesets_from_follow_chart_release_chart_version;

alter table v2_changesets
    rename constraint fk_changesets_to_app_version to fk_v2_changesets_to_app_version;

alter table v2_changesets
    rename constraint fk_changesets_to_chart_version to fk_v2_changesets_to_chart_version;

alter table v2_changesets
    rename constraint fk_changesets_to_follow_chart_release_app_version to fk_v2_changesets_to_follow_chart_release_app_version;

alter table v2_changesets
    rename constraint fk_changesets_to_follow_chart_release_chart_version to fk_v2_changesets_to_follow_chart_release_chart_version;

-- indexes
alter index idx_changesets_deleted_at rename to idx_v2_changesets_deleted_at;

/*
 * chart_releases
 */
alter table chart_releases
    rename to v2_chart_releases;

-- key
alter table v2_chart_releases
    rename constraint chart_releases_pkey to v2_chart_releases_pkey;

-- foreign keys
alter table v2_chart_releases
    rename constraint fk_chart_releases_app_version to fk_v2_chart_releases_app_version;

alter table v2_chart_releases
    rename constraint fk_chart_releases_chart to fk_v2_chart_releases_chart;

alter table v2_chart_releases
    rename constraint fk_chart_releases_chart_version to fk_v2_chart_releases_chart_version;

alter table v2_chart_releases
    rename constraint fk_chart_releases_cluster to fk_v2_chart_releases_cluster;

alter table v2_chart_releases
    rename constraint fk_chart_releases_environment to fk_v2_chart_releases_environment;

alter table v2_chart_releases
    rename constraint fk_chart_releases_follow_chart_release_app_version to fk_v2_chart_releases_follow_chart_release_app_version;

alter table v2_chart_releases
    rename constraint fk_chart_releases_follow_chart_release_chart_version to fk_v2_chart_releases_follow_chart_release_chart_version;

alter table v2_chart_releases
    rename constraint fk_chart_releases_pagerduty_integration to fk_v2_chart_releases_pagerduty_integration;

-- indexes
alter index idx_chart_releases_deleted_at rename to idx_v2_chart_releases_deleted_at;

/*
 * chart_versions
 */
alter table chart_versions
    rename to v2_chart_versions;

-- keys
alter table v2_chart_versions
    rename constraint chart_versions_pkey to v2_chart_versions_pkey;

-- foreign keys
alter table v2_chart_versions
    rename constraint fk_chart_versions_chart to fk_v2_chart_versions_chart;

alter table v2_chart_versions
    rename constraint fk_chart_versions_parent_chart_version to fk_v2_chart_versions_parent_chart_version;

-- indexes
alter index idx_chart_versions_deleted_at rename to idx_v2_chart_versions_deleted_at;

/*
 * charts
 */
alter table charts
    rename to v2_charts;

-- keys
alter table v2_charts
    rename constraint charts_pkey to v2_charts_pkey;

alter table v2_charts
    rename constraint charts_name_key to v2_charts_name_key;

-- indexes
alter index idx_charts_deleted_at rename to idx_v2_charts_deleted_at;

/*
 * ci_identifiers
 */
alter table ci_identifiers
    rename to v2_ci_identifiers;

-- keys
alter table v2_ci_identifiers
    rename constraint ci_identifiers_pkey to v2_ci_identifiers_pkey;

-- indexes
alter index idx_ci_identifiers_deleted_at rename to idx_v2_ci_identifiers_deleted_at;

alter index idx_ci_identifiers_polymorphic_index rename to idx_v2_ci_identifiers_polymorphic_index;

/*
 * ci_runs
 */
alter table ci_runs
    rename to v2_ci_runs;

-- keys
alter table v2_ci_runs
    rename constraint ci_runs_pkey to v2_ci_runs_pkey;

-- indexes
alter index idx_ci_runs_deleted_at rename to idx_v2_ci_runs_deleted_at;

/*
 * ci_runs_for_identifiers
 */
alter table ci_runs_for_identifiers
    rename to v2_ci_runs_for_identifiers;

-- keys
alter table v2_ci_runs_for_identifiers
    rename constraint ci_runs_for_identifiers_pkey to v2_ci_runs_for_identifiers_pkey;

-- foreign keys
alter table v2_ci_runs_for_identifiers
    rename constraint fk_ci_runs_for_identifiers_ci_identifier to fk_v2_ci_runs_for_identifiers_ci_identifier;

alter table v2_ci_runs_for_identifiers
    rename constraint fk_ci_runs_for_identifiers_ci_run to fk_v2_ci_runs_for_identifiers_ci_run;

/*
 * clusters
 */
alter table clusters
    rename to v2_clusters;

-- keys
alter table v2_clusters
    rename constraint clusters_pkey to v2_clusters_pkey;

alter table v2_clusters
    rename constraint clusters_name_key to v2_clusters_name_key;

-- indexes
alter index idx_clusters_deleted_at rename to idx_v2_clusters_deleted_at;

/*
 * database_instances
 */
alter table database_instances
    rename to v2_database_instances;

-- keys
alter table v2_database_instances
    rename constraint database_instances_pkey to v2_database_instances_pkey;

-- foreign keys
alter table v2_database_instances
    rename constraint fk_database_instances_chart_release to fk_v2_database_instances_chart_release;

-- indexes
alter index idx_database_instances_deleted_at rename to idx_v2_database_instances_deleted_at;

/*
 * deploy_hook_trigger_configs
 */
alter table deploy_hook_trigger_configs
    rename to v2_deploy_hook_trigger_configs;

-- keys
alter table v2_deploy_hook_trigger_configs
    rename constraint deploy_hook_trigger_configs_pkey to v2_deploy_hook_trigger_configs_pkey;

-- foreign keys
alter table v2_deploy_hook_trigger_configs
    rename constraint fk_deploy_hook_trigger_configs_on_chart_release to fk_v2_deploy_hook_trigger_configs_on_chart_release;

alter table v2_deploy_hook_trigger_configs
    rename constraint fk_deploy_hook_trigger_configs_on_environment to fk_v2_deploy_hook_trigger_configs_on_environment;

--indexes
alter index idx_deploy_hook_trigger_configs_deleted_at rename to idx_v2_deploy_hook_trigger_configs_deleted_at;

/*
 * environments
 */
alter table environments
    rename to v2_environments;

-- keys
alter table v2_environments
    rename constraint environments_pkey to v2_environments_pkey;

-- foreign keys
alter table v2_environments
    rename constraint fk_environments_default_cluster to fk_v2_environments_default_cluster;

alter table v2_environments
    rename constraint fk_environments_owner to fk_v2_environments_owner;

alter table v2_environments
    rename constraint fk_environments_pagerduty_integration to fk_v2_environments_pagerduty_integration;

alter table v2_environments
    rename constraint fk_environments_template_environment to fk_v2_environments_template_environment;

-- indexes
alter index idx_environments_deleted_at rename to idx_v2_environments_deleted_at;

/*
 * github_actions_deploy_hooks
 */
alter table github_actions_deploy_hooks
    rename to v2_github_actions_deploy_hooks;

-- keys
alter table v2_github_actions_deploy_hooks
    rename constraint github_actions_deploy_hooks_pkey to v2_github_actions_deploy_hooks_pkey;

-- indexes
alter index idx_github_actions_deploy_hooks_deleted_at rename to idx_v2_github_actions_deploy_hooks_deleted_at;

/*
 * pagerduty_integration
 */
alter table pagerduty_integrations
    rename to v2_pagerduty_integrations;

-- keys
alter table v2_pagerduty_integrations
    rename constraint pagerduty_integrations_pkey to v2_pagerduty_integrations_pkey;

-- indexes
alter index idx_pagerduty_integrations_deleted_at rename to idx_v2_pagerduty_integrations_deleted_at;

/*
 * slack_deploy_hooks
 */
alter table slack_deploy_hooks
    rename to v2_slack_deploy_hooks;

-- keys
alter table v2_slack_deploy_hooks
    rename constraint slack_deploy_hooks_pkey to v2_slack_deploy_hooks_pkey;

-- indexes
alter index idx_slack_deploy_hooks_deleted_at rename to idx_v2_slack_deploy_hooks_deleted_at;

/*
 * users
 */
alter table users
    rename to v2_users;

-- keys
alter table v2_users
    rename constraint users_pkey to v2_users_pkey;

alter table v2_users
    rename constraint users_email_key to v2_users_email_key;

alter table v2_users
    rename constraint users_google_id_key to v2_users_google_id_key;

-- indexes
alter index idx_users_deleted_at rename to idx_v2_users_deleted_at;

/*
 * sequences
 */
alter sequence app_versions_id_seq rename to v2_app_versions_id_seq;

alter sequence changesets_id_seq rename to v2_changesets_id_seq;

alter sequence chart_releases_id_seq rename to v2_chart_releases_id_seq;

alter sequence chart_versions_id_seq rename to v2_chart_versions_id_seq;

alter sequence charts_id_seq rename to v2_charts_id_seq;

alter sequence ci_identifiers_id_seq rename to v2_ci_identifiers_id_seq;

alter sequence ci_runs_id_seq rename to v2_ci_runs_id_seq;

alter sequence clusters_id_seq rename to v2_clusters_id_seq;

alter sequence database_instances_id_seq rename to v2_database_instances_id_seq;

alter sequence environments_id_seq rename to v2_environments_id_seq;

alter sequence pagerduty_integrations_id_seq rename to v2_pagerduty_integrations_id_seq;

alter sequence users_id_seq rename to v2_users_id_seq;

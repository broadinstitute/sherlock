ALTER TABLE v2_ci_identifiers
    ADD CONSTRAINT resource_present
        CHECK (resource_type <> '' AND resource_id <> 0);

CREATE UNIQUE INDEX ci_identifiers_selector_unique_constraint
    ON v2_ci_identifiers (resource_type, resource_id)
    WHERE deleted_at IS NULL;

ALTER TABLE v2_ci_runs
    ADD CONSTRAINT platform_present
        CHECK ((
                           platform = 'github-actions'
                       AND github_actions_owner <> ''
                       AND github_actions_repo <> ''
                       AND github_actions_run_id <> 0
                       AND github_actions_attempt_number <> 0
                       AND github_actions_workflow_path <> ''
                       AND (argo_workflows_namespace IS NULL OR argo_workflows_namespace = '')
                       AND (argo_workflows_name IS NULL OR argo_workflows_name = '')
                       AND (argo_workflows_template IS NULL OR argo_workflows_template = '')
                   ) OR (
                           platform = 'argo-workflows'
                       AND argo_workflows_namespace <> ''
                       AND argo_workflows_name <> ''
                       AND argo_workflows_template <> ''
                       AND (github_actions_owner IS NULL OR github_actions_owner = '')
                       AND (github_actions_repo IS NULL OR github_actions_repo = '')
                       AND (github_actions_run_id IS NULL OR github_actions_run_id = 0)
                       AND (github_actions_attempt_number IS NULL OR github_actions_attempt_number = 0)
                       AND (github_actions_workflow_path IS NULL OR github_actions_workflow_path = '')
                   ));

ALTER TABLE v2_ci_runs
    ADD CONSTRAINT terminal_status_present
        CHECK (terminal_at IS NULL or status <> '');

CREATE UNIQUE INDEX ci_runs_selector_unique_constraint
    ON v2_ci_runs (
                   platform,
                   github_actions_owner,
                   github_actions_repo,
                   github_actions_run_id,
                   github_actions_attempt_number,
                   argo_workflows_namespace,
                   argo_workflows_name
        )
    WHERE deleted_at IS NULL;

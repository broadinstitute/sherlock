create table v2_deploy_hook_trigger_configs
(
    id                  bigserial
        primary key,
    created_at          timestamp with time zone,
    updated_at          timestamp with time zone,
    deleted_at          timestamp with time zone,

    hook_id             bigint
        constraint hook_id_present
            check (hook_id is not null and hook_id != 0),
    hook_type           text
        constraint hook_type_present
            check (hook_type is not null and hook_type != ''),

    on_environment_id   bigint
        constraint fk_v2_deploy_hook_trigger_configs_on_environment
            references v2_environments,
    on_chart_release_id bigint
        constraint fk_v2_deploy_hook_trigger_configs_on_chart_release
            references v2_chart_releases,
    constraint environment_or_chart_release_present
        check ((on_environment_id is not null and on_environment_id != 0 and
                (on_chart_release_id is null or on_chart_release_id = 0)) or
               (on_chart_release_id is not null and on_chart_release_id != 0 and
                (on_environment_id is null or on_environment_id = 0))),

    on_failure          boolean,
    on_success          boolean
);

create index idx_v2_deploy_hook_trigger_configs_deleted_at
    on v2_deploy_hook_trigger_configs (deleted_at);

create table v2_slack_deploy_hooks
(
    id            bigserial
        primary key,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone,

    slack_channel text
        constraint slack_channel_present
            check (slack_channel is not null and slack_channel != '')
);

create index idx_v2_slack_deploy_hooks_deleted_at
    on v2_slack_deploy_hooks (deleted_at);

create table v2_github_actions_deploy_hooks
(
    id                             bigserial
        primary key,
    created_at                     timestamp with time zone,
    updated_at                     timestamp with time zone,
    deleted_at                     timestamp with time zone,

    github_actions_owner           text
        constraint github_actions_owner_present
            check (github_actions_owner is not null and github_actions_owner != ''),
    github_actions_repo            text
        constraint github_actions_repo_present
            check (github_actions_repo is not null and github_actions_repo != ''),
    github_actions_workflow_path   text
        constraint github_actions_workflow_path_present
            check (github_actions_workflow_path is not null and github_actions_workflow_path != ''),
    github_actions_default_ref     text
        constraint github_actions_default_ref_present
            check (github_actions_default_ref is not null and github_actions_default_ref != ''),
    github_actions_ref_behavior    text
        constraint github_actions_ref_behavior_valid
            check (github_actions_ref_behavior is not null and (
                        github_actions_ref_behavior = 'always-use-default-ref' or
                        github_actions_ref_behavior = 'use-app-version-as-ref' or
                        github_actions_ref_behavior = 'use-app-version-commit-as-ref'
                )),
    github_actions_workflow_inputs jsonb
        constraint github_actions_workflow_inputs_small
            check (octet_length(github_actions_workflow_inputs::text) < 102400) -- 100 KB as arbitrary DOS protection :shrug:
);

create index idx_v2_github_actions_deploy_hooks_deleted_at
    on v2_github_actions_deploy_hooks (deleted_at);

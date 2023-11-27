create table if not exists slack_deploy_hook_states
(
    ci_run_id bigint not null
        constraint fk_slack_deploy_hook_states_ci_runs
            references ci_runs
            on update cascade on delete cascade,
    slack_deploy_hook_id bigint not null
        constraint fk_slack_deploy_hook_states_slack_deploy_hooks
            references slack_deploy_hooks
            on update cascade on delete cascade,
    primary key (ci_run_id, slack_deploy_hook_id),

    message_timestamp text not null,
    message_channel text not null
);

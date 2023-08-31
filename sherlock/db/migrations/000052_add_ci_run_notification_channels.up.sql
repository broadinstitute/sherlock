alter table v2_ci_runs
    add column notify_slack_channels_upon_success text[];

alter table v2_ci_runs
    add column notify_slack_channels_upon_failure text[];

alter table v2_ci_runs
    add column notify_slack_channels_on_success jsonb;

alter table v2_ci_runs
    add column notify_slack_channels_on_failure jsonb;

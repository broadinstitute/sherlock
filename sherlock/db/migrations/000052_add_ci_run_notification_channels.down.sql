alter table v2_ci_runs
    drop column notify_slack_channels_on_success;

alter table v2_ci_runs
    drop column notify_slack_channels_on_failure;

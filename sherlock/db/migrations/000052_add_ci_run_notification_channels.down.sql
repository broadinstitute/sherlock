alter table v2_ci_runs
    drop column notify_slack_channels_upon_success;

alter table v2_ci_runs
    drop column notify_slack_channels_upon_failure;

alter table ci_runs
    add column notify_slack_channels_upon_retry jsonb;

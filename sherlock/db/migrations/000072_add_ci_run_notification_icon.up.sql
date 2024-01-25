alter table ci_runs
    add column if not exists notify_slack_custom_icon text;

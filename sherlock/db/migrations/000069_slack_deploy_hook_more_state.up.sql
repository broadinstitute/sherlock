alter table slack_deploy_hook_states
    add column if not exists failure_alert_sent bool;
alter table slack_deploy_hook_states
    add column if not exists changelog_sent bool;

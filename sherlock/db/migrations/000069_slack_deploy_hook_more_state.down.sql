alter table slack_deploy_hook_states
    drop column if exists failure_alert_sent;
alter table slack_deploy_hook_states
    drop column if exists changelog_sent;

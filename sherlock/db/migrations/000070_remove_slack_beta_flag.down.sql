alter table slack_deploy_hooks
    add column if not exists beta bool;

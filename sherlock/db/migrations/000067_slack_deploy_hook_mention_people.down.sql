alter table slack_deploy_hooks
    drop column if exists mention_people;
alter table slack_deploy_hooks
    drop column if exists beta;
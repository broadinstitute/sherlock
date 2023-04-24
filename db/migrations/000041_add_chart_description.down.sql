alter table v2_charts
    drop column if exists description;

alter table v2_charts
    drop column if exists playbook_url;

alter table v2_charts
    add if not exists description text;

alter table v2_charts
    add if not exists playbook_url text;


alter table charts
    drop column if exists legacy_configs_enabled;

alter table chart_releases
    drop column if exists firecloud_develop_ref;

alter table changesets
    drop column if exists from_firecloud_develop_ref;

alter table changesets
    drop column if exists to_firecloud_develop_ref;

alter table environments
    drop column if exists default_firecloud_develop_ref;

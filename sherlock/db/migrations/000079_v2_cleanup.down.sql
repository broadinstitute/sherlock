alter table charts
    add column if not exists legacy_configs_enabled bool;

alter table chart_releases
    add column if not exists firecloud_develop_ref text;

alter table changesets
    add column if not exists from_firecloud_develop_ref text;

alter table changesets
    add column if not exists to_firecloud_develop_ref text;

alter table environments
    add column if not exists default_firecloud_develop_ref text;

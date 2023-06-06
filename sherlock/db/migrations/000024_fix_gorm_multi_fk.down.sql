alter table v2_changesets
    rename constraint fk_v2_changesets_to_app_version to fk_v2_changesets_app_version;

alter table v2_changesets
    rename constraint fk_v2_changesets_to_chart_version to fk_v2_changesets_chart_version;

alter table v2_changesets
    drop constraint if exists fk_v2_changesets_from_app_version;

alter table v2_changesets
    drop constraint if exists fk_v2_changesets_from_chart_version;

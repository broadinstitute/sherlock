alter table v2_chart_releases
    drop constraint if exists fk_v2_chart_releases_follow_chart_release_app_version;

alter table v2_chart_releases
    drop column if exists app_version_follow_chart_release_id;



alter table v2_chart_releases
    drop constraint if exists fk_v2_chart_releases_follow_chart_release_chart_version;

alter table v2_chart_releases
    drop column if exists chart_version_follow_chart_release_id;



alter table v2_changesets
    drop constraint if exists fk_v2_changesets_to_follow_chart_release_app_version;

alter table v2_changesets
    drop column if exists to_app_version_follow_chart_release_id;

alter table v2_changesets
    drop constraint if exists fk_v2_changesets_from_follow_chart_release_app_version;

alter table v2_changesets
    drop column if exists from_app_version_follow_chart_release_id;



alter table v2_changesets
    drop constraint if exists fk_v2_changesets_to_follow_chart_release_chart_version;

alter table v2_changesets
    drop column if exists to_chart_version_follow_chart_release_id;

alter table v2_changesets
    drop constraint if exists fk_v2_changesets_from_follow_chart_release_chart_version;

alter table v2_changesets
    drop column if exists from_chart_version_follow_chart_release_id;

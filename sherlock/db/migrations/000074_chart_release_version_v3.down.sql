alter table chart_releases
    drop constraint if exists resolved_at_present;

alter table chart_releases
    drop constraint if exists app_version_resolver_valid;

alter table chart_releases
    drop constraint if exists chart_version_resolver_valid;

alter table chart_releases
    drop constraint if exists helmfile_ref_valid;

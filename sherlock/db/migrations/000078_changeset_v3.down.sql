alter table changesets
    drop constraint if exists to_resolved_at_present;

alter table changesets
    drop constraint if exists to_app_version_resolver_valid;

alter table changesets
    drop constraint if exists to_chart_version_resolver_valid;

alter table changesets
    drop constraint if exists to_helmfile_ref_valid;

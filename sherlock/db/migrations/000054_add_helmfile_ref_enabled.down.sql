alter table v2_chart_releases
    drop column if exists helmfile_ref_enabled;


alter table v2_changesets
    drop column if exists from_helmfile_ref_enabled,
    drop column if exists to_helmfile_ref_enabled;

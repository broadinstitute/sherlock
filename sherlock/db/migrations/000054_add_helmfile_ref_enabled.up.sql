alter table v2_chart_releases
    add if not exists helmfile_ref_enabled boolean not null default false;

alter table v2_changesets
    add column if not exists from_helmfile_ref_enabled boolean,
    add column if not exists to_helmfile_ref_enabled boolean;

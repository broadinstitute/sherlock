alter table v2_chart_releases
    add if not exists helmfile_ref_enabled boolean default false;

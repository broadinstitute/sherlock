alter table v2_ci_runs
    add if not exists helmfile_ref_enabled boolean default false;

alter table v2_ci_runs
    add if not exists started_at timestamp with time zone;

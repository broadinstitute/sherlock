alter table ci_runs_for_identifiers
    add column if not exists resource_status text;

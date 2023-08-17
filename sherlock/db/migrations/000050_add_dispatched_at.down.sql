alter table v2_ci_runs
    drop column if exists deploy_hooks_dispatched_at;

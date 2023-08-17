alter table v2_ci_runs
    add if not exists deploy_hooks_dispatched_at text;

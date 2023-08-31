alter table v2_ci_runs
    rename column termination_hooks_dispatched_at to deploy_hooks_dispatched_at;

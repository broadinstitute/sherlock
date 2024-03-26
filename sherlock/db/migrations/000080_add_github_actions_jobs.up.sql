create table if not exists github_actions_jobs
(
    id                            bigserial
        primary key,
    created_at                    timestamp with time zone,
    updated_at                    timestamp with time zone,
    deleted_at                    timestamp with time zone,
    github_actions_owner          text   not null,
    github_actions_repo           text   not null,
    github_actions_run_id         bigint,
    github_actions_attempt_number bigint,
    github_actions_job_id         bigint not null,

    job_created_at                timestamp with time zone,
    job_started_at                timestamp with time zone,
    job_terminal_at               timestamp with time zone,
    status                        text
);

create unique index if not exists github_actions_jobs_selector_unique_constraint
    on github_actions_jobs (
                            github_actions_owner,
                            github_actions_repo,
                            github_actions_job_id
        )
    where deleted_at is null;

create index if not exists idx_github_actions_jobs_deleted_at
    on github_actions_jobs (deleted_at);

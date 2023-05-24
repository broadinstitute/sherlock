create table v2_ci_identifiers
(
    id            bigserial
        primary key,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone,
    resource_type text,
    resource_id   bigint
);

create index idx_v2_ci_identifiers_deleted_at
    on v2_ci_identifiers (deleted_at);

create index idx_v2_ci_identifiers_polymorphic_index
    on v2_ci_identifiers (resource_type, resource_id);

create table v2_ci_runs
(
    id                            bigserial
        primary key,
    created_at                    timestamp with time zone,
    updated_at                    timestamp with time zone,
    deleted_at                    timestamp with time zone,
    platform                      text,
    github_actions_owner          text,
    github_actions_repo           text,
    github_actions_run_id         bigint,
    github_actions_attempt_number bigint,
    github_actions_workflow_path  text,
    argo_workflows_namespace      text,
    argo_workflows_name           text,
    argo_workflows_template       text,
    terminal_at                   timestamp with time zone,
    status                        text
);

create index idx_v2_ci_runs_deleted_at
    on v2_ci_runs (deleted_at);

create table v2_ci_runs_for_identifiers
(
    ci_run_id        bigint not null
        constraint fk_v2_ci_runs_for_identifiers_ci_run
            references v2_ci_runs
            on update cascade on delete cascade,
    ci_identifier_id bigint not null
        constraint fk_v2_ci_runs_for_identifiers_ci_identifier
            references v2_ci_identifiers
            on update cascade on delete cascade,
    primary key (ci_run_id, ci_identifier_id)
);

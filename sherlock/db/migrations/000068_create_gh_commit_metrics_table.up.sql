create table if not exists git_commits
(
    id                  bigserial
        primary key,
    created_at          timestamp with time zone,
    updated_at          timestamp with time zone,
    deleted_at          timestamp with time zone,
    git_repo            text not null,
    git_commit          text not null,
    git_branch          text not null,
    is_main_branch      boolean not null,
    sec_since_prev      bigint,
    committed_at         timestamp with time zone not null
);

create index if not exists idx_git_commits_deleted_at
    on git_commits (deleted_at);

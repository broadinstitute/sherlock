alter table git_commits
    add column if not exists committed_at timestamp with time zone not null;

CREATE TABLE IF NOT EXISTS builds(
    id serial PRIMARY KEY,
    commit_sha TEXT,
    build_url TEXT,
    version_string TEXT NOT NULL,
    service_id integer REFERENCES services (id),
    created_at TIMESTAMP WITH TIME ZONE default now(),
    updated_at TIMESTAMP WITH TIME ZONE default now()
);

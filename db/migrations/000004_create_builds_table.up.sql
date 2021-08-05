CREATE TABLE IF NOT EXISTS builds(
    id serial PRIMARY KEY,
    commit_sha TEXT,
    build_url TEXT,
    built_at TIMESTAMP WITH TIME ZONE,
    version_string TEXT UNIQUE NOT NULL,
    service_id integer REFERENCES services (id),
    created_at TIMESTAMP WITH TIME ZONE default now(),
    updated_at TIMESTAMP WITH TIME ZONE default now()
);

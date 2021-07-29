CREATE TABLE IF NOT EXISTS builds(
    id serial PRIMARY KEY,
    commit_sha TEXT,
    build_url TEXT,
    built_at TIMESTAMP,
    version_string TEXT UNIQUE NOT NULL,
    service_id integer REFERENCES services (id)
);

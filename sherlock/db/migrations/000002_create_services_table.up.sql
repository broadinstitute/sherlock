CREATE TABLE IF NOT EXISTS services(
    id serial PRIMARY KEY,
    repo_url TEXT,
    name VARCHAR (45) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE default now(),
    updated_at TIMESTAMP WITH TIME ZONE default now()
);

CREATE INDEX services_name_idx ON services (name);

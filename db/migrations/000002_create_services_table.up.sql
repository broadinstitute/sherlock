CREATE TABLE IF NOT EXISTS services(
    id serial PRIMARY KEY,
    repo_url TEXT,
    name VARCHAR (45) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE default now(),
    deleted_at TIMESTAMP WITH TIME ZONE default now(),
    updated_at TIMESTAMP WITH TIME ZONE default now()
);

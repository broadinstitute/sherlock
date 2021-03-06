CREATE TABLE IF NOT EXISTS clusters(
    id                        serial PRIMARY KEY,
    name                      text UNIQUE NOT NULL,
    google_project            text,
    created_at                timestamp WITH TIME ZONE default now(),
    updated_at                timestamp WITH TIME ZONE default now()
);

CREATE UNIQUE INDEX name_idx ON clusters (name);

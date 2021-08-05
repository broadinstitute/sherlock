CREATE TABLE IF NOT EXISTS environments(
    id serial PRIMARY KEY,
    name VARCHAR (45) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE default now(),
    deleted_at TIMESTAMP WITH TIME ZONE default now(),
    updated_at TIMESTAMP WITH TIME ZONE default now()
);

CREATE TABLE IF NOT EXISTS allocation_pools(
    id                        serial PRIMARY KEY,
    name                      text UNIQUE NOT NULL,
    created_at                timestamp WITH TIME ZONE default now(),
    updated_at                timestamp WITH TIME ZONE default now()
);


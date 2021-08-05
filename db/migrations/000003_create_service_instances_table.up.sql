CREATE TABLE IF NOT EXISTS service_instances(
    id serial PRIMARY KEY,
    environment_id integer REFERENCES environments (id),
    service_id integer REFERENCES services,
    created_at TIMESTAMP WITH TIME ZONE default now(),
    deleted_at TIMESTAMP WITH TIME ZONE default now(),
    updated_at TIMESTAMP WITH TIME ZONE default now()
);

CREATE TABLE IF NOT EXISTS deploys(
    id serial PRIMARY KEY,
    service_instance_id integer REFERENCES service_instances (id),
    deployed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE default now(),
    updated_at TIMESTAMP WITH TIME ZONE default now()
);

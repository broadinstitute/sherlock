CREATE TABLE IF NOT EXISTS service_instances(
    id serial PRIMARY KEY,
    environment_id integer REFERENCES environments (id),
    service_id integer REFERENCES services
);

CREATE TABLE IF NOT EXISTS environment_reservations(
    id                serial PRIMARY KEY,
    environment_id    integer REFERENCES environments (id),
    name              text,
    source            text,
    url               text,
    ttl               interval, 
    mutex             boolean,
    metadata          jsonb,
    created_at        timestamp WITH TIME ZONE default now(),
    updated_at        timestamp WITH TIME ZONE default now()
);

CREATE INDEX environment_reservations_environment_id_fkey ON environment_reservations (environment_id);

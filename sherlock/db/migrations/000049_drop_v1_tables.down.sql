-- 000001_create_environments_table.up.sql

CREATE TABLE IF NOT EXISTS environments
(
    id         serial PRIMARY KEY,
    name       VARCHAR(45) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE default now(),
    updated_at TIMESTAMP WITH TIME ZONE default now()
);

-- 000002_create_services_table.up.sql

CREATE TABLE IF NOT EXISTS services
(
    id         serial PRIMARY KEY,
    repo_url   TEXT,
    name       VARCHAR(45) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE default now(),
    updated_at TIMESTAMP WITH TIME ZONE default now()
);

CREATE INDEX services_name_idx ON services (name);

-- 000003_create_service_instances_table.up.sql

CREATE TABLE IF NOT EXISTS service_instances
(
    id             serial PRIMARY KEY,
    environment_id integer REFERENCES environments (id),
    service_id     integer REFERENCES services (id),
    created_at     TIMESTAMP WITH TIME ZONE default now(),
    updated_at     TIMESTAMP WITH TIME ZONE default now()
);

-- 000004_create_builds_table.up.sql

CREATE TABLE IF NOT EXISTS builds
(
    id             serial PRIMARY KEY,
    commit_sha     TEXT,
    build_url      TEXT,
    built_at       TIMESTAMP WITH TIME ZONE,
    version_string TEXT UNIQUE NOT NULL,
    service_id     integer REFERENCES services (id),
    created_at     TIMESTAMP WITH TIME ZONE default now(),
    updated_at     TIMESTAMP WITH TIME ZONE default now()
);

-- 000005_create_deploys_table.up.sql

CREATE TABLE IF NOT EXISTS deploys
(
    id                  serial PRIMARY KEY,
    service_instance_id integer REFERENCES service_instances (id),
    deployed_at         TIMESTAMP WITH TIME ZONE,
    created_at          TIMESTAMP WITH TIME ZONE default now(),
    updated_at          TIMESTAMP WITH TIME ZONE default now()
);

-- 000006_create_clusters_table.up.sql

CREATE TABLE IF NOT EXISTS clusters
(
    id             serial PRIMARY KEY,
    name           text UNIQUE NOT NULL,
    google_project text,
    created_at     timestamp WITH TIME ZONE default now(),
    updated_at     timestamp WITH TIME ZONE default now()
);

CREATE UNIQUE INDEX name_idx ON clusters (name);

-- 000007_create_allocation_pools_table.up.sql

CREATE TABLE IF NOT EXISTS allocation_pools
(
    id         serial PRIMARY KEY,
    name       text UNIQUE NOT NULL,
    created_at timestamp WITH TIME ZONE default now(),
    updated_at timestamp WITH TIME ZONE default now()
);


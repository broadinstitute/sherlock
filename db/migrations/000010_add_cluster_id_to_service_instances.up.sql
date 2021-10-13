ALTER TABLE service_instances
    ADD COLUMN cluster_id integer REFERENCES clusters (id);

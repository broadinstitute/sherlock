ALTER TABLE deploys
    ADD COLUMN build_id integer REFERENCES builds (id);

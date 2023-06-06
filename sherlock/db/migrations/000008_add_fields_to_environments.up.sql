ALTER TABLE environments
ADD COLUMN destroyed_at timestamp WITH TIME ZONE,
ADD COLUMN is_permanent boolean,
ADD COLUMN requester text,
ADD COLUMN allocation_pool_id integer REFERENCES allocation_pools (id);

-- follow standard psql naming: https://stackoverflow.com/questions/4107915/postgresql-default-constraint-names
CREATE INDEX environments_name_idx ON environments (name);
CREATE INDEX environments_allocation_pool_id_fkey ON environments (allocation_pool_id);

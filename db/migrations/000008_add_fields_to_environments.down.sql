ALTER TABLE environments
DROP COLUMN destroyed_at,
DROP COLUMN is_permanent,
DROP COLUMN requester,
DROP COLUMN allocation_pool_id;

DROP INDEX environments_name_idx;
DROP INDEX environments_allocation_pool_id_fkey;

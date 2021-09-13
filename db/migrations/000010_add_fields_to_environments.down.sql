
ALTER TABLE environments
DROP COLUMN destroyed_at,
DROP COLUMN is_permanent,
DROP COLUMN requester,
DROP COLUMN namespace,
DROP COLUMN customization,
DROP COLUMN cluster_id,
DROP COLUMN allocation_pool_id;

DROP INDEX environments_cluster_id_idx;
DROP INDEX environments_allocation_pool_id_idx;

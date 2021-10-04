ALTER TABLE service_instances
    ADD CONSTRAINT uq_service_environment UNIQUE(service_id, environment_id);

CREATE UNIQUE INDEX app_versions_selector_unique_constraint
    ON v2_app_versions (chart_id, app_version)
    WHERE deleted_at IS NULL;

ALTER TABLE v2_app_versions
    ADD CONSTRAINT app_version_present
        CHECK (app_version IS NOT NULL AND app_version != '');

DROP INDEX IF EXISTS app_versions_selector_unique_constraint;

ALTER TABLE v2_app_versions
    DROP CONSTRAINT IF EXISTS app_version_present;

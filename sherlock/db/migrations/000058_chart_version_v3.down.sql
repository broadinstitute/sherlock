DROP INDEX IF EXISTS chart_versions_selector_unique_constraint;

ALTER TABLE v2_chart_versions
    DROP CONSTRAINT IF EXISTS chart_version_present;

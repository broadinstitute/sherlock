CREATE UNIQUE INDEX chart_versions_selector_unique_constraint
ON v2_chart_versions (chart_id, chart_version)
WHERE deleted_at IS NULL;

ALTER TABLE v2_chart_versions
ADD CONSTRAINT chart_version_present
CHECK (chart_version IS NOT NULL AND chart_version != '');

ALTER TABLE v2_charts
    DROP CONSTRAINT IF EXISTS name_present;

ALTER TABLE v2_charts
    DROP CONSTRAINT IF EXISTS chart_repo_present;

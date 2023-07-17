ALTER TABLE v2_ci_identifiers
    DROP CONSTRAINT IF EXISTS resource_present;

DROP INDEX IF EXISTS ci_identifiers_selector_unique_constraint;

ALTER TABLE v2_ci_runs
    DROP CONSTRAINT IF EXISTS platform_present;

ALTER TABLE v2_ci_runs
    DROP CONSTRAINT IF EXISTS terminal_status_present;

DROP INDEX IF EXISTS ci_runs_selector_unique_constraint;

ALTER TABLE v2_charts
    ADD CONSTRAINT name_present
        CHECK (name IS NOT NULL AND name != '');

ALTER TABLE v2_charts
    ADD CONSTRAINT chart_repo_present
        CHECK (chart_repo IS NOT NULL AND chart_repo != '');

ALTER TABLE v2_users
    ADD CONSTRAINT github_info_together
        CHECK (
                (github_username IS NULL AND github_id IS NULL)
                OR
                (github_username IS NOT NULL AND github_id IS NOT NULL)
            );

ALTER TABLE v2_users
    ADD CONSTRAINT email_format
        CHECK (email LIKE '%_@_%_.__%');

CREATE UNIQUE INDEX users_github_username_unique_constraint
    ON v2_users (github_username)
    WHERE deleted_at IS NULL AND github_username IS NOT NULL;

CREATE UNIQUE INDEX users_github_id_unique_constraint
    ON v2_users (github_id)
    WHERE deleted_at IS NULL AND github_id IS NOT NULL;

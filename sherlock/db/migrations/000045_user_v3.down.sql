ALTER TABLE v2_users
    DROP CONSTRAINT IF EXISTS github_info_together;

ALTER TABLE v2_users
    DROP CONSTRAINT IF EXISTS email_format;

DROP INDEX IF EXISTS users_github_username_unique_constraint;

DROP INDEX IF EXISTS users_github_id_unique_constraint;

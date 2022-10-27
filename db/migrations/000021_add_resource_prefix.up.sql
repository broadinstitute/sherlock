-- No, this isn't unique, but older environments would've been created
-- without this value so it is meaningless anyway. Sherlock can handle
-- selector conflicts without completely blowing up, and prod is empty
-- right now.
alter table v2_environments
    add if not exists unique_resource_prefix text DEFAULT 'aaaa' NOT NULL;

alter table v2_environments
    alter column unique_resource_prefix DROP DEFAULT;

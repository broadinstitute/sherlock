-- Rather than carefully converting between the SQL internal type and ISO-8601, we just clear out the fields.
-- We clear out the enabled field too to keep everything valid--the one user (Jack) will need to reconfigure the
-- schedule for his BEE in Beehive but since he's writing this, that seems easier than testing some parsing nonsense.

alter table v2_environments
    drop column if exists offline_schedule_begin_enabled;

alter table v2_environments
    add column offline_schedule_begin_enabled boolean;

alter table v2_environments
    drop column if exists offline_schedule_begin_time;

alter table v2_environments
    add column offline_schedule_begin_time text;

alter table v2_environments
    drop column if exists offline_schedule_end_enabled;

alter table v2_environments
    add column offline_schedule_end_enabled boolean;

alter table v2_environments
    drop column if exists offline_schedule_end_time;

alter table v2_environments
    add column offline_schedule_end_time text;

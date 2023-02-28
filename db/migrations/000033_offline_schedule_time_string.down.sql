alter table v2_environments
    drop column if exists offline_schedule_begin_enabled;

alter table v2_environments
    add column offline_schedule_begin_enabled boolean;

alter table v2_environments
    drop column if exists offline_schedule_begin_time;

alter table v2_environments
    add column offline_schedule_begin_time timestamp with time zone;

alter table v2_environments
    drop column if exists offline_schedule_end_enabled;

alter table v2_environments
    add column offline_schedule_end_enabled boolean;

alter table v2_environments
    drop column if exists offline_schedule_end_time;

alter table v2_environments
    add column offline_schedule_end_time timestamp with time zone;

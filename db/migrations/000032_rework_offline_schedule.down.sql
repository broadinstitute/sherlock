alter table v2_environments
    rename column offline_schedule_begin_enabled to offline_schedule_enabled;

alter table v2_environments
    rename column offline_schedule_begin_time to offline_schedule_begin;

alter table v2_environments
    drop column if exists offline_schedule_end_enabled;

alter table v2_environments
    rename column offline_schedule_end_time to offline_schedule_end;

alter table v2_environments
    rename column offline_schedule_end_weekends to offline_schedule_end_weekday_only;

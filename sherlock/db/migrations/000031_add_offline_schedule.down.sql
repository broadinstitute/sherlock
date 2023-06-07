alter table v2_environments
    drop column if exists offline_schedule_enabled;

alter table v2_environments
    drop column if exists offline_schedule_begin;

alter table v2_environments
    drop column if exists offline_schedule_end;

alter table v2_environments
    drop column if exists offline_schedule_end_weekday_only;

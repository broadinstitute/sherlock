alter table v2_environments
    rename column offline_schedule_enabled to offline_schedule_begin_enabled;

alter table v2_environments
    rename column offline_schule_begin to offline_schedule_begin_time;

alter table v2_environments
    add if not exists offline_schedule_end_enabled boolean;

alter table v2_environments
    rename column offline_schedule_end to offline_schedule_end_time;

alter table v2_environments
    rename column offline_schedule_end_weekday_only to offline_schedule_end_weekends;

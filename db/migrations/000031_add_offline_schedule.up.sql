alter table v2_environments
    add if not exists offline_schedule_enabled boolean;

alter table v2_environments
    add if not exists offline_schedule_begin timestamp with time zone;

alter table v2_environments
    add if not exists offline_schedule_end timestamp with time zone;

alter table v2_environments
    add if not exists offline_schedule_weekday_only boolean;

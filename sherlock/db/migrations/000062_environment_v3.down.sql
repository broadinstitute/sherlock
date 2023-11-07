alter table environments
    drop constraint if exists name_valid;

alter table environments
    drop constraint if exists owner_id_present;

alter table environments
    drop constraint if exists lifecycle_valid;

alter table environments
    drop constraint if exists default_namespace_present;

alter table environments
    drop constraint if exists helmfile_ref_present;

alter table environments
    drop constraint if exists default_firecloud_develop_ref_present;

alter table environments
    drop constraint if exists unique_resource_prefix_present;

alter table environments
    drop constraint if exists prevent_deletion_valid;

alter table environments
    drop constraint if exists delete_after_valid;

alter table environments
    drop constraint if exists offline_valid;

alter table environments
    drop constraint if exists offline_schedule_begin_time_present;

alter table environments
    drop constraint if exists offline_schedule_end_time_present;

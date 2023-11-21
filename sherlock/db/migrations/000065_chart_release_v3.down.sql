alter table chart_releases
    drop constraint if exists name_present;

alter table chart_releases
    drop constraint if exists chart_id_present;

alter table chart_releases
    drop constraint if exists destination_type_valid;

alter table chart_releases
    drop constraint if exists cluster_id_namespace_valid;

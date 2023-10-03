alter table v2_clusters
    drop constraint if exists name_present;

alter table v2_clusters
    drop constraint if exists provider_present;

alter table v2_clusters
    drop constraint if exists base_present;

alter table v2_clusters
    drop constraint if exists address_present;

alter table v2_clusters
    drop constraint if exists location_present;

alter table v2_clusters
    drop constraint if exists helmfile_ref_present;

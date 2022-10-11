alter table v2_chart_releases
    drop column if exists firecloud_develop_ref;

alter table v2_changeset
    drop column if exists from_firecloud_develop_ref;

alter table v2_changeset
    drop column if exists to_firecloud_develop_ref;

alter table v2_clusters
    drop column if exists helmfile_ref;

alter table v2_environments
    drop column if exists helmfile_ref;

alter table v2_clusters
    drop column if exists location;

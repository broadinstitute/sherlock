alter table v2_chart_releases
    add if not exists firecloud_develop_ref text;

alter table v2_changesets
    add if not exists from_firecloud_develop_ref text;

alter table v2_changesets
    add if not exists to_firecloud_develop_ref text;

alter table v2_clusters
    add if not exists helmfile_ref text NOT NULL;

alter table v2_environments
    add if not exists helmfile_ref text NOT NULL;

alter table v2_clusters
    add if not exists location text NOT NULL;

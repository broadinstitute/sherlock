alter table v2_chart_releases
    add if not exists firecloud_develop_ref text;

alter table v2_changesets
    add if not exists from_firecloud_develop_ref text;

alter table v2_changesets
    add if not exists to_firecloud_develop_ref text;

alter table v2_clusters
    add if not exists helmfile_ref text DEFAULT 'HEAD' NOT NULL;

alter table v2_clusters
    alter column helmfile_ref DROP DEFAULT;

alter table v2_clusters
    add if not exists location text DEFAULT 'us-central1-a' NOT NULL;

alter table v2_clusters
    alter column location DROP DEFAULT;
    
alter table v2_environments
    add if not exists helmfile_ref text DEFAUlT 'HEAD' NOT NULL;

alter table v2_environments
   alter column helmfile_ref DROP DEFAULT;



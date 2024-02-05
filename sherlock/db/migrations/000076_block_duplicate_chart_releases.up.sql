create unique index chart_releases_name_unique
    on chart_releases (name)
    where deleted_at is null;

create unique index chart_releases_environment_chart_unique
    on chart_releases (environment_id, chart_id)
    where deleted_at is null and environment_id is not null;

create unique index chart_releases_cluster_namespace_chart_unique
    on chart_releases (cluster_id, namespace, chart_id)
    where deleted_at is null and cluster_id is not null and namespace is not null;

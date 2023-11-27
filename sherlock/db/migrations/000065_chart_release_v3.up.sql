alter table chart_releases
    add constraint name_present
        check (name is not null and name != '');

alter table chart_releases
    add constraint chart_id_present
        check (chart_id != 0);

alter table  chart_releases
    add constraint destination_type_valid
        check ((destination_type = 'environment' and
                environment_id is not null and environment_id != 0) or
               (destination_type = 'cluster' and
                cluster_id is not null and cluster_id != 0 and
                environment_id is null));

alter table chart_releases
    add constraint cluster_id_namespace_valid
        check ((cluster_id is null and
                (namespace = '' or namespace is null)) or
               (cluster_id is not null and
                (namespace != '' or namespace is not null)));


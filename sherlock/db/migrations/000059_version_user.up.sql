alter table v2_app_versions
    add if not exists authored_by_id bigint;

alter table v2_app_versions
    add constraint fk_v2_app_versions_authored_by
        foreign key (authored_by_id) references v2_users;

alter table v2_chart_versions
    add if not exists authored_by_id bigint;

alter table v2_chart_versions
    add constraint fk_v2_chart_versions_authored_by
        foreign key (authored_by_id) references v2_users;

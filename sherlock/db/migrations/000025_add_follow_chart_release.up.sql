alter table v2_chart_releases
    add if not exists app_version_follow_chart_release_id bigint;

alter table v2_chart_releases
    add constraint fk_v2_chart_releases_follow_chart_release_app_version
        foreign key (app_version_follow_chart_release_id) references v2_chart_releases;



alter table v2_chart_releases
    add if not exists chart_version_follow_chart_release_id bigint;

alter table v2_chart_releases
    add constraint fk_v2_chart_releases_follow_chart_release_chart_version
        foreign key (chart_version_follow_chart_release_id) references v2_chart_releases;



alter table v2_changesets
    add if not exists to_app_version_follow_chart_release_id bigint;

alter table v2_changesets
    add constraint fk_v2_changesets_to_follow_chart_release_app_version
        foreign key (to_app_version_follow_chart_release_id) references v2_chart_releases;

alter table v2_changesets
    add if not exists from_app_version_follow_chart_release_id bigint;

alter table v2_changesets
    add constraint fk_v2_changesets_from_follow_chart_release_app_version
        foreign key (from_app_version_follow_chart_release_id) references v2_chart_releases;



alter table v2_changesets
    add if not exists to_chart_version_follow_chart_release_id bigint;

alter table v2_changesets
    add constraint fk_v2_changesets_to_follow_chart_release_chart_version
        foreign key (to_chart_version_follow_chart_release_id) references v2_chart_releases;

alter table v2_changesets
    add if not exists from_chart_version_follow_chart_release_id bigint;

alter table v2_changesets
    add constraint fk_v2_changesets_from_follow_chart_release_chart_version
        foreign key (from_chart_version_follow_chart_release_id) references v2_chart_releases;

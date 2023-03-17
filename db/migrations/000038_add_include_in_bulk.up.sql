alter table v2_chart_releases
    add if not exists include_in_bulk_changesets bool DEFAULT true;

alter table v2_chart_releases
    alter column include_in_bulk_changesets DROP DEFAULT;

alter table v2_chart_releases
    add if not exists include_in_bulk_changesets bool;
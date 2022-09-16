alter table v2_charts
    add if not exists chart_exposes_endpoint boolean;

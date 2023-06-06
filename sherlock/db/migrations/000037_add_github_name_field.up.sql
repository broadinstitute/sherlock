alter table v2_users
    add if not exists name_inferred_from_github bool;

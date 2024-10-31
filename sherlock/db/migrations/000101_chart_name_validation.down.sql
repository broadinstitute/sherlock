alter table charts
    add constraint name_present
        check (name is not null and name != '');

alter table environments
    drop constraint if exists name_valid;

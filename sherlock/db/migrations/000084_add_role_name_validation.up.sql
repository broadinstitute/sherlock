alter table roles
    add constraint name_valid
        check (name is not null and name != ''  and name similar to '[a-z0-9]([-a-z0-9]*[a-z0-9])?');

alter table environments
    add column if not exists enable_janitor boolean;

update environments
    set enable_janitor = (case when lifecycle = 'static' then false else true end)
    where enable_janitor is null;

alter table environments
    alter column enable_janitor set not null;

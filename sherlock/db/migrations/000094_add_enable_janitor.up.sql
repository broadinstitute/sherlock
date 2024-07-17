alter table environments
    add column if not exists enable_janitor boolean;

update environments
    set enable_janitor = (case when lifecycle = 'static' then true else false end)
    where enable_janitor is null;

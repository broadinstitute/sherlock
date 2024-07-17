alter table environments
    add column if not exists enable_janitor boolean not null default (case when lifecycle = 'static' then false else true end);

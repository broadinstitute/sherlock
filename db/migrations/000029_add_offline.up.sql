alter table v2_environments
    add if not exists offline boolean DEFAULT false;

alter table v2_environments
    alter column offline DROP DEFAULT;

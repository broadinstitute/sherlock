create table if not exists suitabilities
(
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,

    email       text    not null
        unique primary key,
    suitable    boolean not null,
    description text    not null
);

create index if not exists idx_suitabilities_email
    on suitabilities (email);

create index if not exists idx_users_email
    on users (email);

create table clients
(
    id                               text not null
        primary key,
    client_secret_hash               bytea,
    client_secret_salt               bytea,
    client_secret_iterations         bigint,
    client_redirect_uris             text,
    client_post_logout_redirect_uris text,
    client_application_type          text,
    client_auth_method               text,
    client_id_token_lifetime         bigint,
    client_dev_mode                  boolean,
    client_clock_skew                bigint
);

create table auth_requests
(
    id                    text not null
        primary key,
    created_at            timestamp with time zone,
    done_at               timestamp with time zone,
    client_id             text
        constraint fk_auth_requests_client
            references clients
            on update cascade on delete cascade,
    nonce                 text,
    redirect_uri          text,
    response_type         text,
    response_mode         text,
    scopes                text,
    state                 text,
    code_challenge        text,
    code_challenge_method text,
    user_id               bigint
        constraint fk_auth_requests_user
            references users
            on update cascade on delete cascade
);

create table auth_request_codes
(
    code            text not null
        primary key,
    created_at      timestamp with time zone,
    auth_request_id text
        constraint fk_auth_request_codes_auth_request
            references auth_requests
            on update cascade on delete cascade
);

create table refresh_tokens
(
    id               text not null
        primary key,
    created_at       timestamp with time zone,
    token_hash       bytea unique,
    client_id        text
        constraint fk_refresh_tokens_client
            references clients
            on update cascade on delete cascade,
    scopes           text,
    original_auth_at timestamp with time zone,
    user_id          bigint
        constraint fk_refresh_tokens_user
            references users
            on update cascade on delete cascade
);

create table tokens
(
    id               text not null
        primary key,
    created_at       timestamp with time zone,
    refresh_token_id text
        constraint fk_tokens_refresh_token
            references refresh_tokens
            on update cascade on delete cascade,
    client_id        text
        constraint fk_tokens_client
            references clients
            on update cascade on delete cascade,
    scopes           text,
    expiry           timestamp with time zone,
    user_id          bigint
        constraint fk_tokens_user
            references users
            on update cascade on delete cascade
);

create table signing_keys
(
    id          text not null
        primary key,
    created_at  timestamp with time zone,
    public_key  bytea,
    private_key bytea
);

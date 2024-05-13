alter table role_operations
    alter from_name drop not null,
    alter to_name drop not null;

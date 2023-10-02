alter table v2_clusters
    add constraint name_present
        check (name is not null and name != '');

-- name uniqueness is already done with a simple `unique` constraint on the field (ignores soft-delete)

alter table v2_clusters
    add constraint provider_present
        check ((provider = 'google' and
                google_project is not null and google_project != '') or
               (provider = 'azure' and
                azure_subscription is not null and azure_subscription != ''));

alter table v2_clusters
    add constraint base_present
        check (base is not null and base != '');

alter table v2_clusters
    add constraint address_present
        check (address is not null and address != '');

alter table v2_clusters
    add constraint location_present
        check (location is not null and location != '');

-- ignore requires_suitability (already has a `not null` constraint that covers it)

alter table v2_clusters
    add constraint helmfile_ref_present
        check (helmfile_ref is not null and helmfile_ref != '');

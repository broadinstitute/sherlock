alter table changesets
    add constraint to_resolved_at_present
        check (to_resolved_at is not null);

alter table changesets
    add constraint to_app_version_resolver_valid
        check ((to_app_version_resolver is not null and to_app_version_resolver != '') and

               ((to_app_version_resolver = 'branch' and
                 to_app_version_branch is not null and to_app_version_branch != '' and
                 to_app_version_id is not null and
                 to_app_version_commit is not null and to_app_version_commit != '' and
                 to_app_version_exact is not null and to_app_version_exact != '') or

                (to_app_version_resolver = 'commit' and
                 to_app_version_commit is not null and to_app_version_commit != '' and
                 to_app_version_exact is not null and to_app_version_exact != '') or

                (to_app_version_resolver = 'exact' and
                 to_app_version_exact is not null and to_app_version_exact != '') or

                (to_app_version_resolver = 'follow' and to_app_version_follow_chart_release_id is not null) or

                (to_app_version_resolver = 'none' and
                 (to_app_version_branch is null or to_app_version_branch = '') and
                 (to_app_version_commit is null or to_app_version_commit = '') and
                 (to_app_version_exact is null or to_app_version_exact = '') and
                 to_app_version_id is null and
                 to_app_version_follow_chart_release_id is null)));

alter table changesets
    add constraint to_chart_version_resolver_valid
        check ((to_chart_version_resolver is not null and to_chart_version_resolver != '' and
                to_chart_version_exact is not null and to_chart_version_exact != '') and

               ((to_chart_version_resolver = 'latest' and
                 to_chart_version_id is not null) or

                (to_chart_version_resolver = 'exact') or

                (to_chart_version_resolver = 'follow' and
                 to_chart_version_follow_chart_release_id is not null)));

alter table changesets
    add constraint to_helmfile_ref_valid
        check ((to_helmfile_ref_enabled is null or to_helmfile_ref_enabled is false) or
               (to_helmfile_ref is not null and to_helmfile_ref != ''));

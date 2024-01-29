alter table chart_releases
    add constraint resolved_at_present
        check (resolved_at is not null);

alter table chart_releases
    add constraint app_version_resolver_valid
        check ((app_version_resolver is not null and app_version_resolver != '') and

               ((app_version_resolver = 'branch' and
                app_version_branch is not null and app_version_branch != '' and
                app_version_id is not null and
                app_version_commit is not null and app_version_commit != '' and
                app_version_exact is not null and app_version_exact != '') or

               (app_version_resolver = 'commit' and
                app_version_commit is not null and app_version_commit != '' and
                app_version_exact is not null and app_version_exact != '') or

               (app_version_resolver = 'exact' and
                app_version_exact is not null and app_version_exact != '') or

               (app_version_resolver = 'follow' and app_version_follow_chart_release_id is not null) or

               (app_version_resolver = 'none' and
                (app_version_branch is null or app_version_branch = '') and
                (app_version_commit is null or app_version_commit = '') and
                (app_version_exact is null or app_version_exact = '') and
                app_version_id is null and
                app_version_follow_chart_release_id is null)));

alter table chart_releases
    add constraint chart_version_resolver_valid
        check ((chart_version_resolver is not null and chart_version_resolver != '' and
                chart_version_exact is not null and chart_version_exact != '') and

               ((chart_version_resolver = 'latest' and
                 chart_version_id is not null) or

                (chart_version_resolver = 'exact' and
                 chart_version_id is null) or

                (chart_version_resolver = 'follow' and
                 chart_version_follow_chart_release_id is not null)));

alter table chart_releases
    add constraint helmfile_ref_valid
        check ((helmfile_ref_enabled is null or helmfile_ref_enabled is false) or
               (helmfile_ref is not null and helmfile_ref != ''));

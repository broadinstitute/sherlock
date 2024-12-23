# coding: utf-8

"""
    Sherlock

    The Data Science Platform's source-of-truth service. Note: this API will try to load and return associations in responses, so clients won't need to make as many requests. This behavior isn't recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).

    The version of the OpenAPI document: development
    Contact: dsp-devops@broadinstitute.org
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from sherlock_python_client.models.sherlock_environment_v3 import SherlockEnvironmentV3

class TestSherlockEnvironmentV3(unittest.TestCase):
    """SherlockEnvironmentV3 unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockEnvironmentV3:
        """Test SherlockEnvironmentV3
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockEnvironmentV3`
        """
        model = SherlockEnvironmentV3()
        if include_optional:
            return SherlockEnvironmentV3(
                auto_populate_chart_releases = True,
                base = '',
                base_domain = 'bee.envs-terra.bio',
                ci_identifier = sherlock_python_client.models.sherlock/ci_identifier_v3.sherlock.CiIdentifierV3(
                    ci_runs = [
                        sherlock_python_client.models.sherlock/ci_run_v3.sherlock.CiRunV3(
                            argo_workflows_name = '', 
                            argo_workflows_namespace = '', 
                            argo_workflows_template = '', 
                            created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            github_actions_attempt_number = 56, 
                            github_actions_owner = '', 
                            github_actions_repo = '', 
                            github_actions_run_id = 56, 
                            github_actions_workflow_path = '', 
                            id = 56, 
                            notify_slack_channels_upon_failure = [
                                ''
                                ], 
                            notify_slack_channels_upon_retry = [
                                ''
                                ], 
                            notify_slack_channels_upon_success = [
                                ''
                                ], 
                            notify_slack_custom_icon = '', 
                            platform = '', 
                            related_resources = [
                                sherlock_python_client.models.sherlock/ci_identifier_v3.sherlock.CiIdentifierV3(
                                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                    id = 56, 
                                    resource_id = 56, 
                                    resource_status = '', 
                                    resource_type = '', 
                                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                                ], 
                            resource_status = '', 
                            started_at = '', 
                            status = '', 
                            terminal_at = '', 
                            termination_hooks_dispatched_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                        ], 
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    id = 56, 
                    resource_id = 56, 
                    resource_status = '', 
                    resource_type = '', 
                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ),
                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                default_cluster = '',
                default_cluster_info = sherlock_python_client.models.sherlock/cluster_v3.sherlock.ClusterV3(
                    address = '', 
                    azure_subscription = '', 
                    base = '', 
                    ci_identifier = sherlock_python_client.models.sherlock/ci_identifier_v3.sherlock.CiIdentifierV3(
                        ci_runs = [
                            sherlock_python_client.models.sherlock/ci_run_v3.sherlock.CiRunV3(
                                argo_workflows_name = '', 
                                argo_workflows_namespace = '', 
                                argo_workflows_template = '', 
                                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                github_actions_attempt_number = 56, 
                                github_actions_owner = '', 
                                github_actions_repo = '', 
                                github_actions_run_id = 56, 
                                github_actions_workflow_path = '', 
                                id = 56, 
                                notify_slack_channels_upon_failure = [
                                    ''
                                    ], 
                                notify_slack_channels_upon_retry = [
                                    ''
                                    ], 
                                notify_slack_channels_upon_success = [
                                    ''
                                    ], 
                                notify_slack_custom_icon = '', 
                                platform = '', 
                                related_resources = [
                                    sherlock_python_client.models.sherlock/ci_identifier_v3.sherlock.CiIdentifierV3(
                                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                        id = 56, 
                                        resource_id = 56, 
                                        resource_status = '', 
                                        resource_type = '', 
                                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                                    ], 
                                resource_status = '', 
                                started_at = '', 
                                status = '', 
                                terminal_at = '', 
                                termination_hooks_dispatched_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                            ], 
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        id = 56, 
                        resource_id = 56, 
                        resource_status = '', 
                        resource_type = '', 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ), 
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    google_project = '', 
                    helmfile_ref = 'HEAD', 
                    id = 56, 
                    location = 'us-central1-a', 
                    name = '', 
                    provider = 'google', 
                    required_role = '', 
                    required_role_info = sherlock_python_client.models.sherlock/role_v3.sherlock.RoleV3(
                        assignments = [
                            sherlock_python_client.models.sherlock/role_assignment_v3.sherlock.RoleAssignmentV3(
                                expires_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                                expires_in = '', 
                                role_info = sherlock_python_client.models.role_info.roleInfo(), 
                                suspended = True, 
                                user_info = sherlock_python_client.models.user_info.userInfo(), )
                            ], 
                        auto_assign_all_users = True, 
                        can_be_glass_broken_by_role = 56, 
                        can_be_glass_broken_by_role_info = sherlock_python_client.models.can_be_glass_broken_by_role_info.canBeGlassBrokenByRoleInfo(), 
                        created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        default_glass_break_duration = '', 
                        grants_broad_institute_group = '', 
                        grants_dev_azure_account = True, 
                        grants_dev_azure_directory_roles = True, 
                        grants_dev_azure_group = '', 
                        grants_dev_firecloud_folder_owner = '', 
                        grants_dev_firecloud_group = '', 
                        grants_prod_azure_account = True, 
                        grants_prod_azure_directory_roles = True, 
                        grants_prod_azure_group = '', 
                        grants_prod_firecloud_folder_owner = '', 
                        grants_prod_firecloud_group = '', 
                        grants_qa_firecloud_folder_owner = '', 
                        grants_qa_firecloud_group = '', 
                        grants_sherlock_super_admin = True, 
                        id = 56, 
                        name = '', 
                        suspend_non_suitable_users = True, 
                        updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ), 
                    requires_suitability = True, 
                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ),
                default_namespace = '',
                delete_after = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                description = '',
                enable_janitor = True,
                helmfile_ref = 'HEAD',
                id = 56,
                lifecycle = 'dynamic',
                name = '',
                name_prefixes_domain = True,
                offline = True,
                offline_schedule_begin_enabled = True,
                offline_schedule_begin_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                offline_schedule_end_enabled = True,
                offline_schedule_end_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                offline_schedule_end_weekends = True,
                owner = '',
                owner_info = sherlock_python_client.models.sherlock/user_v3.sherlock.UserV3(
                    assignments = [
                        sherlock_python_client.models.sherlock/role_assignment_v3.sherlock.RoleAssignmentV3(
                            expires_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            expires_in = '', 
                            role_info = sherlock_python_client.models.role_info.roleInfo(), 
                            suspended = True, 
                            user_info = sherlock_python_client.models.user_info.userInfo(), )
                        ], 
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    deactivated_at = '', 
                    email = '', 
                    github_id = '', 
                    github_username = '', 
                    google_id = '', 
                    id = 56, 
                    name = '', 
                    name_from = 'sherlock', 
                    slack_id = '', 
                    slack_username = '', 
                    suitability_description = '', 
                    suitable = True, 
                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ),
                pact_identifier = '',
                pagerduty_integration = '',
                pagerduty_integration_info = sherlock_python_client.models.sherlock/pagerduty_integration_v3.sherlock.PagerdutyIntegrationV3(
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    id = 56, 
                    name = '', 
                    pagerduty_id = '', 
                    type = '', 
                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ),
                prevent_deletion = True,
                required_role = '',
                required_role_info = sherlock_python_client.models.sherlock/role_v3.sherlock.RoleV3(
                    assignments = [
                        sherlock_python_client.models.sherlock/role_assignment_v3.sherlock.RoleAssignmentV3(
                            expires_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            expires_in = '', 
                            role_info = sherlock_python_client.models.role_info.roleInfo(), 
                            suspended = True, 
                            user_info = sherlock_python_client.models.user_info.userInfo(), )
                        ], 
                    auto_assign_all_users = True, 
                    can_be_glass_broken_by_role = 56, 
                    can_be_glass_broken_by_role_info = sherlock_python_client.models.can_be_glass_broken_by_role_info.canBeGlassBrokenByRoleInfo(), 
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    default_glass_break_duration = '', 
                    grants_broad_institute_group = '', 
                    grants_dev_azure_account = True, 
                    grants_dev_azure_directory_roles = True, 
                    grants_dev_azure_group = '', 
                    grants_dev_firecloud_folder_owner = '', 
                    grants_dev_firecloud_group = '', 
                    grants_prod_azure_account = True, 
                    grants_prod_azure_directory_roles = True, 
                    grants_prod_azure_group = '', 
                    grants_prod_firecloud_folder_owner = '', 
                    grants_prod_firecloud_group = '', 
                    grants_qa_firecloud_folder_owner = '', 
                    grants_qa_firecloud_group = '', 
                    grants_sherlock_super_admin = True, 
                    id = 56, 
                    name = '', 
                    suspend_non_suitable_users = True, 
                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ),
                requires_suitability = True,
                service_banner_bucket = '',
                template_environment = '',
                template_environment_info = sherlock_python_client.models.template_environment_info.templateEnvironmentInfo(),
                unique_resource_prefix = '',
                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                values_name = ''
            )
        else:
            return SherlockEnvironmentV3(
        )
        """

    def testSherlockEnvironmentV3(self):
        """Test SherlockEnvironmentV3"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

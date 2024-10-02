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

from sherlock_python_client.models.sherlock_chart_version_v3 import SherlockChartVersionV3

class TestSherlockChartVersionV3(unittest.TestCase):
    """SherlockChartVersionV3 unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockChartVersionV3:
        """Test SherlockChartVersionV3
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockChartVersionV3`
        """
        model = SherlockChartVersionV3()
        if include_optional:
            return SherlockChartVersionV3(
                authored_by = '',
                authored_by_info = sherlock_python_client.models.sherlock/user_v3.sherlock.UserV3(
                    assignments = [
                        sherlock_python_client.models.sherlock/role_assignment_v3.sherlock.RoleAssignmentV3(
                            expires_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            expires_in = '', 
                            role_info = sherlock_python_client.models.role_info.roleInfo(), 
                            suspended = True, 
                            user_info = sherlock_python_client.models.user_info.userInfo(), )
                        ], 
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
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
                chart = '',
                chart_info = sherlock_python_client.models.sherlock/chart_v3.sherlock.ChartV3(
                    app_image_git_main_branch = '', 
                    app_image_git_repo = '', 
                    chart_exposes_endpoint = True, 
                    chart_repo = 'terra-helm', 
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
                    default_port = 56, 
                    default_protocol = 'https', 
                    default_subdomain = '', 
                    description = '', 
                    id = 56, 
                    name = '', 
                    pact_participant = True, 
                    playbook_url = '', 
                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ),
                chart_version = '',
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
                description = '',
                id = 56,
                parent_chart_version = '',
                parent_chart_version_info = sherlock_python_client.models.parent_chart_version_info.parentChartVersionInfo(),
                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else:
            return SherlockChartVersionV3(
        )
        """

    def testSherlockChartVersionV3(self):
        """Test SherlockChartVersionV3"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

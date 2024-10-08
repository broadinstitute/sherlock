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

from sherlock_python_client.models.sherlock_changeset_v3_plan_request import SherlockChangesetV3PlanRequest

class TestSherlockChangesetV3PlanRequest(unittest.TestCase):
    """SherlockChangesetV3PlanRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockChangesetV3PlanRequest:
        """Test SherlockChangesetV3PlanRequest
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockChangesetV3PlanRequest`
        """
        model = SherlockChangesetV3PlanRequest()
        if include_optional:
            return SherlockChangesetV3PlanRequest(
                chart_releases = [
                    sherlock_python_client.models.sherlock/changeset_v3_plan_request_chart_release_entry.sherlock.ChangesetV3PlanRequestChartReleaseEntry(
                        chart_release = '', 
                        follow_versions_from_other_chart_release = '', 
                        to_app_version_branch = '', 
                        to_app_version_commit = '', 
                        to_app_version_exact = '', 
                        to_app_version_follow_chart_release = '', 
                        to_app_version_resolver = '', 
                        to_chart_version_exact = '', 
                        to_chart_version_follow_chart_release = '', 
                        to_chart_version_resolver = '', 
                        to_helmfile_ref = '', 
                        to_helmfile_ref_enabled = True, 
                        use_exact_versions_from_other_chart_release = '', )
                    ],
                environments = [
                    sherlock_python_client.models.sherlock/changeset_v3_plan_request_environment_entry.sherlock.ChangesetV3PlanRequestEnvironmentEntry(
                        environment = '', 
                        exclude_charts = [
                            ''
                            ], 
                        filter_to_matching_branches = True, 
                        follow_versions_from_other_environment = '', 
                        include_charts = [
                            ''
                            ], 
                        use_exact_versions_from_other_environment = '', )
                    ],
                recreate_changesets = [
                    56
                    ]
            )
        else:
            return SherlockChangesetV3PlanRequest(
        )
        """

    def testSherlockChangesetV3PlanRequest(self):
        """Test SherlockChangesetV3PlanRequest"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

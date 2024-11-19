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

from sherlock_python_client.models.sherlock_environment_v3_edit import SherlockEnvironmentV3Edit

class TestSherlockEnvironmentV3Edit(unittest.TestCase):
    """SherlockEnvironmentV3Edit unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockEnvironmentV3Edit:
        """Test SherlockEnvironmentV3Edit
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockEnvironmentV3Edit`
        """
        model = SherlockEnvironmentV3Edit()
        if include_optional:
            return SherlockEnvironmentV3Edit(
                base_domain = 'bee.envs-terra.bio',
                default_cluster = '',
                delete_after = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                description = '',
                enable_janitor = True,
                helmfile_ref = 'HEAD',
                name_prefixes_domain = True,
                offline = True,
                offline_schedule_begin_enabled = True,
                offline_schedule_begin_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                offline_schedule_end_enabled = True,
                offline_schedule_end_time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                offline_schedule_end_weekends = True,
                owner = '',
                pact_identifier = '',
                pagerduty_integration = '',
                prevent_deletion = True,
                required_role = '',
                requires_suitability = True,
                service_banner_bucket = ''
            )
        else:
            return SherlockEnvironmentV3Edit(
        )
        """

    def testSherlockEnvironmentV3Edit(self):
        """Test SherlockEnvironmentV3Edit"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

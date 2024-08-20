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

from sherlock_python_client.models.sherlock_user_v3 import SherlockUserV3

class TestSherlockUserV3(unittest.TestCase):
    """SherlockUserV3 unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockUserV3:
        """Test SherlockUserV3
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockUserV3`
        """
        model = SherlockUserV3()
        if include_optional:
            return SherlockUserV3(
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
                name_inferred_from_github = True,
                slack_id = '',
                slack_username = '',
                suitability_description = '',
                suitable = True,
                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else:
            return SherlockUserV3(
        )
        """

    def testSherlockUserV3(self):
        """Test SherlockUserV3"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

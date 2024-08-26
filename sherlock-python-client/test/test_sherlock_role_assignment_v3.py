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

from sherlock_python_client.models.sherlock_role_assignment_v3 import SherlockRoleAssignmentV3

class TestSherlockRoleAssignmentV3(unittest.TestCase):
    """SherlockRoleAssignmentV3 unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockRoleAssignmentV3:
        """Test SherlockRoleAssignmentV3
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockRoleAssignmentV3`
        """
        model = SherlockRoleAssignmentV3()
        if include_optional:
            return SherlockRoleAssignmentV3(
                expires_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                expires_in = '',
                role_info = sherlock_python_client.models.role_info.roleInfo(),
                suspended = True,
                user_info = sherlock_python_client.models.user_info.userInfo()
            )
        else:
            return SherlockRoleAssignmentV3(
        )
        """

    def testSherlockRoleAssignmentV3(self):
        """Test SherlockRoleAssignmentV3"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
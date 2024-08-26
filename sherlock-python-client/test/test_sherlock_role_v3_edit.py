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

from sherlock_python_client.models.sherlock_role_v3_edit import SherlockRoleV3Edit

class TestSherlockRoleV3Edit(unittest.TestCase):
    """SherlockRoleV3Edit unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockRoleV3Edit:
        """Test SherlockRoleV3Edit
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockRoleV3Edit`
        """
        model = SherlockRoleV3Edit()
        if include_optional:
            return SherlockRoleV3Edit(
                auto_assign_all_users = True,
                can_be_glass_broken_by_role = 56,
                default_glass_break_duration = '',
                grants_broad_institute_group = '',
                grants_dev_azure_group = '',
                grants_dev_firecloud_group = '',
                grants_prod_azure_group = '',
                grants_prod_firecloud_group = '',
                grants_qa_firecloud_group = '',
                grants_sherlock_super_admin = True,
                name = '',
                suspend_non_suitable_users = True
            )
        else:
            return SherlockRoleV3Edit(
        )
        """

    def testSherlockRoleV3Edit(self):
        """Test SherlockRoleV3Edit"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
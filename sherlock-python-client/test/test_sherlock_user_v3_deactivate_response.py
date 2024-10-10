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

from sherlock_python_client.models.sherlock_user_v3_deactivate_response import SherlockUserV3DeactivateResponse

class TestSherlockUserV3DeactivateResponse(unittest.TestCase):
    """SherlockUserV3DeactivateResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockUserV3DeactivateResponse:
        """Test SherlockUserV3DeactivateResponse
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockUserV3DeactivateResponse`
        """
        model = SherlockUserV3DeactivateResponse()
        if include_optional:
            return SherlockUserV3DeactivateResponse(
                already_deactivated_emails = [
                    ''
                    ],
                newly_deactivated_emails = [
                    ''
                    ],
                not_found_emails = [
                    ''
                    ]
            )
        else:
            return SherlockUserV3DeactivateResponse(
        )
        """

    def testSherlockUserV3DeactivateResponse(self):
        """Test SherlockUserV3DeactivateResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

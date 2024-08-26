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

from sherlock_python_client.models.sherlock_pagerduty_integration_v3 import SherlockPagerdutyIntegrationV3

class TestSherlockPagerdutyIntegrationV3(unittest.TestCase):
    """SherlockPagerdutyIntegrationV3 unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockPagerdutyIntegrationV3:
        """Test SherlockPagerdutyIntegrationV3
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockPagerdutyIntegrationV3`
        """
        model = SherlockPagerdutyIntegrationV3()
        if include_optional:
            return SherlockPagerdutyIntegrationV3(
                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                id = 56,
                name = '',
                pagerduty_id = '',
                type = '',
                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else:
            return SherlockPagerdutyIntegrationV3(
        )
        """

    def testSherlockPagerdutyIntegrationV3(self):
        """Test SherlockPagerdutyIntegrationV3"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
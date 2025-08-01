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

from sherlock_python_client.models.sherlock_service_alert_v3 import SherlockServiceAlertV3

class TestSherlockServiceAlertV3(unittest.TestCase):
    """SherlockServiceAlertV3 unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockServiceAlertV3:
        """Test SherlockServiceAlertV3
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockServiceAlertV3`
        """
        model = SherlockServiceAlertV3()
        if include_optional:
            return SherlockServiceAlertV3(
                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                created_by = '',
                deleted_by = '',
                id = 56,
                link = '',
                message = '',
                on_environment = '',
                severity = 'blocker',
                title = '',
                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                updated_by = '',
                uuid = ''
            )
        else:
            return SherlockServiceAlertV3(
        )
        """

    def testSherlockServiceAlertV3(self):
        """Test SherlockServiceAlertV3"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

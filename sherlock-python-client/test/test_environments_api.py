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

from sherlock_python_client.api.environments_api import EnvironmentsApi


class TestEnvironmentsApi(unittest.TestCase):
    """EnvironmentsApi unit test stubs"""

    def setUp(self) -> None:
        self.api = EnvironmentsApi()

    def tearDown(self) -> None:
        pass

    def test_api_environments_v3_get(self) -> None:
        """Test case for api_environments_v3_get

        List Environments matching a filter
        """
        pass

    def test_api_environments_v3_post(self) -> None:
        """Test case for api_environments_v3_post

        Create a Environment
        """
        pass

    def test_api_environments_v3_selector_delete(self) -> None:
        """Test case for api_environments_v3_selector_delete

        Delete an individual Environment
        """
        pass

    def test_api_environments_v3_selector_get(self) -> None:
        """Test case for api_environments_v3_selector_get

        Get an individual Environment
        """
        pass

    def test_api_environments_v3_selector_patch(self) -> None:
        """Test case for api_environments_v3_selector_patch

        Edit an individual Environment
        """
        pass


if __name__ == '__main__':
    unittest.main()
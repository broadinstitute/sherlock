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

from sherlock_python_client.api.git_commits_api import GitCommitsApi


class TestGitCommitsApi(unittest.TestCase):
    """GitCommitsApi unit test stubs"""

    def setUp(self) -> None:
        self.api = GitCommitsApi()

    def tearDown(self) -> None:
        pass

    def test_api_git_commits_v3_put(self) -> None:
        """Test case for api_git_commits_v3_put

        Upsert a GitCommit
        """
        pass


if __name__ == '__main__':
    unittest.main()

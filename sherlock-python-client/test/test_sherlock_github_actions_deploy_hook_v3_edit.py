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

from sherlock_python_client.models.sherlock_github_actions_deploy_hook_v3_edit import SherlockGithubActionsDeployHookV3Edit

class TestSherlockGithubActionsDeployHookV3Edit(unittest.TestCase):
    """SherlockGithubActionsDeployHookV3Edit unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SherlockGithubActionsDeployHookV3Edit:
        """Test SherlockGithubActionsDeployHookV3Edit
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SherlockGithubActionsDeployHookV3Edit`
        """
        model = SherlockGithubActionsDeployHookV3Edit()
        if include_optional:
            return SherlockGithubActionsDeployHookV3Edit(
                github_actions_default_ref = '',
                github_actions_owner = '',
                github_actions_ref_behavior = 'always-use-default-ref',
                github_actions_repo = '',
                github_actions_workflow_inputs = None,
                github_actions_workflow_path = '',
                on_failure = True,
                on_success = True
            )
        else:
            return SherlockGithubActionsDeployHookV3Edit(
        )
        """

    def testSherlockGithubActionsDeployHookV3Edit(self):
        """Test SherlockGithubActionsDeployHookV3Edit"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
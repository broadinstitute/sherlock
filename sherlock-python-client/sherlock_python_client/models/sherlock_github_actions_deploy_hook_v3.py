# coding: utf-8

"""
    Sherlock

    The Data Science Platform's source-of-truth service. Note: this API will try to load and return associations in responses, so clients won't need to make as many requests. This behavior isn't recursive, though, so associations of associations are *not* fully loaded (even if it might seem that way from looking at the data types).

    The version of the OpenAPI document: development
    Contact: dsp-devops@broadinstitute.org
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictInt, StrictStr, field_validator
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class SherlockGithubActionsDeployHookV3(BaseModel):
    """
    SherlockGithubActionsDeployHookV3
    """ # noqa: E501
    created_at: Optional[datetime] = Field(default=None, alias="createdAt")
    github_actions_default_ref: Optional[StrictStr] = Field(default=None, alias="githubActionsDefaultRef")
    github_actions_owner: Optional[StrictStr] = Field(default=None, alias="githubActionsOwner")
    github_actions_ref_behavior: Optional[StrictStr] = Field(default='always-use-default-ref', description="This field determines what git ref the workflow will be run on. The default of always-use-default-ref always uses the default ref; use-app-version-as-ref will use the app version (when available) as the ref, useful when versions are always commit hashes or tags; use-app-version-commit-as-ref will use the app version's commit (when available) as the ref, useful when the repo is configured to fully report app versions to Sherlock.", alias="githubActionsRefBehavior")
    github_actions_repo: Optional[StrictStr] = Field(default=None, alias="githubActionsRepo")
    github_actions_workflow_inputs: Optional[Dict[str, Any]] = Field(default=None, description="These workflow inputs will be passed statically as-is to GitHub's workflow dispatch API (https://docs.github.com/en/rest/actions/workflows#create-a-workflow-dispatch-event) as the `inputs` parameter object.", alias="githubActionsWorkflowInputs")
    github_actions_workflow_path: Optional[StrictStr] = Field(default=None, alias="githubActionsWorkflowPath")
    id: Optional[StrictInt] = None
    on_chart_release: Optional[StrictStr] = Field(default=None, alias="onChartRelease")
    on_environment: Optional[StrictStr] = Field(default=None, alias="onEnvironment")
    on_failure: Optional[StrictBool] = Field(default=None, alias="onFailure")
    on_success: Optional[StrictBool] = Field(default=None, alias="onSuccess")
    updated_at: Optional[datetime] = Field(default=None, alias="updatedAt")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["createdAt", "githubActionsDefaultRef", "githubActionsOwner", "githubActionsRefBehavior", "githubActionsRepo", "githubActionsWorkflowInputs", "githubActionsWorkflowPath", "id", "onChartRelease", "onEnvironment", "onFailure", "onSuccess", "updatedAt"]

    @field_validator('github_actions_ref_behavior')
    def github_actions_ref_behavior_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['always-use-default-ref', 'use-app-version-as-ref', 'use-app-version-commit-as-ref']):
            raise ValueError("must be one of enum values ('always-use-default-ref', 'use-app-version-as-ref', 'use-app-version-commit-as-ref')")
        return value

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of SherlockGithubActionsDeployHookV3 from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        * Fields in `self.additional_properties` are added to the output dict.
        """
        excluded_fields: Set[str] = set([
            "additional_properties",
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of SherlockGithubActionsDeployHookV3 from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "createdAt": obj.get("createdAt"),
            "githubActionsDefaultRef": obj.get("githubActionsDefaultRef"),
            "githubActionsOwner": obj.get("githubActionsOwner"),
            "githubActionsRefBehavior": obj.get("githubActionsRefBehavior") if obj.get("githubActionsRefBehavior") is not None else 'always-use-default-ref',
            "githubActionsRepo": obj.get("githubActionsRepo"),
            "githubActionsWorkflowInputs": obj.get("githubActionsWorkflowInputs"),
            "githubActionsWorkflowPath": obj.get("githubActionsWorkflowPath"),
            "id": obj.get("id"),
            "onChartRelease": obj.get("onChartRelease"),
            "onEnvironment": obj.get("onEnvironment"),
            "onFailure": obj.get("onFailure"),
            "onSuccess": obj.get("onSuccess"),
            "updatedAt": obj.get("updatedAt")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj



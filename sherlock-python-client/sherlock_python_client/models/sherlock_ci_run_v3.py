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
from pydantic import BaseModel, ConfigDict, Field, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class SherlockCiRunV3(BaseModel):
    """
    SherlockCiRunV3
    """ # noqa: E501
    argo_workflows_name: Optional[StrictStr] = Field(default=None, alias="argoWorkflowsName")
    argo_workflows_namespace: Optional[StrictStr] = Field(default=None, alias="argoWorkflowsNamespace")
    argo_workflows_template: Optional[StrictStr] = Field(default=None, alias="argoWorkflowsTemplate")
    created_at: Optional[datetime] = Field(default=None, alias="createdAt")
    github_actions_attempt_number: Optional[StrictInt] = Field(default=None, alias="githubActionsAttemptNumber")
    github_actions_owner: Optional[StrictStr] = Field(default=None, alias="githubActionsOwner")
    github_actions_repo: Optional[StrictStr] = Field(default=None, alias="githubActionsRepo")
    github_actions_run_id: Optional[StrictInt] = Field(default=None, alias="githubActionsRunID")
    github_actions_workflow_path: Optional[StrictStr] = Field(default=None, alias="githubActionsWorkflowPath")
    id: Optional[StrictInt] = None
    notify_slack_channels_upon_failure: Optional[List[StrictStr]] = Field(default=None, description="Slack channels to notify if this CiRun fails. This field is always appended to when mutated.", alias="notifySlackChannelsUponFailure")
    notify_slack_channels_upon_retry: Optional[List[StrictStr]] = Field(default=None, description="Slack channels to notify if this CiRun is retried. This field is always appended to when mutated. It will de-dupe with the other notify fields.", alias="notifySlackChannelsUponRetry")
    notify_slack_channels_upon_success: Optional[List[StrictStr]] = Field(default=None, description="Slack channels to notify if this CiRun succeeds. This field is always appended to when mutated.", alias="notifySlackChannelsUponSuccess")
    notify_slack_custom_icon: Optional[StrictStr] = Field(default=None, description="Icon to use for success or failure Slack notifications. Can be given either as a URL to an image or as a Slack emoji (using colon shortcodes, like :smiley:). An empty string is ignored to facilitate calling from GitHub Actions (where it's easier to pass an empty string than not send the field at all).", alias="notifySlackCustomIcon")
    platform: Optional[StrictStr] = None
    related_resources: Optional[List[SherlockCiIdentifierV3]] = Field(default=None, alias="relatedResources")
    resource_status: Optional[StrictStr] = Field(default=None, description="Available only when querying a CiRun via a CiIdentifier, indicates the status of the run for that resource", alias="resourceStatus")
    started_at: Optional[StrictStr] = Field(default=None, alias="startedAt")
    status: Optional[StrictStr] = None
    terminal_at: Optional[StrictStr] = Field(default=None, alias="terminalAt")
    termination_hooks_dispatched_at: Optional[datetime] = Field(default=None, alias="terminationHooksDispatchedAt")
    updated_at: Optional[datetime] = Field(default=None, alias="updatedAt")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["argoWorkflowsName", "argoWorkflowsNamespace", "argoWorkflowsTemplate", "createdAt", "githubActionsAttemptNumber", "githubActionsOwner", "githubActionsRepo", "githubActionsRunID", "githubActionsWorkflowPath", "id", "notifySlackChannelsUponFailure", "notifySlackChannelsUponRetry", "notifySlackChannelsUponSuccess", "notifySlackCustomIcon", "platform", "relatedResources", "resourceStatus", "startedAt", "status", "terminalAt", "terminationHooksDispatchedAt", "updatedAt"]

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
        """Create an instance of SherlockCiRunV3 from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in related_resources (list)
        _items = []
        if self.related_resources:
            for _item_related_resources in self.related_resources:
                if _item_related_resources:
                    _items.append(_item_related_resources.to_dict())
            _dict['relatedResources'] = _items
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of SherlockCiRunV3 from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "argoWorkflowsName": obj.get("argoWorkflowsName"),
            "argoWorkflowsNamespace": obj.get("argoWorkflowsNamespace"),
            "argoWorkflowsTemplate": obj.get("argoWorkflowsTemplate"),
            "createdAt": obj.get("createdAt"),
            "githubActionsAttemptNumber": obj.get("githubActionsAttemptNumber"),
            "githubActionsOwner": obj.get("githubActionsOwner"),
            "githubActionsRepo": obj.get("githubActionsRepo"),
            "githubActionsRunID": obj.get("githubActionsRunID"),
            "githubActionsWorkflowPath": obj.get("githubActionsWorkflowPath"),
            "id": obj.get("id"),
            "notifySlackChannelsUponFailure": obj.get("notifySlackChannelsUponFailure"),
            "notifySlackChannelsUponRetry": obj.get("notifySlackChannelsUponRetry"),
            "notifySlackChannelsUponSuccess": obj.get("notifySlackChannelsUponSuccess"),
            "notifySlackCustomIcon": obj.get("notifySlackCustomIcon"),
            "platform": obj.get("platform"),
            "relatedResources": [SherlockCiIdentifierV3.from_dict(_item) for _item in obj["relatedResources"]] if obj.get("relatedResources") is not None else None,
            "resourceStatus": obj.get("resourceStatus"),
            "startedAt": obj.get("startedAt"),
            "status": obj.get("status"),
            "terminalAt": obj.get("terminalAt"),
            "terminationHooksDispatchedAt": obj.get("terminationHooksDispatchedAt"),
            "updatedAt": obj.get("updatedAt")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj

from sherlock_python_client.models.sherlock_ci_identifier_v3 import SherlockCiIdentifierV3
# TODO: Rewrite to not use raise_errors
SherlockCiRunV3.model_rebuild(raise_errors=False)


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
from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class SherlockEnvironmentV3Edit(BaseModel):
    """
    SherlockEnvironmentV3Edit
    """ # noqa: E501
    base_domain: Optional[StrictStr] = Field(default='bee.envs-terra.bio', alias="baseDomain")
    default_cluster: Optional[StrictStr] = Field(default=None, alias="defaultCluster")
    delete_after: Optional[datetime] = Field(default=None, description="If set, the BEE will be automatically deleted after this time. Can be set to \"\" or Go's zero time value to clear the field.", alias="deleteAfter")
    description: Optional[StrictStr] = None
    enable_janitor: Optional[StrictBool] = Field(default=None, description="If true, janitor resource cleanup will be enabled for this environment. BEEs default to template's value, templates default to true, and static/live environments default to false.", alias="enableJanitor")
    helmfile_ref: Optional[StrictStr] = Field(default='HEAD', alias="helmfileRef")
    name_prefixes_domain: Optional[StrictBool] = Field(default=True, alias="namePrefixesDomain")
    offline: Optional[StrictBool] = Field(default=False, description="Applicable for BEEs only, whether Thelma should render the BEE as \"offline\" zero replicas (this field is a target state, not a status)")
    offline_schedule_begin_enabled: Optional[StrictBool] = Field(default=None, description="When enabled, the BEE will be slated to go offline around the begin time each day", alias="offlineScheduleBeginEnabled")
    offline_schedule_begin_time: Optional[datetime] = Field(default=None, description="Stored with timezone to determine day of the week", alias="offlineScheduleBeginTime")
    offline_schedule_end_enabled: Optional[StrictBool] = Field(default=None, description="When enabled, the BEE will be slated to come online around the end time each weekday (each day if weekends enabled)", alias="offlineScheduleEndEnabled")
    offline_schedule_end_time: Optional[datetime] = Field(default=None, description="Stored with timezone to determine day of the week", alias="offlineScheduleEndTime")
    offline_schedule_end_weekends: Optional[StrictBool] = Field(default=None, alias="offlineScheduleEndWeekends")
    owner: Optional[StrictStr] = Field(default=None, description="When creating, will default to you")
    pact_identifier: Optional[StrictStr] = Field(default=None, alias="pactIdentifier")
    pagerduty_integration: Optional[StrictStr] = Field(default=None, alias="pagerdutyIntegration")
    prevent_deletion: Optional[StrictBool] = Field(default=False, description="Used to protect specific BEEs from deletion (thelma checks this field)", alias="preventDeletion")
    required_role: Optional[StrictStr] = Field(default=None, description="If present, requires membership in the given role for mutations. Set to an empty string to clear.", alias="requiredRole")
    requires_suitability: Optional[StrictBool] = Field(default=None, alias="requiresSuitability")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["baseDomain", "defaultCluster", "deleteAfter", "description", "enableJanitor", "helmfileRef", "namePrefixesDomain", "offline", "offlineScheduleBeginEnabled", "offlineScheduleBeginTime", "offlineScheduleEndEnabled", "offlineScheduleEndTime", "offlineScheduleEndWeekends", "owner", "pactIdentifier", "pagerdutyIntegration", "preventDeletion", "requiredRole", "requiresSuitability"]

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
        """Create an instance of SherlockEnvironmentV3Edit from a JSON string"""
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
        """Create an instance of SherlockEnvironmentV3Edit from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "baseDomain": obj.get("baseDomain") if obj.get("baseDomain") is not None else 'bee.envs-terra.bio',
            "defaultCluster": obj.get("defaultCluster"),
            "deleteAfter": obj.get("deleteAfter"),
            "description": obj.get("description"),
            "enableJanitor": obj.get("enableJanitor"),
            "helmfileRef": obj.get("helmfileRef") if obj.get("helmfileRef") is not None else 'HEAD',
            "namePrefixesDomain": obj.get("namePrefixesDomain") if obj.get("namePrefixesDomain") is not None else True,
            "offline": obj.get("offline") if obj.get("offline") is not None else False,
            "offlineScheduleBeginEnabled": obj.get("offlineScheduleBeginEnabled"),
            "offlineScheduleBeginTime": obj.get("offlineScheduleBeginTime"),
            "offlineScheduleEndEnabled": obj.get("offlineScheduleEndEnabled"),
            "offlineScheduleEndTime": obj.get("offlineScheduleEndTime"),
            "offlineScheduleEndWeekends": obj.get("offlineScheduleEndWeekends"),
            "owner": obj.get("owner"),
            "pactIdentifier": obj.get("pactIdentifier"),
            "pagerdutyIntegration": obj.get("pagerdutyIntegration"),
            "preventDeletion": obj.get("preventDeletion") if obj.get("preventDeletion") is not None else False,
            "requiredRole": obj.get("requiredRole"),
            "requiresSuitability": obj.get("requiresSuitability")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj



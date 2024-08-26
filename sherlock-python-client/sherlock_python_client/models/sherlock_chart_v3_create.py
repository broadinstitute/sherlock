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

from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class SherlockChartV3Create(BaseModel):
    """
    SherlockChartV3Create
    """ # noqa: E501
    app_image_git_main_branch: Optional[StrictStr] = Field(default=None, alias="appImageGitMainBranch")
    app_image_git_repo: Optional[StrictStr] = Field(default=None, alias="appImageGitRepo")
    chart_exposes_endpoint: Optional[StrictBool] = Field(default=False, description="Indicates if the default subdomain, protocol, and port fields are relevant for this chart", alias="chartExposesEndpoint")
    chart_repo: Optional[StrictStr] = Field(default='terra-helm', alias="chartRepo")
    default_port: Optional[StrictInt] = Field(default=None, alias="defaultPort")
    default_protocol: Optional[StrictStr] = Field(default='https', alias="defaultProtocol")
    default_subdomain: Optional[StrictStr] = Field(default=None, description="When creating, will default to the name of the chart", alias="defaultSubdomain")
    description: Optional[StrictStr] = None
    name: Optional[StrictStr] = Field(default=None, description="Required when creating")
    pact_participant: Optional[StrictBool] = Field(default=False, alias="pactParticipant")
    playbook_url: Optional[StrictStr] = Field(default=None, alias="playbookURL")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["appImageGitMainBranch", "appImageGitRepo", "chartExposesEndpoint", "chartRepo", "defaultPort", "defaultProtocol", "defaultSubdomain", "description", "name", "pactParticipant", "playbookURL"]

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
        """Create an instance of SherlockChartV3Create from a JSON string"""
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
        """Create an instance of SherlockChartV3Create from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "appImageGitMainBranch": obj.get("appImageGitMainBranch"),
            "appImageGitRepo": obj.get("appImageGitRepo"),
            "chartExposesEndpoint": obj.get("chartExposesEndpoint") if obj.get("chartExposesEndpoint") is not None else False,
            "chartRepo": obj.get("chartRepo") if obj.get("chartRepo") is not None else 'terra-helm',
            "defaultPort": obj.get("defaultPort"),
            "defaultProtocol": obj.get("defaultProtocol") if obj.get("defaultProtocol") is not None else 'https',
            "defaultSubdomain": obj.get("defaultSubdomain"),
            "description": obj.get("description"),
            "name": obj.get("name"),
            "pactParticipant": obj.get("pactParticipant") if obj.get("pactParticipant") is not None else False,
            "playbookURL": obj.get("playbookURL")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


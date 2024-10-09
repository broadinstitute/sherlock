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
from sherlock_python_client.models.sherlock_role_assignment_v3 import SherlockRoleAssignmentV3
from typing import Optional, Set
from typing_extensions import Self

class SherlockUserV3(BaseModel):
    """
    SherlockUserV3
    """ # noqa: E501
    assignments: Optional[List[SherlockRoleAssignmentV3]] = None
    created_at: Optional[datetime] = Field(default=None, alias="createdAt")
    deactivated_at: Optional[StrictStr] = Field(default=None, description="If set, indicates that the user is currently deactivated", alias="deactivatedAt")
    email: Optional[StrictStr] = None
    github_id: Optional[StrictStr] = Field(default=None, alias="githubID")
    github_username: Optional[StrictStr] = Field(default=None, alias="githubUsername")
    google_id: Optional[StrictStr] = Field(default=None, alias="googleID")
    id: Optional[StrictInt] = None
    name: Optional[StrictStr] = None
    name_from: Optional[StrictStr] = Field(default=None, alias="nameFrom")
    slack_id: Optional[StrictStr] = Field(default=None, alias="slackID")
    slack_username: Optional[StrictStr] = Field(default=None, alias="slackUsername")
    suitability_description: Optional[StrictStr] = Field(default=None, description="Available only in responses; describes the user's production-suitability", alias="suitabilityDescription")
    suitable: Optional[StrictBool] = Field(default=None, description="Available only in responses; indicates whether the user is production-suitable")
    updated_at: Optional[datetime] = Field(default=None, alias="updatedAt")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["assignments", "createdAt", "deactivatedAt", "email", "githubID", "githubUsername", "googleID", "id", "name", "nameFrom", "slackID", "slackUsername", "suitabilityDescription", "suitable", "updatedAt"]

    @field_validator('name_from')
    def name_from_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in set(['sherlock', 'github', 'slack']):
            raise ValueError("must be one of enum values ('sherlock', 'github', 'slack')")
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
        """Create an instance of SherlockUserV3 from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in assignments (list)
        _items = []
        if self.assignments:
            for _item_assignments in self.assignments:
                if _item_assignments:
                    _items.append(_item_assignments.to_dict())
            _dict['assignments'] = _items
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of SherlockUserV3 from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "assignments": [SherlockRoleAssignmentV3.from_dict(_item) for _item in obj["assignments"]] if obj.get("assignments") is not None else None,
            "createdAt": obj.get("createdAt"),
            "deactivatedAt": obj.get("deactivatedAt"),
            "email": obj.get("email"),
            "githubID": obj.get("githubID"),
            "githubUsername": obj.get("githubUsername"),
            "googleID": obj.get("googleID"),
            "id": obj.get("id"),
            "name": obj.get("name"),
            "nameFrom": obj.get("nameFrom"),
            "slackID": obj.get("slackID"),
            "slackUsername": obj.get("slackUsername"),
            "suitabilityDescription": obj.get("suitabilityDescription"),
            "suitable": obj.get("suitable"),
            "updatedAt": obj.get("updatedAt")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj



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

from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class SherlockChangesetV3PlanRequestChartReleaseEntry(BaseModel):
    """
    SherlockChangesetV3PlanRequestChartReleaseEntry
    """ # noqa: E501
    chart_release: Optional[StrictStr] = Field(default=None, alias="chartRelease")
    follow_versions_from_other_chart_release: Optional[StrictStr] = Field(default=None, alias="followVersionsFromOtherChartRelease")
    to_app_version_branch: Optional[StrictStr] = Field(default=None, alias="toAppVersionBranch")
    to_app_version_commit: Optional[StrictStr] = Field(default=None, alias="toAppVersionCommit")
    to_app_version_exact: Optional[StrictStr] = Field(default=None, alias="toAppVersionExact")
    to_app_version_follow_chart_release: Optional[StrictStr] = Field(default=None, alias="toAppVersionFollowChartRelease")
    to_app_version_resolver: Optional[StrictStr] = Field(default=None, alias="toAppVersionResolver")
    to_chart_version_exact: Optional[StrictStr] = Field(default=None, alias="toChartVersionExact")
    to_chart_version_follow_chart_release: Optional[StrictStr] = Field(default=None, alias="toChartVersionFollowChartRelease")
    to_chart_version_resolver: Optional[StrictStr] = Field(default=None, alias="toChartVersionResolver")
    to_helmfile_ref: Optional[StrictStr] = Field(default=None, alias="toHelmfileRef")
    to_helmfile_ref_enabled: Optional[StrictBool] = Field(default=None, alias="toHelmfileRefEnabled")
    use_exact_versions_from_other_chart_release: Optional[StrictStr] = Field(default=None, alias="useExactVersionsFromOtherChartRelease")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["chartRelease", "followVersionsFromOtherChartRelease", "toAppVersionBranch", "toAppVersionCommit", "toAppVersionExact", "toAppVersionFollowChartRelease", "toAppVersionResolver", "toChartVersionExact", "toChartVersionFollowChartRelease", "toChartVersionResolver", "toHelmfileRef", "toHelmfileRefEnabled", "useExactVersionsFromOtherChartRelease"]

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
        """Create an instance of SherlockChangesetV3PlanRequestChartReleaseEntry from a JSON string"""
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
        """Create an instance of SherlockChangesetV3PlanRequestChartReleaseEntry from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "chartRelease": obj.get("chartRelease"),
            "followVersionsFromOtherChartRelease": obj.get("followVersionsFromOtherChartRelease"),
            "toAppVersionBranch": obj.get("toAppVersionBranch"),
            "toAppVersionCommit": obj.get("toAppVersionCommit"),
            "toAppVersionExact": obj.get("toAppVersionExact"),
            "toAppVersionFollowChartRelease": obj.get("toAppVersionFollowChartRelease"),
            "toAppVersionResolver": obj.get("toAppVersionResolver"),
            "toChartVersionExact": obj.get("toChartVersionExact"),
            "toChartVersionFollowChartRelease": obj.get("toChartVersionFollowChartRelease"),
            "toChartVersionResolver": obj.get("toChartVersionResolver"),
            "toHelmfileRef": obj.get("toHelmfileRef"),
            "toHelmfileRefEnabled": obj.get("toHelmfileRefEnabled"),
            "useExactVersionsFromOtherChartRelease": obj.get("useExactVersionsFromOtherChartRelease")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


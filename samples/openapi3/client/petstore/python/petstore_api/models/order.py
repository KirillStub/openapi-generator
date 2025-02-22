# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

    The version of the OpenAPI document: 1.0.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from typing import Any, ClassVar, Dict, List, Optional
from pydantic import BaseModel, StrictBool, StrictInt, StrictStr, field_validator
from pydantic import Field

class Order(BaseModel):
    """
    Order
    """
    id: Optional[StrictInt] = None
    pet_id: Optional[StrictInt] = Field(default=None, alias="petId")
    quantity: Optional[StrictInt] = None
    ship_date: Optional[datetime] = Field(default=None, alias="shipDate")
    status: Optional[StrictStr] = Field(default=None, description="Order Status")
    complete: Optional[StrictBool] = False
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["id", "petId", "quantity", "shipDate", "status", "complete"]

    @field_validator('status')
    def status_validate_enum(cls, value):
        """Validates the enum"""
        if value is None:
            return value

        if value not in ('placed', 'approved', 'delivered'):
            raise ValueError("must be one of enum values ('placed', 'approved', 'delivered')")
        return value

    model_config = {
        "populate_by_name": True,
        "validate_assignment": True
    }


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Order:
        """Create an instance of Order from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.model_dump(by_alias=True,
                          exclude={
                            "additional_properties"
                          },
                          exclude_none=True)
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> Order:
        """Create an instance of Order from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return Order.model_validate(obj)

        _obj = Order.model_validate({
            "id": obj.get("id"),
            "petId": obj.get("petId"),
            "quantity": obj.get("quantity"),
            "shipDate": obj.get("shipDate"),
            "status": obj.get("status"),
            "complete": obj.get("complete") if obj.get("complete") is not None else False
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj



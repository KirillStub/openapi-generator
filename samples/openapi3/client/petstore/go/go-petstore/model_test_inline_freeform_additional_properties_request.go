/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// checks if the TestInlineFreeformAdditionalPropertiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestInlineFreeformAdditionalPropertiesRequest{}

// TestInlineFreeformAdditionalPropertiesRequest struct for TestInlineFreeformAdditionalPropertiesRequest
type TestInlineFreeformAdditionalPropertiesRequest struct {
	SomeProperty *string `json:"someProperty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestInlineFreeformAdditionalPropertiesRequest TestInlineFreeformAdditionalPropertiesRequest

// NewTestInlineFreeformAdditionalPropertiesRequest instantiates a new TestInlineFreeformAdditionalPropertiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestInlineFreeformAdditionalPropertiesRequest() *TestInlineFreeformAdditionalPropertiesRequest {
	this := TestInlineFreeformAdditionalPropertiesRequest{}
	return &this
}

// NewTestInlineFreeformAdditionalPropertiesRequestWithDefaults instantiates a new TestInlineFreeformAdditionalPropertiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestInlineFreeformAdditionalPropertiesRequestWithDefaults() *TestInlineFreeformAdditionalPropertiesRequest {
	this := TestInlineFreeformAdditionalPropertiesRequest{}
	return &this
}

// GetSomeProperty returns the SomeProperty field value if set, zero value otherwise.
func (o *TestInlineFreeformAdditionalPropertiesRequest) GetSomeProperty() string {
	if o == nil || IsNil(o.SomeProperty) {
		var ret string
		return ret
	}
	return *o.SomeProperty
}

// GetSomePropertyOk returns a tuple with the SomeProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInlineFreeformAdditionalPropertiesRequest) GetSomePropertyOk() (*string, bool) {
	if o == nil || IsNil(o.SomeProperty) {
		return nil, false
	}
	return o.SomeProperty, true
}

// HasSomeProperty returns a boolean if a field has been set.
func (o *TestInlineFreeformAdditionalPropertiesRequest) HasSomeProperty() bool {
	if o != nil && !IsNil(o.SomeProperty) {
		return true
	}

	return false
}

// SetSomeProperty gets a reference to the given string and assigns it to the SomeProperty field.
func (o *TestInlineFreeformAdditionalPropertiesRequest) SetSomeProperty(v string) {
	o.SomeProperty = &v
}

func (o TestInlineFreeformAdditionalPropertiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestInlineFreeformAdditionalPropertiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SomeProperty) {
		toSerialize["someProperty"] = o.SomeProperty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestInlineFreeformAdditionalPropertiesRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTestInlineFreeformAdditionalPropertiesRequest := _TestInlineFreeformAdditionalPropertiesRequest{}

	err = json.Unmarshal(bytes, &varTestInlineFreeformAdditionalPropertiesRequest)

	if err != nil {
		return err
	}

	*o = TestInlineFreeformAdditionalPropertiesRequest(varTestInlineFreeformAdditionalPropertiesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "someProperty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestInlineFreeformAdditionalPropertiesRequest struct {
	value *TestInlineFreeformAdditionalPropertiesRequest
	isSet bool
}

func (v NullableTestInlineFreeformAdditionalPropertiesRequest) Get() *TestInlineFreeformAdditionalPropertiesRequest {
	return v.value
}

func (v *NullableTestInlineFreeformAdditionalPropertiesRequest) Set(val *TestInlineFreeformAdditionalPropertiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInlineFreeformAdditionalPropertiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInlineFreeformAdditionalPropertiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInlineFreeformAdditionalPropertiesRequest(val *TestInlineFreeformAdditionalPropertiesRequest) *NullableTestInlineFreeformAdditionalPropertiesRequest {
	return &NullableTestInlineFreeformAdditionalPropertiesRequest{value: val, isSet: true}
}

func (v NullableTestInlineFreeformAdditionalPropertiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInlineFreeformAdditionalPropertiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



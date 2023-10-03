package conekta

import (
	"encoding/json"
)

// checks if the ApiKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyRequest{}

// ApiKeyRequest struct for ApiKeyRequest
type ApiKeyRequest struct {
	// A name or brief explanation of what this api key is used for
	Description *string `json:"description,omitempty"`
	Role        string  `json:"role"`
}

// NewApiKeyRequest instantiates a new ApiKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyRequest(role string) *ApiKeyRequest {
	this := ApiKeyRequest{}
	this.Role = role
	return &this
}

// NewApiKeyRequestWithDefaults instantiates a new ApiKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyRequestWithDefaults() *ApiKeyRequest {
	this := ApiKeyRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiKeyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiKeyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiKeyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRole returns the Role field value
func (o *ApiKeyRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ApiKeyRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ApiKeyRequest) SetRole(v string) {
	o.Role = v
}

func (o ApiKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullableApiKeyRequest struct {
	value *ApiKeyRequest
	isSet bool
}

func (v NullableApiKeyRequest) Get() *ApiKeyRequest {
	return v.value
}

func (v *NullableApiKeyRequest) Set(val *ApiKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyRequest(val *ApiKeyRequest) *NullableApiKeyRequest {
	return &NullableApiKeyRequest{value: val, isSet: true}
}

func (v NullableApiKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

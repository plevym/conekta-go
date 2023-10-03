package conekta

import (
	"encoding/json"
)

// checks if the ApiKeyUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyUpdateRequest{}

// ApiKeyUpdateRequest struct for ApiKeyUpdateRequest
type ApiKeyUpdateRequest struct {
	// Indicates if the webhook key is active
	Active *bool `json:"active,omitempty"`
	// A name or brief explanation of what this api key is used for
	Description *string `json:"description,omitempty"`
}

// NewApiKeyUpdateRequest instantiates a new ApiKeyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyUpdateRequest() *ApiKeyUpdateRequest {
	this := ApiKeyUpdateRequest{}
	return &this
}

// NewApiKeyUpdateRequestWithDefaults instantiates a new ApiKeyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyUpdateRequestWithDefaults() *ApiKeyUpdateRequest {
	this := ApiKeyUpdateRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApiKeyUpdateRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyUpdateRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApiKeyUpdateRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApiKeyUpdateRequest) SetActive(v bool) {
	o.Active = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiKeyUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiKeyUpdateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiKeyUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o ApiKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableApiKeyUpdateRequest struct {
	value *ApiKeyUpdateRequest
	isSet bool
}

func (v NullableApiKeyUpdateRequest) Get() *ApiKeyUpdateRequest {
	return v.value
}

func (v *NullableApiKeyUpdateRequest) Set(val *ApiKeyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyUpdateRequest(val *ApiKeyUpdateRequest) *NullableApiKeyUpdateRequest {
	return &NullableApiKeyUpdateRequest{value: val, isSet: true}
}

func (v NullableApiKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

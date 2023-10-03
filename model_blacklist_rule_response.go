package conekta

import (
	"encoding/json"
)

// checks if the BlacklistRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlacklistRuleResponse{}

// BlacklistRuleResponse struct for BlacklistRuleResponse
type BlacklistRuleResponse struct {
	// Blacklist rule id
	Id *string `json:"id,omitempty"`
	// field used for blacklists rule
	Field *string `json:"field,omitempty"`
	// value used for blacklists rule
	Value *string `json:"value,omitempty"`
	// use an description for blacklisted rule
	Description *string `json:"description,omitempty"`
}

// NewBlacklistRuleResponse instantiates a new BlacklistRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlacklistRuleResponse() *BlacklistRuleResponse {
	this := BlacklistRuleResponse{}
	return &this
}

// NewBlacklistRuleResponseWithDefaults instantiates a new BlacklistRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlacklistRuleResponseWithDefaults() *BlacklistRuleResponse {
	this := BlacklistRuleResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlacklistRuleResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlacklistRuleResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlacklistRuleResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BlacklistRuleResponse) SetId(v string) {
	o.Id = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *BlacklistRuleResponse) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlacklistRuleResponse) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *BlacklistRuleResponse) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *BlacklistRuleResponse) SetField(v string) {
	o.Field = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BlacklistRuleResponse) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlacklistRuleResponse) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BlacklistRuleResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BlacklistRuleResponse) SetValue(v string) {
	o.Value = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BlacklistRuleResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlacklistRuleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BlacklistRuleResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BlacklistRuleResponse) SetDescription(v string) {
	o.Description = &v
}

func (o BlacklistRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlacklistRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableBlacklistRuleResponse struct {
	value *BlacklistRuleResponse
	isSet bool
}

func (v NullableBlacklistRuleResponse) Get() *BlacklistRuleResponse {
	return v.value
}

func (v *NullableBlacklistRuleResponse) Set(val *BlacklistRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBlacklistRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBlacklistRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlacklistRuleResponse(val *BlacklistRuleResponse) *NullableBlacklistRuleResponse {
	return &NullableBlacklistRuleResponse{value: val, isSet: true}
}

func (v NullableBlacklistRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlacklistRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

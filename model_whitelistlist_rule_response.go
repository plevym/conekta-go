package conekta

import (
	"encoding/json"
)

// checks if the WhitelistlistRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WhitelistlistRuleResponse{}

// WhitelistlistRuleResponse struct for WhitelistlistRuleResponse
type WhitelistlistRuleResponse struct {
	// Whitelist rule id
	Id *string `json:"id,omitempty"`
	// field used for whitelists rule
	Field *string `json:"field,omitempty"`
	// value used for whitelists rule
	Value *string `json:"value,omitempty"`
	// use an description for whitelisted rule
	Description *string `json:"description,omitempty"`
}

// NewWhitelistlistRuleResponse instantiates a new WhitelistlistRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhitelistlistRuleResponse() *WhitelistlistRuleResponse {
	this := WhitelistlistRuleResponse{}
	return &this
}

// NewWhitelistlistRuleResponseWithDefaults instantiates a new WhitelistlistRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhitelistlistRuleResponseWithDefaults() *WhitelistlistRuleResponse {
	this := WhitelistlistRuleResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WhitelistlistRuleResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistlistRuleResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WhitelistlistRuleResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WhitelistlistRuleResponse) SetId(v string) {
	o.Id = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *WhitelistlistRuleResponse) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistlistRuleResponse) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *WhitelistlistRuleResponse) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *WhitelistlistRuleResponse) SetField(v string) {
	o.Field = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WhitelistlistRuleResponse) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistlistRuleResponse) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WhitelistlistRuleResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WhitelistlistRuleResponse) SetValue(v string) {
	o.Value = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WhitelistlistRuleResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhitelistlistRuleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WhitelistlistRuleResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WhitelistlistRuleResponse) SetDescription(v string) {
	o.Description = &v
}

func (o WhitelistlistRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WhitelistlistRuleResponse) ToMap() (map[string]interface{}, error) {
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

type NullableWhitelistlistRuleResponse struct {
	value *WhitelistlistRuleResponse
	isSet bool
}

func (v NullableWhitelistlistRuleResponse) Get() *WhitelistlistRuleResponse {
	return v.value
}

func (v *NullableWhitelistlistRuleResponse) Set(val *WhitelistlistRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWhitelistlistRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWhitelistlistRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhitelistlistRuleResponse(val *WhitelistlistRuleResponse) *NullableWhitelistlistRuleResponse {
	return &NullableWhitelistlistRuleResponse{value: val, isSet: true}
}

func (v NullableWhitelistlistRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhitelistlistRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

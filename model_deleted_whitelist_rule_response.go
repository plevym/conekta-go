package conekta

import (
	"encoding/json"
)

// checks if the DeletedWhitelistRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletedWhitelistRuleResponse{}

// DeletedWhitelistRuleResponse struct for DeletedWhitelistRuleResponse
type DeletedWhitelistRuleResponse struct {
	// Whitelist rule id
	Id *string `json:"id,omitempty"`
	// field used for whitelists rule deleted
	Field *string `json:"field,omitempty"`
	// value used for whitelists rule deleted
	Value *string `json:"value,omitempty"`
	// use an description for whitelisted rule
	Description *string `json:"description,omitempty"`
}

// NewDeletedWhitelistRuleResponse instantiates a new DeletedWhitelistRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletedWhitelistRuleResponse() *DeletedWhitelistRuleResponse {
	this := DeletedWhitelistRuleResponse{}
	return &this
}

// NewDeletedWhitelistRuleResponseWithDefaults instantiates a new DeletedWhitelistRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletedWhitelistRuleResponseWithDefaults() *DeletedWhitelistRuleResponse {
	this := DeletedWhitelistRuleResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeletedWhitelistRuleResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedWhitelistRuleResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeletedWhitelistRuleResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeletedWhitelistRuleResponse) SetId(v string) {
	o.Id = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *DeletedWhitelistRuleResponse) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedWhitelistRuleResponse) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *DeletedWhitelistRuleResponse) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *DeletedWhitelistRuleResponse) SetField(v string) {
	o.Field = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeletedWhitelistRuleResponse) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedWhitelistRuleResponse) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeletedWhitelistRuleResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DeletedWhitelistRuleResponse) SetValue(v string) {
	o.Value = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeletedWhitelistRuleResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedWhitelistRuleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeletedWhitelistRuleResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeletedWhitelistRuleResponse) SetDescription(v string) {
	o.Description = &v
}

func (o DeletedWhitelistRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletedWhitelistRuleResponse) ToMap() (map[string]interface{}, error) {
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

type NullableDeletedWhitelistRuleResponse struct {
	value *DeletedWhitelistRuleResponse
	isSet bool
}

func (v NullableDeletedWhitelistRuleResponse) Get() *DeletedWhitelistRuleResponse {
	return v.value
}

func (v *NullableDeletedWhitelistRuleResponse) Set(val *DeletedWhitelistRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletedWhitelistRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletedWhitelistRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletedWhitelistRuleResponse(val *DeletedWhitelistRuleResponse) *NullableDeletedWhitelistRuleResponse {
	return &NullableDeletedWhitelistRuleResponse{value: val, isSet: true}
}

func (v NullableDeletedWhitelistRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletedWhitelistRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

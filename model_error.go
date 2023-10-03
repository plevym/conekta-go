package conekta

import (
	"encoding/json"
)

// checks if the ModelError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelError{}

// ModelError err model
type ModelError struct {
	Details []DetailsError `json:"details,omitempty"`
	// log id
	LogId  NullableString `json:"log_id,omitempty"`
	Type   *string        `json:"type,omitempty"`
	Object *string        `json:"object,omitempty"`
}

// NewModelError instantiates a new ModelError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelError() *ModelError {
	this := ModelError{}
	return &this
}

// NewModelErrorWithDefaults instantiates a new ModelError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelErrorWithDefaults() *ModelError {
	this := ModelError{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ModelError) GetDetails() []DetailsError {
	if o == nil || IsNil(o.Details) {
		var ret []DetailsError
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelError) GetDetailsOk() ([]DetailsError, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ModelError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []DetailsError and assigns it to the Details field.
func (o *ModelError) SetDetails(v []DetailsError) {
	o.Details = v
}

// GetLogId returns the LogId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelError) GetLogId() string {
	if o == nil || IsNil(o.LogId.Get()) {
		var ret string
		return ret
	}
	return *o.LogId.Get()
}

// GetLogIdOk returns a tuple with the LogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelError) GetLogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogId.Get(), o.LogId.IsSet()
}

// HasLogId returns a boolean if a field has been set.
func (o *ModelError) HasLogId() bool {
	if o != nil && o.LogId.IsSet() {
		return true
	}

	return false
}

// SetLogId gets a reference to the given NullableString and assigns it to the LogId field.
func (o *ModelError) SetLogId(v string) {
	o.LogId.Set(&v)
}

// SetLogIdNil sets the value for LogId to be an explicit nil
func (o *ModelError) SetLogIdNil() {
	o.LogId.Set(nil)
}

// UnsetLogId ensures that no value is present for LogId, not even an explicit nil
func (o *ModelError) UnsetLogId() {
	o.LogId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelError) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelError) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelError) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModelError) SetType(v string) {
	o.Type = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ModelError) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelError) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ModelError) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ModelError) SetObject(v string) {
	o.Object = &v
}

func (o ModelError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if o.LogId.IsSet() {
		toSerialize["log_id"] = o.LogId.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableModelError struct {
	value *ModelError
	isSet bool
}

func (v NullableModelError) Get() *ModelError {
	return v.value
}

func (v *NullableModelError) Set(val *ModelError) {
	v.value = val
	v.isSet = true
}

func (v NullableModelError) IsSet() bool {
	return v.isSet
}

func (v *NullableModelError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelError(val *ModelError) *NullableModelError {
	return &NullableModelError{value: val, isSet: true}
}

func (v NullableModelError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

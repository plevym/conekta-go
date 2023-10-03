package conekta

import (
	"encoding/json"
)

// checks if the ErrorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorAllOf{}

// ErrorAllOf struct for ErrorAllOf
type ErrorAllOf struct {
	// log id
	LogId  NullableString `json:"log_id,omitempty"`
	Type   *string        `json:"type,omitempty"`
	Object *string        `json:"object,omitempty"`
}

// NewErrorAllOf instantiates a new ErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorAllOf() *ErrorAllOf {
	this := ErrorAllOf{}
	return &this
}

// NewErrorAllOfWithDefaults instantiates a new ErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorAllOfWithDefaults() *ErrorAllOf {
	this := ErrorAllOf{}
	return &this
}

// GetLogId returns the LogId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorAllOf) GetLogId() string {
	if o == nil || IsNil(o.LogId.Get()) {
		var ret string
		return ret
	}
	return *o.LogId.Get()
}

// GetLogIdOk returns a tuple with the LogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorAllOf) GetLogIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogId.Get(), o.LogId.IsSet()
}

// HasLogId returns a boolean if a field has been set.
func (o *ErrorAllOf) HasLogId() bool {
	if o != nil && o.LogId.IsSet() {
		return true
	}

	return false
}

// SetLogId gets a reference to the given NullableString and assigns it to the LogId field.
func (o *ErrorAllOf) SetLogId(v string) {
	o.LogId.Set(&v)
}

// SetLogIdNil sets the value for LogId to be an explicit nil
func (o *ErrorAllOf) SetLogIdNil() {
	o.LogId.Set(nil)
}

// UnsetLogId ensures that no value is present for LogId, not even an explicit nil
func (o *ErrorAllOf) UnsetLogId() {
	o.LogId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrorAllOf) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorAllOf) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErrorAllOf) SetType(v string) {
	o.Type = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ErrorAllOf) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorAllOf) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ErrorAllOf) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ErrorAllOf) SetObject(v string) {
	o.Object = &v
}

func (o ErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableErrorAllOf struct {
	value *ErrorAllOf
	isSet bool
}

func (v NullableErrorAllOf) Get() *ErrorAllOf {
	return v.value
}

func (v *NullableErrorAllOf) Set(val *ErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorAllOf(val *ErrorAllOf) *NullableErrorAllOf {
	return &NullableErrorAllOf{value: val, isSet: true}
}

func (v NullableErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

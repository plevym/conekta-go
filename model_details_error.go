package conekta

import (
	"encoding/json"
)

// checks if the DetailsError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailsError{}

// DetailsError struct for DetailsError
type DetailsError struct {
	Code         *string        `json:"code,omitempty"`
	Param        NullableString `json:"param,omitempty"`
	Message      *string        `json:"message,omitempty"`
	DebugMessage *string        `json:"debug_message,omitempty"`
}

// NewDetailsError instantiates a new DetailsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailsError() *DetailsError {
	this := DetailsError{}
	return &this
}

// NewDetailsErrorWithDefaults instantiates a new DetailsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailsErrorWithDefaults() *DetailsError {
	this := DetailsError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DetailsError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailsError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DetailsError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DetailsError) SetCode(v string) {
	o.Code = &v
}

// GetParam returns the Param field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailsError) GetParam() string {
	if o == nil || IsNil(o.Param.Get()) {
		var ret string
		return ret
	}
	return *o.Param.Get()
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailsError) GetParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Param.Get(), o.Param.IsSet()
}

// HasParam returns a boolean if a field has been set.
func (o *DetailsError) HasParam() bool {
	if o != nil && o.Param.IsSet() {
		return true
	}

	return false
}

// SetParam gets a reference to the given NullableString and assigns it to the Param field.
func (o *DetailsError) SetParam(v string) {
	o.Param.Set(&v)
}

// SetParamNil sets the value for Param to be an explicit nil
func (o *DetailsError) SetParamNil() {
	o.Param.Set(nil)
}

// UnsetParam ensures that no value is present for Param, not even an explicit nil
func (o *DetailsError) UnsetParam() {
	o.Param.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DetailsError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailsError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DetailsError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DetailsError) SetMessage(v string) {
	o.Message = &v
}

// GetDebugMessage returns the DebugMessage field value if set, zero value otherwise.
func (o *DetailsError) GetDebugMessage() string {
	if o == nil || IsNil(o.DebugMessage) {
		var ret string
		return ret
	}
	return *o.DebugMessage
}

// GetDebugMessageOk returns a tuple with the DebugMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailsError) GetDebugMessageOk() (*string, bool) {
	if o == nil || IsNil(o.DebugMessage) {
		return nil, false
	}
	return o.DebugMessage, true
}

// HasDebugMessage returns a boolean if a field has been set.
func (o *DetailsError) HasDebugMessage() bool {
	if o != nil && !IsNil(o.DebugMessage) {
		return true
	}

	return false
}

// SetDebugMessage gets a reference to the given string and assigns it to the DebugMessage field.
func (o *DetailsError) SetDebugMessage(v string) {
	o.DebugMessage = &v
}

func (o DetailsError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailsError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if o.Param.IsSet() {
		toSerialize["param"] = o.Param.Get()
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.DebugMessage) {
		toSerialize["debug_message"] = o.DebugMessage
	}
	return toSerialize, nil
}

type NullableDetailsError struct {
	value *DetailsError
	isSet bool
}

func (v NullableDetailsError) Get() *DetailsError {
	return v.value
}

func (v *NullableDetailsError) Set(val *DetailsError) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailsError) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailsError(val *DetailsError) *NullableDetailsError {
	return &NullableDetailsError{value: val, isSet: true}
}

func (v NullableDetailsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

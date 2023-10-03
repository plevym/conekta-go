package conekta

import (
	"encoding/json"
)

// checks if the UpdatePaymentMethods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePaymentMethods{}

// UpdatePaymentMethods struct for UpdatePaymentMethods
type UpdatePaymentMethods struct {
	Name *string `json:"name,omitempty"`
}

// NewUpdatePaymentMethods instantiates a new UpdatePaymentMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePaymentMethods() *UpdatePaymentMethods {
	this := UpdatePaymentMethods{}
	return &this
}

// NewUpdatePaymentMethodsWithDefaults instantiates a new UpdatePaymentMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePaymentMethodsWithDefaults() *UpdatePaymentMethods {
	this := UpdatePaymentMethods{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatePaymentMethods) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentMethods) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdatePaymentMethods) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatePaymentMethods) SetName(v string) {
	o.Name = &v
}

func (o UpdatePaymentMethods) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePaymentMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdatePaymentMethods struct {
	value *UpdatePaymentMethods
	isSet bool
}

func (v NullableUpdatePaymentMethods) Get() *UpdatePaymentMethods {
	return v.value
}

func (v *NullableUpdatePaymentMethods) Set(val *UpdatePaymentMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePaymentMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePaymentMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePaymentMethods(val *UpdatePaymentMethods) *NullableUpdatePaymentMethods {
	return &NullableUpdatePaymentMethods{value: val, isSet: true}
}

func (v NullableUpdatePaymentMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePaymentMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

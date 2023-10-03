package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseFiscalEntityAddressAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseFiscalEntityAddressAllOf{}

// OrderResponseFiscalEntityAddressAllOf struct for OrderResponseFiscalEntityAddressAllOf
type OrderResponseFiscalEntityAddressAllOf struct {
	Object *string `json:"object,omitempty"`
}

// NewOrderResponseFiscalEntityAddressAllOf instantiates a new OrderResponseFiscalEntityAddressAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseFiscalEntityAddressAllOf() *OrderResponseFiscalEntityAddressAllOf {
	this := OrderResponseFiscalEntityAddressAllOf{}
	return &this
}

// NewOrderResponseFiscalEntityAddressAllOfWithDefaults instantiates a new OrderResponseFiscalEntityAddressAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseFiscalEntityAddressAllOfWithDefaults() *OrderResponseFiscalEntityAddressAllOf {
	this := OrderResponseFiscalEntityAddressAllOf{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntityAddressAllOf) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddressAllOf) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntityAddressAllOf) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderResponseFiscalEntityAddressAllOf) SetObject(v string) {
	o.Object = &v
}

func (o OrderResponseFiscalEntityAddressAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseFiscalEntityAddressAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableOrderResponseFiscalEntityAddressAllOf struct {
	value *OrderResponseFiscalEntityAddressAllOf
	isSet bool
}

func (v NullableOrderResponseFiscalEntityAddressAllOf) Get() *OrderResponseFiscalEntityAddressAllOf {
	return v.value
}

func (v *NullableOrderResponseFiscalEntityAddressAllOf) Set(val *OrderResponseFiscalEntityAddressAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseFiscalEntityAddressAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseFiscalEntityAddressAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseFiscalEntityAddressAllOf(val *OrderResponseFiscalEntityAddressAllOf) *NullableOrderResponseFiscalEntityAddressAllOf {
	return &NullableOrderResponseFiscalEntityAddressAllOf{value: val, isSet: true}
}

func (v NullableOrderResponseFiscalEntityAddressAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseFiscalEntityAddressAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

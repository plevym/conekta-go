package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseCustomerInfoAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseCustomerInfoAllOf{}

// OrderResponseCustomerInfoAllOf struct for OrderResponseCustomerInfoAllOf
type OrderResponseCustomerInfoAllOf struct {
	Object *string `json:"object,omitempty"`
}

// NewOrderResponseCustomerInfoAllOf instantiates a new OrderResponseCustomerInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseCustomerInfoAllOf() *OrderResponseCustomerInfoAllOf {
	this := OrderResponseCustomerInfoAllOf{}
	return &this
}

// NewOrderResponseCustomerInfoAllOfWithDefaults instantiates a new OrderResponseCustomerInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseCustomerInfoAllOfWithDefaults() *OrderResponseCustomerInfoAllOf {
	this := OrderResponseCustomerInfoAllOf{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderResponseCustomerInfoAllOf) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCustomerInfoAllOf) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderResponseCustomerInfoAllOf) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderResponseCustomerInfoAllOf) SetObject(v string) {
	o.Object = &v
}

func (o OrderResponseCustomerInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseCustomerInfoAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableOrderResponseCustomerInfoAllOf struct {
	value *OrderResponseCustomerInfoAllOf
	isSet bool
}

func (v NullableOrderResponseCustomerInfoAllOf) Get() *OrderResponseCustomerInfoAllOf {
	return v.value
}

func (v *NullableOrderResponseCustomerInfoAllOf) Set(val *OrderResponseCustomerInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseCustomerInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseCustomerInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseCustomerInfoAllOf(val *OrderResponseCustomerInfoAllOf) *NullableOrderResponseCustomerInfoAllOf {
	return &NullableOrderResponseCustomerInfoAllOf{value: val, isSet: true}
}

func (v NullableOrderResponseCustomerInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseCustomerInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseProductsAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseProductsAllOf{}

// OrderResponseProductsAllOf struct for OrderResponseProductsAllOf
type OrderResponseProductsAllOf struct {
	Data []ProductDataResponse `json:"data,omitempty"`
}

// NewOrderResponseProductsAllOf instantiates a new OrderResponseProductsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseProductsAllOf() *OrderResponseProductsAllOf {
	this := OrderResponseProductsAllOf{}
	return &this
}

// NewOrderResponseProductsAllOfWithDefaults instantiates a new OrderResponseProductsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseProductsAllOfWithDefaults() *OrderResponseProductsAllOf {
	this := OrderResponseProductsAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResponseProductsAllOf) GetData() []ProductDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []ProductDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseProductsAllOf) GetDataOk() ([]ProductDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResponseProductsAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ProductDataResponse and assigns it to the Data field.
func (o *OrderResponseProductsAllOf) SetData(v []ProductDataResponse) {
	o.Data = v
}

func (o OrderResponseProductsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseProductsAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderResponseProductsAllOf struct {
	value *OrderResponseProductsAllOf
	isSet bool
}

func (v NullableOrderResponseProductsAllOf) Get() *OrderResponseProductsAllOf {
	return v.value
}

func (v *NullableOrderResponseProductsAllOf) Set(val *OrderResponseProductsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseProductsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseProductsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseProductsAllOf(val *OrderResponseProductsAllOf) *NullableOrderResponseProductsAllOf {
	return &NullableOrderResponseProductsAllOf{value: val, isSet: true}
}

func (v NullableOrderResponseProductsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseProductsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

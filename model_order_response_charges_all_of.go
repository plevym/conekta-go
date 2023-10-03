package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseChargesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseChargesAllOf{}

// OrderResponseChargesAllOf struct for OrderResponseChargesAllOf
type OrderResponseChargesAllOf struct {
	Data []ChargesDataResponse `json:"data,omitempty"`
}

// NewOrderResponseChargesAllOf instantiates a new OrderResponseChargesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseChargesAllOf() *OrderResponseChargesAllOf {
	this := OrderResponseChargesAllOf{}
	return &this
}

// NewOrderResponseChargesAllOfWithDefaults instantiates a new OrderResponseChargesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseChargesAllOfWithDefaults() *OrderResponseChargesAllOf {
	this := OrderResponseChargesAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResponseChargesAllOf) GetData() []ChargesDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []ChargesDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseChargesAllOf) GetDataOk() ([]ChargesDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResponseChargesAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ChargesDataResponse and assigns it to the Data field.
func (o *OrderResponseChargesAllOf) SetData(v []ChargesDataResponse) {
	o.Data = v
}

func (o OrderResponseChargesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseChargesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderResponseChargesAllOf struct {
	value *OrderResponseChargesAllOf
	isSet bool
}

func (v NullableOrderResponseChargesAllOf) Get() *OrderResponseChargesAllOf {
	return v.value
}

func (v *NullableOrderResponseChargesAllOf) Set(val *OrderResponseChargesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseChargesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseChargesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseChargesAllOf(val *OrderResponseChargesAllOf) *NullableOrderResponseChargesAllOf {
	return &NullableOrderResponseChargesAllOf{value: val, isSet: true}
}

func (v NullableOrderResponseChargesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseChargesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseDiscountLinesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseDiscountLinesAllOf{}

// OrderResponseDiscountLinesAllOf struct for OrderResponseDiscountLinesAllOf
type OrderResponseDiscountLinesAllOf struct {
	Data []DiscountLinesDataResponse `json:"data,omitempty"`
}

// NewOrderResponseDiscountLinesAllOf instantiates a new OrderResponseDiscountLinesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseDiscountLinesAllOf() *OrderResponseDiscountLinesAllOf {
	this := OrderResponseDiscountLinesAllOf{}
	return &this
}

// NewOrderResponseDiscountLinesAllOfWithDefaults instantiates a new OrderResponseDiscountLinesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseDiscountLinesAllOfWithDefaults() *OrderResponseDiscountLinesAllOf {
	this := OrderResponseDiscountLinesAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResponseDiscountLinesAllOf) GetData() []DiscountLinesDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []DiscountLinesDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseDiscountLinesAllOf) GetDataOk() ([]DiscountLinesDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResponseDiscountLinesAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DiscountLinesDataResponse and assigns it to the Data field.
func (o *OrderResponseDiscountLinesAllOf) SetData(v []DiscountLinesDataResponse) {
	o.Data = v
}

func (o OrderResponseDiscountLinesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseDiscountLinesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderResponseDiscountLinesAllOf struct {
	value *OrderResponseDiscountLinesAllOf
	isSet bool
}

func (v NullableOrderResponseDiscountLinesAllOf) Get() *OrderResponseDiscountLinesAllOf {
	return v.value
}

func (v *NullableOrderResponseDiscountLinesAllOf) Set(val *OrderResponseDiscountLinesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseDiscountLinesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseDiscountLinesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseDiscountLinesAllOf(val *OrderResponseDiscountLinesAllOf) *NullableOrderResponseDiscountLinesAllOf {
	return &NullableOrderResponseDiscountLinesAllOf{value: val, isSet: true}
}

func (v NullableOrderResponseDiscountLinesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseDiscountLinesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

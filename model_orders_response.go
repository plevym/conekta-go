package conekta

import (
	"encoding/json"
)

// checks if the OrdersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrdersResponse{}

// OrdersResponse struct for OrdersResponse
type OrdersResponse struct {
	Data []OrderResponse `json:"data"`
}

// NewOrdersResponse instantiates a new OrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersResponse(data []OrderResponse) *OrdersResponse {
	this := OrdersResponse{}
	this.Data = data
	return &this
}

// NewOrdersResponseWithDefaults instantiates a new OrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersResponseWithDefaults() *OrdersResponse {
	this := OrdersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *OrdersResponse) GetData() []OrderResponse {
	if o == nil {
		var ret []OrderResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrdersResponse) GetDataOk() ([]OrderResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *OrdersResponse) SetData(v []OrderResponse) {
	o.Data = v
}

func (o OrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableOrdersResponse struct {
	value *OrdersResponse
	isSet bool
}

func (v NullableOrdersResponse) Get() *OrdersResponse {
	return v.value
}

func (v *NullableOrdersResponse) Set(val *OrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersResponse(val *OrdersResponse) *NullableOrdersResponse {
	return &NullableOrdersResponse{value: val, isSet: true}
}

func (v NullableOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

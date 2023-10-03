package conekta

import (
	"encoding/json"
)

// checks if the OrderCaptureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCaptureRequest{}

// OrderCaptureRequest struct for OrderCaptureRequest
type OrderCaptureRequest struct {
	// Amount to capture
	Amount int64 `json:"amount"`
}

// NewOrderCaptureRequest instantiates a new OrderCaptureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCaptureRequest(amount int64) *OrderCaptureRequest {
	this := OrderCaptureRequest{}
	this.Amount = amount
	return &this
}

// NewOrderCaptureRequestWithDefaults instantiates a new OrderCaptureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCaptureRequestWithDefaults() *OrderCaptureRequest {
	this := OrderCaptureRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *OrderCaptureRequest) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OrderCaptureRequest) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OrderCaptureRequest) SetAmount(v int64) {
	o.Amount = v
}

func (o OrderCaptureRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCaptureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

type NullableOrderCaptureRequest struct {
	value *OrderCaptureRequest
	isSet bool
}

func (v NullableOrderCaptureRequest) Get() *OrderCaptureRequest {
	return v.value
}

func (v *NullableOrderCaptureRequest) Set(val *OrderCaptureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCaptureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCaptureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCaptureRequest(val *OrderCaptureRequest) *NullableOrderCaptureRequest {
	return &NullableOrderCaptureRequest{value: val, isSet: true}
}

func (v NullableOrderCaptureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCaptureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

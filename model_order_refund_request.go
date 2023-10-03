package conekta

import (
	"encoding/json"
)

// checks if the OrderRefundRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderRefundRequest{}

// OrderRefundRequest struct for OrderRefundRequest
type OrderRefundRequest struct {
	Amount    int32         `json:"amount"`
	ExpiresAt NullableInt64 `json:"expires_at,omitempty"`
	Reason    string        `json:"reason"`
}

// NewOrderRefundRequest instantiates a new OrderRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRefundRequest(amount int32, reason string) *OrderRefundRequest {
	this := OrderRefundRequest{}
	this.Amount = amount
	this.Reason = reason
	return &this
}

// NewOrderRefundRequestWithDefaults instantiates a new OrderRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRefundRequestWithDefaults() *OrderRefundRequest {
	this := OrderRefundRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *OrderRefundRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OrderRefundRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OrderRefundRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderRefundRequest) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderRefundRequest) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OrderRefundRequest) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableInt64 and assigns it to the ExpiresAt field.
func (o *OrderRefundRequest) SetExpiresAt(v int64) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *OrderRefundRequest) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *OrderRefundRequest) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetReason returns the Reason field value
func (o *OrderRefundRequest) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *OrderRefundRequest) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *OrderRefundRequest) SetReason(v string) {
	o.Reason = v
}

func (o OrderRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

type NullableOrderRefundRequest struct {
	value *OrderRefundRequest
	isSet bool
}

func (v NullableOrderRefundRequest) Get() *OrderRefundRequest {
	return v.value
}

func (v *NullableOrderRefundRequest) Set(val *OrderRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRefundRequest(val *OrderRefundRequest) *NullableOrderRefundRequest {
	return &NullableOrderRefundRequest{value: val, isSet: true}
}

func (v NullableOrderRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

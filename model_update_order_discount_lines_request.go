package conekta

import (
	"encoding/json"
)

// checks if the UpdateOrderDiscountLinesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrderDiscountLinesRequest{}

// UpdateOrderDiscountLinesRequest List of discounts that apply to the order.
type UpdateOrderDiscountLinesRequest struct {
	Amount *int64 `json:"amount,omitempty"`
	// Discount code.
	Code *string `json:"code,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewUpdateOrderDiscountLinesRequest instantiates a new UpdateOrderDiscountLinesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrderDiscountLinesRequest() *UpdateOrderDiscountLinesRequest {
	this := UpdateOrderDiscountLinesRequest{}
	return &this
}

// NewUpdateOrderDiscountLinesRequestWithDefaults instantiates a new UpdateOrderDiscountLinesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrderDiscountLinesRequestWithDefaults() *UpdateOrderDiscountLinesRequest {
	this := UpdateOrderDiscountLinesRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UpdateOrderDiscountLinesRequest) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderDiscountLinesRequest) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UpdateOrderDiscountLinesRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UpdateOrderDiscountLinesRequest) SetAmount(v int64) {
	o.Amount = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateOrderDiscountLinesRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderDiscountLinesRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateOrderDiscountLinesRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateOrderDiscountLinesRequest) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateOrderDiscountLinesRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderDiscountLinesRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateOrderDiscountLinesRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateOrderDiscountLinesRequest) SetType(v string) {
	o.Type = &v
}

func (o UpdateOrderDiscountLinesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrderDiscountLinesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUpdateOrderDiscountLinesRequest struct {
	value *UpdateOrderDiscountLinesRequest
	isSet bool
}

func (v NullableUpdateOrderDiscountLinesRequest) Get() *UpdateOrderDiscountLinesRequest {
	return v.value
}

func (v *NullableUpdateOrderDiscountLinesRequest) Set(val *UpdateOrderDiscountLinesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrderDiscountLinesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrderDiscountLinesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrderDiscountLinesRequest(val *UpdateOrderDiscountLinesRequest) *NullableUpdateOrderDiscountLinesRequest {
	return &NullableUpdateOrderDiscountLinesRequest{value: val, isSet: true}
}

func (v NullableUpdateOrderDiscountLinesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrderDiscountLinesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

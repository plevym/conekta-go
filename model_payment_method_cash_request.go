package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCashRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCashRequest{}

// PaymentMethodCashRequest struct for PaymentMethodCashRequest
type PaymentMethodCashRequest struct {
	// Type of payment method
	Type      string `json:"type"`
	ExpiresAt *int64 `json:"expires_at,omitempty"`
}

// NewPaymentMethodCashRequest instantiates a new PaymentMethodCashRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCashRequest(type_ string) *PaymentMethodCashRequest {
	this := PaymentMethodCashRequest{}
	this.Type = type_
	return &this
}

// NewPaymentMethodCashRequestWithDefaults instantiates a new PaymentMethodCashRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCashRequestWithDefaults() *PaymentMethodCashRequest {
	this := PaymentMethodCashRequest{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentMethodCashRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodCashRequest) SetType(v string) {
	o.Type = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PaymentMethodCashRequest) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashRequest) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentMethodCashRequest) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *PaymentMethodCashRequest) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

func (o PaymentMethodCashRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCashRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullablePaymentMethodCashRequest struct {
	value *PaymentMethodCashRequest
	isSet bool
}

func (v NullablePaymentMethodCashRequest) Get() *PaymentMethodCashRequest {
	return v.value
}

func (v *NullablePaymentMethodCashRequest) Set(val *PaymentMethodCashRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCashRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCashRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCashRequest(val *PaymentMethodCashRequest) *NullablePaymentMethodCashRequest {
	return &NullablePaymentMethodCashRequest{value: val, isSet: true}
}

func (v NullablePaymentMethodCashRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCashRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

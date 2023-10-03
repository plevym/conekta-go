package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCashRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCashRequestAllOf{}

// PaymentMethodCashRequestAllOf struct for PaymentMethodCashRequestAllOf
type PaymentMethodCashRequestAllOf struct {
	ExpiresAt *int64 `json:"expires_at,omitempty"`
}

// NewPaymentMethodCashRequestAllOf instantiates a new PaymentMethodCashRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCashRequestAllOf() *PaymentMethodCashRequestAllOf {
	this := PaymentMethodCashRequestAllOf{}
	return &this
}

// NewPaymentMethodCashRequestAllOfWithDefaults instantiates a new PaymentMethodCashRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCashRequestAllOfWithDefaults() *PaymentMethodCashRequestAllOf {
	this := PaymentMethodCashRequestAllOf{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PaymentMethodCashRequestAllOf) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashRequestAllOf) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentMethodCashRequestAllOf) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *PaymentMethodCashRequestAllOf) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

func (o PaymentMethodCashRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCashRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullablePaymentMethodCashRequestAllOf struct {
	value *PaymentMethodCashRequestAllOf
	isSet bool
}

func (v NullablePaymentMethodCashRequestAllOf) Get() *PaymentMethodCashRequestAllOf {
	return v.value
}

func (v *NullablePaymentMethodCashRequestAllOf) Set(val *PaymentMethodCashRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCashRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCashRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCashRequestAllOf(val *PaymentMethodCashRequestAllOf) *NullablePaymentMethodCashRequestAllOf {
	return &NullablePaymentMethodCashRequestAllOf{value: val, isSet: true}
}

func (v NullablePaymentMethodCashRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCashRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

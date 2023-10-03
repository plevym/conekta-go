package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCardRequestAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCardRequestAllOf{}

// PaymentMethodCardRequestAllOf struct for PaymentMethodCardRequestAllOf
type PaymentMethodCardRequestAllOf struct {
	// Token id that will be used to create a \"card\" type payment method. See the (subscriptions)[https://developers.conekta.com/v2.1.0/reference/createsubscription] tutorial for more information on how to tokenize cards.
	TokenId *string `json:"token_id,omitempty"`
}

// NewPaymentMethodCardRequestAllOf instantiates a new PaymentMethodCardRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCardRequestAllOf() *PaymentMethodCardRequestAllOf {
	this := PaymentMethodCardRequestAllOf{}
	return &this
}

// NewPaymentMethodCardRequestAllOfWithDefaults instantiates a new PaymentMethodCardRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCardRequestAllOfWithDefaults() *PaymentMethodCardRequestAllOf {
	this := PaymentMethodCardRequestAllOf{}
	return &this
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *PaymentMethodCardRequestAllOf) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardRequestAllOf) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *PaymentMethodCardRequestAllOf) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *PaymentMethodCardRequestAllOf) SetTokenId(v string) {
	o.TokenId = &v
}

func (o PaymentMethodCardRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCardRequestAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	return toSerialize, nil
}

type NullablePaymentMethodCardRequestAllOf struct {
	value *PaymentMethodCardRequestAllOf
	isSet bool
}

func (v NullablePaymentMethodCardRequestAllOf) Get() *PaymentMethodCardRequestAllOf {
	return v.value
}

func (v *NullablePaymentMethodCardRequestAllOf) Set(val *PaymentMethodCardRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCardRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCardRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCardRequestAllOf(val *PaymentMethodCardRequestAllOf) *NullablePaymentMethodCardRequestAllOf {
	return &NullablePaymentMethodCardRequestAllOf{value: val, isSet: true}
}

func (v NullablePaymentMethodCardRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCardRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

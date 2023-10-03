package conekta

import (
	"encoding/json"
)

// checks if the Token type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Token{}

// Token a token
type Token struct {
	Card NullableTokenCard `json:"card,omitempty"`
	// Deprecated
	Checkout NullableTokenCheckout `json:"checkout,omitempty"`
}

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken() *Token {
	this := Token{}
	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	return &this
}

// GetCard returns the Card field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Token) GetCard() TokenCard {
	if o == nil || IsNil(o.Card.Get()) {
		var ret TokenCard
		return ret
	}
	return *o.Card.Get()
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetCardOk() (*TokenCard, bool) {
	if o == nil {
		return nil, false
	}
	return o.Card.Get(), o.Card.IsSet()
}

// HasCard returns a boolean if a field has been set.
func (o *Token) HasCard() bool {
	if o != nil && o.Card.IsSet() {
		return true
	}

	return false
}

// SetCard gets a reference to the given NullableTokenCard and assigns it to the Card field.
func (o *Token) SetCard(v TokenCard) {
	o.Card.Set(&v)
}

// SetCardNil sets the value for Card to be an explicit nil
func (o *Token) SetCardNil() {
	o.Card.Set(nil)
}

// UnsetCard ensures that no value is present for Card, not even an explicit nil
func (o *Token) UnsetCard() {
	o.Card.Unset()
}

// GetCheckout returns the Checkout field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *Token) GetCheckout() TokenCheckout {
	if o == nil || IsNil(o.Checkout.Get()) {
		var ret TokenCheckout
		return ret
	}
	return *o.Checkout.Get()
}

// GetCheckoutOk returns a tuple with the Checkout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *Token) GetCheckoutOk() (*TokenCheckout, bool) {
	if o == nil {
		return nil, false
	}
	return o.Checkout.Get(), o.Checkout.IsSet()
}

// HasCheckout returns a boolean if a field has been set.
func (o *Token) HasCheckout() bool {
	if o != nil && o.Checkout.IsSet() {
		return true
	}

	return false
}

// SetCheckout gets a reference to the given NullableTokenCheckout and assigns it to the Checkout field.
// Deprecated
func (o *Token) SetCheckout(v TokenCheckout) {
	o.Checkout.Set(&v)
}

// SetCheckoutNil sets the value for Checkout to be an explicit nil
func (o *Token) SetCheckoutNil() {
	o.Checkout.Set(nil)
}

// UnsetCheckout ensures that no value is present for Checkout, not even an explicit nil
func (o *Token) UnsetCheckout() {
	o.Checkout.Unset()
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Token) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Card.IsSet() {
		toSerialize["card"] = o.Card.Get()
	}
	if o.Checkout.IsSet() {
		toSerialize["checkout"] = o.Checkout.Get()
	}
	return toSerialize, nil
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

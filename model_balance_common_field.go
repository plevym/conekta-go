package conekta

import (
	"encoding/json"
)

// checks if the BalanceCommonField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BalanceCommonField{}

// BalanceCommonField balance common fields model
type BalanceCommonField struct {
	// The balance's amount
	Amount *int64 `json:"amount,omitempty"`
	// The balance's currency
	Currency *string `json:"currency,omitempty"`
}

// NewBalanceCommonField instantiates a new BalanceCommonField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceCommonField() *BalanceCommonField {
	this := BalanceCommonField{}
	return &this
}

// NewBalanceCommonFieldWithDefaults instantiates a new BalanceCommonField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceCommonFieldWithDefaults() *BalanceCommonField {
	this := BalanceCommonField{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BalanceCommonField) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceCommonField) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BalanceCommonField) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *BalanceCommonField) SetAmount(v int64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BalanceCommonField) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceCommonField) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BalanceCommonField) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BalanceCommonField) SetCurrency(v string) {
	o.Currency = &v
}

func (o BalanceCommonField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceCommonField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableBalanceCommonField struct {
	value *BalanceCommonField
	isSet bool
}

func (v NullableBalanceCommonField) Get() *BalanceCommonField {
	return v.value
}

func (v *NullableBalanceCommonField) Set(val *BalanceCommonField) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceCommonField) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceCommonField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceCommonField(val *BalanceCommonField) *NullableBalanceCommonField {
	return &NullableBalanceCommonField{value: val, isSet: true}
}

func (v NullableBalanceCommonField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceCommonField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

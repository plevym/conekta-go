package conekta

import (
	"encoding/json"
)

// checks if the TokenCheckout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenCheckout{}

// TokenCheckout struct for TokenCheckout
type TokenCheckout struct {
	// It is a value that allows identifying the returns control on.
	ReturnsControlOn *string `json:"returns_control_on,omitempty"`
}

// NewTokenCheckout instantiates a new TokenCheckout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCheckout() *TokenCheckout {
	this := TokenCheckout{}
	return &this
}

// NewTokenCheckoutWithDefaults instantiates a new TokenCheckout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCheckoutWithDefaults() *TokenCheckout {
	this := TokenCheckout{}
	return &this
}

// GetReturnsControlOn returns the ReturnsControlOn field value if set, zero value otherwise.
func (o *TokenCheckout) GetReturnsControlOn() string {
	if o == nil || IsNil(o.ReturnsControlOn) {
		var ret string
		return ret
	}
	return *o.ReturnsControlOn
}

// GetReturnsControlOnOk returns a tuple with the ReturnsControlOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCheckout) GetReturnsControlOnOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnsControlOn) {
		return nil, false
	}
	return o.ReturnsControlOn, true
}

// HasReturnsControlOn returns a boolean if a field has been set.
func (o *TokenCheckout) HasReturnsControlOn() bool {
	if o != nil && !IsNil(o.ReturnsControlOn) {
		return true
	}

	return false
}

// SetReturnsControlOn gets a reference to the given string and assigns it to the ReturnsControlOn field.
func (o *TokenCheckout) SetReturnsControlOn(v string) {
	o.ReturnsControlOn = &v
}

func (o TokenCheckout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenCheckout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnsControlOn) {
		toSerialize["returns_control_on"] = o.ReturnsControlOn
	}
	return toSerialize, nil
}

type NullableTokenCheckout struct {
	value *TokenCheckout
	isSet bool
}

func (v NullableTokenCheckout) Get() *TokenCheckout {
	return v.value
}

func (v *NullableTokenCheckout) Set(val *TokenCheckout) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenCheckout) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenCheckout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenCheckout(val *TokenCheckout) *NullableTokenCheckout {
	return &NullableTokenCheckout{value: val, isSet: true}
}

func (v NullableTokenCheckout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenCheckout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

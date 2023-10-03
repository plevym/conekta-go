package conekta

import (
	"encoding/json"
)

// checks if the TokenCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenCard{}

// TokenCard struct for TokenCard
type TokenCard struct {
	// It is a value that allows identifying the security code of the card.
	Cvc string `json:"cvc"`
	// It is a value that allows identifying the device fingerprint.
	DeviceFingerprint *string `json:"device_fingerprint,omitempty"`
	// It is a value that allows identifying the expiration month of the card.
	ExpMonth string `json:"exp_month"`
	// It is a value that allows identifying the expiration year of the card.
	ExpYear string `json:"exp_year"`
	// It is a value that allows identifying the name of the cardholder.
	Name string `json:"name"`
	// It is a value that allows identifying the number of the card.
	Number string `json:"number"`
}

// NewTokenCard instantiates a new TokenCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCard(cvc string, expMonth string, expYear string, name string, number string) *TokenCard {
	this := TokenCard{}
	this.Cvc = cvc
	this.ExpMonth = expMonth
	this.ExpYear = expYear
	this.Name = name
	this.Number = number
	return &this
}

// NewTokenCardWithDefaults instantiates a new TokenCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCardWithDefaults() *TokenCard {
	this := TokenCard{}
	return &this
}

// GetCvc returns the Cvc field value
func (o *TokenCard) GetCvc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cvc
}

// GetCvcOk returns a tuple with the Cvc field value
// and a boolean to check if the value has been set.
func (o *TokenCard) GetCvcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cvc, true
}

// SetCvc sets field value
func (o *TokenCard) SetCvc(v string) {
	o.Cvc = v
}

// GetDeviceFingerprint returns the DeviceFingerprint field value if set, zero value otherwise.
func (o *TokenCard) GetDeviceFingerprint() string {
	if o == nil || IsNil(o.DeviceFingerprint) {
		var ret string
		return ret
	}
	return *o.DeviceFingerprint
}

// GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCard) GetDeviceFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceFingerprint) {
		return nil, false
	}
	return o.DeviceFingerprint, true
}

// HasDeviceFingerprint returns a boolean if a field has been set.
func (o *TokenCard) HasDeviceFingerprint() bool {
	if o != nil && !IsNil(o.DeviceFingerprint) {
		return true
	}

	return false
}

// SetDeviceFingerprint gets a reference to the given string and assigns it to the DeviceFingerprint field.
func (o *TokenCard) SetDeviceFingerprint(v string) {
	o.DeviceFingerprint = &v
}

// GetExpMonth returns the ExpMonth field value
func (o *TokenCard) GetExpMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpMonth
}

// GetExpMonthOk returns a tuple with the ExpMonth field value
// and a boolean to check if the value has been set.
func (o *TokenCard) GetExpMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpMonth, true
}

// SetExpMonth sets field value
func (o *TokenCard) SetExpMonth(v string) {
	o.ExpMonth = v
}

// GetExpYear returns the ExpYear field value
func (o *TokenCard) GetExpYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpYear
}

// GetExpYearOk returns a tuple with the ExpYear field value
// and a boolean to check if the value has been set.
func (o *TokenCard) GetExpYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpYear, true
}

// SetExpYear sets field value
func (o *TokenCard) SetExpYear(v string) {
	o.ExpYear = v
}

// GetName returns the Name field value
func (o *TokenCard) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenCard) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenCard) SetName(v string) {
	o.Name = v
}

// GetNumber returns the Number field value
func (o *TokenCard) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *TokenCard) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *TokenCard) SetNumber(v string) {
	o.Number = v
}

func (o TokenCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cvc"] = o.Cvc
	if !IsNil(o.DeviceFingerprint) {
		toSerialize["device_fingerprint"] = o.DeviceFingerprint
	}
	toSerialize["exp_month"] = o.ExpMonth
	toSerialize["exp_year"] = o.ExpYear
	toSerialize["name"] = o.Name
	toSerialize["number"] = o.Number
	return toSerialize, nil
}

type NullableTokenCard struct {
	value *TokenCard
	isSet bool
}

func (v NullableTokenCard) Get() *TokenCard {
	return v.value
}

func (v *NullableTokenCard) Set(val *TokenCard) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenCard) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenCard(val *TokenCard) *NullableTokenCard {
	return &NullableTokenCard{value: val, isSet: true}
}

func (v NullableTokenCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

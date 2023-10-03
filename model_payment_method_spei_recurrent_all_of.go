package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodSpeiRecurrentAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodSpeiRecurrentAllOf{}

// PaymentMethodSpeiRecurrentAllOf use for spei responses
type PaymentMethodSpeiRecurrentAllOf struct {
	Reference *string `json:"reference,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
}

// NewPaymentMethodSpeiRecurrentAllOf instantiates a new PaymentMethodSpeiRecurrentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodSpeiRecurrentAllOf() *PaymentMethodSpeiRecurrentAllOf {
	this := PaymentMethodSpeiRecurrentAllOf{}
	return &this
}

// NewPaymentMethodSpeiRecurrentAllOfWithDefaults instantiates a new PaymentMethodSpeiRecurrentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodSpeiRecurrentAllOfWithDefaults() *PaymentMethodSpeiRecurrentAllOf {
	this := PaymentMethodSpeiRecurrentAllOf{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentMethodSpeiRecurrentAllOf) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSpeiRecurrentAllOf) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentMethodSpeiRecurrentAllOf) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentMethodSpeiRecurrentAllOf) SetReference(v string) {
	o.Reference = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PaymentMethodSpeiRecurrentAllOf) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSpeiRecurrentAllOf) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentMethodSpeiRecurrentAllOf) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *PaymentMethodSpeiRecurrentAllOf) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o PaymentMethodSpeiRecurrentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodSpeiRecurrentAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullablePaymentMethodSpeiRecurrentAllOf struct {
	value *PaymentMethodSpeiRecurrentAllOf
	isSet bool
}

func (v NullablePaymentMethodSpeiRecurrentAllOf) Get() *PaymentMethodSpeiRecurrentAllOf {
	return v.value
}

func (v *NullablePaymentMethodSpeiRecurrentAllOf) Set(val *PaymentMethodSpeiRecurrentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodSpeiRecurrentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodSpeiRecurrentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodSpeiRecurrentAllOf(val *PaymentMethodSpeiRecurrentAllOf) *NullablePaymentMethodSpeiRecurrentAllOf {
	return &NullablePaymentMethodSpeiRecurrentAllOf{value: val, isSet: true}
}

func (v NullablePaymentMethodSpeiRecurrentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodSpeiRecurrentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

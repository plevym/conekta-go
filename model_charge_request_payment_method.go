package conekta

import (
	"encoding/json"
)

// checks if the ChargeRequestPaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeRequestPaymentMethod{}

// ChargeRequestPaymentMethod Payment method used in the charge. Go to the [payment methods](https://developers.conekta.com/reference/m%C3%A9todos-de-pago) section for more details
type ChargeRequestPaymentMethod struct {
	// Method expiration date as unix timestamp
	ExpiresAt       *int64  `json:"expires_at,omitempty"`
	Type            string  `json:"type"`
	TokenId         *string `json:"token_id,omitempty"`
	PaymentSourceId *string `json:"payment_source_id,omitempty"`
	// Optional id sent to indicate the bank contract for recurrent card charges.
	ContractId *string `json:"contract_id,omitempty"`
}

// NewChargeRequestPaymentMethod instantiates a new ChargeRequestPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeRequestPaymentMethod(type_ string) *ChargeRequestPaymentMethod {
	this := ChargeRequestPaymentMethod{}
	this.Type = type_
	return &this
}

// NewChargeRequestPaymentMethodWithDefaults instantiates a new ChargeRequestPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeRequestPaymentMethodWithDefaults() *ChargeRequestPaymentMethod {
	this := ChargeRequestPaymentMethod{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ChargeRequestPaymentMethod) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRequestPaymentMethod) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ChargeRequestPaymentMethod) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *ChargeRequestPaymentMethod) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetType returns the Type field value
func (o *ChargeRequestPaymentMethod) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChargeRequestPaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChargeRequestPaymentMethod) SetType(v string) {
	o.Type = v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ChargeRequestPaymentMethod) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRequestPaymentMethod) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ChargeRequestPaymentMethod) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *ChargeRequestPaymentMethod) SetTokenId(v string) {
	o.TokenId = &v
}

// GetPaymentSourceId returns the PaymentSourceId field value if set, zero value otherwise.
func (o *ChargeRequestPaymentMethod) GetPaymentSourceId() string {
	if o == nil || IsNil(o.PaymentSourceId) {
		var ret string
		return ret
	}
	return *o.PaymentSourceId
}

// GetPaymentSourceIdOk returns a tuple with the PaymentSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRequestPaymentMethod) GetPaymentSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentSourceId) {
		return nil, false
	}
	return o.PaymentSourceId, true
}

// HasPaymentSourceId returns a boolean if a field has been set.
func (o *ChargeRequestPaymentMethod) HasPaymentSourceId() bool {
	if o != nil && !IsNil(o.PaymentSourceId) {
		return true
	}

	return false
}

// SetPaymentSourceId gets a reference to the given string and assigns it to the PaymentSourceId field.
func (o *ChargeRequestPaymentMethod) SetPaymentSourceId(v string) {
	o.PaymentSourceId = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *ChargeRequestPaymentMethod) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRequestPaymentMethod) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *ChargeRequestPaymentMethod) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *ChargeRequestPaymentMethod) SetContractId(v string) {
	o.ContractId = &v
}

func (o ChargeRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeRequestPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.PaymentSourceId) {
		toSerialize["payment_source_id"] = o.PaymentSourceId
	}
	if !IsNil(o.ContractId) {
		toSerialize["contract_id"] = o.ContractId
	}
	return toSerialize, nil
}

type NullableChargeRequestPaymentMethod struct {
	value *ChargeRequestPaymentMethod
	isSet bool
}

func (v NullableChargeRequestPaymentMethod) Get() *ChargeRequestPaymentMethod {
	return v.value
}

func (v *NullableChargeRequestPaymentMethod) Set(val *ChargeRequestPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeRequestPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeRequestPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeRequestPaymentMethod(val *ChargeRequestPaymentMethod) *NullableChargeRequestPaymentMethod {
	return &NullableChargeRequestPaymentMethod{value: val, isSet: true}
}

func (v NullableChargeRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeRequestPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

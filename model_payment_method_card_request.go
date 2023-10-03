package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCardRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCardRequest{}

// PaymentMethodCardRequest struct for PaymentMethodCardRequest
type PaymentMethodCardRequest struct {
	// Type of payment method
	Type string `json:"type"`
	// Token id that will be used to create a \"card\" type payment method. See the (subscriptions)[https://developers.conekta.com/v2.1.0/reference/createsubscription] tutorial for more information on how to tokenize cards.
	TokenId string `json:"token_id"`
}

// NewPaymentMethodCardRequest instantiates a new PaymentMethodCardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCardRequest(type_ string, tokenId string) *PaymentMethodCardRequest {
	this := PaymentMethodCardRequest{}
	this.Type = type_
	this.TokenId = tokenId
	return &this
}

// NewPaymentMethodCardRequestWithDefaults instantiates a new PaymentMethodCardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCardRequestWithDefaults() *PaymentMethodCardRequest {
	this := PaymentMethodCardRequest{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentMethodCardRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodCardRequest) SetType(v string) {
	o.Type = v
}

// GetTokenId returns the TokenId field value
func (o *PaymentMethodCardRequest) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCardRequest) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *PaymentMethodCardRequest) SetTokenId(v string) {
	o.TokenId = v
}

func (o PaymentMethodCardRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCardRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["token_id"] = o.TokenId
	return toSerialize, nil
}

type NullablePaymentMethodCardRequest struct {
	value *PaymentMethodCardRequest
	isSet bool
}

func (v NullablePaymentMethodCardRequest) Get() *PaymentMethodCardRequest {
	return v.value
}

func (v *NullablePaymentMethodCardRequest) Set(val *PaymentMethodCardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCardRequest(val *PaymentMethodCardRequest) *NullablePaymentMethodCardRequest {
	return &NullablePaymentMethodCardRequest{value: val, isSet: true}
}

func (v NullablePaymentMethodCardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

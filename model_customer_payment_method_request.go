package conekta

import (
	"encoding/json"
)

// checks if the CustomerPaymentMethodRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPaymentMethodRequest{}

// CustomerPaymentMethodRequest Contains details of the payment methods that the customer has active or has used in Conekta
type CustomerPaymentMethodRequest struct {
	// Type of payment method
	Type string `json:"type"`
}

// NewCustomerPaymentMethodRequest instantiates a new CustomerPaymentMethodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentMethodRequest(type_ string) *CustomerPaymentMethodRequest {
	this := CustomerPaymentMethodRequest{}
	this.Type = type_
	return &this
}

// NewCustomerPaymentMethodRequestWithDefaults instantiates a new CustomerPaymentMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentMethodRequestWithDefaults() *CustomerPaymentMethodRequest {
	this := CustomerPaymentMethodRequest{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerPaymentMethodRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentMethodRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerPaymentMethodRequest) SetType(v string) {
	o.Type = v
}

func (o CustomerPaymentMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPaymentMethodRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCustomerPaymentMethodRequest struct {
	value *CustomerPaymentMethodRequest
	isSet bool
}

func (v NullableCustomerPaymentMethodRequest) Get() *CustomerPaymentMethodRequest {
	return v.value
}

func (v *NullableCustomerPaymentMethodRequest) Set(val *CustomerPaymentMethodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentMethodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentMethodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentMethodRequest(val *CustomerPaymentMethodRequest) *NullableCustomerPaymentMethodRequest {
	return &NullableCustomerPaymentMethodRequest{value: val, isSet: true}
}

func (v NullableCustomerPaymentMethodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentMethodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

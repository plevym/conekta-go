package conekta

import (
	"encoding/json"
)

// checks if the CustomerPaymentMethods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPaymentMethods{}

// CustomerPaymentMethods struct for CustomerPaymentMethods
type CustomerPaymentMethods struct {
	Data []CustomerPaymentMethodsData `json:"data,omitempty"`
}

// NewCustomerPaymentMethods instantiates a new CustomerPaymentMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentMethods() *CustomerPaymentMethods {
	this := CustomerPaymentMethods{}
	return &this
}

// NewCustomerPaymentMethodsWithDefaults instantiates a new CustomerPaymentMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentMethodsWithDefaults() *CustomerPaymentMethods {
	this := CustomerPaymentMethods{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerPaymentMethods) GetData() []CustomerPaymentMethodsData {
	if o == nil || IsNil(o.Data) {
		var ret []CustomerPaymentMethodsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentMethods) GetDataOk() ([]CustomerPaymentMethodsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerPaymentMethods) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CustomerPaymentMethodsData and assigns it to the Data field.
func (o *CustomerPaymentMethods) SetData(v []CustomerPaymentMethodsData) {
	o.Data = v
}

func (o CustomerPaymentMethods) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPaymentMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomerPaymentMethods struct {
	value *CustomerPaymentMethods
	isSet bool
}

func (v NullableCustomerPaymentMethods) Get() *CustomerPaymentMethods {
	return v.value
}

func (v *NullableCustomerPaymentMethods) Set(val *CustomerPaymentMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentMethods(val *CustomerPaymentMethods) *NullableCustomerPaymentMethods {
	return &NullableCustomerPaymentMethods{value: val, isSet: true}
}

func (v NullableCustomerPaymentMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

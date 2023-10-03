package conekta

import (
	"encoding/json"
)

// checks if the CustomerInfoJustCustomerId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerInfoJustCustomerId{}

// CustomerInfoJustCustomerId struct for CustomerInfoJustCustomerId
type CustomerInfoJustCustomerId struct {
	CustomerId string `json:"customer_id"`
}

// NewCustomerInfoJustCustomerId instantiates a new CustomerInfoJustCustomerId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerInfoJustCustomerId(customerId string) *CustomerInfoJustCustomerId {
	this := CustomerInfoJustCustomerId{}
	this.CustomerId = customerId
	return &this
}

// NewCustomerInfoJustCustomerIdWithDefaults instantiates a new CustomerInfoJustCustomerId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerInfoJustCustomerIdWithDefaults() *CustomerInfoJustCustomerId {
	this := CustomerInfoJustCustomerId{}
	return &this
}

// GetCustomerId returns the CustomerId field value
func (o *CustomerInfoJustCustomerId) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CustomerInfoJustCustomerId) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CustomerInfoJustCustomerId) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o CustomerInfoJustCustomerId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerInfoJustCustomerId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer_id"] = o.CustomerId
	return toSerialize, nil
}

type NullableCustomerInfoJustCustomerId struct {
	value *CustomerInfoJustCustomerId
	isSet bool
}

func (v NullableCustomerInfoJustCustomerId) Get() *CustomerInfoJustCustomerId {
	return v.value
}

func (v *NullableCustomerInfoJustCustomerId) Set(val *CustomerInfoJustCustomerId) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerInfoJustCustomerId) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerInfoJustCustomerId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerInfoJustCustomerId(val *CustomerInfoJustCustomerId) *NullableCustomerInfoJustCustomerId {
	return &NullableCustomerInfoJustCustomerId{value: val, isSet: true}
}

func (v NullableCustomerInfoJustCustomerId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerInfoJustCustomerId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

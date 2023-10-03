package conekta

import (
	"encoding/json"
)

// checks if the CustomerResponseShippingContactsAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerResponseShippingContactsAllOf{}

// CustomerResponseShippingContactsAllOf struct for CustomerResponseShippingContactsAllOf
type CustomerResponseShippingContactsAllOf struct {
	Data []CustomerShippingContactsDataResponse `json:"data,omitempty"`
}

// NewCustomerResponseShippingContactsAllOf instantiates a new CustomerResponseShippingContactsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerResponseShippingContactsAllOf() *CustomerResponseShippingContactsAllOf {
	this := CustomerResponseShippingContactsAllOf{}
	return &this
}

// NewCustomerResponseShippingContactsAllOfWithDefaults instantiates a new CustomerResponseShippingContactsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerResponseShippingContactsAllOfWithDefaults() *CustomerResponseShippingContactsAllOf {
	this := CustomerResponseShippingContactsAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerResponseShippingContactsAllOf) GetData() []CustomerShippingContactsDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []CustomerShippingContactsDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponseShippingContactsAllOf) GetDataOk() ([]CustomerShippingContactsDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerResponseShippingContactsAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CustomerShippingContactsDataResponse and assigns it to the Data field.
func (o *CustomerResponseShippingContactsAllOf) SetData(v []CustomerShippingContactsDataResponse) {
	o.Data = v
}

func (o CustomerResponseShippingContactsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerResponseShippingContactsAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomerResponseShippingContactsAllOf struct {
	value *CustomerResponseShippingContactsAllOf
	isSet bool
}

func (v NullableCustomerResponseShippingContactsAllOf) Get() *CustomerResponseShippingContactsAllOf {
	return v.value
}

func (v *NullableCustomerResponseShippingContactsAllOf) Set(val *CustomerResponseShippingContactsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerResponseShippingContactsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerResponseShippingContactsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerResponseShippingContactsAllOf(val *CustomerResponseShippingContactsAllOf) *NullableCustomerResponseShippingContactsAllOf {
	return &NullableCustomerResponseShippingContactsAllOf{value: val, isSet: true}
}

func (v NullableCustomerResponseShippingContactsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerResponseShippingContactsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

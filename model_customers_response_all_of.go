package conekta

import (
	"encoding/json"
)

// checks if the CustomersResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomersResponseAllOf{}

// CustomersResponseAllOf struct for CustomersResponseAllOf
type CustomersResponseAllOf struct {
	Data []CustomerResponse `json:"data,omitempty"`
}

// NewCustomersResponseAllOf instantiates a new CustomersResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomersResponseAllOf() *CustomersResponseAllOf {
	this := CustomersResponseAllOf{}
	return &this
}

// NewCustomersResponseAllOfWithDefaults instantiates a new CustomersResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomersResponseAllOfWithDefaults() *CustomersResponseAllOf {
	this := CustomersResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomersResponseAllOf) GetData() []CustomerResponse {
	if o == nil || IsNil(o.Data) {
		var ret []CustomerResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomersResponseAllOf) GetDataOk() ([]CustomerResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomersResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CustomerResponse and assigns it to the Data field.
func (o *CustomersResponseAllOf) SetData(v []CustomerResponse) {
	o.Data = v
}

func (o CustomersResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomersResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomersResponseAllOf struct {
	value *CustomersResponseAllOf
	isSet bool
}

func (v NullableCustomersResponseAllOf) Get() *CustomersResponseAllOf {
	return v.value
}

func (v *NullableCustomersResponseAllOf) Set(val *CustomersResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomersResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomersResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomersResponseAllOf(val *CustomersResponseAllOf) *NullableCustomersResponseAllOf {
	return &NullableCustomersResponseAllOf{value: val, isSet: true}
}

func (v NullableCustomersResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomersResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the CustomerFiscalEntitiesResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerFiscalEntitiesResponseAllOf{}

// CustomerFiscalEntitiesResponseAllOf struct for CustomerFiscalEntitiesResponseAllOf
type CustomerFiscalEntitiesResponseAllOf struct {
	Data []CustomerFiscalEntitiesDataResponse `json:"data,omitempty"`
}

// NewCustomerFiscalEntitiesResponseAllOf instantiates a new CustomerFiscalEntitiesResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFiscalEntitiesResponseAllOf() *CustomerFiscalEntitiesResponseAllOf {
	this := CustomerFiscalEntitiesResponseAllOf{}
	return &this
}

// NewCustomerFiscalEntitiesResponseAllOfWithDefaults instantiates a new CustomerFiscalEntitiesResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFiscalEntitiesResponseAllOfWithDefaults() *CustomerFiscalEntitiesResponseAllOf {
	this := CustomerFiscalEntitiesResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesResponseAllOf) GetData() []CustomerFiscalEntitiesDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []CustomerFiscalEntitiesDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesResponseAllOf) GetDataOk() ([]CustomerFiscalEntitiesDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CustomerFiscalEntitiesDataResponse and assigns it to the Data field.
func (o *CustomerFiscalEntitiesResponseAllOf) SetData(v []CustomerFiscalEntitiesDataResponse) {
	o.Data = v
}

func (o CustomerFiscalEntitiesResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerFiscalEntitiesResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomerFiscalEntitiesResponseAllOf struct {
	value *CustomerFiscalEntitiesResponseAllOf
	isSet bool
}

func (v NullableCustomerFiscalEntitiesResponseAllOf) Get() *CustomerFiscalEntitiesResponseAllOf {
	return v.value
}

func (v *NullableCustomerFiscalEntitiesResponseAllOf) Set(val *CustomerFiscalEntitiesResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFiscalEntitiesResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFiscalEntitiesResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFiscalEntitiesResponseAllOf(val *CustomerFiscalEntitiesResponseAllOf) *NullableCustomerFiscalEntitiesResponseAllOf {
	return &NullableCustomerFiscalEntitiesResponseAllOf{value: val, isSet: true}
}

func (v NullableCustomerFiscalEntitiesResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFiscalEntitiesResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

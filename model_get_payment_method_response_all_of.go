package conekta

import (
	"encoding/json"
)

// checks if the GetPaymentMethodResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPaymentMethodResponseAllOf{}

// GetPaymentMethodResponseAllOf struct for GetPaymentMethodResponseAllOf
type GetPaymentMethodResponseAllOf struct {
	Data []GetCustomerPaymentMethodDataResponse `json:"data,omitempty"`
}

// NewGetPaymentMethodResponseAllOf instantiates a new GetPaymentMethodResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPaymentMethodResponseAllOf() *GetPaymentMethodResponseAllOf {
	this := GetPaymentMethodResponseAllOf{}
	return &this
}

// NewGetPaymentMethodResponseAllOfWithDefaults instantiates a new GetPaymentMethodResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPaymentMethodResponseAllOfWithDefaults() *GetPaymentMethodResponseAllOf {
	this := GetPaymentMethodResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPaymentMethodResponseAllOf) GetData() []GetCustomerPaymentMethodDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []GetCustomerPaymentMethodDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPaymentMethodResponseAllOf) GetDataOk() ([]GetCustomerPaymentMethodDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPaymentMethodResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetCustomerPaymentMethodDataResponse and assigns it to the Data field.
func (o *GetPaymentMethodResponseAllOf) SetData(v []GetCustomerPaymentMethodDataResponse) {
	o.Data = v
}

func (o GetPaymentMethodResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPaymentMethodResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetPaymentMethodResponseAllOf struct {
	value *GetPaymentMethodResponseAllOf
	isSet bool
}

func (v NullableGetPaymentMethodResponseAllOf) Get() *GetPaymentMethodResponseAllOf {
	return v.value
}

func (v *NullableGetPaymentMethodResponseAllOf) Set(val *GetPaymentMethodResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPaymentMethodResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPaymentMethodResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPaymentMethodResponseAllOf(val *GetPaymentMethodResponseAllOf) *NullableGetPaymentMethodResponseAllOf {
	return &NullableGetPaymentMethodResponseAllOf{value: val, isSet: true}
}

func (v NullableGetPaymentMethodResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPaymentMethodResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

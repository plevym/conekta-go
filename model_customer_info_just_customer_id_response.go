package conekta

import (
	"encoding/json"
)

// checks if the CustomerInfoJustCustomerIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerInfoJustCustomerIdResponse{}

// CustomerInfoJustCustomerIdResponse struct for CustomerInfoJustCustomerIdResponse
type CustomerInfoJustCustomerIdResponse struct {
	CustomerId *string `json:"customer_id,omitempty"`
}

// NewCustomerInfoJustCustomerIdResponse instantiates a new CustomerInfoJustCustomerIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerInfoJustCustomerIdResponse() *CustomerInfoJustCustomerIdResponse {
	this := CustomerInfoJustCustomerIdResponse{}
	return &this
}

// NewCustomerInfoJustCustomerIdResponseWithDefaults instantiates a new CustomerInfoJustCustomerIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerInfoJustCustomerIdResponseWithDefaults() *CustomerInfoJustCustomerIdResponse {
	this := CustomerInfoJustCustomerIdResponse{}
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CustomerInfoJustCustomerIdResponse) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfoJustCustomerIdResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CustomerInfoJustCustomerIdResponse) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CustomerInfoJustCustomerIdResponse) SetCustomerId(v string) {
	o.CustomerId = &v
}

func (o CustomerInfoJustCustomerIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerInfoJustCustomerIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	return toSerialize, nil
}

type NullableCustomerInfoJustCustomerIdResponse struct {
	value *CustomerInfoJustCustomerIdResponse
	isSet bool
}

func (v NullableCustomerInfoJustCustomerIdResponse) Get() *CustomerInfoJustCustomerIdResponse {
	return v.value
}

func (v *NullableCustomerInfoJustCustomerIdResponse) Set(val *CustomerInfoJustCustomerIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerInfoJustCustomerIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerInfoJustCustomerIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerInfoJustCustomerIdResponse(val *CustomerInfoJustCustomerIdResponse) *NullableCustomerInfoJustCustomerIdResponse {
	return &NullableCustomerInfoJustCustomerIdResponse{value: val, isSet: true}
}

func (v NullableCustomerInfoJustCustomerIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerInfoJustCustomerIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

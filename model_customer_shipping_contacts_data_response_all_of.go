package conekta

import (
	"encoding/json"
)

// checks if the CustomerShippingContactsDataResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerShippingContactsDataResponseAllOf{}

// CustomerShippingContactsDataResponseAllOf struct for CustomerShippingContactsDataResponseAllOf
type CustomerShippingContactsDataResponseAllOf struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	CreatedAt int64  `json:"created_at"`
}

// NewCustomerShippingContactsDataResponseAllOf instantiates a new CustomerShippingContactsDataResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerShippingContactsDataResponseAllOf(id string, object string, createdAt int64) *CustomerShippingContactsDataResponseAllOf {
	this := CustomerShippingContactsDataResponseAllOf{}
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	return &this
}

// NewCustomerShippingContactsDataResponseAllOfWithDefaults instantiates a new CustomerShippingContactsDataResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerShippingContactsDataResponseAllOfWithDefaults() *CustomerShippingContactsDataResponseAllOf {
	this := CustomerShippingContactsDataResponseAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *CustomerShippingContactsDataResponseAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerShippingContactsDataResponseAllOf) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *CustomerShippingContactsDataResponseAllOf) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponseAllOf) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CustomerShippingContactsDataResponseAllOf) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomerShippingContactsDataResponseAllOf) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponseAllOf) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomerShippingContactsDataResponseAllOf) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

func (o CustomerShippingContactsDataResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerShippingContactsDataResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullableCustomerShippingContactsDataResponseAllOf struct {
	value *CustomerShippingContactsDataResponseAllOf
	isSet bool
}

func (v NullableCustomerShippingContactsDataResponseAllOf) Get() *CustomerShippingContactsDataResponseAllOf {
	return v.value
}

func (v *NullableCustomerShippingContactsDataResponseAllOf) Set(val *CustomerShippingContactsDataResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerShippingContactsDataResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerShippingContactsDataResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerShippingContactsDataResponseAllOf(val *CustomerShippingContactsDataResponseAllOf) *NullableCustomerShippingContactsDataResponseAllOf {
	return &NullableCustomerShippingContactsDataResponseAllOf{value: val, isSet: true}
}

func (v NullableCustomerShippingContactsDataResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerShippingContactsDataResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

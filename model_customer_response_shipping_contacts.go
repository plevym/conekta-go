package conekta

import (
	"encoding/json"
)

// checks if the CustomerResponseShippingContacts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerResponseShippingContacts{}

// CustomerResponseShippingContacts struct for CustomerResponseShippingContacts
type CustomerResponseShippingContacts struct {
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string                                 `json:"object"`
	Data   []CustomerShippingContactsDataResponse `json:"data,omitempty"`
}

// NewCustomerResponseShippingContacts instantiates a new CustomerResponseShippingContacts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerResponseShippingContacts(hasMore bool, object string) *CustomerResponseShippingContacts {
	this := CustomerResponseShippingContacts{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewCustomerResponseShippingContactsWithDefaults instantiates a new CustomerResponseShippingContacts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerResponseShippingContactsWithDefaults() *CustomerResponseShippingContacts {
	this := CustomerResponseShippingContacts{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *CustomerResponseShippingContacts) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *CustomerResponseShippingContacts) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *CustomerResponseShippingContacts) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *CustomerResponseShippingContacts) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CustomerResponseShippingContacts) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CustomerResponseShippingContacts) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerResponseShippingContacts) GetData() []CustomerShippingContactsDataResponse {
	if o == nil || IsNil(o.Data) {
		var ret []CustomerShippingContactsDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponseShippingContacts) GetDataOk() ([]CustomerShippingContactsDataResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerResponseShippingContacts) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CustomerShippingContactsDataResponse and assigns it to the Data field.
func (o *CustomerResponseShippingContacts) SetData(v []CustomerShippingContactsDataResponse) {
	o.Data = v
}

func (o CustomerResponseShippingContacts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerResponseShippingContacts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_more"] = o.HasMore
	toSerialize["object"] = o.Object
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomerResponseShippingContacts struct {
	value *CustomerResponseShippingContacts
	isSet bool
}

func (v NullableCustomerResponseShippingContacts) Get() *CustomerResponseShippingContacts {
	return v.value
}

func (v *NullableCustomerResponseShippingContacts) Set(val *CustomerResponseShippingContacts) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerResponseShippingContacts) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerResponseShippingContacts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerResponseShippingContacts(val *CustomerResponseShippingContacts) *NullableCustomerResponseShippingContacts {
	return &NullableCustomerResponseShippingContacts{value: val, isSet: true}
}

func (v NullableCustomerResponseShippingContacts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerResponseShippingContacts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

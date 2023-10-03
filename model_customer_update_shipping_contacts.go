package conekta

import (
	"encoding/json"
)

// checks if the CustomerUpdateShippingContacts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerUpdateShippingContacts{}

// CustomerUpdateShippingContacts [Shipping](https://developers.conekta.com/v2.1.0/reference/createcustomershippingcontacts) details, required in case of sending a shipping. If we do not receive a shipping_contact on the order, the default shipping_contact of the customer will be used.
type CustomerUpdateShippingContacts struct {
	// Phone contact
	Phone *string `json:"phone,omitempty"`
	// Name of the person who will receive the order
	Receiver *string `json:"receiver,omitempty"`
	// The street names between which the order will be delivered.
	BetweenStreets *string                          `json:"between_streets,omitempty"`
	Address        *CustomerShippingContactsAddress `json:"address,omitempty"`
	ParentId       *string                          `json:"parent_id,omitempty"`
	Default        NullableBool                     `json:"default,omitempty"`
	Deleted        NullableBool                     `json:"deleted,omitempty"`
}

// NewCustomerUpdateShippingContacts instantiates a new CustomerUpdateShippingContacts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerUpdateShippingContacts() *CustomerUpdateShippingContacts {
	this := CustomerUpdateShippingContacts{}
	return &this
}

// NewCustomerUpdateShippingContactsWithDefaults instantiates a new CustomerUpdateShippingContacts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerUpdateShippingContactsWithDefaults() *CustomerUpdateShippingContacts {
	this := CustomerUpdateShippingContacts{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerUpdateShippingContacts) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateShippingContacts) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerUpdateShippingContacts) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerUpdateShippingContacts) SetPhone(v string) {
	o.Phone = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *CustomerUpdateShippingContacts) GetReceiver() string {
	if o == nil || IsNil(o.Receiver) {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateShippingContacts) GetReceiverOk() (*string, bool) {
	if o == nil || IsNil(o.Receiver) {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *CustomerUpdateShippingContacts) HasReceiver() bool {
	if o != nil && !IsNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *CustomerUpdateShippingContacts) SetReceiver(v string) {
	o.Receiver = &v
}

// GetBetweenStreets returns the BetweenStreets field value if set, zero value otherwise.
func (o *CustomerUpdateShippingContacts) GetBetweenStreets() string {
	if o == nil || IsNil(o.BetweenStreets) {
		var ret string
		return ret
	}
	return *o.BetweenStreets
}

// GetBetweenStreetsOk returns a tuple with the BetweenStreets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateShippingContacts) GetBetweenStreetsOk() (*string, bool) {
	if o == nil || IsNil(o.BetweenStreets) {
		return nil, false
	}
	return o.BetweenStreets, true
}

// HasBetweenStreets returns a boolean if a field has been set.
func (o *CustomerUpdateShippingContacts) HasBetweenStreets() bool {
	if o != nil && !IsNil(o.BetweenStreets) {
		return true
	}

	return false
}

// SetBetweenStreets gets a reference to the given string and assigns it to the BetweenStreets field.
func (o *CustomerUpdateShippingContacts) SetBetweenStreets(v string) {
	o.BetweenStreets = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CustomerUpdateShippingContacts) GetAddress() CustomerShippingContactsAddress {
	if o == nil || IsNil(o.Address) {
		var ret CustomerShippingContactsAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateShippingContacts) GetAddressOk() (*CustomerShippingContactsAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CustomerUpdateShippingContacts) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CustomerShippingContactsAddress and assigns it to the Address field.
func (o *CustomerUpdateShippingContacts) SetAddress(v CustomerShippingContactsAddress) {
	o.Address = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CustomerUpdateShippingContacts) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateShippingContacts) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CustomerUpdateShippingContacts) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *CustomerUpdateShippingContacts) SetParentId(v string) {
	o.ParentId = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerUpdateShippingContacts) GetDefault() bool {
	if o == nil || IsNil(o.Default.Get()) {
		var ret bool
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerUpdateShippingContacts) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *CustomerUpdateShippingContacts) HasDefault() bool {
	if o != nil && o.Default.IsSet() {
		return true
	}

	return false
}

// SetDefault gets a reference to the given NullableBool and assigns it to the Default field.
func (o *CustomerUpdateShippingContacts) SetDefault(v bool) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil
func (o *CustomerUpdateShippingContacts) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil
func (o *CustomerUpdateShippingContacts) UnsetDefault() {
	o.Default.Unset()
}

// GetDeleted returns the Deleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerUpdateShippingContacts) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted.Get()) {
		var ret bool
		return ret
	}
	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerUpdateShippingContacts) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// HasDeleted returns a boolean if a field has been set.
func (o *CustomerUpdateShippingContacts) HasDeleted() bool {
	if o != nil && o.Deleted.IsSet() {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given NullableBool and assigns it to the Deleted field.
func (o *CustomerUpdateShippingContacts) SetDeleted(v bool) {
	o.Deleted.Set(&v)
}

// SetDeletedNil sets the value for Deleted to be an explicit nil
func (o *CustomerUpdateShippingContacts) SetDeletedNil() {
	o.Deleted.Set(nil)
}

// UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
func (o *CustomerUpdateShippingContacts) UnsetDeleted() {
	o.Deleted.Unset()
}

func (o CustomerUpdateShippingContacts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerUpdateShippingContacts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Receiver) {
		toSerialize["receiver"] = o.Receiver
	}
	if !IsNil(o.BetweenStreets) {
		toSerialize["between_streets"] = o.BetweenStreets
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if o.Deleted.IsSet() {
		toSerialize["deleted"] = o.Deleted.Get()
	}
	return toSerialize, nil
}

type NullableCustomerUpdateShippingContacts struct {
	value *CustomerUpdateShippingContacts
	isSet bool
}

func (v NullableCustomerUpdateShippingContacts) Get() *CustomerUpdateShippingContacts {
	return v.value
}

func (v *NullableCustomerUpdateShippingContacts) Set(val *CustomerUpdateShippingContacts) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerUpdateShippingContacts) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerUpdateShippingContacts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerUpdateShippingContacts(val *CustomerUpdateShippingContacts) *NullableCustomerUpdateShippingContacts {
	return &NullableCustomerUpdateShippingContacts{value: val, isSet: true}
}

func (v NullableCustomerUpdateShippingContacts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerUpdateShippingContacts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseShippingContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseShippingContact{}

// OrderResponseShippingContact struct for OrderResponseShippingContact
type OrderResponseShippingContact struct {
	CreatedAt      *int64                                   `json:"created_at,omitempty"`
	Id             *string                                  `json:"id,omitempty"`
	Object         *string                                  `json:"object,omitempty"`
	Phone          *string                                  `json:"phone,omitempty"`
	Receiver       *string                                  `json:"receiver,omitempty"`
	BetweenStreets NullableString                           `json:"between_streets,omitempty"`
	Address        *CustomerShippingContactsResponseAddress `json:"address,omitempty"`
	ParentId       *string                                  `json:"parent_id,omitempty"`
	Default        *bool                                    `json:"default,omitempty"`
	Deleted        *bool                                    `json:"deleted,omitempty"`
}

// NewOrderResponseShippingContact instantiates a new OrderResponseShippingContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseShippingContact() *OrderResponseShippingContact {
	this := OrderResponseShippingContact{}
	return &this
}

// NewOrderResponseShippingContactWithDefaults instantiates a new OrderResponseShippingContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseShippingContactWithDefaults() *OrderResponseShippingContact {
	this := OrderResponseShippingContact{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderResponseShippingContact) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContact) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *OrderResponseShippingContact) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderResponseShippingContact) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContact) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderResponseShippingContact) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderResponseShippingContact) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContact) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderResponseShippingContact) SetObject(v string) {
	o.Object = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *OrderResponseShippingContact) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContact) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *OrderResponseShippingContact) SetPhone(v string) {
	o.Phone = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *OrderResponseShippingContact) GetReceiver() string {
	if o == nil || IsNil(o.Receiver) {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContact) GetReceiverOk() (*string, bool) {
	if o == nil || IsNil(o.Receiver) {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasReceiver() bool {
	if o != nil && !IsNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *OrderResponseShippingContact) SetReceiver(v string) {
	o.Receiver = &v
}

// GetBetweenStreets returns the BetweenStreets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponseShippingContact) GetBetweenStreets() string {
	if o == nil || IsNil(o.BetweenStreets.Get()) {
		var ret string
		return ret
	}
	return *o.BetweenStreets.Get()
}

// GetBetweenStreetsOk returns a tuple with the BetweenStreets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponseShippingContact) GetBetweenStreetsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BetweenStreets.Get(), o.BetweenStreets.IsSet()
}

// HasBetweenStreets returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasBetweenStreets() bool {
	if o != nil && o.BetweenStreets.IsSet() {
		return true
	}

	return false
}

// SetBetweenStreets gets a reference to the given NullableString and assigns it to the BetweenStreets field.
func (o *OrderResponseShippingContact) SetBetweenStreets(v string) {
	o.BetweenStreets.Set(&v)
}

// SetBetweenStreetsNil sets the value for BetweenStreets to be an explicit nil
func (o *OrderResponseShippingContact) SetBetweenStreetsNil() {
	o.BetweenStreets.Set(nil)
}

// UnsetBetweenStreets ensures that no value is present for BetweenStreets, not even an explicit nil
func (o *OrderResponseShippingContact) UnsetBetweenStreets() {
	o.BetweenStreets.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrderResponseShippingContact) GetAddress() CustomerShippingContactsResponseAddress {
	if o == nil || IsNil(o.Address) {
		var ret CustomerShippingContactsResponseAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContact) GetAddressOk() (*CustomerShippingContactsResponseAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CustomerShippingContactsResponseAddress and assigns it to the Address field.
func (o *OrderResponseShippingContact) SetAddress(v CustomerShippingContactsResponseAddress) {
	o.Address = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *OrderResponseShippingContact) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContact) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *OrderResponseShippingContact) SetParentId(v string) {
	o.ParentId = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *OrderResponseShippingContact) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContact) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *OrderResponseShippingContact) SetDefault(v bool) {
	o.Default = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *OrderResponseShippingContact) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseShippingContact) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *OrderResponseShippingContact) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *OrderResponseShippingContact) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o OrderResponseShippingContact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseShippingContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Receiver) {
		toSerialize["receiver"] = o.Receiver
	}
	if o.BetweenStreets.IsSet() {
		toSerialize["between_streets"] = o.BetweenStreets.Get()
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

type NullableOrderResponseShippingContact struct {
	value *OrderResponseShippingContact
	isSet bool
}

func (v NullableOrderResponseShippingContact) Get() *OrderResponseShippingContact {
	return v.value
}

func (v *NullableOrderResponseShippingContact) Set(val *OrderResponseShippingContact) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseShippingContact) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseShippingContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseShippingContact(val *OrderResponseShippingContact) *NullableOrderResponseShippingContact {
	return &NullableOrderResponseShippingContact{value: val, isSet: true}
}

func (v NullableOrderResponseShippingContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseShippingContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

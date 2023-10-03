package conekta

import (
	"encoding/json"
)

// checks if the CustomerShippingContactsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerShippingContactsResponse{}

// CustomerShippingContactsResponse Contains the detail of the shipping addresses that the client has active or has used in Conekta
type CustomerShippingContactsResponse struct {
	Phone          *string                                  `json:"phone,omitempty"`
	Receiver       *string                                  `json:"receiver,omitempty"`
	BetweenStreets NullableString                           `json:"between_streets,omitempty"`
	Address        *CustomerShippingContactsResponseAddress `json:"address,omitempty"`
	ParentId       *string                                  `json:"parent_id,omitempty"`
	Default        *bool                                    `json:"default,omitempty"`
	Id             *string                                  `json:"id,omitempty"`
	CreatedAt      *int64                                   `json:"created_at,omitempty"`
	Object         *string                                  `json:"object,omitempty"`
	Deleted        *bool                                    `json:"deleted,omitempty"`
}

// NewCustomerShippingContactsResponse instantiates a new CustomerShippingContactsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerShippingContactsResponse() *CustomerShippingContactsResponse {
	this := CustomerShippingContactsResponse{}
	return &this
}

// NewCustomerShippingContactsResponseWithDefaults instantiates a new CustomerShippingContactsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerShippingContactsResponseWithDefaults() *CustomerShippingContactsResponse {
	this := CustomerShippingContactsResponse{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponse) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponse) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerShippingContactsResponse) SetPhone(v string) {
	o.Phone = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponse) GetReceiver() string {
	if o == nil || IsNil(o.Receiver) {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponse) GetReceiverOk() (*string, bool) {
	if o == nil || IsNil(o.Receiver) {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasReceiver() bool {
	if o != nil && !IsNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *CustomerShippingContactsResponse) SetReceiver(v string) {
	o.Receiver = &v
}

// GetBetweenStreets returns the BetweenStreets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerShippingContactsResponse) GetBetweenStreets() string {
	if o == nil || IsNil(o.BetweenStreets.Get()) {
		var ret string
		return ret
	}
	return *o.BetweenStreets.Get()
}

// GetBetweenStreetsOk returns a tuple with the BetweenStreets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerShippingContactsResponse) GetBetweenStreetsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BetweenStreets.Get(), o.BetweenStreets.IsSet()
}

// HasBetweenStreets returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasBetweenStreets() bool {
	if o != nil && o.BetweenStreets.IsSet() {
		return true
	}

	return false
}

// SetBetweenStreets gets a reference to the given NullableString and assigns it to the BetweenStreets field.
func (o *CustomerShippingContactsResponse) SetBetweenStreets(v string) {
	o.BetweenStreets.Set(&v)
}

// SetBetweenStreetsNil sets the value for BetweenStreets to be an explicit nil
func (o *CustomerShippingContactsResponse) SetBetweenStreetsNil() {
	o.BetweenStreets.Set(nil)
}

// UnsetBetweenStreets ensures that no value is present for BetweenStreets, not even an explicit nil
func (o *CustomerShippingContactsResponse) UnsetBetweenStreets() {
	o.BetweenStreets.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponse) GetAddress() CustomerShippingContactsResponseAddress {
	if o == nil || IsNil(o.Address) {
		var ret CustomerShippingContactsResponseAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponse) GetAddressOk() (*CustomerShippingContactsResponseAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CustomerShippingContactsResponseAddress and assigns it to the Address field.
func (o *CustomerShippingContactsResponse) SetAddress(v CustomerShippingContactsResponseAddress) {
	o.Address = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *CustomerShippingContactsResponse) SetParentId(v string) {
	o.ParentId = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponse) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponse) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CustomerShippingContactsResponse) SetDefault(v bool) {
	o.Default = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerShippingContactsResponse) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *CustomerShippingContactsResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CustomerShippingContactsResponse) SetObject(v string) {
	o.Object = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponse) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponse) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponse) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *CustomerShippingContactsResponse) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o CustomerShippingContactsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerShippingContactsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

type NullableCustomerShippingContactsResponse struct {
	value *CustomerShippingContactsResponse
	isSet bool
}

func (v NullableCustomerShippingContactsResponse) Get() *CustomerShippingContactsResponse {
	return v.value
}

func (v *NullableCustomerShippingContactsResponse) Set(val *CustomerShippingContactsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerShippingContactsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerShippingContactsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerShippingContactsResponse(val *CustomerShippingContactsResponse) *NullableCustomerShippingContactsResponse {
	return &NullableCustomerShippingContactsResponse{value: val, isSet: true}
}

func (v NullableCustomerShippingContactsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerShippingContactsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the CustomerShippingContactsDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerShippingContactsDataResponse{}

// CustomerShippingContactsDataResponse struct for CustomerShippingContactsDataResponse
type CustomerShippingContactsDataResponse struct {
	// Phone contact
	Phone *string `json:"phone,omitempty"`
	// Name of the person who will receive the order
	Receiver *string `json:"receiver,omitempty"`
	// The street names between which the order will be delivered.
	BetweenStreets *string                         `json:"between_streets,omitempty"`
	Address        CustomerShippingContactsAddress `json:"address"`
	ParentId       *string                         `json:"parent_id,omitempty"`
	Default        NullableBool                    `json:"default,omitempty"`
	Deleted        NullableBool                    `json:"deleted,omitempty"`
	Id             string                          `json:"id"`
	Object         string                          `json:"object"`
	CreatedAt      int64                           `json:"created_at"`
}

// NewCustomerShippingContactsDataResponse instantiates a new CustomerShippingContactsDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerShippingContactsDataResponse(address CustomerShippingContactsAddress, id string, object string, createdAt int64) *CustomerShippingContactsDataResponse {
	this := CustomerShippingContactsDataResponse{}
	this.Address = address
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	return &this
}

// NewCustomerShippingContactsDataResponseWithDefaults instantiates a new CustomerShippingContactsDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerShippingContactsDataResponseWithDefaults() *CustomerShippingContactsDataResponse {
	this := CustomerShippingContactsDataResponse{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerShippingContactsDataResponse) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponse) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerShippingContactsDataResponse) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerShippingContactsDataResponse) SetPhone(v string) {
	o.Phone = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *CustomerShippingContactsDataResponse) GetReceiver() string {
	if o == nil || IsNil(o.Receiver) {
		var ret string
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponse) GetReceiverOk() (*string, bool) {
	if o == nil || IsNil(o.Receiver) {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *CustomerShippingContactsDataResponse) HasReceiver() bool {
	if o != nil && !IsNil(o.Receiver) {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given string and assigns it to the Receiver field.
func (o *CustomerShippingContactsDataResponse) SetReceiver(v string) {
	o.Receiver = &v
}

// GetBetweenStreets returns the BetweenStreets field value if set, zero value otherwise.
func (o *CustomerShippingContactsDataResponse) GetBetweenStreets() string {
	if o == nil || IsNil(o.BetweenStreets) {
		var ret string
		return ret
	}
	return *o.BetweenStreets
}

// GetBetweenStreetsOk returns a tuple with the BetweenStreets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponse) GetBetweenStreetsOk() (*string, bool) {
	if o == nil || IsNil(o.BetweenStreets) {
		return nil, false
	}
	return o.BetweenStreets, true
}

// HasBetweenStreets returns a boolean if a field has been set.
func (o *CustomerShippingContactsDataResponse) HasBetweenStreets() bool {
	if o != nil && !IsNil(o.BetweenStreets) {
		return true
	}

	return false
}

// SetBetweenStreets gets a reference to the given string and assigns it to the BetweenStreets field.
func (o *CustomerShippingContactsDataResponse) SetBetweenStreets(v string) {
	o.BetweenStreets = &v
}

// GetAddress returns the Address field value
func (o *CustomerShippingContactsDataResponse) GetAddress() CustomerShippingContactsAddress {
	if o == nil {
		var ret CustomerShippingContactsAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponse) GetAddressOk() (*CustomerShippingContactsAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CustomerShippingContactsDataResponse) SetAddress(v CustomerShippingContactsAddress) {
	o.Address = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CustomerShippingContactsDataResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CustomerShippingContactsDataResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *CustomerShippingContactsDataResponse) SetParentId(v string) {
	o.ParentId = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerShippingContactsDataResponse) GetDefault() bool {
	if o == nil || IsNil(o.Default.Get()) {
		var ret bool
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerShippingContactsDataResponse) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *CustomerShippingContactsDataResponse) HasDefault() bool {
	if o != nil && o.Default.IsSet() {
		return true
	}

	return false
}

// SetDefault gets a reference to the given NullableBool and assigns it to the Default field.
func (o *CustomerShippingContactsDataResponse) SetDefault(v bool) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil
func (o *CustomerShippingContactsDataResponse) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil
func (o *CustomerShippingContactsDataResponse) UnsetDefault() {
	o.Default.Unset()
}

// GetDeleted returns the Deleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerShippingContactsDataResponse) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted.Get()) {
		var ret bool
		return ret
	}
	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerShippingContactsDataResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// HasDeleted returns a boolean if a field has been set.
func (o *CustomerShippingContactsDataResponse) HasDeleted() bool {
	if o != nil && o.Deleted.IsSet() {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given NullableBool and assigns it to the Deleted field.
func (o *CustomerShippingContactsDataResponse) SetDeleted(v bool) {
	o.Deleted.Set(&v)
}

// SetDeletedNil sets the value for Deleted to be an explicit nil
func (o *CustomerShippingContactsDataResponse) SetDeletedNil() {
	o.Deleted.Set(nil)
}

// UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
func (o *CustomerShippingContactsDataResponse) UnsetDeleted() {
	o.Deleted.Unset()
}

// GetId returns the Id field value
func (o *CustomerShippingContactsDataResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerShippingContactsDataResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *CustomerShippingContactsDataResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CustomerShippingContactsDataResponse) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomerShippingContactsDataResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsDataResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomerShippingContactsDataResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

func (o CustomerShippingContactsDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerShippingContactsDataResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["address"] = o.Address
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if o.Deleted.IsSet() {
		toSerialize["deleted"] = o.Deleted.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

type NullableCustomerShippingContactsDataResponse struct {
	value *CustomerShippingContactsDataResponse
	isSet bool
}

func (v NullableCustomerShippingContactsDataResponse) Get() *CustomerShippingContactsDataResponse {
	return v.value
}

func (v *NullableCustomerShippingContactsDataResponse) Set(val *CustomerShippingContactsDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerShippingContactsDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerShippingContactsDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerShippingContactsDataResponse(val *CustomerShippingContactsDataResponse) *NullableCustomerShippingContactsDataResponse {
	return &NullableCustomerShippingContactsDataResponse{value: val, isSet: true}
}

func (v NullableCustomerShippingContactsDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerShippingContactsDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

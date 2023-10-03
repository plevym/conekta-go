package conekta

import (
	"encoding/json"
)

// checks if the UpdateCustomerAntifraudInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomerAntifraudInfo{}

// UpdateCustomerAntifraudInfo struct for UpdateCustomerAntifraudInfo
type UpdateCustomerAntifraudInfo struct {
	AccountCreatedAt *int64 `json:"account_created_at,omitempty"`
	FirstPaidAt      *int32 `json:"first_paid_at,omitempty"`
}

// NewUpdateCustomerAntifraudInfo instantiates a new UpdateCustomerAntifraudInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomerAntifraudInfo() *UpdateCustomerAntifraudInfo {
	this := UpdateCustomerAntifraudInfo{}
	return &this
}

// NewUpdateCustomerAntifraudInfoWithDefaults instantiates a new UpdateCustomerAntifraudInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomerAntifraudInfoWithDefaults() *UpdateCustomerAntifraudInfo {
	this := UpdateCustomerAntifraudInfo{}
	return &this
}

// GetAccountCreatedAt returns the AccountCreatedAt field value if set, zero value otherwise.
func (o *UpdateCustomerAntifraudInfo) GetAccountCreatedAt() int64 {
	if o == nil || IsNil(o.AccountCreatedAt) {
		var ret int64
		return ret
	}
	return *o.AccountCreatedAt
}

// GetAccountCreatedAtOk returns a tuple with the AccountCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomerAntifraudInfo) GetAccountCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountCreatedAt) {
		return nil, false
	}
	return o.AccountCreatedAt, true
}

// HasAccountCreatedAt returns a boolean if a field has been set.
func (o *UpdateCustomerAntifraudInfo) HasAccountCreatedAt() bool {
	if o != nil && !IsNil(o.AccountCreatedAt) {
		return true
	}

	return false
}

// SetAccountCreatedAt gets a reference to the given int64 and assigns it to the AccountCreatedAt field.
func (o *UpdateCustomerAntifraudInfo) SetAccountCreatedAt(v int64) {
	o.AccountCreatedAt = &v
}

// GetFirstPaidAt returns the FirstPaidAt field value if set, zero value otherwise.
func (o *UpdateCustomerAntifraudInfo) GetFirstPaidAt() int32 {
	if o == nil || IsNil(o.FirstPaidAt) {
		var ret int32
		return ret
	}
	return *o.FirstPaidAt
}

// GetFirstPaidAtOk returns a tuple with the FirstPaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomerAntifraudInfo) GetFirstPaidAtOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstPaidAt) {
		return nil, false
	}
	return o.FirstPaidAt, true
}

// HasFirstPaidAt returns a boolean if a field has been set.
func (o *UpdateCustomerAntifraudInfo) HasFirstPaidAt() bool {
	if o != nil && !IsNil(o.FirstPaidAt) {
		return true
	}

	return false
}

// SetFirstPaidAt gets a reference to the given int32 and assigns it to the FirstPaidAt field.
func (o *UpdateCustomerAntifraudInfo) SetFirstPaidAt(v int32) {
	o.FirstPaidAt = &v
}

func (o UpdateCustomerAntifraudInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomerAntifraudInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountCreatedAt) {
		toSerialize["account_created_at"] = o.AccountCreatedAt
	}
	if !IsNil(o.FirstPaidAt) {
		toSerialize["first_paid_at"] = o.FirstPaidAt
	}
	return toSerialize, nil
}

type NullableUpdateCustomerAntifraudInfo struct {
	value *UpdateCustomerAntifraudInfo
	isSet bool
}

func (v NullableUpdateCustomerAntifraudInfo) Get() *UpdateCustomerAntifraudInfo {
	return v.value
}

func (v *NullableUpdateCustomerAntifraudInfo) Set(val *UpdateCustomerAntifraudInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomerAntifraudInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomerAntifraudInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomerAntifraudInfo(val *UpdateCustomerAntifraudInfo) *NullableUpdateCustomerAntifraudInfo {
	return &NullableUpdateCustomerAntifraudInfo{value: val, isSet: true}
}

func (v NullableUpdateCustomerAntifraudInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomerAntifraudInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

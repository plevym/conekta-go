package conekta

import (
	"encoding/json"
)

// checks if the CustomerAntifraudInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerAntifraudInfo{}

// CustomerAntifraudInfo struct for CustomerAntifraudInfo
type CustomerAntifraudInfo struct {
	AccountCreatedAt *int64 `json:"account_created_at,omitempty"`
	FirstPaidAt      *int64 `json:"first_paid_at,omitempty"`
}

// NewCustomerAntifraudInfo instantiates a new CustomerAntifraudInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAntifraudInfo() *CustomerAntifraudInfo {
	this := CustomerAntifraudInfo{}
	return &this
}

// NewCustomerAntifraudInfoWithDefaults instantiates a new CustomerAntifraudInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAntifraudInfoWithDefaults() *CustomerAntifraudInfo {
	this := CustomerAntifraudInfo{}
	return &this
}

// GetAccountCreatedAt returns the AccountCreatedAt field value if set, zero value otherwise.
func (o *CustomerAntifraudInfo) GetAccountCreatedAt() int64 {
	if o == nil || IsNil(o.AccountCreatedAt) {
		var ret int64
		return ret
	}
	return *o.AccountCreatedAt
}

// GetAccountCreatedAtOk returns a tuple with the AccountCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAntifraudInfo) GetAccountCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountCreatedAt) {
		return nil, false
	}
	return o.AccountCreatedAt, true
}

// HasAccountCreatedAt returns a boolean if a field has been set.
func (o *CustomerAntifraudInfo) HasAccountCreatedAt() bool {
	if o != nil && !IsNil(o.AccountCreatedAt) {
		return true
	}

	return false
}

// SetAccountCreatedAt gets a reference to the given int64 and assigns it to the AccountCreatedAt field.
func (o *CustomerAntifraudInfo) SetAccountCreatedAt(v int64) {
	o.AccountCreatedAt = &v
}

// GetFirstPaidAt returns the FirstPaidAt field value if set, zero value otherwise.
func (o *CustomerAntifraudInfo) GetFirstPaidAt() int64 {
	if o == nil || IsNil(o.FirstPaidAt) {
		var ret int64
		return ret
	}
	return *o.FirstPaidAt
}

// GetFirstPaidAtOk returns a tuple with the FirstPaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAntifraudInfo) GetFirstPaidAtOk() (*int64, bool) {
	if o == nil || IsNil(o.FirstPaidAt) {
		return nil, false
	}
	return o.FirstPaidAt, true
}

// HasFirstPaidAt returns a boolean if a field has been set.
func (o *CustomerAntifraudInfo) HasFirstPaidAt() bool {
	if o != nil && !IsNil(o.FirstPaidAt) {
		return true
	}

	return false
}

// SetFirstPaidAt gets a reference to the given int64 and assigns it to the FirstPaidAt field.
func (o *CustomerAntifraudInfo) SetFirstPaidAt(v int64) {
	o.FirstPaidAt = &v
}

func (o CustomerAntifraudInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerAntifraudInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountCreatedAt) {
		toSerialize["account_created_at"] = o.AccountCreatedAt
	}
	if !IsNil(o.FirstPaidAt) {
		toSerialize["first_paid_at"] = o.FirstPaidAt
	}
	return toSerialize, nil
}

type NullableCustomerAntifraudInfo struct {
	value *CustomerAntifraudInfo
	isSet bool
}

func (v NullableCustomerAntifraudInfo) Get() *CustomerAntifraudInfo {
	return v.value
}

func (v *NullableCustomerAntifraudInfo) Set(val *CustomerAntifraudInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAntifraudInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAntifraudInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAntifraudInfo(val *CustomerAntifraudInfo) *NullableCustomerAntifraudInfo {
	return &NullableCustomerAntifraudInfo{value: val, isSet: true}
}

func (v NullableCustomerAntifraudInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAntifraudInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

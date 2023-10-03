package conekta

import (
	"encoding/json"
)

// checks if the CustomerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerInfo{}

// CustomerInfo struct for CustomerInfo
type CustomerInfo struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Phone     string  `json:"phone"`
	Corporate *bool   `json:"corporate,omitempty"`
	Object    *string `json:"object,omitempty"`
}

// NewCustomerInfo instantiates a new CustomerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerInfo(name string, email string, phone string) *CustomerInfo {
	this := CustomerInfo{}
	this.Name = name
	this.Email = email
	this.Phone = phone
	return &this
}

// NewCustomerInfoWithDefaults instantiates a new CustomerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerInfoWithDefaults() *CustomerInfo {
	this := CustomerInfo{}
	return &this
}

// GetName returns the Name field value
func (o *CustomerInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomerInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomerInfo) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *CustomerInfo) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CustomerInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CustomerInfo) SetEmail(v string) {
	o.Email = v
}

// GetPhone returns the Phone field value
func (o *CustomerInfo) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *CustomerInfo) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *CustomerInfo) SetPhone(v string) {
	o.Phone = v
}

// GetCorporate returns the Corporate field value if set, zero value otherwise.
func (o *CustomerInfo) GetCorporate() bool {
	if o == nil || IsNil(o.Corporate) {
		var ret bool
		return ret
	}
	return *o.Corporate
}

// GetCorporateOk returns a tuple with the Corporate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfo) GetCorporateOk() (*bool, bool) {
	if o == nil || IsNil(o.Corporate) {
		return nil, false
	}
	return o.Corporate, true
}

// HasCorporate returns a boolean if a field has been set.
func (o *CustomerInfo) HasCorporate() bool {
	if o != nil && !IsNil(o.Corporate) {
		return true
	}

	return false
}

// SetCorporate gets a reference to the given bool and assigns it to the Corporate field.
func (o *CustomerInfo) SetCorporate(v bool) {
	o.Corporate = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CustomerInfo) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfo) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CustomerInfo) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CustomerInfo) SetObject(v string) {
	o.Object = &v
}

func (o CustomerInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["email"] = o.Email
	toSerialize["phone"] = o.Phone
	if !IsNil(o.Corporate) {
		toSerialize["corporate"] = o.Corporate
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableCustomerInfo struct {
	value *CustomerInfo
	isSet bool
}

func (v NullableCustomerInfo) Get() *CustomerInfo {
	return v.value
}

func (v *NullableCustomerInfo) Set(val *CustomerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerInfo(val *CustomerInfo) *NullableCustomerInfo {
	return &NullableCustomerInfo{value: val, isSet: true}
}

func (v NullableCustomerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseCustomerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseCustomerInfo{}

// OrderResponseCustomerInfo struct for OrderResponseCustomerInfo
type OrderResponseCustomerInfo struct {
	Object     *string `json:"object,omitempty"`
	Name       *string `json:"name,omitempty"`
	Email      *string `json:"email,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Corporate  *bool   `json:"corporate,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
}

// NewOrderResponseCustomerInfo instantiates a new OrderResponseCustomerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseCustomerInfo() *OrderResponseCustomerInfo {
	this := OrderResponseCustomerInfo{}
	var corporate bool = false
	this.Corporate = &corporate
	return &this
}

// NewOrderResponseCustomerInfoWithDefaults instantiates a new OrderResponseCustomerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseCustomerInfoWithDefaults() *OrderResponseCustomerInfo {
	this := OrderResponseCustomerInfo{}
	var corporate bool = false
	this.Corporate = &corporate
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderResponseCustomerInfo) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCustomerInfo) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderResponseCustomerInfo) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderResponseCustomerInfo) SetObject(v string) {
	o.Object = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrderResponseCustomerInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCustomerInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrderResponseCustomerInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrderResponseCustomerInfo) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrderResponseCustomerInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCustomerInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrderResponseCustomerInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrderResponseCustomerInfo) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *OrderResponseCustomerInfo) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCustomerInfo) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *OrderResponseCustomerInfo) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *OrderResponseCustomerInfo) SetPhone(v string) {
	o.Phone = &v
}

// GetCorporate returns the Corporate field value if set, zero value otherwise.
func (o *OrderResponseCustomerInfo) GetCorporate() bool {
	if o == nil || IsNil(o.Corporate) {
		var ret bool
		return ret
	}
	return *o.Corporate
}

// GetCorporateOk returns a tuple with the Corporate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCustomerInfo) GetCorporateOk() (*bool, bool) {
	if o == nil || IsNil(o.Corporate) {
		return nil, false
	}
	return o.Corporate, true
}

// HasCorporate returns a boolean if a field has been set.
func (o *OrderResponseCustomerInfo) HasCorporate() bool {
	if o != nil && !IsNil(o.Corporate) {
		return true
	}

	return false
}

// SetCorporate gets a reference to the given bool and assigns it to the Corporate field.
func (o *OrderResponseCustomerInfo) SetCorporate(v bool) {
	o.Corporate = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *OrderResponseCustomerInfo) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCustomerInfo) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *OrderResponseCustomerInfo) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *OrderResponseCustomerInfo) SetCustomerId(v string) {
	o.CustomerId = &v
}

func (o OrderResponseCustomerInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseCustomerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Corporate) {
		toSerialize["corporate"] = o.Corporate
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	return toSerialize, nil
}

type NullableOrderResponseCustomerInfo struct {
	value *OrderResponseCustomerInfo
	isSet bool
}

func (v NullableOrderResponseCustomerInfo) Get() *OrderResponseCustomerInfo {
	return v.value
}

func (v *NullableOrderResponseCustomerInfo) Set(val *OrderResponseCustomerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseCustomerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseCustomerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseCustomerInfo(val *OrderResponseCustomerInfo) *NullableOrderResponseCustomerInfo {
	return &NullableOrderResponseCustomerInfo{value: val, isSet: true}
}

func (v NullableOrderResponseCustomerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseCustomerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the CustomerInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerInfoResponse{}

// CustomerInfoResponse struct for CustomerInfoResponse
type CustomerInfoResponse struct {
	Name      *string `json:"name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Corporate *bool   `json:"corporate,omitempty"`
	Object    *string `json:"object,omitempty"`
}

// NewCustomerInfoResponse instantiates a new CustomerInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerInfoResponse() *CustomerInfoResponse {
	this := CustomerInfoResponse{}
	var corporate bool = false
	this.Corporate = &corporate
	return &this
}

// NewCustomerInfoResponseWithDefaults instantiates a new CustomerInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerInfoResponseWithDefaults() *CustomerInfoResponse {
	this := CustomerInfoResponse{}
	var corporate bool = false
	this.Corporate = &corporate
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerInfoResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfoResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerInfoResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerInfoResponse) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerInfoResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfoResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerInfoResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerInfoResponse) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerInfoResponse) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfoResponse) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerInfoResponse) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerInfoResponse) SetPhone(v string) {
	o.Phone = &v
}

// GetCorporate returns the Corporate field value if set, zero value otherwise.
func (o *CustomerInfoResponse) GetCorporate() bool {
	if o == nil || IsNil(o.Corporate) {
		var ret bool
		return ret
	}
	return *o.Corporate
}

// GetCorporateOk returns a tuple with the Corporate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfoResponse) GetCorporateOk() (*bool, bool) {
	if o == nil || IsNil(o.Corporate) {
		return nil, false
	}
	return o.Corporate, true
}

// HasCorporate returns a boolean if a field has been set.
func (o *CustomerInfoResponse) HasCorporate() bool {
	if o != nil && !IsNil(o.Corporate) {
		return true
	}

	return false
}

// SetCorporate gets a reference to the given bool and assigns it to the Corporate field.
func (o *CustomerInfoResponse) SetCorporate(v bool) {
	o.Corporate = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CustomerInfoResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerInfoResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CustomerInfoResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CustomerInfoResponse) SetObject(v string) {
	o.Object = &v
}

func (o CustomerInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableCustomerInfoResponse struct {
	value *CustomerInfoResponse
	isSet bool
}

func (v NullableCustomerInfoResponse) Get() *CustomerInfoResponse {
	return v.value
}

func (v *NullableCustomerInfoResponse) Set(val *CustomerInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerInfoResponse(val *CustomerInfoResponse) *NullableCustomerInfoResponse {
	return &NullableCustomerInfoResponse{value: val, isSet: true}
}

func (v NullableCustomerInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

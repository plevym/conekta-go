package conekta

import (
	"encoding/json"
)

// checks if the CustomerShippingContactsResponseAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerShippingContactsResponseAddress{}

// CustomerShippingContactsResponseAddress struct for CustomerShippingContactsResponseAddress
type CustomerShippingContactsResponseAddress struct {
	Object      *string `json:"object,omitempty"`
	Street1     *string `json:"street1,omitempty"`
	Street2     *string `json:"street2,omitempty"`
	PostalCode  *string `json:"postal_code,omitempty"`
	City        *string `json:"city,omitempty"`
	State       *string `json:"state,omitempty"`
	Country     *string `json:"country,omitempty"`
	Residential *bool   `json:"residential,omitempty"`
}

// NewCustomerShippingContactsResponseAddress instantiates a new CustomerShippingContactsResponseAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerShippingContactsResponseAddress() *CustomerShippingContactsResponseAddress {
	this := CustomerShippingContactsResponseAddress{}
	return &this
}

// NewCustomerShippingContactsResponseAddressWithDefaults instantiates a new CustomerShippingContactsResponseAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerShippingContactsResponseAddressWithDefaults() *CustomerShippingContactsResponseAddress {
	this := CustomerShippingContactsResponseAddress{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponseAddress) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponseAddress) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponseAddress) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CustomerShippingContactsResponseAddress) SetObject(v string) {
	o.Object = &v
}

// GetStreet1 returns the Street1 field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponseAddress) GetStreet1() string {
	if o == nil || IsNil(o.Street1) {
		var ret string
		return ret
	}
	return *o.Street1
}

// GetStreet1Ok returns a tuple with the Street1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponseAddress) GetStreet1Ok() (*string, bool) {
	if o == nil || IsNil(o.Street1) {
		return nil, false
	}
	return o.Street1, true
}

// HasStreet1 returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponseAddress) HasStreet1() bool {
	if o != nil && !IsNil(o.Street1) {
		return true
	}

	return false
}

// SetStreet1 gets a reference to the given string and assigns it to the Street1 field.
func (o *CustomerShippingContactsResponseAddress) SetStreet1(v string) {
	o.Street1 = &v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponseAddress) GetStreet2() string {
	if o == nil || IsNil(o.Street2) {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponseAddress) GetStreet2Ok() (*string, bool) {
	if o == nil || IsNil(o.Street2) {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponseAddress) HasStreet2() bool {
	if o != nil && !IsNil(o.Street2) {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *CustomerShippingContactsResponseAddress) SetStreet2(v string) {
	o.Street2 = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponseAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponseAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponseAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CustomerShippingContactsResponseAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponseAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponseAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponseAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CustomerShippingContactsResponseAddress) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponseAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponseAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponseAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CustomerShippingContactsResponseAddress) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponseAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponseAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponseAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerShippingContactsResponseAddress) SetCountry(v string) {
	o.Country = &v
}

// GetResidential returns the Residential field value if set, zero value otherwise.
func (o *CustomerShippingContactsResponseAddress) GetResidential() bool {
	if o == nil || IsNil(o.Residential) {
		var ret bool
		return ret
	}
	return *o.Residential
}

// GetResidentialOk returns a tuple with the Residential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsResponseAddress) GetResidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Residential) {
		return nil, false
	}
	return o.Residential, true
}

// HasResidential returns a boolean if a field has been set.
func (o *CustomerShippingContactsResponseAddress) HasResidential() bool {
	if o != nil && !IsNil(o.Residential) {
		return true
	}

	return false
}

// SetResidential gets a reference to the given bool and assigns it to the Residential field.
func (o *CustomerShippingContactsResponseAddress) SetResidential(v bool) {
	o.Residential = &v
}

func (o CustomerShippingContactsResponseAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerShippingContactsResponseAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Street1) {
		toSerialize["street1"] = o.Street1
	}
	if !IsNil(o.Street2) {
		toSerialize["street2"] = o.Street2
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Residential) {
		toSerialize["residential"] = o.Residential
	}
	return toSerialize, nil
}

type NullableCustomerShippingContactsResponseAddress struct {
	value *CustomerShippingContactsResponseAddress
	isSet bool
}

func (v NullableCustomerShippingContactsResponseAddress) Get() *CustomerShippingContactsResponseAddress {
	return v.value
}

func (v *NullableCustomerShippingContactsResponseAddress) Set(val *CustomerShippingContactsResponseAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerShippingContactsResponseAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerShippingContactsResponseAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerShippingContactsResponseAddress(val *CustomerShippingContactsResponseAddress) *NullableCustomerShippingContactsResponseAddress {
	return &NullableCustomerShippingContactsResponseAddress{value: val, isSet: true}
}

func (v NullableCustomerShippingContactsResponseAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerShippingContactsResponseAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

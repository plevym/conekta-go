package conekta

import (
	"encoding/json"
)

// checks if the CustomerShippingContactsAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerShippingContactsAddress{}

// CustomerShippingContactsAddress Address of the person who will receive the order
type CustomerShippingContactsAddress struct {
	Street1    *string `json:"street1,omitempty"`
	Street2    *string `json:"street2,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	City       *string `json:"city,omitempty"`
	State      *string `json:"state,omitempty"`
	// this field follows the [ISO 3166-1 alpha-2 standard](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Country     *string      `json:"country,omitempty"`
	Residential NullableBool `json:"residential,omitempty"`
}

// NewCustomerShippingContactsAddress instantiates a new CustomerShippingContactsAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerShippingContactsAddress() *CustomerShippingContactsAddress {
	this := CustomerShippingContactsAddress{}
	return &this
}

// NewCustomerShippingContactsAddressWithDefaults instantiates a new CustomerShippingContactsAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerShippingContactsAddressWithDefaults() *CustomerShippingContactsAddress {
	this := CustomerShippingContactsAddress{}
	return &this
}

// GetStreet1 returns the Street1 field value if set, zero value otherwise.
func (o *CustomerShippingContactsAddress) GetStreet1() string {
	if o == nil || IsNil(o.Street1) {
		var ret string
		return ret
	}
	return *o.Street1
}

// GetStreet1Ok returns a tuple with the Street1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsAddress) GetStreet1Ok() (*string, bool) {
	if o == nil || IsNil(o.Street1) {
		return nil, false
	}
	return o.Street1, true
}

// HasStreet1 returns a boolean if a field has been set.
func (o *CustomerShippingContactsAddress) HasStreet1() bool {
	if o != nil && !IsNil(o.Street1) {
		return true
	}

	return false
}

// SetStreet1 gets a reference to the given string and assigns it to the Street1 field.
func (o *CustomerShippingContactsAddress) SetStreet1(v string) {
	o.Street1 = &v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *CustomerShippingContactsAddress) GetStreet2() string {
	if o == nil || IsNil(o.Street2) {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsAddress) GetStreet2Ok() (*string, bool) {
	if o == nil || IsNil(o.Street2) {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *CustomerShippingContactsAddress) HasStreet2() bool {
	if o != nil && !IsNil(o.Street2) {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *CustomerShippingContactsAddress) SetStreet2(v string) {
	o.Street2 = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CustomerShippingContactsAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CustomerShippingContactsAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CustomerShippingContactsAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CustomerShippingContactsAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerShippingContactsAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CustomerShippingContactsAddress) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CustomerShippingContactsAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CustomerShippingContactsAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CustomerShippingContactsAddress) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerShippingContactsAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerShippingContactsAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerShippingContactsAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerShippingContactsAddress) SetCountry(v string) {
	o.Country = &v
}

// GetResidential returns the Residential field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerShippingContactsAddress) GetResidential() bool {
	if o == nil || IsNil(o.Residential.Get()) {
		var ret bool
		return ret
	}
	return *o.Residential.Get()
}

// GetResidentialOk returns a tuple with the Residential field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerShippingContactsAddress) GetResidentialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Residential.Get(), o.Residential.IsSet()
}

// HasResidential returns a boolean if a field has been set.
func (o *CustomerShippingContactsAddress) HasResidential() bool {
	if o != nil && o.Residential.IsSet() {
		return true
	}

	return false
}

// SetResidential gets a reference to the given NullableBool and assigns it to the Residential field.
func (o *CustomerShippingContactsAddress) SetResidential(v bool) {
	o.Residential.Set(&v)
}

// SetResidentialNil sets the value for Residential to be an explicit nil
func (o *CustomerShippingContactsAddress) SetResidentialNil() {
	o.Residential.Set(nil)
}

// UnsetResidential ensures that no value is present for Residential, not even an explicit nil
func (o *CustomerShippingContactsAddress) UnsetResidential() {
	o.Residential.Unset()
}

func (o CustomerShippingContactsAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerShippingContactsAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Residential.IsSet() {
		toSerialize["residential"] = o.Residential.Get()
	}
	return toSerialize, nil
}

type NullableCustomerShippingContactsAddress struct {
	value *CustomerShippingContactsAddress
	isSet bool
}

func (v NullableCustomerShippingContactsAddress) Get() *CustomerShippingContactsAddress {
	return v.value
}

func (v *NullableCustomerShippingContactsAddress) Set(val *CustomerShippingContactsAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerShippingContactsAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerShippingContactsAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerShippingContactsAddress(val *CustomerShippingContactsAddress) *NullableCustomerShippingContactsAddress {
	return &NullableCustomerShippingContactsAddress{value: val, isSet: true}
}

func (v NullableCustomerShippingContactsAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerShippingContactsAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

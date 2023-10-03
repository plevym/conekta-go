package conekta

import (
	"encoding/json"
)

// checks if the CustomerAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerAddress{}

// CustomerAddress struct for CustomerAddress
type CustomerAddress struct {
	Street1    string  `json:"street1"`
	Street2    *string `json:"street2,omitempty"`
	PostalCode string  `json:"postal_code"`
	City       string  `json:"city"`
	State      *string `json:"state,omitempty"`
	// this field follows the [ISO 3166-1 alpha-2 standard](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Country        *string `json:"country,omitempty"`
	Residential    *bool   `json:"residential,omitempty"`
	ExternalNumber *string `json:"external_number,omitempty"`
}

// NewCustomerAddress instantiates a new CustomerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddress(street1 string, postalCode string, city string) *CustomerAddress {
	this := CustomerAddress{}
	this.Street1 = street1
	this.PostalCode = postalCode
	this.City = city
	return &this
}

// NewCustomerAddressWithDefaults instantiates a new CustomerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressWithDefaults() *CustomerAddress {
	this := CustomerAddress{}
	return &this
}

// GetStreet1 returns the Street1 field value
func (o *CustomerAddress) GetStreet1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street1
}

// GetStreet1Ok returns a tuple with the Street1 field value
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetStreet1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Street1, true
}

// SetStreet1 sets field value
func (o *CustomerAddress) SetStreet1(v string) {
	o.Street1 = v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *CustomerAddress) GetStreet2() string {
	if o == nil || IsNil(o.Street2) {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetStreet2Ok() (*string, bool) {
	if o == nil || IsNil(o.Street2) {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *CustomerAddress) HasStreet2() bool {
	if o != nil && !IsNil(o.Street2) {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *CustomerAddress) SetStreet2(v string) {
	o.Street2 = &v
}

// GetPostalCode returns the PostalCode field value
func (o *CustomerAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *CustomerAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCity returns the City field value
func (o *CustomerAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *CustomerAddress) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CustomerAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CustomerAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CustomerAddress) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerAddress) SetCountry(v string) {
	o.Country = &v
}

// GetResidential returns the Residential field value if set, zero value otherwise.
func (o *CustomerAddress) GetResidential() bool {
	if o == nil || IsNil(o.Residential) {
		var ret bool
		return ret
	}
	return *o.Residential
}

// GetResidentialOk returns a tuple with the Residential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetResidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Residential) {
		return nil, false
	}
	return o.Residential, true
}

// HasResidential returns a boolean if a field has been set.
func (o *CustomerAddress) HasResidential() bool {
	if o != nil && !IsNil(o.Residential) {
		return true
	}

	return false
}

// SetResidential gets a reference to the given bool and assigns it to the Residential field.
func (o *CustomerAddress) SetResidential(v bool) {
	o.Residential = &v
}

// GetExternalNumber returns the ExternalNumber field value if set, zero value otherwise.
func (o *CustomerAddress) GetExternalNumber() string {
	if o == nil || IsNil(o.ExternalNumber) {
		var ret string
		return ret
	}
	return *o.ExternalNumber
}

// GetExternalNumberOk returns a tuple with the ExternalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddress) GetExternalNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalNumber) {
		return nil, false
	}
	return o.ExternalNumber, true
}

// HasExternalNumber returns a boolean if a field has been set.
func (o *CustomerAddress) HasExternalNumber() bool {
	if o != nil && !IsNil(o.ExternalNumber) {
		return true
	}

	return false
}

// SetExternalNumber gets a reference to the given string and assigns it to the ExternalNumber field.
func (o *CustomerAddress) SetExternalNumber(v string) {
	o.ExternalNumber = &v
}

func (o CustomerAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["street1"] = o.Street1
	if !IsNil(o.Street2) {
		toSerialize["street2"] = o.Street2
	}
	toSerialize["postal_code"] = o.PostalCode
	toSerialize["city"] = o.City
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Residential) {
		toSerialize["residential"] = o.Residential
	}
	if !IsNil(o.ExternalNumber) {
		toSerialize["external_number"] = o.ExternalNumber
	}
	return toSerialize, nil
}

type NullableCustomerAddress struct {
	value *CustomerAddress
	isSet bool
}

func (v NullableCustomerAddress) Get() *CustomerAddress {
	return v.value
}

func (v *NullableCustomerAddress) Set(val *CustomerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddress(val *CustomerAddress) *NullableCustomerAddress {
	return &NullableCustomerAddress{value: val, isSet: true}
}

func (v NullableCustomerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

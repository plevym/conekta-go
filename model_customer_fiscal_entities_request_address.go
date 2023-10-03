package conekta

import (
	"encoding/json"
)

// checks if the CustomerFiscalEntitiesRequestAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerFiscalEntitiesRequestAddress{}

// CustomerFiscalEntitiesRequestAddress struct for CustomerFiscalEntitiesRequestAddress
type CustomerFiscalEntitiesRequestAddress struct {
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

// NewCustomerFiscalEntitiesRequestAddress instantiates a new CustomerFiscalEntitiesRequestAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFiscalEntitiesRequestAddress(street1 string, postalCode string, city string) *CustomerFiscalEntitiesRequestAddress {
	this := CustomerFiscalEntitiesRequestAddress{}
	this.Street1 = street1
	this.PostalCode = postalCode
	this.City = city
	return &this
}

// NewCustomerFiscalEntitiesRequestAddressWithDefaults instantiates a new CustomerFiscalEntitiesRequestAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFiscalEntitiesRequestAddressWithDefaults() *CustomerFiscalEntitiesRequestAddress {
	this := CustomerFiscalEntitiesRequestAddress{}
	return &this
}

// GetStreet1 returns the Street1 field value
func (o *CustomerFiscalEntitiesRequestAddress) GetStreet1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street1
}

// GetStreet1Ok returns a tuple with the Street1 field value
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequestAddress) GetStreet1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Street1, true
}

// SetStreet1 sets field value
func (o *CustomerFiscalEntitiesRequestAddress) SetStreet1(v string) {
	o.Street1 = v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequestAddress) GetStreet2() string {
	if o == nil || IsNil(o.Street2) {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequestAddress) GetStreet2Ok() (*string, bool) {
	if o == nil || IsNil(o.Street2) {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequestAddress) HasStreet2() bool {
	if o != nil && !IsNil(o.Street2) {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *CustomerFiscalEntitiesRequestAddress) SetStreet2(v string) {
	o.Street2 = &v
}

// GetPostalCode returns the PostalCode field value
func (o *CustomerFiscalEntitiesRequestAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequestAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *CustomerFiscalEntitiesRequestAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCity returns the City field value
func (o *CustomerFiscalEntitiesRequestAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequestAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *CustomerFiscalEntitiesRequestAddress) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequestAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequestAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequestAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CustomerFiscalEntitiesRequestAddress) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequestAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequestAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequestAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerFiscalEntitiesRequestAddress) SetCountry(v string) {
	o.Country = &v
}

// GetResidential returns the Residential field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequestAddress) GetResidential() bool {
	if o == nil || IsNil(o.Residential) {
		var ret bool
		return ret
	}
	return *o.Residential
}

// GetResidentialOk returns a tuple with the Residential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequestAddress) GetResidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Residential) {
		return nil, false
	}
	return o.Residential, true
}

// HasResidential returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequestAddress) HasResidential() bool {
	if o != nil && !IsNil(o.Residential) {
		return true
	}

	return false
}

// SetResidential gets a reference to the given bool and assigns it to the Residential field.
func (o *CustomerFiscalEntitiesRequestAddress) SetResidential(v bool) {
	o.Residential = &v
}

// GetExternalNumber returns the ExternalNumber field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequestAddress) GetExternalNumber() string {
	if o == nil || IsNil(o.ExternalNumber) {
		var ret string
		return ret
	}
	return *o.ExternalNumber
}

// GetExternalNumberOk returns a tuple with the ExternalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequestAddress) GetExternalNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalNumber) {
		return nil, false
	}
	return o.ExternalNumber, true
}

// HasExternalNumber returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequestAddress) HasExternalNumber() bool {
	if o != nil && !IsNil(o.ExternalNumber) {
		return true
	}

	return false
}

// SetExternalNumber gets a reference to the given string and assigns it to the ExternalNumber field.
func (o *CustomerFiscalEntitiesRequestAddress) SetExternalNumber(v string) {
	o.ExternalNumber = &v
}

func (o CustomerFiscalEntitiesRequestAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerFiscalEntitiesRequestAddress) ToMap() (map[string]interface{}, error) {
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

type NullableCustomerFiscalEntitiesRequestAddress struct {
	value *CustomerFiscalEntitiesRequestAddress
	isSet bool
}

func (v NullableCustomerFiscalEntitiesRequestAddress) Get() *CustomerFiscalEntitiesRequestAddress {
	return v.value
}

func (v *NullableCustomerFiscalEntitiesRequestAddress) Set(val *CustomerFiscalEntitiesRequestAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFiscalEntitiesRequestAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFiscalEntitiesRequestAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFiscalEntitiesRequestAddress(val *CustomerFiscalEntitiesRequestAddress) *NullableCustomerFiscalEntitiesRequestAddress {
	return &NullableCustomerFiscalEntitiesRequestAddress{value: val, isSet: true}
}

func (v NullableCustomerFiscalEntitiesRequestAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFiscalEntitiesRequestAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

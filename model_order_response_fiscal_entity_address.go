package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseFiscalEntityAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseFiscalEntityAddress{}

// OrderResponseFiscalEntityAddress struct for OrderResponseFiscalEntityAddress
type OrderResponseFiscalEntityAddress struct {
	Street1    string  `json:"street1"`
	Street2    *string `json:"street2,omitempty"`
	PostalCode string  `json:"postal_code"`
	City       string  `json:"city"`
	State      *string `json:"state,omitempty"`
	// this field follows the [ISO 3166-1 alpha-2 standard](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Country        *string `json:"country,omitempty"`
	Residential    *bool   `json:"residential,omitempty"`
	ExternalNumber *string `json:"external_number,omitempty"`
	Object         *string `json:"object,omitempty"`
}

// NewOrderResponseFiscalEntityAddress instantiates a new OrderResponseFiscalEntityAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseFiscalEntityAddress(street1 string, postalCode string, city string) *OrderResponseFiscalEntityAddress {
	this := OrderResponseFiscalEntityAddress{}
	this.Street1 = street1
	this.PostalCode = postalCode
	this.City = city
	return &this
}

// NewOrderResponseFiscalEntityAddressWithDefaults instantiates a new OrderResponseFiscalEntityAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseFiscalEntityAddressWithDefaults() *OrderResponseFiscalEntityAddress {
	this := OrderResponseFiscalEntityAddress{}
	return &this
}

// GetStreet1 returns the Street1 field value
func (o *OrderResponseFiscalEntityAddress) GetStreet1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street1
}

// GetStreet1Ok returns a tuple with the Street1 field value
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddress) GetStreet1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Street1, true
}

// SetStreet1 sets field value
func (o *OrderResponseFiscalEntityAddress) SetStreet1(v string) {
	o.Street1 = v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntityAddress) GetStreet2() string {
	if o == nil || IsNil(o.Street2) {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddress) GetStreet2Ok() (*string, bool) {
	if o == nil || IsNil(o.Street2) {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntityAddress) HasStreet2() bool {
	if o != nil && !IsNil(o.Street2) {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *OrderResponseFiscalEntityAddress) SetStreet2(v string) {
	o.Street2 = &v
}

// GetPostalCode returns the PostalCode field value
func (o *OrderResponseFiscalEntityAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *OrderResponseFiscalEntityAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCity returns the City field value
func (o *OrderResponseFiscalEntityAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *OrderResponseFiscalEntityAddress) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntityAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntityAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrderResponseFiscalEntityAddress) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntityAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntityAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *OrderResponseFiscalEntityAddress) SetCountry(v string) {
	o.Country = &v
}

// GetResidential returns the Residential field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntityAddress) GetResidential() bool {
	if o == nil || IsNil(o.Residential) {
		var ret bool
		return ret
	}
	return *o.Residential
}

// GetResidentialOk returns a tuple with the Residential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddress) GetResidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Residential) {
		return nil, false
	}
	return o.Residential, true
}

// HasResidential returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntityAddress) HasResidential() bool {
	if o != nil && !IsNil(o.Residential) {
		return true
	}

	return false
}

// SetResidential gets a reference to the given bool and assigns it to the Residential field.
func (o *OrderResponseFiscalEntityAddress) SetResidential(v bool) {
	o.Residential = &v
}

// GetExternalNumber returns the ExternalNumber field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntityAddress) GetExternalNumber() string {
	if o == nil || IsNil(o.ExternalNumber) {
		var ret string
		return ret
	}
	return *o.ExternalNumber
}

// GetExternalNumberOk returns a tuple with the ExternalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddress) GetExternalNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalNumber) {
		return nil, false
	}
	return o.ExternalNumber, true
}

// HasExternalNumber returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntityAddress) HasExternalNumber() bool {
	if o != nil && !IsNil(o.ExternalNumber) {
		return true
	}

	return false
}

// SetExternalNumber gets a reference to the given string and assigns it to the ExternalNumber field.
func (o *OrderResponseFiscalEntityAddress) SetExternalNumber(v string) {
	o.ExternalNumber = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderResponseFiscalEntityAddress) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseFiscalEntityAddress) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderResponseFiscalEntityAddress) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderResponseFiscalEntityAddress) SetObject(v string) {
	o.Object = &v
}

func (o OrderResponseFiscalEntityAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseFiscalEntityAddress) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableOrderResponseFiscalEntityAddress struct {
	value *OrderResponseFiscalEntityAddress
	isSet bool
}

func (v NullableOrderResponseFiscalEntityAddress) Get() *OrderResponseFiscalEntityAddress {
	return v.value
}

func (v *NullableOrderResponseFiscalEntityAddress) Set(val *OrderResponseFiscalEntityAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseFiscalEntityAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseFiscalEntityAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseFiscalEntityAddress(val *OrderResponseFiscalEntityAddress) *NullableOrderResponseFiscalEntityAddress {
	return &NullableOrderResponseFiscalEntityAddress{value: val, isSet: true}
}

func (v NullableOrderResponseFiscalEntityAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseFiscalEntityAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

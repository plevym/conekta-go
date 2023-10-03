package conekta

import (
	"encoding/json"
)

// checks if the CompanyFiscalInfoAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyFiscalInfoAddressResponse{}

// CompanyFiscalInfoAddressResponse Company fiscal info address model
type CompanyFiscalInfoAddressResponse struct {
	// The resource's type
	Object *string `json:"object,omitempty"`
	// Street Address
	Street1 *string `json:"street1,omitempty"`
	// Colonia
	Street2 *string `json:"street2,omitempty"`
	// City
	City *string `json:"city,omitempty"`
	// State
	State *string `json:"state,omitempty"`
	// Country
	Country *string `json:"country,omitempty"`
	// Postal code
	PostalCode *string `json:"postal_code,omitempty"`
	// Street number
	ExternalNumber *string `json:"external_number,omitempty"`
	// Unit / apartment number
	InternalNumber *string `json:"internal_number,omitempty"`
}

// NewCompanyFiscalInfoAddressResponse instantiates a new CompanyFiscalInfoAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyFiscalInfoAddressResponse() *CompanyFiscalInfoAddressResponse {
	this := CompanyFiscalInfoAddressResponse{}
	return &this
}

// NewCompanyFiscalInfoAddressResponseWithDefaults instantiates a new CompanyFiscalInfoAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyFiscalInfoAddressResponseWithDefaults() *CompanyFiscalInfoAddressResponse {
	this := CompanyFiscalInfoAddressResponse{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CompanyFiscalInfoAddressResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoAddressResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CompanyFiscalInfoAddressResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CompanyFiscalInfoAddressResponse) SetObject(v string) {
	o.Object = &v
}

// GetStreet1 returns the Street1 field value if set, zero value otherwise.
func (o *CompanyFiscalInfoAddressResponse) GetStreet1() string {
	if o == nil || IsNil(o.Street1) {
		var ret string
		return ret
	}
	return *o.Street1
}

// GetStreet1Ok returns a tuple with the Street1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoAddressResponse) GetStreet1Ok() (*string, bool) {
	if o == nil || IsNil(o.Street1) {
		return nil, false
	}
	return o.Street1, true
}

// HasStreet1 returns a boolean if a field has been set.
func (o *CompanyFiscalInfoAddressResponse) HasStreet1() bool {
	if o != nil && !IsNil(o.Street1) {
		return true
	}

	return false
}

// SetStreet1 gets a reference to the given string and assigns it to the Street1 field.
func (o *CompanyFiscalInfoAddressResponse) SetStreet1(v string) {
	o.Street1 = &v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *CompanyFiscalInfoAddressResponse) GetStreet2() string {
	if o == nil || IsNil(o.Street2) {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoAddressResponse) GetStreet2Ok() (*string, bool) {
	if o == nil || IsNil(o.Street2) {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *CompanyFiscalInfoAddressResponse) HasStreet2() bool {
	if o != nil && !IsNil(o.Street2) {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *CompanyFiscalInfoAddressResponse) SetStreet2(v string) {
	o.Street2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CompanyFiscalInfoAddressResponse) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoAddressResponse) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CompanyFiscalInfoAddressResponse) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CompanyFiscalInfoAddressResponse) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CompanyFiscalInfoAddressResponse) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoAddressResponse) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CompanyFiscalInfoAddressResponse) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CompanyFiscalInfoAddressResponse) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CompanyFiscalInfoAddressResponse) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoAddressResponse) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CompanyFiscalInfoAddressResponse) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CompanyFiscalInfoAddressResponse) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CompanyFiscalInfoAddressResponse) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoAddressResponse) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CompanyFiscalInfoAddressResponse) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CompanyFiscalInfoAddressResponse) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetExternalNumber returns the ExternalNumber field value if set, zero value otherwise.
func (o *CompanyFiscalInfoAddressResponse) GetExternalNumber() string {
	if o == nil || IsNil(o.ExternalNumber) {
		var ret string
		return ret
	}
	return *o.ExternalNumber
}

// GetExternalNumberOk returns a tuple with the ExternalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoAddressResponse) GetExternalNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalNumber) {
		return nil, false
	}
	return o.ExternalNumber, true
}

// HasExternalNumber returns a boolean if a field has been set.
func (o *CompanyFiscalInfoAddressResponse) HasExternalNumber() bool {
	if o != nil && !IsNil(o.ExternalNumber) {
		return true
	}

	return false
}

// SetExternalNumber gets a reference to the given string and assigns it to the ExternalNumber field.
func (o *CompanyFiscalInfoAddressResponse) SetExternalNumber(v string) {
	o.ExternalNumber = &v
}

// GetInternalNumber returns the InternalNumber field value if set, zero value otherwise.
func (o *CompanyFiscalInfoAddressResponse) GetInternalNumber() string {
	if o == nil || IsNil(o.InternalNumber) {
		var ret string
		return ret
	}
	return *o.InternalNumber
}

// GetInternalNumberOk returns a tuple with the InternalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoAddressResponse) GetInternalNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InternalNumber) {
		return nil, false
	}
	return o.InternalNumber, true
}

// HasInternalNumber returns a boolean if a field has been set.
func (o *CompanyFiscalInfoAddressResponse) HasInternalNumber() bool {
	if o != nil && !IsNil(o.InternalNumber) {
		return true
	}

	return false
}

// SetInternalNumber gets a reference to the given string and assigns it to the InternalNumber field.
func (o *CompanyFiscalInfoAddressResponse) SetInternalNumber(v string) {
	o.InternalNumber = &v
}

func (o CompanyFiscalInfoAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyFiscalInfoAddressResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.ExternalNumber) {
		toSerialize["external_number"] = o.ExternalNumber
	}
	if !IsNil(o.InternalNumber) {
		toSerialize["internal_number"] = o.InternalNumber
	}
	return toSerialize, nil
}

type NullableCompanyFiscalInfoAddressResponse struct {
	value *CompanyFiscalInfoAddressResponse
	isSet bool
}

func (v NullableCompanyFiscalInfoAddressResponse) Get() *CompanyFiscalInfoAddressResponse {
	return v.value
}

func (v *NullableCompanyFiscalInfoAddressResponse) Set(val *CompanyFiscalInfoAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyFiscalInfoAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyFiscalInfoAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyFiscalInfoAddressResponse(val *CompanyFiscalInfoAddressResponse) *NullableCompanyFiscalInfoAddressResponse {
	return &NullableCompanyFiscalInfoAddressResponse{value: val, isSet: true}
}

func (v NullableCompanyFiscalInfoAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyFiscalInfoAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the CompanyFiscalInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyFiscalInfoResponse{}

// CompanyFiscalInfoResponse Company fiscal info model
type CompanyFiscalInfoResponse struct {
	// The resource's type
	Object *string `json:"object,omitempty"`
	// Tax ID of the company
	TaxId *string `json:"tax_id,omitempty"`
	// Legal name of the company
	LegalEntityName *string `json:"legal_entity_name,omitempty"`
	// Business type of the company
	BusinessType *string `json:"business_type,omitempty"`
	// Phone number of the company
	Phone *string `json:"phone,omitempty"`
	// Business type if 'persona_fisica'
	PhysicalPersonBusinessType *string                           `json:"physical_person_business_type,omitempty"`
	Address                    *CompanyFiscalInfoAddressResponse `json:"address,omitempty"`
}

// NewCompanyFiscalInfoResponse instantiates a new CompanyFiscalInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyFiscalInfoResponse() *CompanyFiscalInfoResponse {
	this := CompanyFiscalInfoResponse{}
	return &this
}

// NewCompanyFiscalInfoResponseWithDefaults instantiates a new CompanyFiscalInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyFiscalInfoResponseWithDefaults() *CompanyFiscalInfoResponse {
	this := CompanyFiscalInfoResponse{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CompanyFiscalInfoResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CompanyFiscalInfoResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CompanyFiscalInfoResponse) SetObject(v string) {
	o.Object = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *CompanyFiscalInfoResponse) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoResponse) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *CompanyFiscalInfoResponse) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *CompanyFiscalInfoResponse) SetTaxId(v string) {
	o.TaxId = &v
}

// GetLegalEntityName returns the LegalEntityName field value if set, zero value otherwise.
func (o *CompanyFiscalInfoResponse) GetLegalEntityName() string {
	if o == nil || IsNil(o.LegalEntityName) {
		var ret string
		return ret
	}
	return *o.LegalEntityName
}

// GetLegalEntityNameOk returns a tuple with the LegalEntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoResponse) GetLegalEntityNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalEntityName) {
		return nil, false
	}
	return o.LegalEntityName, true
}

// HasLegalEntityName returns a boolean if a field has been set.
func (o *CompanyFiscalInfoResponse) HasLegalEntityName() bool {
	if o != nil && !IsNil(o.LegalEntityName) {
		return true
	}

	return false
}

// SetLegalEntityName gets a reference to the given string and assigns it to the LegalEntityName field.
func (o *CompanyFiscalInfoResponse) SetLegalEntityName(v string) {
	o.LegalEntityName = &v
}

// GetBusinessType returns the BusinessType field value if set, zero value otherwise.
func (o *CompanyFiscalInfoResponse) GetBusinessType() string {
	if o == nil || IsNil(o.BusinessType) {
		var ret string
		return ret
	}
	return *o.BusinessType
}

// GetBusinessTypeOk returns a tuple with the BusinessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoResponse) GetBusinessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessType) {
		return nil, false
	}
	return o.BusinessType, true
}

// HasBusinessType returns a boolean if a field has been set.
func (o *CompanyFiscalInfoResponse) HasBusinessType() bool {
	if o != nil && !IsNil(o.BusinessType) {
		return true
	}

	return false
}

// SetBusinessType gets a reference to the given string and assigns it to the BusinessType field.
func (o *CompanyFiscalInfoResponse) SetBusinessType(v string) {
	o.BusinessType = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CompanyFiscalInfoResponse) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoResponse) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CompanyFiscalInfoResponse) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CompanyFiscalInfoResponse) SetPhone(v string) {
	o.Phone = &v
}

// GetPhysicalPersonBusinessType returns the PhysicalPersonBusinessType field value if set, zero value otherwise.
func (o *CompanyFiscalInfoResponse) GetPhysicalPersonBusinessType() string {
	if o == nil || IsNil(o.PhysicalPersonBusinessType) {
		var ret string
		return ret
	}
	return *o.PhysicalPersonBusinessType
}

// GetPhysicalPersonBusinessTypeOk returns a tuple with the PhysicalPersonBusinessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoResponse) GetPhysicalPersonBusinessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalPersonBusinessType) {
		return nil, false
	}
	return o.PhysicalPersonBusinessType, true
}

// HasPhysicalPersonBusinessType returns a boolean if a field has been set.
func (o *CompanyFiscalInfoResponse) HasPhysicalPersonBusinessType() bool {
	if o != nil && !IsNil(o.PhysicalPersonBusinessType) {
		return true
	}

	return false
}

// SetPhysicalPersonBusinessType gets a reference to the given string and assigns it to the PhysicalPersonBusinessType field.
func (o *CompanyFiscalInfoResponse) SetPhysicalPersonBusinessType(v string) {
	o.PhysicalPersonBusinessType = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CompanyFiscalInfoResponse) GetAddress() CompanyFiscalInfoAddressResponse {
	if o == nil || IsNil(o.Address) {
		var ret CompanyFiscalInfoAddressResponse
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyFiscalInfoResponse) GetAddressOk() (*CompanyFiscalInfoAddressResponse, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CompanyFiscalInfoResponse) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CompanyFiscalInfoAddressResponse and assigns it to the Address field.
func (o *CompanyFiscalInfoResponse) SetAddress(v CompanyFiscalInfoAddressResponse) {
	o.Address = &v
}

func (o CompanyFiscalInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyFiscalInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	if !IsNil(o.LegalEntityName) {
		toSerialize["legal_entity_name"] = o.LegalEntityName
	}
	if !IsNil(o.BusinessType) {
		toSerialize["business_type"] = o.BusinessType
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.PhysicalPersonBusinessType) {
		toSerialize["physical_person_business_type"] = o.PhysicalPersonBusinessType
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableCompanyFiscalInfoResponse struct {
	value *CompanyFiscalInfoResponse
	isSet bool
}

func (v NullableCompanyFiscalInfoResponse) Get() *CompanyFiscalInfoResponse {
	return v.value
}

func (v *NullableCompanyFiscalInfoResponse) Set(val *CompanyFiscalInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyFiscalInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyFiscalInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyFiscalInfoResponse(val *CompanyFiscalInfoResponse) *NullableCompanyFiscalInfoResponse {
	return &NullableCompanyFiscalInfoResponse{value: val, isSet: true}
}

func (v NullableCompanyFiscalInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyFiscalInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the CustomerUpdateFiscalEntitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerUpdateFiscalEntitiesRequest{}

// CustomerUpdateFiscalEntitiesRequest struct for CustomerUpdateFiscalEntitiesRequest
type CustomerUpdateFiscalEntitiesRequest struct {
	Address     *CustomerFiscalEntitiesRequestAddress `json:"address,omitempty"`
	TaxId       *string                               `json:"tax_id,omitempty"`
	Email       *string                               `json:"email,omitempty"`
	Phone       *string                               `json:"phone,omitempty"`
	Metadata    map[string]map[string]interface{}     `json:"metadata,omitempty"`
	CompanyName *string                               `json:"company_name,omitempty"`
}

// NewCustomerUpdateFiscalEntitiesRequest instantiates a new CustomerUpdateFiscalEntitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerUpdateFiscalEntitiesRequest() *CustomerUpdateFiscalEntitiesRequest {
	this := CustomerUpdateFiscalEntitiesRequest{}
	return &this
}

// NewCustomerUpdateFiscalEntitiesRequestWithDefaults instantiates a new CustomerUpdateFiscalEntitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerUpdateFiscalEntitiesRequestWithDefaults() *CustomerUpdateFiscalEntitiesRequest {
	this := CustomerUpdateFiscalEntitiesRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CustomerUpdateFiscalEntitiesRequest) GetAddress() CustomerFiscalEntitiesRequestAddress {
	if o == nil || IsNil(o.Address) {
		var ret CustomerFiscalEntitiesRequestAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) GetAddressOk() (*CustomerFiscalEntitiesRequestAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CustomerFiscalEntitiesRequestAddress and assigns it to the Address field.
func (o *CustomerUpdateFiscalEntitiesRequest) SetAddress(v CustomerFiscalEntitiesRequestAddress) {
	o.Address = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *CustomerUpdateFiscalEntitiesRequest) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *CustomerUpdateFiscalEntitiesRequest) SetTaxId(v string) {
	o.TaxId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerUpdateFiscalEntitiesRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerUpdateFiscalEntitiesRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerUpdateFiscalEntitiesRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerUpdateFiscalEntitiesRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomerUpdateFiscalEntitiesRequest) GetMetadata() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) GetMetadataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *CustomerUpdateFiscalEntitiesRequest) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *CustomerUpdateFiscalEntitiesRequest) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *CustomerUpdateFiscalEntitiesRequest) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *CustomerUpdateFiscalEntitiesRequest) SetCompanyName(v string) {
	o.CompanyName = &v
}

func (o CustomerUpdateFiscalEntitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerUpdateFiscalEntitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	return toSerialize, nil
}

type NullableCustomerUpdateFiscalEntitiesRequest struct {
	value *CustomerUpdateFiscalEntitiesRequest
	isSet bool
}

func (v NullableCustomerUpdateFiscalEntitiesRequest) Get() *CustomerUpdateFiscalEntitiesRequest {
	return v.value
}

func (v *NullableCustomerUpdateFiscalEntitiesRequest) Set(val *CustomerUpdateFiscalEntitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerUpdateFiscalEntitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerUpdateFiscalEntitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerUpdateFiscalEntitiesRequest(val *CustomerUpdateFiscalEntitiesRequest) *NullableCustomerUpdateFiscalEntitiesRequest {
	return &NullableCustomerUpdateFiscalEntitiesRequest{value: val, isSet: true}
}

func (v NullableCustomerUpdateFiscalEntitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerUpdateFiscalEntitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

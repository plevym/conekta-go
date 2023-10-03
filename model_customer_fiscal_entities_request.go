package conekta

import (
	"encoding/json"
)

// checks if the CustomerFiscalEntitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerFiscalEntitiesRequest{}

// CustomerFiscalEntitiesRequest struct for CustomerFiscalEntitiesRequest
type CustomerFiscalEntitiesRequest struct {
	Address     CustomerFiscalEntitiesRequestAddress `json:"address"`
	TaxId       *string                              `json:"tax_id,omitempty"`
	Email       *string                              `json:"email,omitempty"`
	Phone       *string                              `json:"phone,omitempty"`
	Metadata    map[string]map[string]interface{}    `json:"metadata,omitempty"`
	CompanyName *string                              `json:"company_name,omitempty"`
}

// NewCustomerFiscalEntitiesRequest instantiates a new CustomerFiscalEntitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFiscalEntitiesRequest(address CustomerFiscalEntitiesRequestAddress) *CustomerFiscalEntitiesRequest {
	this := CustomerFiscalEntitiesRequest{}
	this.Address = address
	return &this
}

// NewCustomerFiscalEntitiesRequestWithDefaults instantiates a new CustomerFiscalEntitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFiscalEntitiesRequestWithDefaults() *CustomerFiscalEntitiesRequest {
	this := CustomerFiscalEntitiesRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *CustomerFiscalEntitiesRequest) GetAddress() CustomerFiscalEntitiesRequestAddress {
	if o == nil {
		var ret CustomerFiscalEntitiesRequestAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequest) GetAddressOk() (*CustomerFiscalEntitiesRequestAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CustomerFiscalEntitiesRequest) SetAddress(v CustomerFiscalEntitiesRequestAddress) {
	o.Address = v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequest) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequest) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequest) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *CustomerFiscalEntitiesRequest) SetTaxId(v string) {
	o.TaxId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerFiscalEntitiesRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequest) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerFiscalEntitiesRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequest) GetMetadata() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequest) GetMetadataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *CustomerFiscalEntitiesRequest) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *CustomerFiscalEntitiesRequest) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFiscalEntitiesRequest) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *CustomerFiscalEntitiesRequest) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *CustomerFiscalEntitiesRequest) SetCompanyName(v string) {
	o.CompanyName = &v
}

func (o CustomerFiscalEntitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerFiscalEntitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
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

type NullableCustomerFiscalEntitiesRequest struct {
	value *CustomerFiscalEntitiesRequest
	isSet bool
}

func (v NullableCustomerFiscalEntitiesRequest) Get() *CustomerFiscalEntitiesRequest {
	return v.value
}

func (v *NullableCustomerFiscalEntitiesRequest) Set(val *CustomerFiscalEntitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFiscalEntitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFiscalEntitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFiscalEntitiesRequest(val *CustomerFiscalEntitiesRequest) *NullableCustomerFiscalEntitiesRequest {
	return &NullableCustomerFiscalEntitiesRequest{value: val, isSet: true}
}

func (v NullableCustomerFiscalEntitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFiscalEntitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

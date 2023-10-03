package conekta

import (
	"encoding/json"
)

// checks if the CreateCustomerFiscalEntitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomerFiscalEntitiesResponse{}

// CreateCustomerFiscalEntitiesResponse struct for CreateCustomerFiscalEntitiesResponse
type CreateCustomerFiscalEntitiesResponse struct {
	Address     CustomerFiscalEntitiesRequestAddress `json:"address"`
	TaxId       *string                              `json:"tax_id,omitempty"`
	Email       *string                              `json:"email,omitempty"`
	Phone       *string                              `json:"phone,omitempty"`
	Metadata    map[string]map[string]interface{}    `json:"metadata,omitempty"`
	CompanyName *string                              `json:"company_name,omitempty"`
	Id          string                               `json:"id"`
	Object      string                               `json:"object"`
	CreatedAt   int64                                `json:"created_at"`
	ParentId    *string                              `json:"parent_id,omitempty"`
	Default     *bool                                `json:"default,omitempty"`
}

// NewCreateCustomerFiscalEntitiesResponse instantiates a new CreateCustomerFiscalEntitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomerFiscalEntitiesResponse(address CustomerFiscalEntitiesRequestAddress, id string, object string, createdAt int64) *CreateCustomerFiscalEntitiesResponse {
	this := CreateCustomerFiscalEntitiesResponse{}
	this.Address = address
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	return &this
}

// NewCreateCustomerFiscalEntitiesResponseWithDefaults instantiates a new CreateCustomerFiscalEntitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomerFiscalEntitiesResponseWithDefaults() *CreateCustomerFiscalEntitiesResponse {
	this := CreateCustomerFiscalEntitiesResponse{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateCustomerFiscalEntitiesResponse) GetAddress() CustomerFiscalEntitiesRequestAddress {
	if o == nil {
		var ret CustomerFiscalEntitiesRequestAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetAddressOk() (*CustomerFiscalEntitiesRequestAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateCustomerFiscalEntitiesResponse) SetAddress(v CustomerFiscalEntitiesRequestAddress) {
	o.Address = v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *CreateCustomerFiscalEntitiesResponse) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *CreateCustomerFiscalEntitiesResponse) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *CreateCustomerFiscalEntitiesResponse) SetTaxId(v string) {
	o.TaxId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateCustomerFiscalEntitiesResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateCustomerFiscalEntitiesResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreateCustomerFiscalEntitiesResponse) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CreateCustomerFiscalEntitiesResponse) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CreateCustomerFiscalEntitiesResponse) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CreateCustomerFiscalEntitiesResponse) SetPhone(v string) {
	o.Phone = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateCustomerFiscalEntitiesResponse) GetMetadata() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetMetadataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateCustomerFiscalEntitiesResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *CreateCustomerFiscalEntitiesResponse) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *CreateCustomerFiscalEntitiesResponse) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *CreateCustomerFiscalEntitiesResponse) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *CreateCustomerFiscalEntitiesResponse) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetId returns the Id field value
func (o *CreateCustomerFiscalEntitiesResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateCustomerFiscalEntitiesResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *CreateCustomerFiscalEntitiesResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CreateCustomerFiscalEntitiesResponse) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreateCustomerFiscalEntitiesResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreateCustomerFiscalEntitiesResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CreateCustomerFiscalEntitiesResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CreateCustomerFiscalEntitiesResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *CreateCustomerFiscalEntitiesResponse) SetParentId(v string) {
	o.ParentId = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CreateCustomerFiscalEntitiesResponse) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponse) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CreateCustomerFiscalEntitiesResponse) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CreateCustomerFiscalEntitiesResponse) SetDefault(v bool) {
	o.Default = &v
}

func (o CreateCustomerFiscalEntitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomerFiscalEntitiesResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

type NullableCreateCustomerFiscalEntitiesResponse struct {
	value *CreateCustomerFiscalEntitiesResponse
	isSet bool
}

func (v NullableCreateCustomerFiscalEntitiesResponse) Get() *CreateCustomerFiscalEntitiesResponse {
	return v.value
}

func (v *NullableCreateCustomerFiscalEntitiesResponse) Set(val *CreateCustomerFiscalEntitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerFiscalEntitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerFiscalEntitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerFiscalEntitiesResponse(val *CreateCustomerFiscalEntitiesResponse) *NullableCreateCustomerFiscalEntitiesResponse {
	return &NullableCreateCustomerFiscalEntitiesResponse{value: val, isSet: true}
}

func (v NullableCreateCustomerFiscalEntitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerFiscalEntitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

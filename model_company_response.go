package conekta

import (
	"encoding/json"
)

// checks if the CompanyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyResponse{}

// CompanyResponse Company model
type CompanyResponse struct {
	// The child company's unique identifier
	Id *string `json:"id,omitempty"`
	// The resource's creation date (unix timestamp)
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The child company's name
	Name *string `json:"name,omitempty"`
	// The resource's type
	Object *string `json:"object,omitempty"`
	// Id of the parent company
	ParentCompanyId *string `json:"parent_company_id,omitempty"`
	// Whether the parent company's fiscal data is to be used for liquidation and tax purposes
	UseParentFiscalData *bool                             `json:"use_parent_fiscal_data,omitempty"`
	PayoutDestination   *CompanyPayoutDestinationResponse `json:"payout_destination,omitempty"`
	FiscalInfo          *CompanyFiscalInfoResponse        `json:"fiscal_info,omitempty"`
}

// NewCompanyResponse instantiates a new CompanyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyResponse() *CompanyResponse {
	this := CompanyResponse{}
	return &this
}

// NewCompanyResponseWithDefaults instantiates a new CompanyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyResponseWithDefaults() *CompanyResponse {
	this := CompanyResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompanyResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompanyResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompanyResponse) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CompanyResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CompanyResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *CompanyResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompanyResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompanyResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompanyResponse) SetName(v string) {
	o.Name = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CompanyResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CompanyResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CompanyResponse) SetObject(v string) {
	o.Object = &v
}

// GetParentCompanyId returns the ParentCompanyId field value if set, zero value otherwise.
func (o *CompanyResponse) GetParentCompanyId() string {
	if o == nil || IsNil(o.ParentCompanyId) {
		var ret string
		return ret
	}
	return *o.ParentCompanyId
}

// GetParentCompanyIdOk returns a tuple with the ParentCompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetParentCompanyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCompanyId) {
		return nil, false
	}
	return o.ParentCompanyId, true
}

// HasParentCompanyId returns a boolean if a field has been set.
func (o *CompanyResponse) HasParentCompanyId() bool {
	if o != nil && !IsNil(o.ParentCompanyId) {
		return true
	}

	return false
}

// SetParentCompanyId gets a reference to the given string and assigns it to the ParentCompanyId field.
func (o *CompanyResponse) SetParentCompanyId(v string) {
	o.ParentCompanyId = &v
}

// GetUseParentFiscalData returns the UseParentFiscalData field value if set, zero value otherwise.
func (o *CompanyResponse) GetUseParentFiscalData() bool {
	if o == nil || IsNil(o.UseParentFiscalData) {
		var ret bool
		return ret
	}
	return *o.UseParentFiscalData
}

// GetUseParentFiscalDataOk returns a tuple with the UseParentFiscalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetUseParentFiscalDataOk() (*bool, bool) {
	if o == nil || IsNil(o.UseParentFiscalData) {
		return nil, false
	}
	return o.UseParentFiscalData, true
}

// HasUseParentFiscalData returns a boolean if a field has been set.
func (o *CompanyResponse) HasUseParentFiscalData() bool {
	if o != nil && !IsNil(o.UseParentFiscalData) {
		return true
	}

	return false
}

// SetUseParentFiscalData gets a reference to the given bool and assigns it to the UseParentFiscalData field.
func (o *CompanyResponse) SetUseParentFiscalData(v bool) {
	o.UseParentFiscalData = &v
}

// GetPayoutDestination returns the PayoutDestination field value if set, zero value otherwise.
func (o *CompanyResponse) GetPayoutDestination() CompanyPayoutDestinationResponse {
	if o == nil || IsNil(o.PayoutDestination) {
		var ret CompanyPayoutDestinationResponse
		return ret
	}
	return *o.PayoutDestination
}

// GetPayoutDestinationOk returns a tuple with the PayoutDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetPayoutDestinationOk() (*CompanyPayoutDestinationResponse, bool) {
	if o == nil || IsNil(o.PayoutDestination) {
		return nil, false
	}
	return o.PayoutDestination, true
}

// HasPayoutDestination returns a boolean if a field has been set.
func (o *CompanyResponse) HasPayoutDestination() bool {
	if o != nil && !IsNil(o.PayoutDestination) {
		return true
	}

	return false
}

// SetPayoutDestination gets a reference to the given CompanyPayoutDestinationResponse and assigns it to the PayoutDestination field.
func (o *CompanyResponse) SetPayoutDestination(v CompanyPayoutDestinationResponse) {
	o.PayoutDestination = &v
}

// GetFiscalInfo returns the FiscalInfo field value if set, zero value otherwise.
func (o *CompanyResponse) GetFiscalInfo() CompanyFiscalInfoResponse {
	if o == nil || IsNil(o.FiscalInfo) {
		var ret CompanyFiscalInfoResponse
		return ret
	}
	return *o.FiscalInfo
}

// GetFiscalInfoOk returns a tuple with the FiscalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyResponse) GetFiscalInfoOk() (*CompanyFiscalInfoResponse, bool) {
	if o == nil || IsNil(o.FiscalInfo) {
		return nil, false
	}
	return o.FiscalInfo, true
}

// HasFiscalInfo returns a boolean if a field has been set.
func (o *CompanyResponse) HasFiscalInfo() bool {
	if o != nil && !IsNil(o.FiscalInfo) {
		return true
	}

	return false
}

// SetFiscalInfo gets a reference to the given CompanyFiscalInfoResponse and assigns it to the FiscalInfo field.
func (o *CompanyResponse) SetFiscalInfo(v CompanyFiscalInfoResponse) {
	o.FiscalInfo = &v
}

func (o CompanyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.ParentCompanyId) {
		toSerialize["parent_company_id"] = o.ParentCompanyId
	}
	if !IsNil(o.UseParentFiscalData) {
		toSerialize["use_parent_fiscal_data"] = o.UseParentFiscalData
	}
	if !IsNil(o.PayoutDestination) {
		toSerialize["payout_destination"] = o.PayoutDestination
	}
	if !IsNil(o.FiscalInfo) {
		toSerialize["fiscal_info"] = o.FiscalInfo
	}
	return toSerialize, nil
}

type NullableCompanyResponse struct {
	value *CompanyResponse
	isSet bool
}

func (v NullableCompanyResponse) Get() *CompanyResponse {
	return v.value
}

func (v *NullableCompanyResponse) Set(val *CompanyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyResponse(val *CompanyResponse) *NullableCompanyResponse {
	return &NullableCompanyResponse{value: val, isSet: true}
}

func (v NullableCompanyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the CustomerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerResponse{}

// CustomerResponse customer response
type CustomerResponse struct {
	AntifraudInfo            NullableCustomerAntifraudInfoResponse `json:"antifraud_info,omitempty"`
	Corporate                *bool                                 `json:"corporate,omitempty"`
	CreatedAt                int64                                 `json:"created_at"`
	CustomReference          *string                               `json:"custom_reference,omitempty"`
	DefaultFiscalEntityId    NullableString                        `json:"default_fiscal_entity_id,omitempty"`
	DefaultShippingContactId *string                               `json:"default_shipping_contact_id,omitempty"`
	DefaultPaymentSourceId   NullableString                        `json:"default_payment_source_id,omitempty"`
	Email                    *string                               `json:"email,omitempty"`
	FiscalEntities           *CustomerFiscalEntitiesResponse       `json:"fiscal_entities,omitempty"`
	Id                       string                                `json:"id"`
	Livemode                 bool                                  `json:"livemode"`
	Name                     *string                               `json:"name,omitempty"`
	Object                   string                                `json:"object"`
	PaymentSources           *CustomerPaymentMethodsResponse       `json:"payment_sources,omitempty"`
	Phone                    *string                               `json:"phone,omitempty"`
	ShippingContacts         *CustomerResponseShippingContacts     `json:"shipping_contacts,omitempty"`
	Subscription             *SubscriptionResponse                 `json:"subscription,omitempty"`
}

// NewCustomerResponse instantiates a new CustomerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerResponse(createdAt int64, id string, livemode bool, object string) *CustomerResponse {
	this := CustomerResponse{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Livemode = livemode
	this.Object = object
	return &this
}

// NewCustomerResponseWithDefaults instantiates a new CustomerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerResponseWithDefaults() *CustomerResponse {
	this := CustomerResponse{}
	return &this
}

// GetAntifraudInfo returns the AntifraudInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetAntifraudInfo() CustomerAntifraudInfoResponse {
	if o == nil || IsNil(o.AntifraudInfo.Get()) {
		var ret CustomerAntifraudInfoResponse
		return ret
	}
	return *o.AntifraudInfo.Get()
}

// GetAntifraudInfoOk returns a tuple with the AntifraudInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetAntifraudInfoOk() (*CustomerAntifraudInfoResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AntifraudInfo.Get(), o.AntifraudInfo.IsSet()
}

// HasAntifraudInfo returns a boolean if a field has been set.
func (o *CustomerResponse) HasAntifraudInfo() bool {
	if o != nil && o.AntifraudInfo.IsSet() {
		return true
	}

	return false
}

// SetAntifraudInfo gets a reference to the given NullableCustomerAntifraudInfoResponse and assigns it to the AntifraudInfo field.
func (o *CustomerResponse) SetAntifraudInfo(v CustomerAntifraudInfoResponse) {
	o.AntifraudInfo.Set(&v)
}

// SetAntifraudInfoNil sets the value for AntifraudInfo to be an explicit nil
func (o *CustomerResponse) SetAntifraudInfoNil() {
	o.AntifraudInfo.Set(nil)
}

// UnsetAntifraudInfo ensures that no value is present for AntifraudInfo, not even an explicit nil
func (o *CustomerResponse) UnsetAntifraudInfo() {
	o.AntifraudInfo.Unset()
}

// GetCorporate returns the Corporate field value if set, zero value otherwise.
func (o *CustomerResponse) GetCorporate() bool {
	if o == nil || IsNil(o.Corporate) {
		var ret bool
		return ret
	}
	return *o.Corporate
}

// GetCorporateOk returns a tuple with the Corporate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetCorporateOk() (*bool, bool) {
	if o == nil || IsNil(o.Corporate) {
		return nil, false
	}
	return o.Corporate, true
}

// HasCorporate returns a boolean if a field has been set.
func (o *CustomerResponse) HasCorporate() bool {
	if o != nil && !IsNil(o.Corporate) {
		return true
	}

	return false
}

// SetCorporate gets a reference to the given bool and assigns it to the Corporate field.
func (o *CustomerResponse) SetCorporate(v bool) {
	o.Corporate = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomerResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomerResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCustomReference returns the CustomReference field value if set, zero value otherwise.
func (o *CustomerResponse) GetCustomReference() string {
	if o == nil || IsNil(o.CustomReference) {
		var ret string
		return ret
	}
	return *o.CustomReference
}

// GetCustomReferenceOk returns a tuple with the CustomReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetCustomReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomReference) {
		return nil, false
	}
	return o.CustomReference, true
}

// HasCustomReference returns a boolean if a field has been set.
func (o *CustomerResponse) HasCustomReference() bool {
	if o != nil && !IsNil(o.CustomReference) {
		return true
	}

	return false
}

// SetCustomReference gets a reference to the given string and assigns it to the CustomReference field.
func (o *CustomerResponse) SetCustomReference(v string) {
	o.CustomReference = &v
}

// GetDefaultFiscalEntityId returns the DefaultFiscalEntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetDefaultFiscalEntityId() string {
	if o == nil || IsNil(o.DefaultFiscalEntityId.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultFiscalEntityId.Get()
}

// GetDefaultFiscalEntityIdOk returns a tuple with the DefaultFiscalEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetDefaultFiscalEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultFiscalEntityId.Get(), o.DefaultFiscalEntityId.IsSet()
}

// HasDefaultFiscalEntityId returns a boolean if a field has been set.
func (o *CustomerResponse) HasDefaultFiscalEntityId() bool {
	if o != nil && o.DefaultFiscalEntityId.IsSet() {
		return true
	}

	return false
}

// SetDefaultFiscalEntityId gets a reference to the given NullableString and assigns it to the DefaultFiscalEntityId field.
func (o *CustomerResponse) SetDefaultFiscalEntityId(v string) {
	o.DefaultFiscalEntityId.Set(&v)
}

// SetDefaultFiscalEntityIdNil sets the value for DefaultFiscalEntityId to be an explicit nil
func (o *CustomerResponse) SetDefaultFiscalEntityIdNil() {
	o.DefaultFiscalEntityId.Set(nil)
}

// UnsetDefaultFiscalEntityId ensures that no value is present for DefaultFiscalEntityId, not even an explicit nil
func (o *CustomerResponse) UnsetDefaultFiscalEntityId() {
	o.DefaultFiscalEntityId.Unset()
}

// GetDefaultShippingContactId returns the DefaultShippingContactId field value if set, zero value otherwise.
func (o *CustomerResponse) GetDefaultShippingContactId() string {
	if o == nil || IsNil(o.DefaultShippingContactId) {
		var ret string
		return ret
	}
	return *o.DefaultShippingContactId
}

// GetDefaultShippingContactIdOk returns a tuple with the DefaultShippingContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetDefaultShippingContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultShippingContactId) {
		return nil, false
	}
	return o.DefaultShippingContactId, true
}

// HasDefaultShippingContactId returns a boolean if a field has been set.
func (o *CustomerResponse) HasDefaultShippingContactId() bool {
	if o != nil && !IsNil(o.DefaultShippingContactId) {
		return true
	}

	return false
}

// SetDefaultShippingContactId gets a reference to the given string and assigns it to the DefaultShippingContactId field.
func (o *CustomerResponse) SetDefaultShippingContactId(v string) {
	o.DefaultShippingContactId = &v
}

// GetDefaultPaymentSourceId returns the DefaultPaymentSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerResponse) GetDefaultPaymentSourceId() string {
	if o == nil || IsNil(o.DefaultPaymentSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultPaymentSourceId.Get()
}

// GetDefaultPaymentSourceIdOk returns a tuple with the DefaultPaymentSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerResponse) GetDefaultPaymentSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultPaymentSourceId.Get(), o.DefaultPaymentSourceId.IsSet()
}

// HasDefaultPaymentSourceId returns a boolean if a field has been set.
func (o *CustomerResponse) HasDefaultPaymentSourceId() bool {
	if o != nil && o.DefaultPaymentSourceId.IsSet() {
		return true
	}

	return false
}

// SetDefaultPaymentSourceId gets a reference to the given NullableString and assigns it to the DefaultPaymentSourceId field.
func (o *CustomerResponse) SetDefaultPaymentSourceId(v string) {
	o.DefaultPaymentSourceId.Set(&v)
}

// SetDefaultPaymentSourceIdNil sets the value for DefaultPaymentSourceId to be an explicit nil
func (o *CustomerResponse) SetDefaultPaymentSourceIdNil() {
	o.DefaultPaymentSourceId.Set(nil)
}

// UnsetDefaultPaymentSourceId ensures that no value is present for DefaultPaymentSourceId, not even an explicit nil
func (o *CustomerResponse) UnsetDefaultPaymentSourceId() {
	o.DefaultPaymentSourceId.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerResponse) SetEmail(v string) {
	o.Email = &v
}

// GetFiscalEntities returns the FiscalEntities field value if set, zero value otherwise.
func (o *CustomerResponse) GetFiscalEntities() CustomerFiscalEntitiesResponse {
	if o == nil || IsNil(o.FiscalEntities) {
		var ret CustomerFiscalEntitiesResponse
		return ret
	}
	return *o.FiscalEntities
}

// GetFiscalEntitiesOk returns a tuple with the FiscalEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetFiscalEntitiesOk() (*CustomerFiscalEntitiesResponse, bool) {
	if o == nil || IsNil(o.FiscalEntities) {
		return nil, false
	}
	return o.FiscalEntities, true
}

// HasFiscalEntities returns a boolean if a field has been set.
func (o *CustomerResponse) HasFiscalEntities() bool {
	if o != nil && !IsNil(o.FiscalEntities) {
		return true
	}

	return false
}

// SetFiscalEntities gets a reference to the given CustomerFiscalEntitiesResponse and assigns it to the FiscalEntities field.
func (o *CustomerResponse) SetFiscalEntities(v CustomerFiscalEntitiesResponse) {
	o.FiscalEntities = &v
}

// GetId returns the Id field value
func (o *CustomerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerResponse) SetId(v string) {
	o.Id = v
}

// GetLivemode returns the Livemode field value
func (o *CustomerResponse) GetLivemode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Livemode, true
}

// SetLivemode sets field value
func (o *CustomerResponse) SetLivemode(v bool) {
	o.Livemode = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerResponse) SetName(v string) {
	o.Name = &v
}

// GetObject returns the Object field value
func (o *CustomerResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CustomerResponse) SetObject(v string) {
	o.Object = v
}

// GetPaymentSources returns the PaymentSources field value if set, zero value otherwise.
func (o *CustomerResponse) GetPaymentSources() CustomerPaymentMethodsResponse {
	if o == nil || IsNil(o.PaymentSources) {
		var ret CustomerPaymentMethodsResponse
		return ret
	}
	return *o.PaymentSources
}

// GetPaymentSourcesOk returns a tuple with the PaymentSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetPaymentSourcesOk() (*CustomerPaymentMethodsResponse, bool) {
	if o == nil || IsNil(o.PaymentSources) {
		return nil, false
	}
	return o.PaymentSources, true
}

// HasPaymentSources returns a boolean if a field has been set.
func (o *CustomerResponse) HasPaymentSources() bool {
	if o != nil && !IsNil(o.PaymentSources) {
		return true
	}

	return false
}

// SetPaymentSources gets a reference to the given CustomerPaymentMethodsResponse and assigns it to the PaymentSources field.
func (o *CustomerResponse) SetPaymentSources(v CustomerPaymentMethodsResponse) {
	o.PaymentSources = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerResponse) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerResponse) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerResponse) SetPhone(v string) {
	o.Phone = &v
}

// GetShippingContacts returns the ShippingContacts field value if set, zero value otherwise.
func (o *CustomerResponse) GetShippingContacts() CustomerResponseShippingContacts {
	if o == nil || IsNil(o.ShippingContacts) {
		var ret CustomerResponseShippingContacts
		return ret
	}
	return *o.ShippingContacts
}

// GetShippingContactsOk returns a tuple with the ShippingContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetShippingContactsOk() (*CustomerResponseShippingContacts, bool) {
	if o == nil || IsNil(o.ShippingContacts) {
		return nil, false
	}
	return o.ShippingContacts, true
}

// HasShippingContacts returns a boolean if a field has been set.
func (o *CustomerResponse) HasShippingContacts() bool {
	if o != nil && !IsNil(o.ShippingContacts) {
		return true
	}

	return false
}

// SetShippingContacts gets a reference to the given CustomerResponseShippingContacts and assigns it to the ShippingContacts field.
func (o *CustomerResponse) SetShippingContacts(v CustomerResponseShippingContacts) {
	o.ShippingContacts = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *CustomerResponse) GetSubscription() SubscriptionResponse {
	if o == nil || IsNil(o.Subscription) {
		var ret SubscriptionResponse
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponse) GetSubscriptionOk() (*SubscriptionResponse, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *CustomerResponse) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given SubscriptionResponse and assigns it to the Subscription field.
func (o *CustomerResponse) SetSubscription(v SubscriptionResponse) {
	o.Subscription = &v
}

func (o CustomerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AntifraudInfo.IsSet() {
		toSerialize["antifraud_info"] = o.AntifraudInfo.Get()
	}
	if !IsNil(o.Corporate) {
		toSerialize["corporate"] = o.Corporate
	}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.CustomReference) {
		toSerialize["custom_reference"] = o.CustomReference
	}
	if o.DefaultFiscalEntityId.IsSet() {
		toSerialize["default_fiscal_entity_id"] = o.DefaultFiscalEntityId.Get()
	}
	if !IsNil(o.DefaultShippingContactId) {
		toSerialize["default_shipping_contact_id"] = o.DefaultShippingContactId
	}
	if o.DefaultPaymentSourceId.IsSet() {
		toSerialize["default_payment_source_id"] = o.DefaultPaymentSourceId.Get()
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FiscalEntities) {
		toSerialize["fiscal_entities"] = o.FiscalEntities
	}
	toSerialize["id"] = o.Id
	toSerialize["livemode"] = o.Livemode
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["object"] = o.Object
	if !IsNil(o.PaymentSources) {
		toSerialize["payment_sources"] = o.PaymentSources
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.ShippingContacts) {
		toSerialize["shipping_contacts"] = o.ShippingContacts
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableCustomerResponse struct {
	value *CustomerResponse
	isSet bool
}

func (v NullableCustomerResponse) Get() *CustomerResponse {
	return v.value
}

func (v *NullableCustomerResponse) Set(val *CustomerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerResponse(val *CustomerResponse) *NullableCustomerResponse {
	return &NullableCustomerResponse{value: val, isSet: true}
}

func (v NullableCustomerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the UpdateCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomer{}

// UpdateCustomer update customer
type UpdateCustomer struct {
	AntifraudInfo NullableUpdateCustomerAntifraudInfo `json:"antifraud_info,omitempty"`
	// It is a parameter that allows to identify in the response, the Conekta ID of a payment method (payment_id)
	DefaultPaymentSourceId *string `json:"default_payment_source_id,omitempty"`
	// An email address is a series of customizable characters followed by a universal Internet symbol, the at symbol (@), the name of a host server, and a web domain ending (.mx, .com, .org, . net, etc).
	Email *string `json:"email,omitempty"`
	// Client's name
	Name *string `json:"name,omitempty"`
	// Is the customer's phone number
	Phone *string `json:"phone,omitempty"`
	// Contains the ID of a plan, which could together with name, email and phone create a client directly to a subscription
	PlanId *string `json:"plan_id,omitempty"`
	// It is a parameter that allows to identify in the response, the Conekta ID of the shipping address (shipping_contact)
	DefaultShippingContactId *string `json:"default_shipping_contact_id,omitempty"`
	// It is a value that allows identifying if the email is corporate or not.
	Corporate *bool `json:"corporate,omitempty"`
	// It is an undefined value.
	CustomReference *string                         `json:"custom_reference,omitempty"`
	FiscalEntities  []CustomerFiscalEntitiesRequest `json:"fiscal_entities,omitempty"`
	Metadata        map[string]interface{}          `json:"metadata,omitempty"`
	// Contains details of the payment methods that the customer has active or has used in Conekta
	PaymentSources []CustomerPaymentMethodsRequest `json:"payment_sources,omitempty"`
	// Contains the detail of the shipping addresses that the client has active or has used in Conekta
	ShippingContacts []CustomerShippingContacts `json:"shipping_contacts,omitempty"`
	Subscription     *SubscriptionRequest       `json:"subscription,omitempty"`
}

// NewUpdateCustomer instantiates a new UpdateCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomer() *UpdateCustomer {
	this := UpdateCustomer{}
	var corporate bool = false
	this.Corporate = &corporate
	return &this
}

// NewUpdateCustomerWithDefaults instantiates a new UpdateCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomerWithDefaults() *UpdateCustomer {
	this := UpdateCustomer{}
	var corporate bool = false
	this.Corporate = &corporate
	return &this
}

// GetAntifraudInfo returns the AntifraudInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCustomer) GetAntifraudInfo() UpdateCustomerAntifraudInfo {
	if o == nil || IsNil(o.AntifraudInfo.Get()) {
		var ret UpdateCustomerAntifraudInfo
		return ret
	}
	return *o.AntifraudInfo.Get()
}

// GetAntifraudInfoOk returns a tuple with the AntifraudInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCustomer) GetAntifraudInfoOk() (*UpdateCustomerAntifraudInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AntifraudInfo.Get(), o.AntifraudInfo.IsSet()
}

// HasAntifraudInfo returns a boolean if a field has been set.
func (o *UpdateCustomer) HasAntifraudInfo() bool {
	if o != nil && o.AntifraudInfo.IsSet() {
		return true
	}

	return false
}

// SetAntifraudInfo gets a reference to the given NullableUpdateCustomerAntifraudInfo and assigns it to the AntifraudInfo field.
func (o *UpdateCustomer) SetAntifraudInfo(v UpdateCustomerAntifraudInfo) {
	o.AntifraudInfo.Set(&v)
}

// SetAntifraudInfoNil sets the value for AntifraudInfo to be an explicit nil
func (o *UpdateCustomer) SetAntifraudInfoNil() {
	o.AntifraudInfo.Set(nil)
}

// UnsetAntifraudInfo ensures that no value is present for AntifraudInfo, not even an explicit nil
func (o *UpdateCustomer) UnsetAntifraudInfo() {
	o.AntifraudInfo.Unset()
}

// GetDefaultPaymentSourceId returns the DefaultPaymentSourceId field value if set, zero value otherwise.
func (o *UpdateCustomer) GetDefaultPaymentSourceId() string {
	if o == nil || IsNil(o.DefaultPaymentSourceId) {
		var ret string
		return ret
	}
	return *o.DefaultPaymentSourceId
}

// GetDefaultPaymentSourceIdOk returns a tuple with the DefaultPaymentSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetDefaultPaymentSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPaymentSourceId) {
		return nil, false
	}
	return o.DefaultPaymentSourceId, true
}

// HasDefaultPaymentSourceId returns a boolean if a field has been set.
func (o *UpdateCustomer) HasDefaultPaymentSourceId() bool {
	if o != nil && !IsNil(o.DefaultPaymentSourceId) {
		return true
	}

	return false
}

// SetDefaultPaymentSourceId gets a reference to the given string and assigns it to the DefaultPaymentSourceId field.
func (o *UpdateCustomer) SetDefaultPaymentSourceId(v string) {
	o.DefaultPaymentSourceId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateCustomer) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateCustomer) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateCustomer) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateCustomer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateCustomer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateCustomer) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UpdateCustomer) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UpdateCustomer) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UpdateCustomer) SetPhone(v string) {
	o.Phone = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *UpdateCustomer) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *UpdateCustomer) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *UpdateCustomer) SetPlanId(v string) {
	o.PlanId = &v
}

// GetDefaultShippingContactId returns the DefaultShippingContactId field value if set, zero value otherwise.
func (o *UpdateCustomer) GetDefaultShippingContactId() string {
	if o == nil || IsNil(o.DefaultShippingContactId) {
		var ret string
		return ret
	}
	return *o.DefaultShippingContactId
}

// GetDefaultShippingContactIdOk returns a tuple with the DefaultShippingContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetDefaultShippingContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultShippingContactId) {
		return nil, false
	}
	return o.DefaultShippingContactId, true
}

// HasDefaultShippingContactId returns a boolean if a field has been set.
func (o *UpdateCustomer) HasDefaultShippingContactId() bool {
	if o != nil && !IsNil(o.DefaultShippingContactId) {
		return true
	}

	return false
}

// SetDefaultShippingContactId gets a reference to the given string and assigns it to the DefaultShippingContactId field.
func (o *UpdateCustomer) SetDefaultShippingContactId(v string) {
	o.DefaultShippingContactId = &v
}

// GetCorporate returns the Corporate field value if set, zero value otherwise.
func (o *UpdateCustomer) GetCorporate() bool {
	if o == nil || IsNil(o.Corporate) {
		var ret bool
		return ret
	}
	return *o.Corporate
}

// GetCorporateOk returns a tuple with the Corporate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetCorporateOk() (*bool, bool) {
	if o == nil || IsNil(o.Corporate) {
		return nil, false
	}
	return o.Corporate, true
}

// HasCorporate returns a boolean if a field has been set.
func (o *UpdateCustomer) HasCorporate() bool {
	if o != nil && !IsNil(o.Corporate) {
		return true
	}

	return false
}

// SetCorporate gets a reference to the given bool and assigns it to the Corporate field.
func (o *UpdateCustomer) SetCorporate(v bool) {
	o.Corporate = &v
}

// GetCustomReference returns the CustomReference field value if set, zero value otherwise.
func (o *UpdateCustomer) GetCustomReference() string {
	if o == nil || IsNil(o.CustomReference) {
		var ret string
		return ret
	}
	return *o.CustomReference
}

// GetCustomReferenceOk returns a tuple with the CustomReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetCustomReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomReference) {
		return nil, false
	}
	return o.CustomReference, true
}

// HasCustomReference returns a boolean if a field has been set.
func (o *UpdateCustomer) HasCustomReference() bool {
	if o != nil && !IsNil(o.CustomReference) {
		return true
	}

	return false
}

// SetCustomReference gets a reference to the given string and assigns it to the CustomReference field.
func (o *UpdateCustomer) SetCustomReference(v string) {
	o.CustomReference = &v
}

// GetFiscalEntities returns the FiscalEntities field value if set, zero value otherwise.
func (o *UpdateCustomer) GetFiscalEntities() []CustomerFiscalEntitiesRequest {
	if o == nil || IsNil(o.FiscalEntities) {
		var ret []CustomerFiscalEntitiesRequest
		return ret
	}
	return o.FiscalEntities
}

// GetFiscalEntitiesOk returns a tuple with the FiscalEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetFiscalEntitiesOk() ([]CustomerFiscalEntitiesRequest, bool) {
	if o == nil || IsNil(o.FiscalEntities) {
		return nil, false
	}
	return o.FiscalEntities, true
}

// HasFiscalEntities returns a boolean if a field has been set.
func (o *UpdateCustomer) HasFiscalEntities() bool {
	if o != nil && !IsNil(o.FiscalEntities) {
		return true
	}

	return false
}

// SetFiscalEntities gets a reference to the given []CustomerFiscalEntitiesRequest and assigns it to the FiscalEntities field.
func (o *UpdateCustomer) SetFiscalEntities(v []CustomerFiscalEntitiesRequest) {
	o.FiscalEntities = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateCustomer) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateCustomer) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UpdateCustomer) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPaymentSources returns the PaymentSources field value if set, zero value otherwise.
func (o *UpdateCustomer) GetPaymentSources() []CustomerPaymentMethodsRequest {
	if o == nil || IsNil(o.PaymentSources) {
		var ret []CustomerPaymentMethodsRequest
		return ret
	}
	return o.PaymentSources
}

// GetPaymentSourcesOk returns a tuple with the PaymentSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetPaymentSourcesOk() ([]CustomerPaymentMethodsRequest, bool) {
	if o == nil || IsNil(o.PaymentSources) {
		return nil, false
	}
	return o.PaymentSources, true
}

// HasPaymentSources returns a boolean if a field has been set.
func (o *UpdateCustomer) HasPaymentSources() bool {
	if o != nil && !IsNil(o.PaymentSources) {
		return true
	}

	return false
}

// SetPaymentSources gets a reference to the given []CustomerPaymentMethodsRequest and assigns it to the PaymentSources field.
func (o *UpdateCustomer) SetPaymentSources(v []CustomerPaymentMethodsRequest) {
	o.PaymentSources = v
}

// GetShippingContacts returns the ShippingContacts field value if set, zero value otherwise.
func (o *UpdateCustomer) GetShippingContacts() []CustomerShippingContacts {
	if o == nil || IsNil(o.ShippingContacts) {
		var ret []CustomerShippingContacts
		return ret
	}
	return o.ShippingContacts
}

// GetShippingContactsOk returns a tuple with the ShippingContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetShippingContactsOk() ([]CustomerShippingContacts, bool) {
	if o == nil || IsNil(o.ShippingContacts) {
		return nil, false
	}
	return o.ShippingContacts, true
}

// HasShippingContacts returns a boolean if a field has been set.
func (o *UpdateCustomer) HasShippingContacts() bool {
	if o != nil && !IsNil(o.ShippingContacts) {
		return true
	}

	return false
}

// SetShippingContacts gets a reference to the given []CustomerShippingContacts and assigns it to the ShippingContacts field.
func (o *UpdateCustomer) SetShippingContacts(v []CustomerShippingContacts) {
	o.ShippingContacts = v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *UpdateCustomer) GetSubscription() SubscriptionRequest {
	if o == nil || IsNil(o.Subscription) {
		var ret SubscriptionRequest
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomer) GetSubscriptionOk() (*SubscriptionRequest, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *UpdateCustomer) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given SubscriptionRequest and assigns it to the Subscription field.
func (o *UpdateCustomer) SetSubscription(v SubscriptionRequest) {
	o.Subscription = &v
}

func (o UpdateCustomer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AntifraudInfo.IsSet() {
		toSerialize["antifraud_info"] = o.AntifraudInfo.Get()
	}
	if !IsNil(o.DefaultPaymentSourceId) {
		toSerialize["default_payment_source_id"] = o.DefaultPaymentSourceId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.DefaultShippingContactId) {
		toSerialize["default_shipping_contact_id"] = o.DefaultShippingContactId
	}
	if !IsNil(o.Corporate) {
		toSerialize["corporate"] = o.Corporate
	}
	if !IsNil(o.CustomReference) {
		toSerialize["custom_reference"] = o.CustomReference
	}
	if !IsNil(o.FiscalEntities) {
		toSerialize["fiscal_entities"] = o.FiscalEntities
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PaymentSources) {
		toSerialize["payment_sources"] = o.PaymentSources
	}
	if !IsNil(o.ShippingContacts) {
		toSerialize["shipping_contacts"] = o.ShippingContacts
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableUpdateCustomer struct {
	value *UpdateCustomer
	isSet bool
}

func (v NullableUpdateCustomer) Get() *UpdateCustomer {
	return v.value
}

func (v *NullableUpdateCustomer) Set(val *UpdateCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomer(val *UpdateCustomer) *NullableUpdateCustomer {
	return &NullableUpdateCustomer{value: val, isSet: true}
}

func (v NullableUpdateCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the Customer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Customer{}

// Customer a customer
type Customer struct {
	AntifraudInfo NullableCustomerAntifraudInfo `json:"antifraud_info,omitempty"`
	// It is a value that allows identifying if the email is corporate or not.
	Corporate *bool `json:"corporate,omitempty"`
	// It is an undefined value.
	CustomReference *string `json:"custom_reference,omitempty"`
	// An email address is a series of customizable characters followed by a universal Internet symbol, the at symbol (@), the name of a host server, and a web domain ending (.mx, .com, .org, . net, etc).
	Email string `json:"email"`
	// It is a parameter that allows to identify in the response, the Conekta ID of a payment method (payment_id)
	DefaultPaymentSourceId *string `json:"default_payment_source_id,omitempty"`
	// It is a parameter that allows to identify in the response, the Conekta ID of the shipping address (shipping_contact)
	DefaultShippingContactId *string                         `json:"default_shipping_contact_id,omitempty"`
	FiscalEntities           []CustomerFiscalEntitiesRequest `json:"fiscal_entities,omitempty"`
	Metadata                 map[string]interface{}          `json:"metadata,omitempty"`
	// Client's name
	Name string `json:"name"`
	// Contains details of the payment methods that the customer has active or has used in Conekta
	PaymentSources []CustomerPaymentMethodsRequest `json:"payment_sources,omitempty"`
	// Is the customer's phone number
	Phone string `json:"phone"`
	// Contains the ID of a plan, which could together with name, email and phone create a client directly to a subscription
	PlanId *string `json:"plan_id,omitempty"`
	// Contains the detail of the shipping addresses that the client has active or has used in Conekta
	ShippingContacts []CustomerShippingContacts `json:"shipping_contacts,omitempty"`
	Subscription     *SubscriptionRequest       `json:"subscription,omitempty"`
}

// NewCustomer instantiates a new Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomer(email string, name string, phone string) *Customer {
	this := Customer{}
	var corporate bool = false
	this.Corporate = &corporate
	this.Email = email
	this.Name = name
	this.Phone = phone
	return &this
}

// NewCustomerWithDefaults instantiates a new Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWithDefaults() *Customer {
	this := Customer{}
	var corporate bool = false
	this.Corporate = &corporate
	return &this
}

// GetAntifraudInfo returns the AntifraudInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Customer) GetAntifraudInfo() CustomerAntifraudInfo {
	if o == nil || IsNil(o.AntifraudInfo.Get()) {
		var ret CustomerAntifraudInfo
		return ret
	}
	return *o.AntifraudInfo.Get()
}

// GetAntifraudInfoOk returns a tuple with the AntifraudInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Customer) GetAntifraudInfoOk() (*CustomerAntifraudInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AntifraudInfo.Get(), o.AntifraudInfo.IsSet()
}

// HasAntifraudInfo returns a boolean if a field has been set.
func (o *Customer) HasAntifraudInfo() bool {
	if o != nil && o.AntifraudInfo.IsSet() {
		return true
	}

	return false
}

// SetAntifraudInfo gets a reference to the given NullableCustomerAntifraudInfo and assigns it to the AntifraudInfo field.
func (o *Customer) SetAntifraudInfo(v CustomerAntifraudInfo) {
	o.AntifraudInfo.Set(&v)
}

// SetAntifraudInfoNil sets the value for AntifraudInfo to be an explicit nil
func (o *Customer) SetAntifraudInfoNil() {
	o.AntifraudInfo.Set(nil)
}

// UnsetAntifraudInfo ensures that no value is present for AntifraudInfo, not even an explicit nil
func (o *Customer) UnsetAntifraudInfo() {
	o.AntifraudInfo.Unset()
}

// GetCorporate returns the Corporate field value if set, zero value otherwise.
func (o *Customer) GetCorporate() bool {
	if o == nil || IsNil(o.Corporate) {
		var ret bool
		return ret
	}
	return *o.Corporate
}

// GetCorporateOk returns a tuple with the Corporate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCorporateOk() (*bool, bool) {
	if o == nil || IsNil(o.Corporate) {
		return nil, false
	}
	return o.Corporate, true
}

// HasCorporate returns a boolean if a field has been set.
func (o *Customer) HasCorporate() bool {
	if o != nil && !IsNil(o.Corporate) {
		return true
	}

	return false
}

// SetCorporate gets a reference to the given bool and assigns it to the Corporate field.
func (o *Customer) SetCorporate(v bool) {
	o.Corporate = &v
}

// GetCustomReference returns the CustomReference field value if set, zero value otherwise.
func (o *Customer) GetCustomReference() string {
	if o == nil || IsNil(o.CustomReference) {
		var ret string
		return ret
	}
	return *o.CustomReference
}

// GetCustomReferenceOk returns a tuple with the CustomReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCustomReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomReference) {
		return nil, false
	}
	return o.CustomReference, true
}

// HasCustomReference returns a boolean if a field has been set.
func (o *Customer) HasCustomReference() bool {
	if o != nil && !IsNil(o.CustomReference) {
		return true
	}

	return false
}

// SetCustomReference gets a reference to the given string and assigns it to the CustomReference field.
func (o *Customer) SetCustomReference(v string) {
	o.CustomReference = &v
}

// GetEmail returns the Email field value
func (o *Customer) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Customer) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Customer) SetEmail(v string) {
	o.Email = v
}

// GetDefaultPaymentSourceId returns the DefaultPaymentSourceId field value if set, zero value otherwise.
func (o *Customer) GetDefaultPaymentSourceId() string {
	if o == nil || IsNil(o.DefaultPaymentSourceId) {
		var ret string
		return ret
	}
	return *o.DefaultPaymentSourceId
}

// GetDefaultPaymentSourceIdOk returns a tuple with the DefaultPaymentSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDefaultPaymentSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPaymentSourceId) {
		return nil, false
	}
	return o.DefaultPaymentSourceId, true
}

// HasDefaultPaymentSourceId returns a boolean if a field has been set.
func (o *Customer) HasDefaultPaymentSourceId() bool {
	if o != nil && !IsNil(o.DefaultPaymentSourceId) {
		return true
	}

	return false
}

// SetDefaultPaymentSourceId gets a reference to the given string and assigns it to the DefaultPaymentSourceId field.
func (o *Customer) SetDefaultPaymentSourceId(v string) {
	o.DefaultPaymentSourceId = &v
}

// GetDefaultShippingContactId returns the DefaultShippingContactId field value if set, zero value otherwise.
func (o *Customer) GetDefaultShippingContactId() string {
	if o == nil || IsNil(o.DefaultShippingContactId) {
		var ret string
		return ret
	}
	return *o.DefaultShippingContactId
}

// GetDefaultShippingContactIdOk returns a tuple with the DefaultShippingContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDefaultShippingContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultShippingContactId) {
		return nil, false
	}
	return o.DefaultShippingContactId, true
}

// HasDefaultShippingContactId returns a boolean if a field has been set.
func (o *Customer) HasDefaultShippingContactId() bool {
	if o != nil && !IsNil(o.DefaultShippingContactId) {
		return true
	}

	return false
}

// SetDefaultShippingContactId gets a reference to the given string and assigns it to the DefaultShippingContactId field.
func (o *Customer) SetDefaultShippingContactId(v string) {
	o.DefaultShippingContactId = &v
}

// GetFiscalEntities returns the FiscalEntities field value if set, zero value otherwise.
func (o *Customer) GetFiscalEntities() []CustomerFiscalEntitiesRequest {
	if o == nil || IsNil(o.FiscalEntities) {
		var ret []CustomerFiscalEntitiesRequest
		return ret
	}
	return o.FiscalEntities
}

// GetFiscalEntitiesOk returns a tuple with the FiscalEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetFiscalEntitiesOk() ([]CustomerFiscalEntitiesRequest, bool) {
	if o == nil || IsNil(o.FiscalEntities) {
		return nil, false
	}
	return o.FiscalEntities, true
}

// HasFiscalEntities returns a boolean if a field has been set.
func (o *Customer) HasFiscalEntities() bool {
	if o != nil && !IsNil(o.FiscalEntities) {
		return true
	}

	return false
}

// SetFiscalEntities gets a reference to the given []CustomerFiscalEntitiesRequest and assigns it to the FiscalEntities field.
func (o *Customer) SetFiscalEntities(v []CustomerFiscalEntitiesRequest) {
	o.FiscalEntities = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Customer) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Customer) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Customer) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *Customer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Customer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Customer) SetName(v string) {
	o.Name = v
}

// GetPaymentSources returns the PaymentSources field value if set, zero value otherwise.
func (o *Customer) GetPaymentSources() []CustomerPaymentMethodsRequest {
	if o == nil || IsNil(o.PaymentSources) {
		var ret []CustomerPaymentMethodsRequest
		return ret
	}
	return o.PaymentSources
}

// GetPaymentSourcesOk returns a tuple with the PaymentSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetPaymentSourcesOk() ([]CustomerPaymentMethodsRequest, bool) {
	if o == nil || IsNil(o.PaymentSources) {
		return nil, false
	}
	return o.PaymentSources, true
}

// HasPaymentSources returns a boolean if a field has been set.
func (o *Customer) HasPaymentSources() bool {
	if o != nil && !IsNil(o.PaymentSources) {
		return true
	}

	return false
}

// SetPaymentSources gets a reference to the given []CustomerPaymentMethodsRequest and assigns it to the PaymentSources field.
func (o *Customer) SetPaymentSources(v []CustomerPaymentMethodsRequest) {
	o.PaymentSources = v
}

// GetPhone returns the Phone field value
func (o *Customer) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *Customer) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *Customer) SetPhone(v string) {
	o.Phone = v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *Customer) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *Customer) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *Customer) SetPlanId(v string) {
	o.PlanId = &v
}

// GetShippingContacts returns the ShippingContacts field value if set, zero value otherwise.
func (o *Customer) GetShippingContacts() []CustomerShippingContacts {
	if o == nil || IsNil(o.ShippingContacts) {
		var ret []CustomerShippingContacts
		return ret
	}
	return o.ShippingContacts
}

// GetShippingContactsOk returns a tuple with the ShippingContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetShippingContactsOk() ([]CustomerShippingContacts, bool) {
	if o == nil || IsNil(o.ShippingContacts) {
		return nil, false
	}
	return o.ShippingContacts, true
}

// HasShippingContacts returns a boolean if a field has been set.
func (o *Customer) HasShippingContacts() bool {
	if o != nil && !IsNil(o.ShippingContacts) {
		return true
	}

	return false
}

// SetShippingContacts gets a reference to the given []CustomerShippingContacts and assigns it to the ShippingContacts field.
func (o *Customer) SetShippingContacts(v []CustomerShippingContacts) {
	o.ShippingContacts = v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *Customer) GetSubscription() SubscriptionRequest {
	if o == nil || IsNil(o.Subscription) {
		var ret SubscriptionRequest
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetSubscriptionOk() (*SubscriptionRequest, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *Customer) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given SubscriptionRequest and assigns it to the Subscription field.
func (o *Customer) SetSubscription(v SubscriptionRequest) {
	o.Subscription = &v
}

func (o Customer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Customer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AntifraudInfo.IsSet() {
		toSerialize["antifraud_info"] = o.AntifraudInfo.Get()
	}
	if !IsNil(o.Corporate) {
		toSerialize["corporate"] = o.Corporate
	}
	if !IsNil(o.CustomReference) {
		toSerialize["custom_reference"] = o.CustomReference
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.DefaultPaymentSourceId) {
		toSerialize["default_payment_source_id"] = o.DefaultPaymentSourceId
	}
	if !IsNil(o.DefaultShippingContactId) {
		toSerialize["default_shipping_contact_id"] = o.DefaultShippingContactId
	}
	if !IsNil(o.FiscalEntities) {
		toSerialize["fiscal_entities"] = o.FiscalEntities
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PaymentSources) {
		toSerialize["payment_sources"] = o.PaymentSources
	}
	toSerialize["phone"] = o.Phone
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.ShippingContacts) {
		toSerialize["shipping_contacts"] = o.ShippingContacts
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableCustomer struct {
	value *Customer
	isSet bool
}

func (v NullableCustomer) Get() *Customer {
	return v.value
}

func (v *NullableCustomer) Set(val *Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomer(val *Customer) *NullableCustomer {
	return &NullableCustomer{value: val, isSet: true}
}

func (v NullableCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

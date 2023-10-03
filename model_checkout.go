package conekta

import (
	"encoding/json"
)

// checks if the Checkout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Checkout{}

// Checkout It is a sub-resource of the Order model that can be stipulated in order to configure its corresponding checkout
type Checkout struct {
	// Those are the payment methods that will be available for the link
	AllowedPaymentMethods []string `json:"allowed_payment_methods"`
	// It is the time when the link will expire. It is expressed in seconds since the Unix epoch. The valid range is from 2 to 365 days (the valid range will be taken from the next day of the creation date at 00:01 hrs)
	ExpiresAt int64 `json:"expires_at"`
	// This flag allows you to specify if months without interest will be active.
	MonthlyInstallmentsEnabled *bool `json:"monthly_installments_enabled,omitempty"`
	// This field allows you to specify the number of months without interest.
	MonthlyInstallmentsOptions []int32 `json:"monthly_installments_options,omitempty"`
	// Reason for charge
	Name string `json:"name"`
	// This flag allows you to fill in the shipping information at checkout.
	NeedsShippingContact *bool `json:"needs_shipping_contact,omitempty"`
	// This flag allows you to specify if the link will be on demand.
	OnDemandEnabled NullableBool          `json:"on_demand_enabled,omitempty"`
	OrderTemplate   CheckoutOrderTemplate `json:"order_template"`
	// It is the number of payments that can be made through the link.
	PaymentsLimitCount *int32 `json:"payments_limit_count,omitempty"`
	// false: single use. true: multiple payments
	Recurrent bool `json:"recurrent"`
	// It is the type of link that will be created. It must be a valid type.
	Type string `json:"type"`
}

// NewCheckout instantiates a new Checkout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckout(allowedPaymentMethods []string, expiresAt int64, name string, orderTemplate CheckoutOrderTemplate, recurrent bool, type_ string) *Checkout {
	this := Checkout{}
	this.AllowedPaymentMethods = allowedPaymentMethods
	this.ExpiresAt = expiresAt
	this.Name = name
	this.OrderTemplate = orderTemplate
	this.Recurrent = recurrent
	this.Type = type_
	return &this
}

// NewCheckoutWithDefaults instantiates a new Checkout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutWithDefaults() *Checkout {
	this := Checkout{}
	return &this
}

// GetAllowedPaymentMethods returns the AllowedPaymentMethods field value
func (o *Checkout) GetAllowedPaymentMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedPaymentMethods
}

// GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field value
// and a boolean to check if the value has been set.
func (o *Checkout) GetAllowedPaymentMethodsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedPaymentMethods, true
}

// SetAllowedPaymentMethods sets field value
func (o *Checkout) SetAllowedPaymentMethods(v []string) {
	o.AllowedPaymentMethods = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *Checkout) GetExpiresAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *Checkout) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *Checkout) SetExpiresAt(v int64) {
	o.ExpiresAt = v
}

// GetMonthlyInstallmentsEnabled returns the MonthlyInstallmentsEnabled field value if set, zero value otherwise.
func (o *Checkout) GetMonthlyInstallmentsEnabled() bool {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		var ret bool
		return ret
	}
	return *o.MonthlyInstallmentsEnabled
}

// GetMonthlyInstallmentsEnabledOk returns a tuple with the MonthlyInstallmentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkout) GetMonthlyInstallmentsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		return nil, false
	}
	return o.MonthlyInstallmentsEnabled, true
}

// HasMonthlyInstallmentsEnabled returns a boolean if a field has been set.
func (o *Checkout) HasMonthlyInstallmentsEnabled() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsEnabled) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsEnabled gets a reference to the given bool and assigns it to the MonthlyInstallmentsEnabled field.
func (o *Checkout) SetMonthlyInstallmentsEnabled(v bool) {
	o.MonthlyInstallmentsEnabled = &v
}

// GetMonthlyInstallmentsOptions returns the MonthlyInstallmentsOptions field value if set, zero value otherwise.
func (o *Checkout) GetMonthlyInstallmentsOptions() []int32 {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		var ret []int32
		return ret
	}
	return o.MonthlyInstallmentsOptions
}

// GetMonthlyInstallmentsOptionsOk returns a tuple with the MonthlyInstallmentsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkout) GetMonthlyInstallmentsOptionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		return nil, false
	}
	return o.MonthlyInstallmentsOptions, true
}

// HasMonthlyInstallmentsOptions returns a boolean if a field has been set.
func (o *Checkout) HasMonthlyInstallmentsOptions() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsOptions) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsOptions gets a reference to the given []int32 and assigns it to the MonthlyInstallmentsOptions field.
func (o *Checkout) SetMonthlyInstallmentsOptions(v []int32) {
	o.MonthlyInstallmentsOptions = v
}

// GetName returns the Name field value
func (o *Checkout) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Checkout) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Checkout) SetName(v string) {
	o.Name = v
}

// GetNeedsShippingContact returns the NeedsShippingContact field value if set, zero value otherwise.
func (o *Checkout) GetNeedsShippingContact() bool {
	if o == nil || IsNil(o.NeedsShippingContact) {
		var ret bool
		return ret
	}
	return *o.NeedsShippingContact
}

// GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkout) GetNeedsShippingContactOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsShippingContact) {
		return nil, false
	}
	return o.NeedsShippingContact, true
}

// HasNeedsShippingContact returns a boolean if a field has been set.
func (o *Checkout) HasNeedsShippingContact() bool {
	if o != nil && !IsNil(o.NeedsShippingContact) {
		return true
	}

	return false
}

// SetNeedsShippingContact gets a reference to the given bool and assigns it to the NeedsShippingContact field.
func (o *Checkout) SetNeedsShippingContact(v bool) {
	o.NeedsShippingContact = &v
}

// GetOnDemandEnabled returns the OnDemandEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Checkout) GetOnDemandEnabled() bool {
	if o == nil || IsNil(o.OnDemandEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.OnDemandEnabled.Get()
}

// GetOnDemandEnabledOk returns a tuple with the OnDemandEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Checkout) GetOnDemandEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnDemandEnabled.Get(), o.OnDemandEnabled.IsSet()
}

// HasOnDemandEnabled returns a boolean if a field has been set.
func (o *Checkout) HasOnDemandEnabled() bool {
	if o != nil && o.OnDemandEnabled.IsSet() {
		return true
	}

	return false
}

// SetOnDemandEnabled gets a reference to the given NullableBool and assigns it to the OnDemandEnabled field.
func (o *Checkout) SetOnDemandEnabled(v bool) {
	o.OnDemandEnabled.Set(&v)
}

// SetOnDemandEnabledNil sets the value for OnDemandEnabled to be an explicit nil
func (o *Checkout) SetOnDemandEnabledNil() {
	o.OnDemandEnabled.Set(nil)
}

// UnsetOnDemandEnabled ensures that no value is present for OnDemandEnabled, not even an explicit nil
func (o *Checkout) UnsetOnDemandEnabled() {
	o.OnDemandEnabled.Unset()
}

// GetOrderTemplate returns the OrderTemplate field value
func (o *Checkout) GetOrderTemplate() CheckoutOrderTemplate {
	if o == nil {
		var ret CheckoutOrderTemplate
		return ret
	}

	return o.OrderTemplate
}

// GetOrderTemplateOk returns a tuple with the OrderTemplate field value
// and a boolean to check if the value has been set.
func (o *Checkout) GetOrderTemplateOk() (*CheckoutOrderTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderTemplate, true
}

// SetOrderTemplate sets field value
func (o *Checkout) SetOrderTemplate(v CheckoutOrderTemplate) {
	o.OrderTemplate = v
}

// GetPaymentsLimitCount returns the PaymentsLimitCount field value if set, zero value otherwise.
func (o *Checkout) GetPaymentsLimitCount() int32 {
	if o == nil || IsNil(o.PaymentsLimitCount) {
		var ret int32
		return ret
	}
	return *o.PaymentsLimitCount
}

// GetPaymentsLimitCountOk returns a tuple with the PaymentsLimitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Checkout) GetPaymentsLimitCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PaymentsLimitCount) {
		return nil, false
	}
	return o.PaymentsLimitCount, true
}

// HasPaymentsLimitCount returns a boolean if a field has been set.
func (o *Checkout) HasPaymentsLimitCount() bool {
	if o != nil && !IsNil(o.PaymentsLimitCount) {
		return true
	}

	return false
}

// SetPaymentsLimitCount gets a reference to the given int32 and assigns it to the PaymentsLimitCount field.
func (o *Checkout) SetPaymentsLimitCount(v int32) {
	o.PaymentsLimitCount = &v
}

// GetRecurrent returns the Recurrent field value
func (o *Checkout) GetRecurrent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Recurrent
}

// GetRecurrentOk returns a tuple with the Recurrent field value
// and a boolean to check if the value has been set.
func (o *Checkout) GetRecurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurrent, true
}

// SetRecurrent sets field value
func (o *Checkout) SetRecurrent(v bool) {
	o.Recurrent = v
}

// GetType returns the Type field value
func (o *Checkout) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Checkout) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Checkout) SetType(v string) {
	o.Type = v
}

func (o Checkout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Checkout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowed_payment_methods"] = o.AllowedPaymentMethods
	toSerialize["expires_at"] = o.ExpiresAt
	if !IsNil(o.MonthlyInstallmentsEnabled) {
		toSerialize["monthly_installments_enabled"] = o.MonthlyInstallmentsEnabled
	}
	if !IsNil(o.MonthlyInstallmentsOptions) {
		toSerialize["monthly_installments_options"] = o.MonthlyInstallmentsOptions
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NeedsShippingContact) {
		toSerialize["needs_shipping_contact"] = o.NeedsShippingContact
	}
	if o.OnDemandEnabled.IsSet() {
		toSerialize["on_demand_enabled"] = o.OnDemandEnabled.Get()
	}
	toSerialize["order_template"] = o.OrderTemplate
	if !IsNil(o.PaymentsLimitCount) {
		toSerialize["payments_limit_count"] = o.PaymentsLimitCount
	}
	toSerialize["recurrent"] = o.Recurrent
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCheckout struct {
	value *Checkout
	isSet bool
}

func (v NullableCheckout) Get() *Checkout {
	return v.value
}

func (v *NullableCheckout) Set(val *Checkout) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckout) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckout(val *Checkout) *NullableCheckout {
	return &NullableCheckout{value: val, isSet: true}
}

func (v NullableCheckout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

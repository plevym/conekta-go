package conekta

import (
	"encoding/json"
)

// checks if the OrderResponseCheckout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderResponseCheckout{}

// OrderResponseCheckout struct for OrderResponseCheckout
type OrderResponseCheckout struct {
	AllowedPaymentMethods      []string                 `json:"allowed_payment_methods,omitempty"`
	CanNotExpire               *bool                    `json:"can_not_expire,omitempty"`
	EmailsSent                 *int32                   `json:"emails_sent,omitempty"`
	ExcludeCardNetworks        []map[string]interface{} `json:"exclude_card_networks,omitempty"`
	ExpiresAt                  *int64                   `json:"expires_at,omitempty"`
	FailureUrl                 *string                  `json:"failure_url,omitempty"`
	Force3dsFlow               *bool                    `json:"force_3ds_flow,omitempty"`
	Id                         *string                  `json:"id,omitempty"`
	IsRedirectOnFailure        *bool                    `json:"is_redirect_on_failure,omitempty"`
	Livemode                   *bool                    `json:"livemode,omitempty"`
	Metadata                   map[string]interface{}   `json:"metadata,omitempty"`
	MonthlyInstallmentsEnabled *bool                    `json:"monthly_installments_enabled,omitempty"`
	MonthlyInstallmentsOptions []int32                  `json:"monthly_installments_options,omitempty"`
	Name                       *string                  `json:"name,omitempty"`
	NeedsShippingContact       *bool                    `json:"needs_shipping_contact,omitempty"`
	Object                     *string                  `json:"object,omitempty"`
	OnDemandEnabled            NullableBool             `json:"on_demand_enabled,omitempty"`
	PaidPaymentsCount          *int32                   `json:"paid_payments_count,omitempty"`
	Recurrent                  *bool                    `json:"recurrent,omitempty"`
	Slug                       *string                  `json:"slug,omitempty"`
	SmsSent                    *int32                   `json:"sms_sent,omitempty"`
	SuccessUrl                 *string                  `json:"success_url,omitempty"`
	StartsAt                   *int32                   `json:"starts_at,omitempty"`
	Status                     *string                  `json:"status,omitempty"`
	Type                       *string                  `json:"type,omitempty"`
	Url                        *string                  `json:"url,omitempty"`
}

// NewOrderResponseCheckout instantiates a new OrderResponseCheckout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponseCheckout() *OrderResponseCheckout {
	this := OrderResponseCheckout{}
	return &this
}

// NewOrderResponseCheckoutWithDefaults instantiates a new OrderResponseCheckout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseCheckoutWithDefaults() *OrderResponseCheckout {
	this := OrderResponseCheckout{}
	return &this
}

// GetAllowedPaymentMethods returns the AllowedPaymentMethods field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetAllowedPaymentMethods() []string {
	if o == nil || IsNil(o.AllowedPaymentMethods) {
		var ret []string
		return ret
	}
	return o.AllowedPaymentMethods
}

// GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetAllowedPaymentMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedPaymentMethods) {
		return nil, false
	}
	return o.AllowedPaymentMethods, true
}

// HasAllowedPaymentMethods returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasAllowedPaymentMethods() bool {
	if o != nil && !IsNil(o.AllowedPaymentMethods) {
		return true
	}

	return false
}

// SetAllowedPaymentMethods gets a reference to the given []string and assigns it to the AllowedPaymentMethods field.
func (o *OrderResponseCheckout) SetAllowedPaymentMethods(v []string) {
	o.AllowedPaymentMethods = v
}

// GetCanNotExpire returns the CanNotExpire field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetCanNotExpire() bool {
	if o == nil || IsNil(o.CanNotExpire) {
		var ret bool
		return ret
	}
	return *o.CanNotExpire
}

// GetCanNotExpireOk returns a tuple with the CanNotExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetCanNotExpireOk() (*bool, bool) {
	if o == nil || IsNil(o.CanNotExpire) {
		return nil, false
	}
	return o.CanNotExpire, true
}

// HasCanNotExpire returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasCanNotExpire() bool {
	if o != nil && !IsNil(o.CanNotExpire) {
		return true
	}

	return false
}

// SetCanNotExpire gets a reference to the given bool and assigns it to the CanNotExpire field.
func (o *OrderResponseCheckout) SetCanNotExpire(v bool) {
	o.CanNotExpire = &v
}

// GetEmailsSent returns the EmailsSent field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetEmailsSent() int32 {
	if o == nil || IsNil(o.EmailsSent) {
		var ret int32
		return ret
	}
	return *o.EmailsSent
}

// GetEmailsSentOk returns a tuple with the EmailsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetEmailsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsSent) {
		return nil, false
	}
	return o.EmailsSent, true
}

// HasEmailsSent returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasEmailsSent() bool {
	if o != nil && !IsNil(o.EmailsSent) {
		return true
	}

	return false
}

// SetEmailsSent gets a reference to the given int32 and assigns it to the EmailsSent field.
func (o *OrderResponseCheckout) SetEmailsSent(v int32) {
	o.EmailsSent = &v
}

// GetExcludeCardNetworks returns the ExcludeCardNetworks field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetExcludeCardNetworks() []map[string]interface{} {
	if o == nil || IsNil(o.ExcludeCardNetworks) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ExcludeCardNetworks
}

// GetExcludeCardNetworksOk returns a tuple with the ExcludeCardNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetExcludeCardNetworksOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExcludeCardNetworks) {
		return nil, false
	}
	return o.ExcludeCardNetworks, true
}

// HasExcludeCardNetworks returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasExcludeCardNetworks() bool {
	if o != nil && !IsNil(o.ExcludeCardNetworks) {
		return true
	}

	return false
}

// SetExcludeCardNetworks gets a reference to the given []map[string]interface{} and assigns it to the ExcludeCardNetworks field.
func (o *OrderResponseCheckout) SetExcludeCardNetworks(v []map[string]interface{}) {
	o.ExcludeCardNetworks = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *OrderResponseCheckout) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetFailureUrl returns the FailureUrl field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetFailureUrl() string {
	if o == nil || IsNil(o.FailureUrl) {
		var ret string
		return ret
	}
	return *o.FailureUrl
}

// GetFailureUrlOk returns a tuple with the FailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetFailureUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FailureUrl) {
		return nil, false
	}
	return o.FailureUrl, true
}

// HasFailureUrl returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasFailureUrl() bool {
	if o != nil && !IsNil(o.FailureUrl) {
		return true
	}

	return false
}

// SetFailureUrl gets a reference to the given string and assigns it to the FailureUrl field.
func (o *OrderResponseCheckout) SetFailureUrl(v string) {
	o.FailureUrl = &v
}

// GetForce3dsFlow returns the Force3dsFlow field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetForce3dsFlow() bool {
	if o == nil || IsNil(o.Force3dsFlow) {
		var ret bool
		return ret
	}
	return *o.Force3dsFlow
}

// GetForce3dsFlowOk returns a tuple with the Force3dsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetForce3dsFlowOk() (*bool, bool) {
	if o == nil || IsNil(o.Force3dsFlow) {
		return nil, false
	}
	return o.Force3dsFlow, true
}

// HasForce3dsFlow returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasForce3dsFlow() bool {
	if o != nil && !IsNil(o.Force3dsFlow) {
		return true
	}

	return false
}

// SetForce3dsFlow gets a reference to the given bool and assigns it to the Force3dsFlow field.
func (o *OrderResponseCheckout) SetForce3dsFlow(v bool) {
	o.Force3dsFlow = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderResponseCheckout) SetId(v string) {
	o.Id = &v
}

// GetIsRedirectOnFailure returns the IsRedirectOnFailure field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetIsRedirectOnFailure() bool {
	if o == nil || IsNil(o.IsRedirectOnFailure) {
		var ret bool
		return ret
	}
	return *o.IsRedirectOnFailure
}

// GetIsRedirectOnFailureOk returns a tuple with the IsRedirectOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetIsRedirectOnFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRedirectOnFailure) {
		return nil, false
	}
	return o.IsRedirectOnFailure, true
}

// HasIsRedirectOnFailure returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasIsRedirectOnFailure() bool {
	if o != nil && !IsNil(o.IsRedirectOnFailure) {
		return true
	}

	return false
}

// SetIsRedirectOnFailure gets a reference to the given bool and assigns it to the IsRedirectOnFailure field.
func (o *OrderResponseCheckout) SetIsRedirectOnFailure(v bool) {
	o.IsRedirectOnFailure = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *OrderResponseCheckout) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *OrderResponseCheckout) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMonthlyInstallmentsEnabled returns the MonthlyInstallmentsEnabled field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetMonthlyInstallmentsEnabled() bool {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		var ret bool
		return ret
	}
	return *o.MonthlyInstallmentsEnabled
}

// GetMonthlyInstallmentsEnabledOk returns a tuple with the MonthlyInstallmentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetMonthlyInstallmentsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		return nil, false
	}
	return o.MonthlyInstallmentsEnabled, true
}

// HasMonthlyInstallmentsEnabled returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasMonthlyInstallmentsEnabled() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsEnabled) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsEnabled gets a reference to the given bool and assigns it to the MonthlyInstallmentsEnabled field.
func (o *OrderResponseCheckout) SetMonthlyInstallmentsEnabled(v bool) {
	o.MonthlyInstallmentsEnabled = &v
}

// GetMonthlyInstallmentsOptions returns the MonthlyInstallmentsOptions field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetMonthlyInstallmentsOptions() []int32 {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		var ret []int32
		return ret
	}
	return o.MonthlyInstallmentsOptions
}

// GetMonthlyInstallmentsOptionsOk returns a tuple with the MonthlyInstallmentsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetMonthlyInstallmentsOptionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		return nil, false
	}
	return o.MonthlyInstallmentsOptions, true
}

// HasMonthlyInstallmentsOptions returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasMonthlyInstallmentsOptions() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsOptions) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsOptions gets a reference to the given []int32 and assigns it to the MonthlyInstallmentsOptions field.
func (o *OrderResponseCheckout) SetMonthlyInstallmentsOptions(v []int32) {
	o.MonthlyInstallmentsOptions = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrderResponseCheckout) SetName(v string) {
	o.Name = &v
}

// GetNeedsShippingContact returns the NeedsShippingContact field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetNeedsShippingContact() bool {
	if o == nil || IsNil(o.NeedsShippingContact) {
		var ret bool
		return ret
	}
	return *o.NeedsShippingContact
}

// GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetNeedsShippingContactOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsShippingContact) {
		return nil, false
	}
	return o.NeedsShippingContact, true
}

// HasNeedsShippingContact returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasNeedsShippingContact() bool {
	if o != nil && !IsNil(o.NeedsShippingContact) {
		return true
	}

	return false
}

// SetNeedsShippingContact gets a reference to the given bool and assigns it to the NeedsShippingContact field.
func (o *OrderResponseCheckout) SetNeedsShippingContact(v bool) {
	o.NeedsShippingContact = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderResponseCheckout) SetObject(v string) {
	o.Object = &v
}

// GetOnDemandEnabled returns the OnDemandEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponseCheckout) GetOnDemandEnabled() bool {
	if o == nil || IsNil(o.OnDemandEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.OnDemandEnabled.Get()
}

// GetOnDemandEnabledOk returns a tuple with the OnDemandEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponseCheckout) GetOnDemandEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnDemandEnabled.Get(), o.OnDemandEnabled.IsSet()
}

// HasOnDemandEnabled returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasOnDemandEnabled() bool {
	if o != nil && o.OnDemandEnabled.IsSet() {
		return true
	}

	return false
}

// SetOnDemandEnabled gets a reference to the given NullableBool and assigns it to the OnDemandEnabled field.
func (o *OrderResponseCheckout) SetOnDemandEnabled(v bool) {
	o.OnDemandEnabled.Set(&v)
}

// SetOnDemandEnabledNil sets the value for OnDemandEnabled to be an explicit nil
func (o *OrderResponseCheckout) SetOnDemandEnabledNil() {
	o.OnDemandEnabled.Set(nil)
}

// UnsetOnDemandEnabled ensures that no value is present for OnDemandEnabled, not even an explicit nil
func (o *OrderResponseCheckout) UnsetOnDemandEnabled() {
	o.OnDemandEnabled.Unset()
}

// GetPaidPaymentsCount returns the PaidPaymentsCount field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetPaidPaymentsCount() int32 {
	if o == nil || IsNil(o.PaidPaymentsCount) {
		var ret int32
		return ret
	}
	return *o.PaidPaymentsCount
}

// GetPaidPaymentsCountOk returns a tuple with the PaidPaymentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetPaidPaymentsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PaidPaymentsCount) {
		return nil, false
	}
	return o.PaidPaymentsCount, true
}

// HasPaidPaymentsCount returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasPaidPaymentsCount() bool {
	if o != nil && !IsNil(o.PaidPaymentsCount) {
		return true
	}

	return false
}

// SetPaidPaymentsCount gets a reference to the given int32 and assigns it to the PaidPaymentsCount field.
func (o *OrderResponseCheckout) SetPaidPaymentsCount(v int32) {
	o.PaidPaymentsCount = &v
}

// GetRecurrent returns the Recurrent field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetRecurrent() bool {
	if o == nil || IsNil(o.Recurrent) {
		var ret bool
		return ret
	}
	return *o.Recurrent
}

// GetRecurrentOk returns a tuple with the Recurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetRecurrentOk() (*bool, bool) {
	if o == nil || IsNil(o.Recurrent) {
		return nil, false
	}
	return o.Recurrent, true
}

// HasRecurrent returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasRecurrent() bool {
	if o != nil && !IsNil(o.Recurrent) {
		return true
	}

	return false
}

// SetRecurrent gets a reference to the given bool and assigns it to the Recurrent field.
func (o *OrderResponseCheckout) SetRecurrent(v bool) {
	o.Recurrent = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *OrderResponseCheckout) SetSlug(v string) {
	o.Slug = &v
}

// GetSmsSent returns the SmsSent field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetSmsSent() int32 {
	if o == nil || IsNil(o.SmsSent) {
		var ret int32
		return ret
	}
	return *o.SmsSent
}

// GetSmsSentOk returns a tuple with the SmsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetSmsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.SmsSent) {
		return nil, false
	}
	return o.SmsSent, true
}

// HasSmsSent returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasSmsSent() bool {
	if o != nil && !IsNil(o.SmsSent) {
		return true
	}

	return false
}

// SetSmsSent gets a reference to the given int32 and assigns it to the SmsSent field.
func (o *OrderResponseCheckout) SetSmsSent(v int32) {
	o.SmsSent = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetSuccessUrl() string {
	if o == nil || IsNil(o.SuccessUrl) {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetSuccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessUrl) {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasSuccessUrl() bool {
	if o != nil && !IsNil(o.SuccessUrl) {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *OrderResponseCheckout) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetStartsAt() int32 {
	if o == nil || IsNil(o.StartsAt) {
		var ret int32
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetStartsAtOk() (*int32, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given int32 and assigns it to the StartsAt field.
func (o *OrderResponseCheckout) SetStartsAt(v int32) {
	o.StartsAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderResponseCheckout) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderResponseCheckout) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OrderResponseCheckout) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponseCheckout) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OrderResponseCheckout) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OrderResponseCheckout) SetUrl(v string) {
	o.Url = &v
}

func (o OrderResponseCheckout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponseCheckout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedPaymentMethods) {
		toSerialize["allowed_payment_methods"] = o.AllowedPaymentMethods
	}
	if !IsNil(o.CanNotExpire) {
		toSerialize["can_not_expire"] = o.CanNotExpire
	}
	if !IsNil(o.EmailsSent) {
		toSerialize["emails_sent"] = o.EmailsSent
	}
	if !IsNil(o.ExcludeCardNetworks) {
		toSerialize["exclude_card_networks"] = o.ExcludeCardNetworks
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.FailureUrl) {
		toSerialize["failure_url"] = o.FailureUrl
	}
	if !IsNil(o.Force3dsFlow) {
		toSerialize["force_3ds_flow"] = o.Force3dsFlow
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsRedirectOnFailure) {
		toSerialize["is_redirect_on_failure"] = o.IsRedirectOnFailure
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MonthlyInstallmentsEnabled) {
		toSerialize["monthly_installments_enabled"] = o.MonthlyInstallmentsEnabled
	}
	if !IsNil(o.MonthlyInstallmentsOptions) {
		toSerialize["monthly_installments_options"] = o.MonthlyInstallmentsOptions
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NeedsShippingContact) {
		toSerialize["needs_shipping_contact"] = o.NeedsShippingContact
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if o.OnDemandEnabled.IsSet() {
		toSerialize["on_demand_enabled"] = o.OnDemandEnabled.Get()
	}
	if !IsNil(o.PaidPaymentsCount) {
		toSerialize["paid_payments_count"] = o.PaidPaymentsCount
	}
	if !IsNil(o.Recurrent) {
		toSerialize["recurrent"] = o.Recurrent
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.SmsSent) {
		toSerialize["sms_sent"] = o.SmsSent
	}
	if !IsNil(o.SuccessUrl) {
		toSerialize["success_url"] = o.SuccessUrl
	}
	if !IsNil(o.StartsAt) {
		toSerialize["starts_at"] = o.StartsAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableOrderResponseCheckout struct {
	value *OrderResponseCheckout
	isSet bool
}

func (v NullableOrderResponseCheckout) Get() *OrderResponseCheckout {
	return v.value
}

func (v *NullableOrderResponseCheckout) Set(val *OrderResponseCheckout) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponseCheckout) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponseCheckout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponseCheckout(val *OrderResponseCheckout) *NullableOrderResponseCheckout {
	return &NullableOrderResponseCheckout{value: val, isSet: true}
}

func (v NullableOrderResponseCheckout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponseCheckout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the TokenResponseCheckout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenResponseCheckout{}

// TokenResponseCheckout struct for TokenResponseCheckout
type TokenResponseCheckout struct {
	AllowedPaymentMethods []string `json:"allowed_payment_methods,omitempty"`
	// Indicates if the checkout can not expire.
	CanNotExpire        *bool    `json:"can_not_expire,omitempty"`
	EmailsSent          *int32   `json:"emails_sent,omitempty"`
	ExcludeCardNetworks []string `json:"exclude_card_networks,omitempty"`
	// Date and time when the checkout expires.
	ExpiresAt *int64 `json:"expires_at,omitempty"`
	// URL to redirect the customer to if the payment process fails.
	FailureUrl *string `json:"failure_url,omitempty"`
	// Indicates if the checkout forces the 3DS flow.
	Force3dsFlow *bool                  `json:"force_3ds_flow,omitempty"`
	Id           *string                `json:"id,omitempty"`
	Livemode     *bool                  `json:"livemode,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	// Indicates if the checkout allows monthly installments.
	MonthlyInstallmentsEnabled *bool `json:"monthly_installments_enabled,omitempty"`
	// List of monthly installments options.
	MonthlyInstallmentsOptions []int32 `json:"monthly_installments_options,omitempty"`
	Name                       *string `json:"name,omitempty"`
	NeedsShippingContact       *bool   `json:"needs_shipping_contact,omitempty"`
	// Indicates the type of object, in this case checkout.
	Object *string `json:"object,omitempty"`
	// Indicates if the checkout allows on demand payments.
	OnDemandEnabled *bool `json:"on_demand_enabled,omitempty"`
	// Number of payments that have been paid.
	PaidPaymentsCount *int32 `json:"paid_payments_count,omitempty"`
	// Indicates if the checkout is recurrent.
	Recurrent *bool  `json:"recurrent,omitempty"`
	SmsSent   *int32 `json:"sms_sent,omitempty"`
	// Date and time when the checkout starts.
	StartsAt *int64 `json:"starts_at,omitempty"`
	// Status of the checkout.
	Status *string `json:"status,omitempty"`
	// URL to redirect the customer to after the payment process is completed.
	SuccessUrl *string `json:"success_url,omitempty"`
	// Type of checkout.
	Type *string `json:"type,omitempty"`
}

// NewTokenResponseCheckout instantiates a new TokenResponseCheckout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponseCheckout() *TokenResponseCheckout {
	this := TokenResponseCheckout{}
	return &this
}

// NewTokenResponseCheckoutWithDefaults instantiates a new TokenResponseCheckout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseCheckoutWithDefaults() *TokenResponseCheckout {
	this := TokenResponseCheckout{}
	return &this
}

// GetAllowedPaymentMethods returns the AllowedPaymentMethods field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetAllowedPaymentMethods() []string {
	if o == nil || IsNil(o.AllowedPaymentMethods) {
		var ret []string
		return ret
	}
	return o.AllowedPaymentMethods
}

// GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetAllowedPaymentMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedPaymentMethods) {
		return nil, false
	}
	return o.AllowedPaymentMethods, true
}

// HasAllowedPaymentMethods returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasAllowedPaymentMethods() bool {
	if o != nil && !IsNil(o.AllowedPaymentMethods) {
		return true
	}

	return false
}

// SetAllowedPaymentMethods gets a reference to the given []string and assigns it to the AllowedPaymentMethods field.
func (o *TokenResponseCheckout) SetAllowedPaymentMethods(v []string) {
	o.AllowedPaymentMethods = v
}

// GetCanNotExpire returns the CanNotExpire field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetCanNotExpire() bool {
	if o == nil || IsNil(o.CanNotExpire) {
		var ret bool
		return ret
	}
	return *o.CanNotExpire
}

// GetCanNotExpireOk returns a tuple with the CanNotExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetCanNotExpireOk() (*bool, bool) {
	if o == nil || IsNil(o.CanNotExpire) {
		return nil, false
	}
	return o.CanNotExpire, true
}

// HasCanNotExpire returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasCanNotExpire() bool {
	if o != nil && !IsNil(o.CanNotExpire) {
		return true
	}

	return false
}

// SetCanNotExpire gets a reference to the given bool and assigns it to the CanNotExpire field.
func (o *TokenResponseCheckout) SetCanNotExpire(v bool) {
	o.CanNotExpire = &v
}

// GetEmailsSent returns the EmailsSent field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetEmailsSent() int32 {
	if o == nil || IsNil(o.EmailsSent) {
		var ret int32
		return ret
	}
	return *o.EmailsSent
}

// GetEmailsSentOk returns a tuple with the EmailsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetEmailsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsSent) {
		return nil, false
	}
	return o.EmailsSent, true
}

// HasEmailsSent returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasEmailsSent() bool {
	if o != nil && !IsNil(o.EmailsSent) {
		return true
	}

	return false
}

// SetEmailsSent gets a reference to the given int32 and assigns it to the EmailsSent field.
func (o *TokenResponseCheckout) SetEmailsSent(v int32) {
	o.EmailsSent = &v
}

// GetExcludeCardNetworks returns the ExcludeCardNetworks field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetExcludeCardNetworks() []string {
	if o == nil || IsNil(o.ExcludeCardNetworks) {
		var ret []string
		return ret
	}
	return o.ExcludeCardNetworks
}

// GetExcludeCardNetworksOk returns a tuple with the ExcludeCardNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetExcludeCardNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeCardNetworks) {
		return nil, false
	}
	return o.ExcludeCardNetworks, true
}

// HasExcludeCardNetworks returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasExcludeCardNetworks() bool {
	if o != nil && !IsNil(o.ExcludeCardNetworks) {
		return true
	}

	return false
}

// SetExcludeCardNetworks gets a reference to the given []string and assigns it to the ExcludeCardNetworks field.
func (o *TokenResponseCheckout) SetExcludeCardNetworks(v []string) {
	o.ExcludeCardNetworks = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *TokenResponseCheckout) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetFailureUrl returns the FailureUrl field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetFailureUrl() string {
	if o == nil || IsNil(o.FailureUrl) {
		var ret string
		return ret
	}
	return *o.FailureUrl
}

// GetFailureUrlOk returns a tuple with the FailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetFailureUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FailureUrl) {
		return nil, false
	}
	return o.FailureUrl, true
}

// HasFailureUrl returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasFailureUrl() bool {
	if o != nil && !IsNil(o.FailureUrl) {
		return true
	}

	return false
}

// SetFailureUrl gets a reference to the given string and assigns it to the FailureUrl field.
func (o *TokenResponseCheckout) SetFailureUrl(v string) {
	o.FailureUrl = &v
}

// GetForce3dsFlow returns the Force3dsFlow field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetForce3dsFlow() bool {
	if o == nil || IsNil(o.Force3dsFlow) {
		var ret bool
		return ret
	}
	return *o.Force3dsFlow
}

// GetForce3dsFlowOk returns a tuple with the Force3dsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetForce3dsFlowOk() (*bool, bool) {
	if o == nil || IsNil(o.Force3dsFlow) {
		return nil, false
	}
	return o.Force3dsFlow, true
}

// HasForce3dsFlow returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasForce3dsFlow() bool {
	if o != nil && !IsNil(o.Force3dsFlow) {
		return true
	}

	return false
}

// SetForce3dsFlow gets a reference to the given bool and assigns it to the Force3dsFlow field.
func (o *TokenResponseCheckout) SetForce3dsFlow(v bool) {
	o.Force3dsFlow = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TokenResponseCheckout) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *TokenResponseCheckout) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *TokenResponseCheckout) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMonthlyInstallmentsEnabled returns the MonthlyInstallmentsEnabled field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetMonthlyInstallmentsEnabled() bool {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		var ret bool
		return ret
	}
	return *o.MonthlyInstallmentsEnabled
}

// GetMonthlyInstallmentsEnabledOk returns a tuple with the MonthlyInstallmentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetMonthlyInstallmentsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		return nil, false
	}
	return o.MonthlyInstallmentsEnabled, true
}

// HasMonthlyInstallmentsEnabled returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasMonthlyInstallmentsEnabled() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsEnabled) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsEnabled gets a reference to the given bool and assigns it to the MonthlyInstallmentsEnabled field.
func (o *TokenResponseCheckout) SetMonthlyInstallmentsEnabled(v bool) {
	o.MonthlyInstallmentsEnabled = &v
}

// GetMonthlyInstallmentsOptions returns the MonthlyInstallmentsOptions field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetMonthlyInstallmentsOptions() []int32 {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		var ret []int32
		return ret
	}
	return o.MonthlyInstallmentsOptions
}

// GetMonthlyInstallmentsOptionsOk returns a tuple with the MonthlyInstallmentsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetMonthlyInstallmentsOptionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		return nil, false
	}
	return o.MonthlyInstallmentsOptions, true
}

// HasMonthlyInstallmentsOptions returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasMonthlyInstallmentsOptions() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsOptions) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsOptions gets a reference to the given []int32 and assigns it to the MonthlyInstallmentsOptions field.
func (o *TokenResponseCheckout) SetMonthlyInstallmentsOptions(v []int32) {
	o.MonthlyInstallmentsOptions = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TokenResponseCheckout) SetName(v string) {
	o.Name = &v
}

// GetNeedsShippingContact returns the NeedsShippingContact field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetNeedsShippingContact() bool {
	if o == nil || IsNil(o.NeedsShippingContact) {
		var ret bool
		return ret
	}
	return *o.NeedsShippingContact
}

// GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetNeedsShippingContactOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsShippingContact) {
		return nil, false
	}
	return o.NeedsShippingContact, true
}

// HasNeedsShippingContact returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasNeedsShippingContact() bool {
	if o != nil && !IsNil(o.NeedsShippingContact) {
		return true
	}

	return false
}

// SetNeedsShippingContact gets a reference to the given bool and assigns it to the NeedsShippingContact field.
func (o *TokenResponseCheckout) SetNeedsShippingContact(v bool) {
	o.NeedsShippingContact = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *TokenResponseCheckout) SetObject(v string) {
	o.Object = &v
}

// GetOnDemandEnabled returns the OnDemandEnabled field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetOnDemandEnabled() bool {
	if o == nil || IsNil(o.OnDemandEnabled) {
		var ret bool
		return ret
	}
	return *o.OnDemandEnabled
}

// GetOnDemandEnabledOk returns a tuple with the OnDemandEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetOnDemandEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OnDemandEnabled) {
		return nil, false
	}
	return o.OnDemandEnabled, true
}

// HasOnDemandEnabled returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasOnDemandEnabled() bool {
	if o != nil && !IsNil(o.OnDemandEnabled) {
		return true
	}

	return false
}

// SetOnDemandEnabled gets a reference to the given bool and assigns it to the OnDemandEnabled field.
func (o *TokenResponseCheckout) SetOnDemandEnabled(v bool) {
	o.OnDemandEnabled = &v
}

// GetPaidPaymentsCount returns the PaidPaymentsCount field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetPaidPaymentsCount() int32 {
	if o == nil || IsNil(o.PaidPaymentsCount) {
		var ret int32
		return ret
	}
	return *o.PaidPaymentsCount
}

// GetPaidPaymentsCountOk returns a tuple with the PaidPaymentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetPaidPaymentsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PaidPaymentsCount) {
		return nil, false
	}
	return o.PaidPaymentsCount, true
}

// HasPaidPaymentsCount returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasPaidPaymentsCount() bool {
	if o != nil && !IsNil(o.PaidPaymentsCount) {
		return true
	}

	return false
}

// SetPaidPaymentsCount gets a reference to the given int32 and assigns it to the PaidPaymentsCount field.
func (o *TokenResponseCheckout) SetPaidPaymentsCount(v int32) {
	o.PaidPaymentsCount = &v
}

// GetRecurrent returns the Recurrent field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetRecurrent() bool {
	if o == nil || IsNil(o.Recurrent) {
		var ret bool
		return ret
	}
	return *o.Recurrent
}

// GetRecurrentOk returns a tuple with the Recurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetRecurrentOk() (*bool, bool) {
	if o == nil || IsNil(o.Recurrent) {
		return nil, false
	}
	return o.Recurrent, true
}

// HasRecurrent returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasRecurrent() bool {
	if o != nil && !IsNil(o.Recurrent) {
		return true
	}

	return false
}

// SetRecurrent gets a reference to the given bool and assigns it to the Recurrent field.
func (o *TokenResponseCheckout) SetRecurrent(v bool) {
	o.Recurrent = &v
}

// GetSmsSent returns the SmsSent field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetSmsSent() int32 {
	if o == nil || IsNil(o.SmsSent) {
		var ret int32
		return ret
	}
	return *o.SmsSent
}

// GetSmsSentOk returns a tuple with the SmsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetSmsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.SmsSent) {
		return nil, false
	}
	return o.SmsSent, true
}

// HasSmsSent returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasSmsSent() bool {
	if o != nil && !IsNil(o.SmsSent) {
		return true
	}

	return false
}

// SetSmsSent gets a reference to the given int32 and assigns it to the SmsSent field.
func (o *TokenResponseCheckout) SetSmsSent(v int32) {
	o.SmsSent = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetStartsAt() int64 {
	if o == nil || IsNil(o.StartsAt) {
		var ret int64
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetStartsAtOk() (*int64, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given int64 and assigns it to the StartsAt field.
func (o *TokenResponseCheckout) SetStartsAt(v int64) {
	o.StartsAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TokenResponseCheckout) SetStatus(v string) {
	o.Status = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetSuccessUrl() string {
	if o == nil || IsNil(o.SuccessUrl) {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetSuccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessUrl) {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasSuccessUrl() bool {
	if o != nil && !IsNil(o.SuccessUrl) {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *TokenResponseCheckout) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TokenResponseCheckout) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseCheckout) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TokenResponseCheckout) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TokenResponseCheckout) SetType(v string) {
	o.Type = &v
}

func (o TokenResponseCheckout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenResponseCheckout) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.OnDemandEnabled) {
		toSerialize["on_demand_enabled"] = o.OnDemandEnabled
	}
	if !IsNil(o.PaidPaymentsCount) {
		toSerialize["paid_payments_count"] = o.PaidPaymentsCount
	}
	if !IsNil(o.Recurrent) {
		toSerialize["recurrent"] = o.Recurrent
	}
	if !IsNil(o.SmsSent) {
		toSerialize["sms_sent"] = o.SmsSent
	}
	if !IsNil(o.StartsAt) {
		toSerialize["starts_at"] = o.StartsAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SuccessUrl) {
		toSerialize["success_url"] = o.SuccessUrl
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTokenResponseCheckout struct {
	value *TokenResponseCheckout
	isSet bool
}

func (v NullableTokenResponseCheckout) Get() *TokenResponseCheckout {
	return v.value
}

func (v *NullableTokenResponseCheckout) Set(val *TokenResponseCheckout) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenResponseCheckout) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenResponseCheckout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenResponseCheckout(val *TokenResponseCheckout) *NullableTokenResponseCheckout {
	return &NullableTokenResponseCheckout{value: val, isSet: true}
}

func (v NullableTokenResponseCheckout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenResponseCheckout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

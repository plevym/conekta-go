package conekta

import (
	"encoding/json"
)

// checks if the CheckoutResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutResponse{}

// CheckoutResponse checkout response
type CheckoutResponse struct {
	AllowedPaymentMethods      []string                 `json:"allowed_payment_methods,omitempty"`
	CanNotExpire               *bool                    `json:"can_not_expire,omitempty"`
	EmailsSent                 *int32                   `json:"emails_sent,omitempty"`
	ExcludeCardNetworks        []map[string]interface{} `json:"exclude_card_networks,omitempty"`
	ExpiresAt                  *int64                   `json:"expires_at,omitempty"`
	FailureUrl                 *string                  `json:"failure_url,omitempty"`
	Force3dsFlow               *bool                    `json:"force_3ds_flow,omitempty"`
	Id                         string                   `json:"id"`
	Livemode                   bool                     `json:"livemode"`
	Metadata                   map[string]interface{}   `json:"metadata,omitempty"`
	MonthlyInstallmentsEnabled *bool                    `json:"monthly_installments_enabled,omitempty"`
	MonthlyInstallmentsOptions []int32                  `json:"monthly_installments_options,omitempty"`
	// Reason for charge
	Name                 string        `json:"name"`
	NeedsShippingContact *bool         `json:"needs_shipping_contact,omitempty"`
	Object               string        `json:"object"`
	PaidPaymentsCount    *int32        `json:"paid_payments_count,omitempty"`
	PaymentsLimitCount   NullableInt32 `json:"payments_limit_count,omitempty"`
	Recurrent            *bool         `json:"recurrent,omitempty"`
	Slug                 *string       `json:"slug,omitempty"`
	SmsSent              *int32        `json:"sms_sent,omitempty"`
	StartsAt             *int32        `json:"starts_at,omitempty"`
	Status               *string       `json:"status,omitempty"`
	SuccessUrl           *string       `json:"success_url,omitempty"`
	Type                 *string       `json:"type,omitempty"`
	Url                  *string       `json:"url,omitempty"`
}

// NewCheckoutResponse instantiates a new CheckoutResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutResponse(id string, livemode bool, name string, object string) *CheckoutResponse {
	this := CheckoutResponse{}
	this.Id = id
	this.Livemode = livemode
	this.Name = name
	this.Object = object
	return &this
}

// NewCheckoutResponseWithDefaults instantiates a new CheckoutResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutResponseWithDefaults() *CheckoutResponse {
	this := CheckoutResponse{}
	return &this
}

// GetAllowedPaymentMethods returns the AllowedPaymentMethods field value if set, zero value otherwise.
func (o *CheckoutResponse) GetAllowedPaymentMethods() []string {
	if o == nil || IsNil(o.AllowedPaymentMethods) {
		var ret []string
		return ret
	}
	return o.AllowedPaymentMethods
}

// GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetAllowedPaymentMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedPaymentMethods) {
		return nil, false
	}
	return o.AllowedPaymentMethods, true
}

// HasAllowedPaymentMethods returns a boolean if a field has been set.
func (o *CheckoutResponse) HasAllowedPaymentMethods() bool {
	if o != nil && !IsNil(o.AllowedPaymentMethods) {
		return true
	}

	return false
}

// SetAllowedPaymentMethods gets a reference to the given []string and assigns it to the AllowedPaymentMethods field.
func (o *CheckoutResponse) SetAllowedPaymentMethods(v []string) {
	o.AllowedPaymentMethods = v
}

// GetCanNotExpire returns the CanNotExpire field value if set, zero value otherwise.
func (o *CheckoutResponse) GetCanNotExpire() bool {
	if o == nil || IsNil(o.CanNotExpire) {
		var ret bool
		return ret
	}
	return *o.CanNotExpire
}

// GetCanNotExpireOk returns a tuple with the CanNotExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetCanNotExpireOk() (*bool, bool) {
	if o == nil || IsNil(o.CanNotExpire) {
		return nil, false
	}
	return o.CanNotExpire, true
}

// HasCanNotExpire returns a boolean if a field has been set.
func (o *CheckoutResponse) HasCanNotExpire() bool {
	if o != nil && !IsNil(o.CanNotExpire) {
		return true
	}

	return false
}

// SetCanNotExpire gets a reference to the given bool and assigns it to the CanNotExpire field.
func (o *CheckoutResponse) SetCanNotExpire(v bool) {
	o.CanNotExpire = &v
}

// GetEmailsSent returns the EmailsSent field value if set, zero value otherwise.
func (o *CheckoutResponse) GetEmailsSent() int32 {
	if o == nil || IsNil(o.EmailsSent) {
		var ret int32
		return ret
	}
	return *o.EmailsSent
}

// GetEmailsSentOk returns a tuple with the EmailsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetEmailsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsSent) {
		return nil, false
	}
	return o.EmailsSent, true
}

// HasEmailsSent returns a boolean if a field has been set.
func (o *CheckoutResponse) HasEmailsSent() bool {
	if o != nil && !IsNil(o.EmailsSent) {
		return true
	}

	return false
}

// SetEmailsSent gets a reference to the given int32 and assigns it to the EmailsSent field.
func (o *CheckoutResponse) SetEmailsSent(v int32) {
	o.EmailsSent = &v
}

// GetExcludeCardNetworks returns the ExcludeCardNetworks field value if set, zero value otherwise.
func (o *CheckoutResponse) GetExcludeCardNetworks() []map[string]interface{} {
	if o == nil || IsNil(o.ExcludeCardNetworks) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ExcludeCardNetworks
}

// GetExcludeCardNetworksOk returns a tuple with the ExcludeCardNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetExcludeCardNetworksOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExcludeCardNetworks) {
		return nil, false
	}
	return o.ExcludeCardNetworks, true
}

// HasExcludeCardNetworks returns a boolean if a field has been set.
func (o *CheckoutResponse) HasExcludeCardNetworks() bool {
	if o != nil && !IsNil(o.ExcludeCardNetworks) {
		return true
	}

	return false
}

// SetExcludeCardNetworks gets a reference to the given []map[string]interface{} and assigns it to the ExcludeCardNetworks field.
func (o *CheckoutResponse) SetExcludeCardNetworks(v []map[string]interface{}) {
	o.ExcludeCardNetworks = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CheckoutResponse) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CheckoutResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *CheckoutResponse) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetFailureUrl returns the FailureUrl field value if set, zero value otherwise.
func (o *CheckoutResponse) GetFailureUrl() string {
	if o == nil || IsNil(o.FailureUrl) {
		var ret string
		return ret
	}
	return *o.FailureUrl
}

// GetFailureUrlOk returns a tuple with the FailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetFailureUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FailureUrl) {
		return nil, false
	}
	return o.FailureUrl, true
}

// HasFailureUrl returns a boolean if a field has been set.
func (o *CheckoutResponse) HasFailureUrl() bool {
	if o != nil && !IsNil(o.FailureUrl) {
		return true
	}

	return false
}

// SetFailureUrl gets a reference to the given string and assigns it to the FailureUrl field.
func (o *CheckoutResponse) SetFailureUrl(v string) {
	o.FailureUrl = &v
}

// GetForce3dsFlow returns the Force3dsFlow field value if set, zero value otherwise.
func (o *CheckoutResponse) GetForce3dsFlow() bool {
	if o == nil || IsNil(o.Force3dsFlow) {
		var ret bool
		return ret
	}
	return *o.Force3dsFlow
}

// GetForce3dsFlowOk returns a tuple with the Force3dsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetForce3dsFlowOk() (*bool, bool) {
	if o == nil || IsNil(o.Force3dsFlow) {
		return nil, false
	}
	return o.Force3dsFlow, true
}

// HasForce3dsFlow returns a boolean if a field has been set.
func (o *CheckoutResponse) HasForce3dsFlow() bool {
	if o != nil && !IsNil(o.Force3dsFlow) {
		return true
	}

	return false
}

// SetForce3dsFlow gets a reference to the given bool and assigns it to the Force3dsFlow field.
func (o *CheckoutResponse) SetForce3dsFlow(v bool) {
	o.Force3dsFlow = &v
}

// GetId returns the Id field value
func (o *CheckoutResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CheckoutResponse) SetId(v string) {
	o.Id = v
}

// GetLivemode returns the Livemode field value
func (o *CheckoutResponse) GetLivemode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Livemode, true
}

// SetLivemode sets field value
func (o *CheckoutResponse) SetLivemode(v bool) {
	o.Livemode = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CheckoutResponse) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CheckoutResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CheckoutResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMonthlyInstallmentsEnabled returns the MonthlyInstallmentsEnabled field value if set, zero value otherwise.
func (o *CheckoutResponse) GetMonthlyInstallmentsEnabled() bool {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		var ret bool
		return ret
	}
	return *o.MonthlyInstallmentsEnabled
}

// GetMonthlyInstallmentsEnabledOk returns a tuple with the MonthlyInstallmentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetMonthlyInstallmentsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		return nil, false
	}
	return o.MonthlyInstallmentsEnabled, true
}

// HasMonthlyInstallmentsEnabled returns a boolean if a field has been set.
func (o *CheckoutResponse) HasMonthlyInstallmentsEnabled() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsEnabled) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsEnabled gets a reference to the given bool and assigns it to the MonthlyInstallmentsEnabled field.
func (o *CheckoutResponse) SetMonthlyInstallmentsEnabled(v bool) {
	o.MonthlyInstallmentsEnabled = &v
}

// GetMonthlyInstallmentsOptions returns the MonthlyInstallmentsOptions field value if set, zero value otherwise.
func (o *CheckoutResponse) GetMonthlyInstallmentsOptions() []int32 {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		var ret []int32
		return ret
	}
	return o.MonthlyInstallmentsOptions
}

// GetMonthlyInstallmentsOptionsOk returns a tuple with the MonthlyInstallmentsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetMonthlyInstallmentsOptionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		return nil, false
	}
	return o.MonthlyInstallmentsOptions, true
}

// HasMonthlyInstallmentsOptions returns a boolean if a field has been set.
func (o *CheckoutResponse) HasMonthlyInstallmentsOptions() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsOptions) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsOptions gets a reference to the given []int32 and assigns it to the MonthlyInstallmentsOptions field.
func (o *CheckoutResponse) SetMonthlyInstallmentsOptions(v []int32) {
	o.MonthlyInstallmentsOptions = v
}

// GetName returns the Name field value
func (o *CheckoutResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CheckoutResponse) SetName(v string) {
	o.Name = v
}

// GetNeedsShippingContact returns the NeedsShippingContact field value if set, zero value otherwise.
func (o *CheckoutResponse) GetNeedsShippingContact() bool {
	if o == nil || IsNil(o.NeedsShippingContact) {
		var ret bool
		return ret
	}
	return *o.NeedsShippingContact
}

// GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetNeedsShippingContactOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsShippingContact) {
		return nil, false
	}
	return o.NeedsShippingContact, true
}

// HasNeedsShippingContact returns a boolean if a field has been set.
func (o *CheckoutResponse) HasNeedsShippingContact() bool {
	if o != nil && !IsNil(o.NeedsShippingContact) {
		return true
	}

	return false
}

// SetNeedsShippingContact gets a reference to the given bool and assigns it to the NeedsShippingContact field.
func (o *CheckoutResponse) SetNeedsShippingContact(v bool) {
	o.NeedsShippingContact = &v
}

// GetObject returns the Object field value
func (o *CheckoutResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CheckoutResponse) SetObject(v string) {
	o.Object = v
}

// GetPaidPaymentsCount returns the PaidPaymentsCount field value if set, zero value otherwise.
func (o *CheckoutResponse) GetPaidPaymentsCount() int32 {
	if o == nil || IsNil(o.PaidPaymentsCount) {
		var ret int32
		return ret
	}
	return *o.PaidPaymentsCount
}

// GetPaidPaymentsCountOk returns a tuple with the PaidPaymentsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetPaidPaymentsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PaidPaymentsCount) {
		return nil, false
	}
	return o.PaidPaymentsCount, true
}

// HasPaidPaymentsCount returns a boolean if a field has been set.
func (o *CheckoutResponse) HasPaidPaymentsCount() bool {
	if o != nil && !IsNil(o.PaidPaymentsCount) {
		return true
	}

	return false
}

// SetPaidPaymentsCount gets a reference to the given int32 and assigns it to the PaidPaymentsCount field.
func (o *CheckoutResponse) SetPaidPaymentsCount(v int32) {
	o.PaidPaymentsCount = &v
}

// GetPaymentsLimitCount returns the PaymentsLimitCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutResponse) GetPaymentsLimitCount() int32 {
	if o == nil || IsNil(o.PaymentsLimitCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PaymentsLimitCount.Get()
}

// GetPaymentsLimitCountOk returns a tuple with the PaymentsLimitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutResponse) GetPaymentsLimitCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentsLimitCount.Get(), o.PaymentsLimitCount.IsSet()
}

// HasPaymentsLimitCount returns a boolean if a field has been set.
func (o *CheckoutResponse) HasPaymentsLimitCount() bool {
	if o != nil && o.PaymentsLimitCount.IsSet() {
		return true
	}

	return false
}

// SetPaymentsLimitCount gets a reference to the given NullableInt32 and assigns it to the PaymentsLimitCount field.
func (o *CheckoutResponse) SetPaymentsLimitCount(v int32) {
	o.PaymentsLimitCount.Set(&v)
}

// SetPaymentsLimitCountNil sets the value for PaymentsLimitCount to be an explicit nil
func (o *CheckoutResponse) SetPaymentsLimitCountNil() {
	o.PaymentsLimitCount.Set(nil)
}

// UnsetPaymentsLimitCount ensures that no value is present for PaymentsLimitCount, not even an explicit nil
func (o *CheckoutResponse) UnsetPaymentsLimitCount() {
	o.PaymentsLimitCount.Unset()
}

// GetRecurrent returns the Recurrent field value if set, zero value otherwise.
func (o *CheckoutResponse) GetRecurrent() bool {
	if o == nil || IsNil(o.Recurrent) {
		var ret bool
		return ret
	}
	return *o.Recurrent
}

// GetRecurrentOk returns a tuple with the Recurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetRecurrentOk() (*bool, bool) {
	if o == nil || IsNil(o.Recurrent) {
		return nil, false
	}
	return o.Recurrent, true
}

// HasRecurrent returns a boolean if a field has been set.
func (o *CheckoutResponse) HasRecurrent() bool {
	if o != nil && !IsNil(o.Recurrent) {
		return true
	}

	return false
}

// SetRecurrent gets a reference to the given bool and assigns it to the Recurrent field.
func (o *CheckoutResponse) SetRecurrent(v bool) {
	o.Recurrent = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *CheckoutResponse) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *CheckoutResponse) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *CheckoutResponse) SetSlug(v string) {
	o.Slug = &v
}

// GetSmsSent returns the SmsSent field value if set, zero value otherwise.
func (o *CheckoutResponse) GetSmsSent() int32 {
	if o == nil || IsNil(o.SmsSent) {
		var ret int32
		return ret
	}
	return *o.SmsSent
}

// GetSmsSentOk returns a tuple with the SmsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetSmsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.SmsSent) {
		return nil, false
	}
	return o.SmsSent, true
}

// HasSmsSent returns a boolean if a field has been set.
func (o *CheckoutResponse) HasSmsSent() bool {
	if o != nil && !IsNil(o.SmsSent) {
		return true
	}

	return false
}

// SetSmsSent gets a reference to the given int32 and assigns it to the SmsSent field.
func (o *CheckoutResponse) SetSmsSent(v int32) {
	o.SmsSent = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *CheckoutResponse) GetStartsAt() int32 {
	if o == nil || IsNil(o.StartsAt) {
		var ret int32
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetStartsAtOk() (*int32, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *CheckoutResponse) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given int32 and assigns it to the StartsAt field.
func (o *CheckoutResponse) SetStartsAt(v int32) {
	o.StartsAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CheckoutResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CheckoutResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CheckoutResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *CheckoutResponse) GetSuccessUrl() string {
	if o == nil || IsNil(o.SuccessUrl) {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetSuccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessUrl) {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *CheckoutResponse) HasSuccessUrl() bool {
	if o != nil && !IsNil(o.SuccessUrl) {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *CheckoutResponse) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CheckoutResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CheckoutResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CheckoutResponse) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CheckoutResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CheckoutResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CheckoutResponse) SetUrl(v string) {
	o.Url = &v
}

func (o CheckoutResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["id"] = o.Id
	toSerialize["livemode"] = o.Livemode
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
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
	toSerialize["object"] = o.Object
	if !IsNil(o.PaidPaymentsCount) {
		toSerialize["paid_payments_count"] = o.PaidPaymentsCount
	}
	if o.PaymentsLimitCount.IsSet() {
		toSerialize["payments_limit_count"] = o.PaymentsLimitCount.Get()
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
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCheckoutResponse struct {
	value *CheckoutResponse
	isSet bool
}

func (v NullableCheckoutResponse) Get() *CheckoutResponse {
	return v.value
}

func (v *NullableCheckoutResponse) Set(val *CheckoutResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutResponse(val *CheckoutResponse) *NullableCheckoutResponse {
	return &NullableCheckoutResponse{value: val, isSet: true}
}

func (v NullableCheckoutResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

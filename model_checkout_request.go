package conekta

import (
	"encoding/json"
)

// checks if the CheckoutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutRequest{}

// CheckoutRequest [Checkout](https://developers.conekta.com/v2.1.0/reference/payment-link) details
type CheckoutRequest struct {
	// Are the payment methods available for this link
	AllowedPaymentMethods []string `json:"allowed_payment_methods"`
	// Unix timestamp of checkout expiration
	ExpiresAt *int64 `json:"expires_at,omitempty"`
	// Redirection url back to the site in case of failed payment, applies only to HostedPayment.
	FailureUrl                 *string `json:"failure_url,omitempty"`
	MonthlyInstallmentsEnabled *bool   `json:"monthly_installments_enabled,omitempty"`
	MonthlyInstallmentsOptions []int32 `json:"monthly_installments_options,omitempty"`
	// Reason for payment
	Name            *string `json:"name,omitempty"`
	OnDemandEnabled *bool   `json:"on_demand_enabled,omitempty"`
	// Redirection url back to the site in case of successful payment, applies only to HostedPayment
	SuccessUrl *string `json:"success_url,omitempty"`
	// This field represents the type of checkout
	Type *string `json:"type,omitempty"`
}

// NewCheckoutRequest instantiates a new CheckoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutRequest(allowedPaymentMethods []string) *CheckoutRequest {
	this := CheckoutRequest{}
	this.AllowedPaymentMethods = allowedPaymentMethods
	return &this
}

// NewCheckoutRequestWithDefaults instantiates a new CheckoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutRequestWithDefaults() *CheckoutRequest {
	this := CheckoutRequest{}
	return &this
}

// GetAllowedPaymentMethods returns the AllowedPaymentMethods field value
func (o *CheckoutRequest) GetAllowedPaymentMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedPaymentMethods
}

// GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field value
// and a boolean to check if the value has been set.
func (o *CheckoutRequest) GetAllowedPaymentMethodsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedPaymentMethods, true
}

// SetAllowedPaymentMethods sets field value
func (o *CheckoutRequest) SetAllowedPaymentMethods(v []string) {
	o.AllowedPaymentMethods = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CheckoutRequest) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRequest) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CheckoutRequest) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *CheckoutRequest) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetFailureUrl returns the FailureUrl field value if set, zero value otherwise.
func (o *CheckoutRequest) GetFailureUrl() string {
	if o == nil || IsNil(o.FailureUrl) {
		var ret string
		return ret
	}
	return *o.FailureUrl
}

// GetFailureUrlOk returns a tuple with the FailureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRequest) GetFailureUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FailureUrl) {
		return nil, false
	}
	return o.FailureUrl, true
}

// HasFailureUrl returns a boolean if a field has been set.
func (o *CheckoutRequest) HasFailureUrl() bool {
	if o != nil && !IsNil(o.FailureUrl) {
		return true
	}

	return false
}

// SetFailureUrl gets a reference to the given string and assigns it to the FailureUrl field.
func (o *CheckoutRequest) SetFailureUrl(v string) {
	o.FailureUrl = &v
}

// GetMonthlyInstallmentsEnabled returns the MonthlyInstallmentsEnabled field value if set, zero value otherwise.
func (o *CheckoutRequest) GetMonthlyInstallmentsEnabled() bool {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		var ret bool
		return ret
	}
	return *o.MonthlyInstallmentsEnabled
}

// GetMonthlyInstallmentsEnabledOk returns a tuple with the MonthlyInstallmentsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRequest) GetMonthlyInstallmentsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsEnabled) {
		return nil, false
	}
	return o.MonthlyInstallmentsEnabled, true
}

// HasMonthlyInstallmentsEnabled returns a boolean if a field has been set.
func (o *CheckoutRequest) HasMonthlyInstallmentsEnabled() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsEnabled) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsEnabled gets a reference to the given bool and assigns it to the MonthlyInstallmentsEnabled field.
func (o *CheckoutRequest) SetMonthlyInstallmentsEnabled(v bool) {
	o.MonthlyInstallmentsEnabled = &v
}

// GetMonthlyInstallmentsOptions returns the MonthlyInstallmentsOptions field value if set, zero value otherwise.
func (o *CheckoutRequest) GetMonthlyInstallmentsOptions() []int32 {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		var ret []int32
		return ret
	}
	return o.MonthlyInstallmentsOptions
}

// GetMonthlyInstallmentsOptionsOk returns a tuple with the MonthlyInstallmentsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRequest) GetMonthlyInstallmentsOptionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MonthlyInstallmentsOptions) {
		return nil, false
	}
	return o.MonthlyInstallmentsOptions, true
}

// HasMonthlyInstallmentsOptions returns a boolean if a field has been set.
func (o *CheckoutRequest) HasMonthlyInstallmentsOptions() bool {
	if o != nil && !IsNil(o.MonthlyInstallmentsOptions) {
		return true
	}

	return false
}

// SetMonthlyInstallmentsOptions gets a reference to the given []int32 and assigns it to the MonthlyInstallmentsOptions field.
func (o *CheckoutRequest) SetMonthlyInstallmentsOptions(v []int32) {
	o.MonthlyInstallmentsOptions = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CheckoutRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CheckoutRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CheckoutRequest) SetName(v string) {
	o.Name = &v
}

// GetOnDemandEnabled returns the OnDemandEnabled field value if set, zero value otherwise.
func (o *CheckoutRequest) GetOnDemandEnabled() bool {
	if o == nil || IsNil(o.OnDemandEnabled) {
		var ret bool
		return ret
	}
	return *o.OnDemandEnabled
}

// GetOnDemandEnabledOk returns a tuple with the OnDemandEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRequest) GetOnDemandEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OnDemandEnabled) {
		return nil, false
	}
	return o.OnDemandEnabled, true
}

// HasOnDemandEnabled returns a boolean if a field has been set.
func (o *CheckoutRequest) HasOnDemandEnabled() bool {
	if o != nil && !IsNil(o.OnDemandEnabled) {
		return true
	}

	return false
}

// SetOnDemandEnabled gets a reference to the given bool and assigns it to the OnDemandEnabled field.
func (o *CheckoutRequest) SetOnDemandEnabled(v bool) {
	o.OnDemandEnabled = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *CheckoutRequest) GetSuccessUrl() string {
	if o == nil || IsNil(o.SuccessUrl) {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRequest) GetSuccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessUrl) {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *CheckoutRequest) HasSuccessUrl() bool {
	if o != nil && !IsNil(o.SuccessUrl) {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *CheckoutRequest) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CheckoutRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CheckoutRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CheckoutRequest) SetType(v string) {
	o.Type = &v
}

func (o CheckoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowed_payment_methods"] = o.AllowedPaymentMethods
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.FailureUrl) {
		toSerialize["failure_url"] = o.FailureUrl
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
	if !IsNil(o.OnDemandEnabled) {
		toSerialize["on_demand_enabled"] = o.OnDemandEnabled
	}
	if !IsNil(o.SuccessUrl) {
		toSerialize["success_url"] = o.SuccessUrl
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCheckoutRequest struct {
	value *CheckoutRequest
	isSet bool
}

func (v NullableCheckoutRequest) Get() *CheckoutRequest {
	return v.value
}

func (v *NullableCheckoutRequest) Set(val *CheckoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutRequest(val *CheckoutRequest) *NullableCheckoutRequest {
	return &NullableCheckoutRequest{value: val, isSet: true}
}

func (v NullableCheckoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

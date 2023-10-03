package conekta

import (
	"encoding/json"
)

// checks if the SubscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionResponse{}

// SubscriptionResponse subscription model
type SubscriptionResponse struct {
	BillingCycleStart       NullableInt64  `json:"billing_cycle_start,omitempty"`
	BillingCycleEnd         NullableInt64  `json:"billing_cycle_end,omitempty"`
	CanceledAt              NullableInt64  `json:"canceled_at,omitempty"`
	CardId                  *string        `json:"card_id,omitempty"`
	ChargeId                NullableString `json:"charge_id,omitempty"`
	CreatedAt               *int64         `json:"created_at,omitempty"`
	CustomerCustomReference *string        `json:"customer_custom_reference,omitempty"`
	CustomerId              *string        `json:"customer_id,omitempty"`
	Id                      *string        `json:"id,omitempty"`
	LastBillingCycleOrderId *string        `json:"last_billing_cycle_order_id,omitempty"`
	Object                  *string        `json:"object,omitempty"`
	PausedAt                NullableInt64  `json:"paused_at,omitempty"`
	PlanId                  *string        `json:"plan_id,omitempty"`
	Status                  *string        `json:"status,omitempty"`
	SubscriptionStart       *int32         `json:"subscription_start,omitempty"`
	TrialStart              NullableInt64  `json:"trial_start,omitempty"`
	TrialEnd                NullableInt64  `json:"trial_end,omitempty"`
}

// NewSubscriptionResponse instantiates a new SubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionResponse() *SubscriptionResponse {
	this := SubscriptionResponse{}
	return &this
}

// NewSubscriptionResponseWithDefaults instantiates a new SubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionResponseWithDefaults() *SubscriptionResponse {
	this := SubscriptionResponse{}
	return &this
}

// GetBillingCycleStart returns the BillingCycleStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionResponse) GetBillingCycleStart() int64 {
	if o == nil || IsNil(o.BillingCycleStart.Get()) {
		var ret int64
		return ret
	}
	return *o.BillingCycleStart.Get()
}

// GetBillingCycleStartOk returns a tuple with the BillingCycleStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionResponse) GetBillingCycleStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingCycleStart.Get(), o.BillingCycleStart.IsSet()
}

// HasBillingCycleStart returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasBillingCycleStart() bool {
	if o != nil && o.BillingCycleStart.IsSet() {
		return true
	}

	return false
}

// SetBillingCycleStart gets a reference to the given NullableInt64 and assigns it to the BillingCycleStart field.
func (o *SubscriptionResponse) SetBillingCycleStart(v int64) {
	o.BillingCycleStart.Set(&v)
}

// SetBillingCycleStartNil sets the value for BillingCycleStart to be an explicit nil
func (o *SubscriptionResponse) SetBillingCycleStartNil() {
	o.BillingCycleStart.Set(nil)
}

// UnsetBillingCycleStart ensures that no value is present for BillingCycleStart, not even an explicit nil
func (o *SubscriptionResponse) UnsetBillingCycleStart() {
	o.BillingCycleStart.Unset()
}

// GetBillingCycleEnd returns the BillingCycleEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionResponse) GetBillingCycleEnd() int64 {
	if o == nil || IsNil(o.BillingCycleEnd.Get()) {
		var ret int64
		return ret
	}
	return *o.BillingCycleEnd.Get()
}

// GetBillingCycleEndOk returns a tuple with the BillingCycleEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionResponse) GetBillingCycleEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingCycleEnd.Get(), o.BillingCycleEnd.IsSet()
}

// HasBillingCycleEnd returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasBillingCycleEnd() bool {
	if o != nil && o.BillingCycleEnd.IsSet() {
		return true
	}

	return false
}

// SetBillingCycleEnd gets a reference to the given NullableInt64 and assigns it to the BillingCycleEnd field.
func (o *SubscriptionResponse) SetBillingCycleEnd(v int64) {
	o.BillingCycleEnd.Set(&v)
}

// SetBillingCycleEndNil sets the value for BillingCycleEnd to be an explicit nil
func (o *SubscriptionResponse) SetBillingCycleEndNil() {
	o.BillingCycleEnd.Set(nil)
}

// UnsetBillingCycleEnd ensures that no value is present for BillingCycleEnd, not even an explicit nil
func (o *SubscriptionResponse) UnsetBillingCycleEnd() {
	o.BillingCycleEnd.Unset()
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionResponse) GetCanceledAt() int64 {
	if o == nil || IsNil(o.CanceledAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CanceledAt.Get()
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionResponse) GetCanceledAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanceledAt.Get(), o.CanceledAt.IsSet()
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasCanceledAt() bool {
	if o != nil && o.CanceledAt.IsSet() {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given NullableInt64 and assigns it to the CanceledAt field.
func (o *SubscriptionResponse) SetCanceledAt(v int64) {
	o.CanceledAt.Set(&v)
}

// SetCanceledAtNil sets the value for CanceledAt to be an explicit nil
func (o *SubscriptionResponse) SetCanceledAtNil() {
	o.CanceledAt.Set(nil)
}

// UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
func (o *SubscriptionResponse) UnsetCanceledAt() {
	o.CanceledAt.Unset()
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetCardId() string {
	if o == nil || IsNil(o.CardId) {
		var ret string
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetCardIdOk() (*string, bool) {
	if o == nil || IsNil(o.CardId) {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasCardId() bool {
	if o != nil && !IsNil(o.CardId) {
		return true
	}

	return false
}

// SetCardId gets a reference to the given string and assigns it to the CardId field.
func (o *SubscriptionResponse) SetCardId(v string) {
	o.CardId = &v
}

// GetChargeId returns the ChargeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionResponse) GetChargeId() string {
	if o == nil || IsNil(o.ChargeId.Get()) {
		var ret string
		return ret
	}
	return *o.ChargeId.Get()
}

// GetChargeIdOk returns a tuple with the ChargeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionResponse) GetChargeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargeId.Get(), o.ChargeId.IsSet()
}

// HasChargeId returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasChargeId() bool {
	if o != nil && o.ChargeId.IsSet() {
		return true
	}

	return false
}

// SetChargeId gets a reference to the given NullableString and assigns it to the ChargeId field.
func (o *SubscriptionResponse) SetChargeId(v string) {
	o.ChargeId.Set(&v)
}

// SetChargeIdNil sets the value for ChargeId to be an explicit nil
func (o *SubscriptionResponse) SetChargeIdNil() {
	o.ChargeId.Set(nil)
}

// UnsetChargeId ensures that no value is present for ChargeId, not even an explicit nil
func (o *SubscriptionResponse) UnsetChargeId() {
	o.ChargeId.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *SubscriptionResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCustomerCustomReference returns the CustomerCustomReference field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetCustomerCustomReference() string {
	if o == nil || IsNil(o.CustomerCustomReference) {
		var ret string
		return ret
	}
	return *o.CustomerCustomReference
}

// GetCustomerCustomReferenceOk returns a tuple with the CustomerCustomReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetCustomerCustomReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerCustomReference) {
		return nil, false
	}
	return o.CustomerCustomReference, true
}

// HasCustomerCustomReference returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasCustomerCustomReference() bool {
	if o != nil && !IsNil(o.CustomerCustomReference) {
		return true
	}

	return false
}

// SetCustomerCustomReference gets a reference to the given string and assigns it to the CustomerCustomReference field.
func (o *SubscriptionResponse) SetCustomerCustomReference(v string) {
	o.CustomerCustomReference = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *SubscriptionResponse) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionResponse) SetId(v string) {
	o.Id = &v
}

// GetLastBillingCycleOrderId returns the LastBillingCycleOrderId field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetLastBillingCycleOrderId() string {
	if o == nil || IsNil(o.LastBillingCycleOrderId) {
		var ret string
		return ret
	}
	return *o.LastBillingCycleOrderId
}

// GetLastBillingCycleOrderIdOk returns a tuple with the LastBillingCycleOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetLastBillingCycleOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastBillingCycleOrderId) {
		return nil, false
	}
	return o.LastBillingCycleOrderId, true
}

// HasLastBillingCycleOrderId returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasLastBillingCycleOrderId() bool {
	if o != nil && !IsNil(o.LastBillingCycleOrderId) {
		return true
	}

	return false
}

// SetLastBillingCycleOrderId gets a reference to the given string and assigns it to the LastBillingCycleOrderId field.
func (o *SubscriptionResponse) SetLastBillingCycleOrderId(v string) {
	o.LastBillingCycleOrderId = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *SubscriptionResponse) SetObject(v string) {
	o.Object = &v
}

// GetPausedAt returns the PausedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionResponse) GetPausedAt() int64 {
	if o == nil || IsNil(o.PausedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.PausedAt.Get()
}

// GetPausedAtOk returns a tuple with the PausedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionResponse) GetPausedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PausedAt.Get(), o.PausedAt.IsSet()
}

// HasPausedAt returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasPausedAt() bool {
	if o != nil && o.PausedAt.IsSet() {
		return true
	}

	return false
}

// SetPausedAt gets a reference to the given NullableInt64 and assigns it to the PausedAt field.
func (o *SubscriptionResponse) SetPausedAt(v int64) {
	o.PausedAt.Set(&v)
}

// SetPausedAtNil sets the value for PausedAt to be an explicit nil
func (o *SubscriptionResponse) SetPausedAtNil() {
	o.PausedAt.Set(nil)
}

// UnsetPausedAt ensures that no value is present for PausedAt, not even an explicit nil
func (o *SubscriptionResponse) UnsetPausedAt() {
	o.PausedAt.Unset()
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SubscriptionResponse) SetPlanId(v string) {
	o.PlanId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SubscriptionResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubscriptionStart returns the SubscriptionStart field value if set, zero value otherwise.
func (o *SubscriptionResponse) GetSubscriptionStart() int32 {
	if o == nil || IsNil(o.SubscriptionStart) {
		var ret int32
		return ret
	}
	return *o.SubscriptionStart
}

// GetSubscriptionStartOk returns a tuple with the SubscriptionStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionResponse) GetSubscriptionStartOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriptionStart) {
		return nil, false
	}
	return o.SubscriptionStart, true
}

// HasSubscriptionStart returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasSubscriptionStart() bool {
	if o != nil && !IsNil(o.SubscriptionStart) {
		return true
	}

	return false
}

// SetSubscriptionStart gets a reference to the given int32 and assigns it to the SubscriptionStart field.
func (o *SubscriptionResponse) SetSubscriptionStart(v int32) {
	o.SubscriptionStart = &v
}

// GetTrialStart returns the TrialStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionResponse) GetTrialStart() int64 {
	if o == nil || IsNil(o.TrialStart.Get()) {
		var ret int64
		return ret
	}
	return *o.TrialStart.Get()
}

// GetTrialStartOk returns a tuple with the TrialStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionResponse) GetTrialStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialStart.Get(), o.TrialStart.IsSet()
}

// HasTrialStart returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasTrialStart() bool {
	if o != nil && o.TrialStart.IsSet() {
		return true
	}

	return false
}

// SetTrialStart gets a reference to the given NullableInt64 and assigns it to the TrialStart field.
func (o *SubscriptionResponse) SetTrialStart(v int64) {
	o.TrialStart.Set(&v)
}

// SetTrialStartNil sets the value for TrialStart to be an explicit nil
func (o *SubscriptionResponse) SetTrialStartNil() {
	o.TrialStart.Set(nil)
}

// UnsetTrialStart ensures that no value is present for TrialStart, not even an explicit nil
func (o *SubscriptionResponse) UnsetTrialStart() {
	o.TrialStart.Unset()
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionResponse) GetTrialEnd() int64 {
	if o == nil || IsNil(o.TrialEnd.Get()) {
		var ret int64
		return ret
	}
	return *o.TrialEnd.Get()
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionResponse) GetTrialEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialEnd.Get(), o.TrialEnd.IsSet()
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *SubscriptionResponse) HasTrialEnd() bool {
	if o != nil && o.TrialEnd.IsSet() {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given NullableInt64 and assigns it to the TrialEnd field.
func (o *SubscriptionResponse) SetTrialEnd(v int64) {
	o.TrialEnd.Set(&v)
}

// SetTrialEndNil sets the value for TrialEnd to be an explicit nil
func (o *SubscriptionResponse) SetTrialEndNil() {
	o.TrialEnd.Set(nil)
}

// UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
func (o *SubscriptionResponse) UnsetTrialEnd() {
	o.TrialEnd.Unset()
}

func (o SubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingCycleStart.IsSet() {
		toSerialize["billing_cycle_start"] = o.BillingCycleStart.Get()
	}
	if o.BillingCycleEnd.IsSet() {
		toSerialize["billing_cycle_end"] = o.BillingCycleEnd.Get()
	}
	if o.CanceledAt.IsSet() {
		toSerialize["canceled_at"] = o.CanceledAt.Get()
	}
	if !IsNil(o.CardId) {
		toSerialize["card_id"] = o.CardId
	}
	if o.ChargeId.IsSet() {
		toSerialize["charge_id"] = o.ChargeId.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CustomerCustomReference) {
		toSerialize["customer_custom_reference"] = o.CustomerCustomReference
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastBillingCycleOrderId) {
		toSerialize["last_billing_cycle_order_id"] = o.LastBillingCycleOrderId
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if o.PausedAt.IsSet() {
		toSerialize["paused_at"] = o.PausedAt.Get()
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscriptionStart) {
		toSerialize["subscription_start"] = o.SubscriptionStart
	}
	if o.TrialStart.IsSet() {
		toSerialize["trial_start"] = o.TrialStart.Get()
	}
	if o.TrialEnd.IsSet() {
		toSerialize["trial_end"] = o.TrialEnd.Get()
	}
	return toSerialize, nil
}

type NullableSubscriptionResponse struct {
	value *SubscriptionResponse
	isSet bool
}

func (v NullableSubscriptionResponse) Get() *SubscriptionResponse {
	return v.value
}

func (v *NullableSubscriptionResponse) Set(val *SubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionResponse(val *SubscriptionResponse) *NullableSubscriptionResponse {
	return &NullableSubscriptionResponse{value: val, isSet: true}
}

func (v NullableSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

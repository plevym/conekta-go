package conekta

import (
	"encoding/json"
)

// checks if the PlanResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanResponse{}

// PlanResponse plans model
type PlanResponse struct {
	Amount          *int32        `json:"amount,omitempty"`
	CreatedAt       *int64        `json:"created_at,omitempty"`
	Currency        *string       `json:"currency,omitempty"`
	ExpiryCount     NullableInt32 `json:"expiry_count,omitempty"`
	Frequency       *int32        `json:"frequency,omitempty"`
	Id              *string       `json:"id,omitempty"`
	Interval        *string       `json:"interval,omitempty"`
	Livemode        *bool         `json:"livemode,omitempty"`
	Name            *string       `json:"name,omitempty"`
	Object          *string       `json:"object,omitempty"`
	TrialPeriodDays NullableInt32 `json:"trial_period_days,omitempty"`
}

// NewPlanResponse instantiates a new PlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanResponse() *PlanResponse {
	this := PlanResponse{}
	return &this
}

// NewPlanResponseWithDefaults instantiates a new PlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanResponseWithDefaults() *PlanResponse {
	this := PlanResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PlanResponse) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanResponse) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PlanResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *PlanResponse) SetAmount(v int32) {
	o.Amount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PlanResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PlanResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *PlanResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PlanResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PlanResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PlanResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetExpiryCount returns the ExpiryCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlanResponse) GetExpiryCount() int32 {
	if o == nil || IsNil(o.ExpiryCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ExpiryCount.Get()
}

// GetExpiryCountOk returns a tuple with the ExpiryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanResponse) GetExpiryCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryCount.Get(), o.ExpiryCount.IsSet()
}

// HasExpiryCount returns a boolean if a field has been set.
func (o *PlanResponse) HasExpiryCount() bool {
	if o != nil && o.ExpiryCount.IsSet() {
		return true
	}

	return false
}

// SetExpiryCount gets a reference to the given NullableInt32 and assigns it to the ExpiryCount field.
func (o *PlanResponse) SetExpiryCount(v int32) {
	o.ExpiryCount.Set(&v)
}

// SetExpiryCountNil sets the value for ExpiryCount to be an explicit nil
func (o *PlanResponse) SetExpiryCountNil() {
	o.ExpiryCount.Set(nil)
}

// UnsetExpiryCount ensures that no value is present for ExpiryCount, not even an explicit nil
func (o *PlanResponse) UnsetExpiryCount() {
	o.ExpiryCount.Unset()
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *PlanResponse) GetFrequency() int32 {
	if o == nil || IsNil(o.Frequency) {
		var ret int32
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanResponse) GetFrequencyOk() (*int32, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *PlanResponse) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given int32 and assigns it to the Frequency field.
func (o *PlanResponse) SetFrequency(v int32) {
	o.Frequency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlanResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlanResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlanResponse) SetId(v string) {
	o.Id = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *PlanResponse) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanResponse) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *PlanResponse) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *PlanResponse) SetInterval(v string) {
	o.Interval = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *PlanResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *PlanResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *PlanResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlanResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlanResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlanResponse) SetName(v string) {
	o.Name = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *PlanResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *PlanResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *PlanResponse) SetObject(v string) {
	o.Object = &v
}

// GetTrialPeriodDays returns the TrialPeriodDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlanResponse) GetTrialPeriodDays() int32 {
	if o == nil || IsNil(o.TrialPeriodDays.Get()) {
		var ret int32
		return ret
	}
	return *o.TrialPeriodDays.Get()
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanResponse) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialPeriodDays.Get(), o.TrialPeriodDays.IsSet()
}

// HasTrialPeriodDays returns a boolean if a field has been set.
func (o *PlanResponse) HasTrialPeriodDays() bool {
	if o != nil && o.TrialPeriodDays.IsSet() {
		return true
	}

	return false
}

// SetTrialPeriodDays gets a reference to the given NullableInt32 and assigns it to the TrialPeriodDays field.
func (o *PlanResponse) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays.Set(&v)
}

// SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil
func (o *PlanResponse) SetTrialPeriodDaysNil() {
	o.TrialPeriodDays.Set(nil)
}

// UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
func (o *PlanResponse) UnsetTrialPeriodDays() {
	o.TrialPeriodDays.Unset()
}

func (o PlanResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if o.ExpiryCount.IsSet() {
		toSerialize["expiry_count"] = o.ExpiryCount.Get()
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if o.TrialPeriodDays.IsSet() {
		toSerialize["trial_period_days"] = o.TrialPeriodDays.Get()
	}
	return toSerialize, nil
}

type NullablePlanResponse struct {
	value *PlanResponse
	isSet bool
}

func (v NullablePlanResponse) Get() *PlanResponse {
	return v.value
}

func (v *NullablePlanResponse) Set(val *PlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanResponse(val *PlanResponse) *NullablePlanResponse {
	return &NullablePlanResponse{value: val, isSet: true}
}

func (v NullablePlanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

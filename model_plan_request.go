package conekta

import (
	"encoding/json"
)

// checks if the PlanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanRequest{}

// PlanRequest a plan
type PlanRequest struct {
	// The amount in cents that will be charged on the interval specified.
	Amount int32 `json:"amount"`
	// ISO 4217 for currencies, for the Mexican peso it is MXN/USD
	Currency *string `json:"currency,omitempty"`
	// Number of repetitions of the frequency NUMBER OF CHARGES TO BE MADE, considering the interval and frequency, this evolves over time, but is subject to the expiration count.
	ExpiryCount *int32 `json:"expiry_count,omitempty"`
	// Frequency of the charge, which together with the interval, can be every 3 weeks, every 4 months, every 2 years, every 5 fortnights
	Frequency int32 `json:"frequency"`
	// internal reference id
	Id *string `json:"id,omitempty"`
	// The interval of time between each charge.
	Interval string `json:"interval"`
	// The name of the plan.
	Name string `json:"name"`
	// The number of days the customer will have a free trial.
	TrialPeriodDays *int32 `json:"trial_period_days,omitempty"`
}

// NewPlanRequest instantiates a new PlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanRequest(amount int32, frequency int32, interval string, name string) *PlanRequest {
	this := PlanRequest{}
	this.Amount = amount
	this.Frequency = frequency
	this.Interval = interval
	this.Name = name
	return &this
}

// NewPlanRequestWithDefaults instantiates a new PlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanRequestWithDefaults() *PlanRequest {
	this := PlanRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PlanRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PlanRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PlanRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PlanRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PlanRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetExpiryCount returns the ExpiryCount field value if set, zero value otherwise.
func (o *PlanRequest) GetExpiryCount() int32 {
	if o == nil || IsNil(o.ExpiryCount) {
		var ret int32
		return ret
	}
	return *o.ExpiryCount
}

// GetExpiryCountOk returns a tuple with the ExpiryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetExpiryCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiryCount) {
		return nil, false
	}
	return o.ExpiryCount, true
}

// HasExpiryCount returns a boolean if a field has been set.
func (o *PlanRequest) HasExpiryCount() bool {
	if o != nil && !IsNil(o.ExpiryCount) {
		return true
	}

	return false
}

// SetExpiryCount gets a reference to the given int32 and assigns it to the ExpiryCount field.
func (o *PlanRequest) SetExpiryCount(v int32) {
	o.ExpiryCount = &v
}

// GetFrequency returns the Frequency field value
func (o *PlanRequest) GetFrequency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetFrequencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *PlanRequest) SetFrequency(v int32) {
	o.Frequency = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PlanRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PlanRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PlanRequest) SetId(v string) {
	o.Id = &v
}

// GetInterval returns the Interval field value
func (o *PlanRequest) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *PlanRequest) SetInterval(v string) {
	o.Interval = v
}

// GetName returns the Name field value
func (o *PlanRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlanRequest) SetName(v string) {
	o.Name = v
}

// GetTrialPeriodDays returns the TrialPeriodDays field value if set, zero value otherwise.
func (o *PlanRequest) GetTrialPeriodDays() int32 {
	if o == nil || IsNil(o.TrialPeriodDays) {
		var ret int32
		return ret
	}
	return *o.TrialPeriodDays
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialPeriodDays) {
		return nil, false
	}
	return o.TrialPeriodDays, true
}

// HasTrialPeriodDays returns a boolean if a field has been set.
func (o *PlanRequest) HasTrialPeriodDays() bool {
	if o != nil && !IsNil(o.TrialPeriodDays) {
		return true
	}

	return false
}

// SetTrialPeriodDays gets a reference to the given int32 and assigns it to the TrialPeriodDays field.
func (o *PlanRequest) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays = &v
}

func (o PlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExpiryCount) {
		toSerialize["expiry_count"] = o.ExpiryCount
	}
	toSerialize["frequency"] = o.Frequency
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["interval"] = o.Interval
	toSerialize["name"] = o.Name
	if !IsNil(o.TrialPeriodDays) {
		toSerialize["trial_period_days"] = o.TrialPeriodDays
	}
	return toSerialize, nil
}

type NullablePlanRequest struct {
	value *PlanRequest
	isSet bool
}

func (v NullablePlanRequest) Get() *PlanRequest {
	return v.value
}

func (v *NullablePlanRequest) Set(val *PlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanRequest(val *PlanRequest) *NullablePlanRequest {
	return &NullablePlanRequest{value: val, isSet: true}
}

func (v NullablePlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

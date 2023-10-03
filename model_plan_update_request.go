package conekta

import (
	"encoding/json"
)

// checks if the PlanUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanUpdateRequest{}

// PlanUpdateRequest a plan
type PlanUpdateRequest struct {
	// The amount in cents that will be charged on the interval specified.
	Amount *int32 `json:"amount,omitempty"`
	// ISO 4217 for currencies, for the Mexican peso it is MXN/USD
	Currency *string `json:"currency,omitempty"`
	// Number of repetitions of the frequency NUMBER OF CHARGES TO BE MADE, considering the interval and frequency, this evolves over time, but is subject to the expiration count.
	ExpiryCount *int32 `json:"expiry_count,omitempty"`
	// The name of the plan.
	Name *string `json:"name,omitempty"`
}

// NewPlanUpdateRequest instantiates a new PlanUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanUpdateRequest() *PlanUpdateRequest {
	this := PlanUpdateRequest{}
	return &this
}

// NewPlanUpdateRequestWithDefaults instantiates a new PlanUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanUpdateRequestWithDefaults() *PlanUpdateRequest {
	this := PlanUpdateRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PlanUpdateRequest) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanUpdateRequest) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PlanUpdateRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *PlanUpdateRequest) SetAmount(v int32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PlanUpdateRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanUpdateRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PlanUpdateRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PlanUpdateRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetExpiryCount returns the ExpiryCount field value if set, zero value otherwise.
func (o *PlanUpdateRequest) GetExpiryCount() int32 {
	if o == nil || IsNil(o.ExpiryCount) {
		var ret int32
		return ret
	}
	return *o.ExpiryCount
}

// GetExpiryCountOk returns a tuple with the ExpiryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanUpdateRequest) GetExpiryCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiryCount) {
		return nil, false
	}
	return o.ExpiryCount, true
}

// HasExpiryCount returns a boolean if a field has been set.
func (o *PlanUpdateRequest) HasExpiryCount() bool {
	if o != nil && !IsNil(o.ExpiryCount) {
		return true
	}

	return false
}

// SetExpiryCount gets a reference to the given int32 and assigns it to the ExpiryCount field.
func (o *PlanUpdateRequest) SetExpiryCount(v int32) {
	o.ExpiryCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlanUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlanUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlanUpdateRequest) SetName(v string) {
	o.Name = &v
}

func (o PlanUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExpiryCount) {
		toSerialize["expiry_count"] = o.ExpiryCount
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePlanUpdateRequest struct {
	value *PlanUpdateRequest
	isSet bool
}

func (v NullablePlanUpdateRequest) Get() *PlanUpdateRequest {
	return v.value
}

func (v *NullablePlanUpdateRequest) Set(val *PlanUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanUpdateRequest(val *PlanUpdateRequest) *NullablePlanUpdateRequest {
	return &NullablePlanUpdateRequest{value: val, isSet: true}
}

func (v NullablePlanUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

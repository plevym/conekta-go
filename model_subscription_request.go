package conekta

import (
	"encoding/json"
)

// checks if the SubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionRequest{}

// SubscriptionRequest It is a parameter that allows to identify in the response, the detailed content of the plans to which the client has subscribed
type SubscriptionRequest struct {
	PlanId   string  `json:"plan_id"`
	CardId   *string `json:"card_id,omitempty"`
	TrialEnd *int32  `json:"trial_end,omitempty"`
}

// NewSubscriptionRequest instantiates a new SubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionRequest(planId string) *SubscriptionRequest {
	this := SubscriptionRequest{}
	this.PlanId = planId
	return &this
}

// NewSubscriptionRequestWithDefaults instantiates a new SubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionRequestWithDefaults() *SubscriptionRequest {
	this := SubscriptionRequest{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *SubscriptionRequest) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionRequest) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *SubscriptionRequest) SetPlanId(v string) {
	o.PlanId = v
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *SubscriptionRequest) GetCardId() string {
	if o == nil || IsNil(o.CardId) {
		var ret string
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRequest) GetCardIdOk() (*string, bool) {
	if o == nil || IsNil(o.CardId) {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *SubscriptionRequest) HasCardId() bool {
	if o != nil && !IsNil(o.CardId) {
		return true
	}

	return false
}

// SetCardId gets a reference to the given string and assigns it to the CardId field.
func (o *SubscriptionRequest) SetCardId(v string) {
	o.CardId = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *SubscriptionRequest) GetTrialEnd() int32 {
	if o == nil || IsNil(o.TrialEnd) {
		var ret int32
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionRequest) GetTrialEndOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialEnd) {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *SubscriptionRequest) HasTrialEnd() bool {
	if o != nil && !IsNil(o.TrialEnd) {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given int32 and assigns it to the TrialEnd field.
func (o *SubscriptionRequest) SetTrialEnd(v int32) {
	o.TrialEnd = &v
}

func (o SubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plan_id"] = o.PlanId
	if !IsNil(o.CardId) {
		toSerialize["card_id"] = o.CardId
	}
	if !IsNil(o.TrialEnd) {
		toSerialize["trial_end"] = o.TrialEnd
	}
	return toSerialize, nil
}

type NullableSubscriptionRequest struct {
	value *SubscriptionRequest
	isSet bool
}

func (v NullableSubscriptionRequest) Get() *SubscriptionRequest {
	return v.value
}

func (v *NullableSubscriptionRequest) Set(val *SubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionRequest(val *SubscriptionRequest) *NullableSubscriptionRequest {
	return &NullableSubscriptionRequest{value: val, isSet: true}
}

func (v NullableSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the SubscriptionUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionUpdateRequest{}

// SubscriptionUpdateRequest You can modify the subscription to change the plan used by your customers.
type SubscriptionUpdateRequest struct {
	PlanId   *string `json:"plan_id,omitempty"`
	CardId   *string `json:"card_id,omitempty"`
	TrialEnd *int32  `json:"trial_end,omitempty"`
}

// NewSubscriptionUpdateRequest instantiates a new SubscriptionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionUpdateRequest() *SubscriptionUpdateRequest {
	this := SubscriptionUpdateRequest{}
	return &this
}

// NewSubscriptionUpdateRequestWithDefaults instantiates a new SubscriptionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionUpdateRequestWithDefaults() *SubscriptionUpdateRequest {
	this := SubscriptionUpdateRequest{}
	return &this
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SubscriptionUpdateRequest) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequest) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SubscriptionUpdateRequest) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SubscriptionUpdateRequest) SetPlanId(v string) {
	o.PlanId = &v
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *SubscriptionUpdateRequest) GetCardId() string {
	if o == nil || IsNil(o.CardId) {
		var ret string
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequest) GetCardIdOk() (*string, bool) {
	if o == nil || IsNil(o.CardId) {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *SubscriptionUpdateRequest) HasCardId() bool {
	if o != nil && !IsNil(o.CardId) {
		return true
	}

	return false
}

// SetCardId gets a reference to the given string and assigns it to the CardId field.
func (o *SubscriptionUpdateRequest) SetCardId(v string) {
	o.CardId = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *SubscriptionUpdateRequest) GetTrialEnd() int32 {
	if o == nil || IsNil(o.TrialEnd) {
		var ret int32
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionUpdateRequest) GetTrialEndOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialEnd) {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *SubscriptionUpdateRequest) HasTrialEnd() bool {
	if o != nil && !IsNil(o.TrialEnd) {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given int32 and assigns it to the TrialEnd field.
func (o *SubscriptionUpdateRequest) SetTrialEnd(v int32) {
	o.TrialEnd = &v
}

func (o SubscriptionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.CardId) {
		toSerialize["card_id"] = o.CardId
	}
	if !IsNil(o.TrialEnd) {
		toSerialize["trial_end"] = o.TrialEnd
	}
	return toSerialize, nil
}

type NullableSubscriptionUpdateRequest struct {
	value *SubscriptionUpdateRequest
	isSet bool
}

func (v NullableSubscriptionUpdateRequest) Get() *SubscriptionUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionUpdateRequest) Set(val *SubscriptionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionUpdateRequest(val *SubscriptionUpdateRequest) *NullableSubscriptionUpdateRequest {
	return &NullableSubscriptionUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

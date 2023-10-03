package conekta

import (
	"encoding/json"
)

// checks if the WebhookKeyUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookKeyUpdateRequest{}

// WebhookKeyUpdateRequest struct for WebhookKeyUpdateRequest
type WebhookKeyUpdateRequest struct {
	// Indicates if the webhook key is active
	Active *bool `json:"active,omitempty"`
}

// NewWebhookKeyUpdateRequest instantiates a new WebhookKeyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookKeyUpdateRequest() *WebhookKeyUpdateRequest {
	this := WebhookKeyUpdateRequest{}
	var active bool = false
	this.Active = &active
	return &this
}

// NewWebhookKeyUpdateRequestWithDefaults instantiates a new WebhookKeyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookKeyUpdateRequestWithDefaults() *WebhookKeyUpdateRequest {
	this := WebhookKeyUpdateRequest{}
	var active bool = false
	this.Active = &active
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WebhookKeyUpdateRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyUpdateRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WebhookKeyUpdateRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WebhookKeyUpdateRequest) SetActive(v bool) {
	o.Active = &v
}

func (o WebhookKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookKeyUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableWebhookKeyUpdateRequest struct {
	value *WebhookKeyUpdateRequest
	isSet bool
}

func (v NullableWebhookKeyUpdateRequest) Get() *WebhookKeyUpdateRequest {
	return v.value
}

func (v *NullableWebhookKeyUpdateRequest) Set(val *WebhookKeyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookKeyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookKeyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookKeyUpdateRequest(val *WebhookKeyUpdateRequest) *NullableWebhookKeyUpdateRequest {
	return &NullableWebhookKeyUpdateRequest{value: val, isSet: true}
}

func (v NullableWebhookKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookKeyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the WebhookUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookUpdateRequest{}

// WebhookUpdateRequest an updated webhook
type WebhookUpdateRequest struct {
	// Here you must place the URL of your Webhook remember that you must program what you will do with the events received. Also do not forget to handle the HTTPS protocol for greater security.
	Url string `json:"url"`
	// It is a value that allows to decide if the events will be synchronous or asynchronous. We recommend asynchronous = false
	Synchronous      *bool    `json:"synchronous,omitempty"`
	SubscribedEvents []string `json:"subscribed_events,omitempty"`
}

// NewWebhookUpdateRequest instantiates a new WebhookUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookUpdateRequest(url string) *WebhookUpdateRequest {
	this := WebhookUpdateRequest{}
	this.Url = url
	var synchronous bool = false
	this.Synchronous = &synchronous
	return &this
}

// NewWebhookUpdateRequestWithDefaults instantiates a new WebhookUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookUpdateRequestWithDefaults() *WebhookUpdateRequest {
	this := WebhookUpdateRequest{}
	var synchronous bool = false
	this.Synchronous = &synchronous
	return &this
}

// GetUrl returns the Url field value
func (o *WebhookUpdateRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdateRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookUpdateRequest) SetUrl(v string) {
	o.Url = v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *WebhookUpdateRequest) GetSynchronous() bool {
	if o == nil || IsNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookUpdateRequest) GetSynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Synchronous) {
		return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *WebhookUpdateRequest) HasSynchronous() bool {
	if o != nil && !IsNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *WebhookUpdateRequest) SetSynchronous(v bool) {
	o.Synchronous = &v
}

// GetSubscribedEvents returns the SubscribedEvents field value if set, zero value otherwise.
func (o *WebhookUpdateRequest) GetSubscribedEvents() []string {
	if o == nil || IsNil(o.SubscribedEvents) {
		var ret []string
		return ret
	}
	return o.SubscribedEvents
}

// GetSubscribedEventsOk returns a tuple with the SubscribedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookUpdateRequest) GetSubscribedEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubscribedEvents) {
		return nil, false
	}
	return o.SubscribedEvents, true
}

// HasSubscribedEvents returns a boolean if a field has been set.
func (o *WebhookUpdateRequest) HasSubscribedEvents() bool {
	if o != nil && !IsNil(o.SubscribedEvents) {
		return true
	}

	return false
}

// SetSubscribedEvents gets a reference to the given []string and assigns it to the SubscribedEvents field.
func (o *WebhookUpdateRequest) SetSubscribedEvents(v []string) {
	o.SubscribedEvents = v
}

func (o WebhookUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	if !IsNil(o.SubscribedEvents) {
		toSerialize["subscribed_events"] = o.SubscribedEvents
	}
	return toSerialize, nil
}

type NullableWebhookUpdateRequest struct {
	value *WebhookUpdateRequest
	isSet bool
}

func (v NullableWebhookUpdateRequest) Get() *WebhookUpdateRequest {
	return v.value
}

func (v *NullableWebhookUpdateRequest) Set(val *WebhookUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookUpdateRequest(val *WebhookUpdateRequest) *NullableWebhookUpdateRequest {
	return &NullableWebhookUpdateRequest{value: val, isSet: true}
}

func (v NullableWebhookUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

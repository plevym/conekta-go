package conekta

import (
	"encoding/json"
)

// checks if the EventsWebhookResendResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsWebhookResendResponse{}

// EventsWebhookResendResponse event model
type EventsWebhookResendResponse struct {
	FailedAttempts         *int32                 `json:"failed_attempts,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	LastAttemptedAt        *int32                 `json:"last_attempted_at,omitempty"`
	LastHttpResponseStatus *int32                 `json:"last_http_response_status,omitempty"`
	ResponseData           map[string]interface{} `json:"response_data,omitempty"`
	Url                    *string                `json:"url,omitempty"`
}

// NewEventsWebhookResendResponse instantiates a new EventsWebhookResendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsWebhookResendResponse() *EventsWebhookResendResponse {
	this := EventsWebhookResendResponse{}
	return &this
}

// NewEventsWebhookResendResponseWithDefaults instantiates a new EventsWebhookResendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsWebhookResendResponseWithDefaults() *EventsWebhookResendResponse {
	this := EventsWebhookResendResponse{}
	return &this
}

// GetFailedAttempts returns the FailedAttempts field value if set, zero value otherwise.
func (o *EventsWebhookResendResponse) GetFailedAttempts() int32 {
	if o == nil || IsNil(o.FailedAttempts) {
		var ret int32
		return ret
	}
	return *o.FailedAttempts
}

// GetFailedAttemptsOk returns a tuple with the FailedAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsWebhookResendResponse) GetFailedAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.FailedAttempts) {
		return nil, false
	}
	return o.FailedAttempts, true
}

// HasFailedAttempts returns a boolean if a field has been set.
func (o *EventsWebhookResendResponse) HasFailedAttempts() bool {
	if o != nil && !IsNil(o.FailedAttempts) {
		return true
	}

	return false
}

// SetFailedAttempts gets a reference to the given int32 and assigns it to the FailedAttempts field.
func (o *EventsWebhookResendResponse) SetFailedAttempts(v int32) {
	o.FailedAttempts = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventsWebhookResendResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsWebhookResendResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventsWebhookResendResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventsWebhookResendResponse) SetId(v string) {
	o.Id = &v
}

// GetLastAttemptedAt returns the LastAttemptedAt field value if set, zero value otherwise.
func (o *EventsWebhookResendResponse) GetLastAttemptedAt() int32 {
	if o == nil || IsNil(o.LastAttemptedAt) {
		var ret int32
		return ret
	}
	return *o.LastAttemptedAt
}

// GetLastAttemptedAtOk returns a tuple with the LastAttemptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsWebhookResendResponse) GetLastAttemptedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.LastAttemptedAt) {
		return nil, false
	}
	return o.LastAttemptedAt, true
}

// HasLastAttemptedAt returns a boolean if a field has been set.
func (o *EventsWebhookResendResponse) HasLastAttemptedAt() bool {
	if o != nil && !IsNil(o.LastAttemptedAt) {
		return true
	}

	return false
}

// SetLastAttemptedAt gets a reference to the given int32 and assigns it to the LastAttemptedAt field.
func (o *EventsWebhookResendResponse) SetLastAttemptedAt(v int32) {
	o.LastAttemptedAt = &v
}

// GetLastHttpResponseStatus returns the LastHttpResponseStatus field value if set, zero value otherwise.
func (o *EventsWebhookResendResponse) GetLastHttpResponseStatus() int32 {
	if o == nil || IsNil(o.LastHttpResponseStatus) {
		var ret int32
		return ret
	}
	return *o.LastHttpResponseStatus
}

// GetLastHttpResponseStatusOk returns a tuple with the LastHttpResponseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsWebhookResendResponse) GetLastHttpResponseStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.LastHttpResponseStatus) {
		return nil, false
	}
	return o.LastHttpResponseStatus, true
}

// HasLastHttpResponseStatus returns a boolean if a field has been set.
func (o *EventsWebhookResendResponse) HasLastHttpResponseStatus() bool {
	if o != nil && !IsNil(o.LastHttpResponseStatus) {
		return true
	}

	return false
}

// SetLastHttpResponseStatus gets a reference to the given int32 and assigns it to the LastHttpResponseStatus field.
func (o *EventsWebhookResendResponse) SetLastHttpResponseStatus(v int32) {
	o.LastHttpResponseStatus = &v
}

// GetResponseData returns the ResponseData field value if set, zero value otherwise.
func (o *EventsWebhookResendResponse) GetResponseData() map[string]interface{} {
	if o == nil || IsNil(o.ResponseData) {
		var ret map[string]interface{}
		return ret
	}
	return o.ResponseData
}

// GetResponseDataOk returns a tuple with the ResponseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsWebhookResendResponse) GetResponseDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResponseData) {
		return map[string]interface{}{}, false
	}
	return o.ResponseData, true
}

// HasResponseData returns a boolean if a field has been set.
func (o *EventsWebhookResendResponse) HasResponseData() bool {
	if o != nil && !IsNil(o.ResponseData) {
		return true
	}

	return false
}

// SetResponseData gets a reference to the given map[string]interface{} and assigns it to the ResponseData field.
func (o *EventsWebhookResendResponse) SetResponseData(v map[string]interface{}) {
	o.ResponseData = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EventsWebhookResendResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsWebhookResendResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EventsWebhookResendResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EventsWebhookResendResponse) SetUrl(v string) {
	o.Url = &v
}

func (o EventsWebhookResendResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsWebhookResendResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailedAttempts) {
		toSerialize["failed_attempts"] = o.FailedAttempts
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastAttemptedAt) {
		toSerialize["last_attempted_at"] = o.LastAttemptedAt
	}
	if !IsNil(o.LastHttpResponseStatus) {
		toSerialize["last_http_response_status"] = o.LastHttpResponseStatus
	}
	if !IsNil(o.ResponseData) {
		toSerialize["response_data"] = o.ResponseData
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableEventsWebhookResendResponse struct {
	value *EventsWebhookResendResponse
	isSet bool
}

func (v NullableEventsWebhookResendResponse) Get() *EventsWebhookResendResponse {
	return v.value
}

func (v *NullableEventsWebhookResendResponse) Set(val *EventsWebhookResendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsWebhookResendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsWebhookResendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsWebhookResendResponse(val *EventsWebhookResendResponse) *NullableEventsWebhookResendResponse {
	return &NullableEventsWebhookResendResponse{value: val, isSet: true}
}

func (v NullableEventsWebhookResendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsWebhookResendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

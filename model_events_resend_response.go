package conekta

import (
	"encoding/json"
)

// checks if the EventsResendResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsResendResponse{}

// EventsResendResponse event model
type EventsResendResponse struct {
	FailedAttempts         *int32                 `json:"failed_attempts,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	LastAttemptedAt        *int32                 `json:"last_attempted_at,omitempty"`
	LastHttpResponseStatus *int32                 `json:"last_http_response_status,omitempty"`
	ResponseData           map[string]interface{} `json:"response_data,omitempty"`
	Url                    *string                `json:"url,omitempty"`
}

// NewEventsResendResponse instantiates a new EventsResendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsResendResponse() *EventsResendResponse {
	this := EventsResendResponse{}
	return &this
}

// NewEventsResendResponseWithDefaults instantiates a new EventsResendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsResendResponseWithDefaults() *EventsResendResponse {
	this := EventsResendResponse{}
	return &this
}

// GetFailedAttempts returns the FailedAttempts field value if set, zero value otherwise.
func (o *EventsResendResponse) GetFailedAttempts() int32 {
	if o == nil || IsNil(o.FailedAttempts) {
		var ret int32
		return ret
	}
	return *o.FailedAttempts
}

// GetFailedAttemptsOk returns a tuple with the FailedAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResendResponse) GetFailedAttemptsOk() (*int32, bool) {
	if o == nil || IsNil(o.FailedAttempts) {
		return nil, false
	}
	return o.FailedAttempts, true
}

// HasFailedAttempts returns a boolean if a field has been set.
func (o *EventsResendResponse) HasFailedAttempts() bool {
	if o != nil && !IsNil(o.FailedAttempts) {
		return true
	}

	return false
}

// SetFailedAttempts gets a reference to the given int32 and assigns it to the FailedAttempts field.
func (o *EventsResendResponse) SetFailedAttempts(v int32) {
	o.FailedAttempts = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventsResendResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResendResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventsResendResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventsResendResponse) SetId(v string) {
	o.Id = &v
}

// GetLastAttemptedAt returns the LastAttemptedAt field value if set, zero value otherwise.
func (o *EventsResendResponse) GetLastAttemptedAt() int32 {
	if o == nil || IsNil(o.LastAttemptedAt) {
		var ret int32
		return ret
	}
	return *o.LastAttemptedAt
}

// GetLastAttemptedAtOk returns a tuple with the LastAttemptedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResendResponse) GetLastAttemptedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.LastAttemptedAt) {
		return nil, false
	}
	return o.LastAttemptedAt, true
}

// HasLastAttemptedAt returns a boolean if a field has been set.
func (o *EventsResendResponse) HasLastAttemptedAt() bool {
	if o != nil && !IsNil(o.LastAttemptedAt) {
		return true
	}

	return false
}

// SetLastAttemptedAt gets a reference to the given int32 and assigns it to the LastAttemptedAt field.
func (o *EventsResendResponse) SetLastAttemptedAt(v int32) {
	o.LastAttemptedAt = &v
}

// GetLastHttpResponseStatus returns the LastHttpResponseStatus field value if set, zero value otherwise.
func (o *EventsResendResponse) GetLastHttpResponseStatus() int32 {
	if o == nil || IsNil(o.LastHttpResponseStatus) {
		var ret int32
		return ret
	}
	return *o.LastHttpResponseStatus
}

// GetLastHttpResponseStatusOk returns a tuple with the LastHttpResponseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResendResponse) GetLastHttpResponseStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.LastHttpResponseStatus) {
		return nil, false
	}
	return o.LastHttpResponseStatus, true
}

// HasLastHttpResponseStatus returns a boolean if a field has been set.
func (o *EventsResendResponse) HasLastHttpResponseStatus() bool {
	if o != nil && !IsNil(o.LastHttpResponseStatus) {
		return true
	}

	return false
}

// SetLastHttpResponseStatus gets a reference to the given int32 and assigns it to the LastHttpResponseStatus field.
func (o *EventsResendResponse) SetLastHttpResponseStatus(v int32) {
	o.LastHttpResponseStatus = &v
}

// GetResponseData returns the ResponseData field value if set, zero value otherwise.
func (o *EventsResendResponse) GetResponseData() map[string]interface{} {
	if o == nil || IsNil(o.ResponseData) {
		var ret map[string]interface{}
		return ret
	}
	return o.ResponseData
}

// GetResponseDataOk returns a tuple with the ResponseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResendResponse) GetResponseDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResponseData) {
		return map[string]interface{}{}, false
	}
	return o.ResponseData, true
}

// HasResponseData returns a boolean if a field has been set.
func (o *EventsResendResponse) HasResponseData() bool {
	if o != nil && !IsNil(o.ResponseData) {
		return true
	}

	return false
}

// SetResponseData gets a reference to the given map[string]interface{} and assigns it to the ResponseData field.
func (o *EventsResendResponse) SetResponseData(v map[string]interface{}) {
	o.ResponseData = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EventsResendResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsResendResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EventsResendResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EventsResendResponse) SetUrl(v string) {
	o.Url = &v
}

func (o EventsResendResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsResendResponse) ToMap() (map[string]interface{}, error) {
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

type NullableEventsResendResponse struct {
	value *EventsResendResponse
	isSet bool
}

func (v NullableEventsResendResponse) Get() *EventsResendResponse {
	return v.value
}

func (v *NullableEventsResendResponse) Set(val *EventsResendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsResendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsResendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsResendResponse(val *EventsResendResponse) *NullableEventsResendResponse {
	return &NullableEventsResendResponse{value: val, isSet: true}
}

func (v NullableEventsResendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsResendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

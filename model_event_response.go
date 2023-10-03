package conekta

import (
	"encoding/json"
)

// checks if the EventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventResponse{}

// EventResponse event model
type EventResponse struct {
	CreatedAt     *int64                 `json:"created_at,omitempty"`
	Data          map[string]interface{} `json:"data,omitempty"`
	Id            *string                `json:"id,omitempty"`
	Livemode      *bool                  `json:"livemode,omitempty"`
	Object        *string                `json:"object,omitempty"`
	Type          *string                `json:"type,omitempty"`
	WebhookLogs   []WebhookLog           `json:"webhook_logs,omitempty"`
	WebhookStatus *string                `json:"webhook_status,omitempty"`
}

// NewEventResponse instantiates a new EventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponse() *EventResponse {
	this := EventResponse{}
	return &this
}

// NewEventResponseWithDefaults instantiates a new EventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseWithDefaults() *EventResponse {
	this := EventResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EventResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EventResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *EventResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventResponse) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *EventResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventResponse) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *EventResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *EventResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *EventResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *EventResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *EventResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *EventResponse) SetObject(v string) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventResponse) SetType(v string) {
	o.Type = &v
}

// GetWebhookLogs returns the WebhookLogs field value if set, zero value otherwise.
func (o *EventResponse) GetWebhookLogs() []WebhookLog {
	if o == nil || IsNil(o.WebhookLogs) {
		var ret []WebhookLog
		return ret
	}
	return o.WebhookLogs
}

// GetWebhookLogsOk returns a tuple with the WebhookLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetWebhookLogsOk() ([]WebhookLog, bool) {
	if o == nil || IsNil(o.WebhookLogs) {
		return nil, false
	}
	return o.WebhookLogs, true
}

// HasWebhookLogs returns a boolean if a field has been set.
func (o *EventResponse) HasWebhookLogs() bool {
	if o != nil && !IsNil(o.WebhookLogs) {
		return true
	}

	return false
}

// SetWebhookLogs gets a reference to the given []WebhookLog and assigns it to the WebhookLogs field.
func (o *EventResponse) SetWebhookLogs(v []WebhookLog) {
	o.WebhookLogs = v
}

// GetWebhookStatus returns the WebhookStatus field value if set, zero value otherwise.
func (o *EventResponse) GetWebhookStatus() string {
	if o == nil || IsNil(o.WebhookStatus) {
		var ret string
		return ret
	}
	return *o.WebhookStatus
}

// GetWebhookStatusOk returns a tuple with the WebhookStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponse) GetWebhookStatusOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookStatus) {
		return nil, false
	}
	return o.WebhookStatus, true
}

// HasWebhookStatus returns a boolean if a field has been set.
func (o *EventResponse) HasWebhookStatus() bool {
	if o != nil && !IsNil(o.WebhookStatus) {
		return true
	}

	return false
}

// SetWebhookStatus gets a reference to the given string and assigns it to the WebhookStatus field.
func (o *EventResponse) SetWebhookStatus(v string) {
	o.WebhookStatus = &v
}

func (o EventResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.WebhookLogs) {
		toSerialize["webhook_logs"] = o.WebhookLogs
	}
	if !IsNil(o.WebhookStatus) {
		toSerialize["webhook_status"] = o.WebhookStatus
	}
	return toSerialize, nil
}

type NullableEventResponse struct {
	value *EventResponse
	isSet bool
}

func (v NullableEventResponse) Get() *EventResponse {
	return v.value
}

func (v *NullableEventResponse) Set(val *EventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponse(val *EventResponse) *NullableEventResponse {
	return &NullableEventResponse{value: val, isSet: true}
}

func (v NullableEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

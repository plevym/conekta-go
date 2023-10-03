package conekta

import (
	"encoding/json"
)

// checks if the WebhookKeyDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookKeyDeleteResponse{}

// WebhookKeyDeleteResponse webhook keys model
type WebhookKeyDeleteResponse struct {
	// Indicates if the webhook key is active
	Active *bool `json:"active,omitempty"`
	// Unix timestamp in seconds with the creation date of the webhook key
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Indicates if the webhook key is deleted
	Deleted *bool `json:"deleted,omitempty"`
	// Unique identifier of the webhook key
	Id *string `json:"id,omitempty"`
	// Indicates if the webhook key is in live mode
	Livemode *bool `json:"livemode,omitempty"`
	// Object name, value is webhook_key
	Object *string `json:"object,omitempty"`
}

// NewWebhookKeyDeleteResponse instantiates a new WebhookKeyDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookKeyDeleteResponse() *WebhookKeyDeleteResponse {
	this := WebhookKeyDeleteResponse{}
	return &this
}

// NewWebhookKeyDeleteResponseWithDefaults instantiates a new WebhookKeyDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookKeyDeleteResponseWithDefaults() *WebhookKeyDeleteResponse {
	this := WebhookKeyDeleteResponse{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WebhookKeyDeleteResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyDeleteResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WebhookKeyDeleteResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WebhookKeyDeleteResponse) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WebhookKeyDeleteResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyDeleteResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WebhookKeyDeleteResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *WebhookKeyDeleteResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *WebhookKeyDeleteResponse) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyDeleteResponse) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *WebhookKeyDeleteResponse) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *WebhookKeyDeleteResponse) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookKeyDeleteResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyDeleteResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookKeyDeleteResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhookKeyDeleteResponse) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *WebhookKeyDeleteResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyDeleteResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *WebhookKeyDeleteResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *WebhookKeyDeleteResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *WebhookKeyDeleteResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyDeleteResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *WebhookKeyDeleteResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *WebhookKeyDeleteResponse) SetObject(v string) {
	o.Object = &v
}

func (o WebhookKeyDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookKeyDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
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
	return toSerialize, nil
}

type NullableWebhookKeyDeleteResponse struct {
	value *WebhookKeyDeleteResponse
	isSet bool
}

func (v NullableWebhookKeyDeleteResponse) Get() *WebhookKeyDeleteResponse {
	return v.value
}

func (v *NullableWebhookKeyDeleteResponse) Set(val *WebhookKeyDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookKeyDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookKeyDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookKeyDeleteResponse(val *WebhookKeyDeleteResponse) *NullableWebhookKeyDeleteResponse {
	return &NullableWebhookKeyDeleteResponse{value: val, isSet: true}
}

func (v NullableWebhookKeyDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookKeyDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

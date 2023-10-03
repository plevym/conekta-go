package conekta

import (
	"encoding/json"
)

// checks if the WebhookKeyCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookKeyCreateResponse{}

// WebhookKeyCreateResponse webhook keys model
type WebhookKeyCreateResponse struct {
	// Indicates if the webhook key is active
	Active *bool `json:"active,omitempty"`
	// Unix timestamp in seconds with the creation date of the webhook key
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Unique identifier of the webhook key
	Id *string `json:"id,omitempty"`
	// Indicates if the webhook key is in live mode
	Livemode *bool `json:"livemode,omitempty"`
	// Object name, value is webhook_key
	Object *string `json:"object,omitempty"`
	// Public key to be used in the webhook
	PublicKey *string `json:"public_key,omitempty"`
}

// NewWebhookKeyCreateResponse instantiates a new WebhookKeyCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookKeyCreateResponse() *WebhookKeyCreateResponse {
	this := WebhookKeyCreateResponse{}
	return &this
}

// NewWebhookKeyCreateResponseWithDefaults instantiates a new WebhookKeyCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookKeyCreateResponseWithDefaults() *WebhookKeyCreateResponse {
	this := WebhookKeyCreateResponse{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WebhookKeyCreateResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyCreateResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WebhookKeyCreateResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WebhookKeyCreateResponse) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WebhookKeyCreateResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyCreateResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WebhookKeyCreateResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *WebhookKeyCreateResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookKeyCreateResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyCreateResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookKeyCreateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhookKeyCreateResponse) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *WebhookKeyCreateResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyCreateResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *WebhookKeyCreateResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *WebhookKeyCreateResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *WebhookKeyCreateResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyCreateResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *WebhookKeyCreateResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *WebhookKeyCreateResponse) SetObject(v string) {
	o.Object = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *WebhookKeyCreateResponse) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyCreateResponse) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *WebhookKeyCreateResponse) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *WebhookKeyCreateResponse) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o WebhookKeyCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookKeyCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
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
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullableWebhookKeyCreateResponse struct {
	value *WebhookKeyCreateResponse
	isSet bool
}

func (v NullableWebhookKeyCreateResponse) Get() *WebhookKeyCreateResponse {
	return v.value
}

func (v *NullableWebhookKeyCreateResponse) Set(val *WebhookKeyCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookKeyCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookKeyCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookKeyCreateResponse(val *WebhookKeyCreateResponse) *NullableWebhookKeyCreateResponse {
	return &NullableWebhookKeyCreateResponse{value: val, isSet: true}
}

func (v NullableWebhookKeyCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookKeyCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

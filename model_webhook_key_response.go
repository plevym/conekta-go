package conekta

import (
	"encoding/json"
)

// checks if the WebhookKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookKeyResponse{}

// WebhookKeyResponse webhook keys model
type WebhookKeyResponse struct {
	// Unique identifier of the webhook key
	Id *string `json:"id,omitempty"`
	// Indicates if the webhook key is active
	Active *bool `json:"active,omitempty"`
	// Unix timestamp in seconds with the creation date of the webhook key
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Unix timestamp in seconds with the deactivation date of the webhook key
	DeactivatedAt NullableInt64 `json:"deactivated_at,omitempty"`
	// Public key to be used in the webhook
	PublicKey *string `json:"public_key,omitempty"`
	// Indicates if the webhook key is in live mode
	Livemode *bool `json:"livemode,omitempty"`
	// Object name, value is webhook_key
	Object *string `json:"object,omitempty"`
}

// NewWebhookKeyResponse instantiates a new WebhookKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookKeyResponse() *WebhookKeyResponse {
	this := WebhookKeyResponse{}
	return &this
}

// NewWebhookKeyResponseWithDefaults instantiates a new WebhookKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookKeyResponseWithDefaults() *WebhookKeyResponse {
	this := WebhookKeyResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookKeyResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookKeyResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhookKeyResponse) SetId(v string) {
	o.Id = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WebhookKeyResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WebhookKeyResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WebhookKeyResponse) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WebhookKeyResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WebhookKeyResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *WebhookKeyResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDeactivatedAt returns the DeactivatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookKeyResponse) GetDeactivatedAt() int64 {
	if o == nil || IsNil(o.DeactivatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.DeactivatedAt.Get()
}

// GetDeactivatedAtOk returns a tuple with the DeactivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookKeyResponse) GetDeactivatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeactivatedAt.Get(), o.DeactivatedAt.IsSet()
}

// HasDeactivatedAt returns a boolean if a field has been set.
func (o *WebhookKeyResponse) HasDeactivatedAt() bool {
	if o != nil && o.DeactivatedAt.IsSet() {
		return true
	}

	return false
}

// SetDeactivatedAt gets a reference to the given NullableInt64 and assigns it to the DeactivatedAt field.
func (o *WebhookKeyResponse) SetDeactivatedAt(v int64) {
	o.DeactivatedAt.Set(&v)
}

// SetDeactivatedAtNil sets the value for DeactivatedAt to be an explicit nil
func (o *WebhookKeyResponse) SetDeactivatedAtNil() {
	o.DeactivatedAt.Set(nil)
}

// UnsetDeactivatedAt ensures that no value is present for DeactivatedAt, not even an explicit nil
func (o *WebhookKeyResponse) UnsetDeactivatedAt() {
	o.DeactivatedAt.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *WebhookKeyResponse) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyResponse) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *WebhookKeyResponse) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *WebhookKeyResponse) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *WebhookKeyResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *WebhookKeyResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *WebhookKeyResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *WebhookKeyResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *WebhookKeyResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *WebhookKeyResponse) SetObject(v string) {
	o.Object = &v
}

func (o WebhookKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DeactivatedAt.IsSet() {
		toSerialize["deactivated_at"] = o.DeactivatedAt.Get()
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableWebhookKeyResponse struct {
	value *WebhookKeyResponse
	isSet bool
}

func (v NullableWebhookKeyResponse) Get() *WebhookKeyResponse {
	return v.value
}

func (v *NullableWebhookKeyResponse) Set(val *WebhookKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookKeyResponse(val *WebhookKeyResponse) *NullableWebhookKeyResponse {
	return &NullableWebhookKeyResponse{value: val, isSet: true}
}

func (v NullableWebhookKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

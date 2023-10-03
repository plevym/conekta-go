package conekta

import (
	"encoding/json"
)

// checks if the WebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResponse{}

// WebhookResponse webhooks model
type WebhookResponse struct {
	Deleted            NullableBool `json:"deleted,omitempty"`
	DevelopmentEnabled *bool        `json:"development_enabled,omitempty"`
	Id                 *string      `json:"id,omitempty"`
	Livemode           *bool        `json:"livemode,omitempty"`
	Object             *string      `json:"object,omitempty"`
	ProductionEnabled  *bool        `json:"production_enabled,omitempty"`
	Status             *string      `json:"status,omitempty"`
	SubscribedEvents   []string     `json:"subscribed_events,omitempty"`
	Synchronous        *bool        `json:"synchronous,omitempty"`
	Url                *string      `json:"url,omitempty"`
}

// NewWebhookResponse instantiates a new WebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponse() *WebhookResponse {
	this := WebhookResponse{}
	return &this
}

// NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseWithDefaults() *WebhookResponse {
	this := WebhookResponse{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookResponse) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted.Get()) {
		var ret bool
		return ret
	}
	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// HasDeleted returns a boolean if a field has been set.
func (o *WebhookResponse) HasDeleted() bool {
	if o != nil && o.Deleted.IsSet() {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given NullableBool and assigns it to the Deleted field.
func (o *WebhookResponse) SetDeleted(v bool) {
	o.Deleted.Set(&v)
}

// SetDeletedNil sets the value for Deleted to be an explicit nil
func (o *WebhookResponse) SetDeletedNil() {
	o.Deleted.Set(nil)
}

// UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
func (o *WebhookResponse) UnsetDeleted() {
	o.Deleted.Unset()
}

// GetDevelopmentEnabled returns the DevelopmentEnabled field value if set, zero value otherwise.
func (o *WebhookResponse) GetDevelopmentEnabled() bool {
	if o == nil || IsNil(o.DevelopmentEnabled) {
		var ret bool
		return ret
	}
	return *o.DevelopmentEnabled
}

// GetDevelopmentEnabledOk returns a tuple with the DevelopmentEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetDevelopmentEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DevelopmentEnabled) {
		return nil, false
	}
	return o.DevelopmentEnabled, true
}

// HasDevelopmentEnabled returns a boolean if a field has been set.
func (o *WebhookResponse) HasDevelopmentEnabled() bool {
	if o != nil && !IsNil(o.DevelopmentEnabled) {
		return true
	}

	return false
}

// SetDevelopmentEnabled gets a reference to the given bool and assigns it to the DevelopmentEnabled field.
func (o *WebhookResponse) SetDevelopmentEnabled(v bool) {
	o.DevelopmentEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhookResponse) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *WebhookResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *WebhookResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *WebhookResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *WebhookResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *WebhookResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *WebhookResponse) SetObject(v string) {
	o.Object = &v
}

// GetProductionEnabled returns the ProductionEnabled field value if set, zero value otherwise.
func (o *WebhookResponse) GetProductionEnabled() bool {
	if o == nil || IsNil(o.ProductionEnabled) {
		var ret bool
		return ret
	}
	return *o.ProductionEnabled
}

// GetProductionEnabledOk returns a tuple with the ProductionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetProductionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProductionEnabled) {
		return nil, false
	}
	return o.ProductionEnabled, true
}

// HasProductionEnabled returns a boolean if a field has been set.
func (o *WebhookResponse) HasProductionEnabled() bool {
	if o != nil && !IsNil(o.ProductionEnabled) {
		return true
	}

	return false
}

// SetProductionEnabled gets a reference to the given bool and assigns it to the ProductionEnabled field.
func (o *WebhookResponse) SetProductionEnabled(v bool) {
	o.ProductionEnabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebhookResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhookResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WebhookResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubscribedEvents returns the SubscribedEvents field value if set, zero value otherwise.
func (o *WebhookResponse) GetSubscribedEvents() []string {
	if o == nil || IsNil(o.SubscribedEvents) {
		var ret []string
		return ret
	}
	return o.SubscribedEvents
}

// GetSubscribedEventsOk returns a tuple with the SubscribedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSubscribedEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubscribedEvents) {
		return nil, false
	}
	return o.SubscribedEvents, true
}

// HasSubscribedEvents returns a boolean if a field has been set.
func (o *WebhookResponse) HasSubscribedEvents() bool {
	if o != nil && !IsNil(o.SubscribedEvents) {
		return true
	}

	return false
}

// SetSubscribedEvents gets a reference to the given []string and assigns it to the SubscribedEvents field.
func (o *WebhookResponse) SetSubscribedEvents(v []string) {
	o.SubscribedEvents = v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *WebhookResponse) GetSynchronous() bool {
	if o == nil || IsNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Synchronous) {
		return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *WebhookResponse) HasSynchronous() bool {
	if o != nil && !IsNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *WebhookResponse) SetSynchronous(v bool) {
	o.Synchronous = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WebhookResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WebhookResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WebhookResponse) SetUrl(v string) {
	o.Url = &v
}

func (o WebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Deleted.IsSet() {
		toSerialize["deleted"] = o.Deleted.Get()
	}
	if !IsNil(o.DevelopmentEnabled) {
		toSerialize["development_enabled"] = o.DevelopmentEnabled
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
	if !IsNil(o.ProductionEnabled) {
		toSerialize["production_enabled"] = o.ProductionEnabled
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscribedEvents) {
		toSerialize["subscribed_events"] = o.SubscribedEvents
	}
	if !IsNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableWebhookResponse struct {
	value *WebhookResponse
	isSet bool
}

func (v NullableWebhookResponse) Get() *WebhookResponse {
	return v.value
}

func (v *NullableWebhookResponse) Set(val *WebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponse(val *WebhookResponse) *NullableWebhookResponse {
	return &NullableWebhookResponse{value: val, isSet: true}
}

func (v NullableWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

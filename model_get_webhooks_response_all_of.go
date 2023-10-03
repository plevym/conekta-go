package conekta

import (
	"encoding/json"
)

// checks if the GetWebhooksResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWebhooksResponseAllOf{}

// GetWebhooksResponseAllOf struct for GetWebhooksResponseAllOf
type GetWebhooksResponseAllOf struct {
	Data []WebhookResponse `json:"data,omitempty"`
}

// NewGetWebhooksResponseAllOf instantiates a new GetWebhooksResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhooksResponseAllOf() *GetWebhooksResponseAllOf {
	this := GetWebhooksResponseAllOf{}
	return &this
}

// NewGetWebhooksResponseAllOfWithDefaults instantiates a new GetWebhooksResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhooksResponseAllOfWithDefaults() *GetWebhooksResponseAllOf {
	this := GetWebhooksResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetWebhooksResponseAllOf) GetData() []WebhookResponse {
	if o == nil || IsNil(o.Data) {
		var ret []WebhookResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhooksResponseAllOf) GetDataOk() ([]WebhookResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetWebhooksResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []WebhookResponse and assigns it to the Data field.
func (o *GetWebhooksResponseAllOf) SetData(v []WebhookResponse) {
	o.Data = v
}

func (o GetWebhooksResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWebhooksResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetWebhooksResponseAllOf struct {
	value *GetWebhooksResponseAllOf
	isSet bool
}

func (v NullableGetWebhooksResponseAllOf) Get() *GetWebhooksResponseAllOf {
	return v.value
}

func (v *NullableGetWebhooksResponseAllOf) Set(val *GetWebhooksResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhooksResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhooksResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhooksResponseAllOf(val *GetWebhooksResponseAllOf) *NullableGetWebhooksResponseAllOf {
	return &NullableGetWebhooksResponseAllOf{value: val, isSet: true}
}

func (v NullableGetWebhooksResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhooksResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

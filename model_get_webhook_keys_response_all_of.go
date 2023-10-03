package conekta

import (
	"encoding/json"
)

// checks if the GetWebhookKeysResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWebhookKeysResponseAllOf{}

// GetWebhookKeysResponseAllOf struct for GetWebhookKeysResponseAllOf
type GetWebhookKeysResponseAllOf struct {
	Data []WebhookKeyResponse `json:"data,omitempty"`
}

// NewGetWebhookKeysResponseAllOf instantiates a new GetWebhookKeysResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhookKeysResponseAllOf() *GetWebhookKeysResponseAllOf {
	this := GetWebhookKeysResponseAllOf{}
	return &this
}

// NewGetWebhookKeysResponseAllOfWithDefaults instantiates a new GetWebhookKeysResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhookKeysResponseAllOfWithDefaults() *GetWebhookKeysResponseAllOf {
	this := GetWebhookKeysResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetWebhookKeysResponseAllOf) GetData() []WebhookKeyResponse {
	if o == nil || IsNil(o.Data) {
		var ret []WebhookKeyResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetWebhookKeysResponseAllOf) GetDataOk() ([]WebhookKeyResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetWebhookKeysResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []WebhookKeyResponse and assigns it to the Data field.
func (o *GetWebhookKeysResponseAllOf) SetData(v []WebhookKeyResponse) {
	o.Data = v
}

func (o GetWebhookKeysResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWebhookKeysResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetWebhookKeysResponseAllOf struct {
	value *GetWebhookKeysResponseAllOf
	isSet bool
}

func (v NullableGetWebhookKeysResponseAllOf) Get() *GetWebhookKeysResponseAllOf {
	return v.value
}

func (v *NullableGetWebhookKeysResponseAllOf) Set(val *GetWebhookKeysResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhookKeysResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhookKeysResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhookKeysResponseAllOf(val *GetWebhookKeysResponseAllOf) *NullableGetWebhookKeysResponseAllOf {
	return &NullableGetWebhookKeysResponseAllOf{value: val, isSet: true}
}

func (v NullableGetWebhookKeysResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhookKeysResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

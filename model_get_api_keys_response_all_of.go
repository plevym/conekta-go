package conekta

import (
	"encoding/json"
)

// checks if the GetApiKeysResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApiKeysResponseAllOf{}

// GetApiKeysResponseAllOf struct for GetApiKeysResponseAllOf
type GetApiKeysResponseAllOf struct {
	Data []ApiKeyResponse `json:"data,omitempty"`
}

// NewGetApiKeysResponseAllOf instantiates a new GetApiKeysResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiKeysResponseAllOf() *GetApiKeysResponseAllOf {
	this := GetApiKeysResponseAllOf{}
	return &this
}

// NewGetApiKeysResponseAllOfWithDefaults instantiates a new GetApiKeysResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiKeysResponseAllOfWithDefaults() *GetApiKeysResponseAllOf {
	this := GetApiKeysResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetApiKeysResponseAllOf) GetData() []ApiKeyResponse {
	if o == nil || IsNil(o.Data) {
		var ret []ApiKeyResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiKeysResponseAllOf) GetDataOk() ([]ApiKeyResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetApiKeysResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ApiKeyResponse and assigns it to the Data field.
func (o *GetApiKeysResponseAllOf) SetData(v []ApiKeyResponse) {
	o.Data = v
}

func (o GetApiKeysResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApiKeysResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetApiKeysResponseAllOf struct {
	value *GetApiKeysResponseAllOf
	isSet bool
}

func (v NullableGetApiKeysResponseAllOf) Get() *GetApiKeysResponseAllOf {
	return v.value
}

func (v *NullableGetApiKeysResponseAllOf) Set(val *GetApiKeysResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiKeysResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiKeysResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiKeysResponseAllOf(val *GetApiKeysResponseAllOf) *NullableGetApiKeysResponseAllOf {
	return &NullableGetApiKeysResponseAllOf{value: val, isSet: true}
}

func (v NullableGetApiKeysResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiKeysResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

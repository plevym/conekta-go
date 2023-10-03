package conekta

import (
	"encoding/json"
)

// checks if the GetChargesResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChargesResponseAllOf{}

// GetChargesResponseAllOf struct for GetChargesResponseAllOf
type GetChargesResponseAllOf struct {
	Data []ChargeResponse `json:"data,omitempty"`
}

// NewGetChargesResponseAllOf instantiates a new GetChargesResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChargesResponseAllOf() *GetChargesResponseAllOf {
	this := GetChargesResponseAllOf{}
	return &this
}

// NewGetChargesResponseAllOfWithDefaults instantiates a new GetChargesResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChargesResponseAllOfWithDefaults() *GetChargesResponseAllOf {
	this := GetChargesResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetChargesResponseAllOf) GetData() []ChargeResponse {
	if o == nil || IsNil(o.Data) {
		var ret []ChargeResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChargesResponseAllOf) GetDataOk() ([]ChargeResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetChargesResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ChargeResponse and assigns it to the Data field.
func (o *GetChargesResponseAllOf) SetData(v []ChargeResponse) {
	o.Data = v
}

func (o GetChargesResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChargesResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetChargesResponseAllOf struct {
	value *GetChargesResponseAllOf
	isSet bool
}

func (v NullableGetChargesResponseAllOf) Get() *GetChargesResponseAllOf {
	return v.value
}

func (v *NullableGetChargesResponseAllOf) Set(val *GetChargesResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChargesResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChargesResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChargesResponseAllOf(val *GetChargesResponseAllOf) *NullableGetChargesResponseAllOf {
	return &NullableGetChargesResponseAllOf{value: val, isSet: true}
}

func (v NullableGetChargesResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChargesResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

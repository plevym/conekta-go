package conekta

import (
	"encoding/json"
)

// checks if the GetPlansResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPlansResponseAllOf{}

// GetPlansResponseAllOf struct for GetPlansResponseAllOf
type GetPlansResponseAllOf struct {
	Data []PlanResponse `json:"data,omitempty"`
}

// NewGetPlansResponseAllOf instantiates a new GetPlansResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPlansResponseAllOf() *GetPlansResponseAllOf {
	this := GetPlansResponseAllOf{}
	return &this
}

// NewGetPlansResponseAllOfWithDefaults instantiates a new GetPlansResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPlansResponseAllOfWithDefaults() *GetPlansResponseAllOf {
	this := GetPlansResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPlansResponseAllOf) GetData() []PlanResponse {
	if o == nil || IsNil(o.Data) {
		var ret []PlanResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPlansResponseAllOf) GetDataOk() ([]PlanResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPlansResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PlanResponse and assigns it to the Data field.
func (o *GetPlansResponseAllOf) SetData(v []PlanResponse) {
	o.Data = v
}

func (o GetPlansResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPlansResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetPlansResponseAllOf struct {
	value *GetPlansResponseAllOf
	isSet bool
}

func (v NullableGetPlansResponseAllOf) Get() *GetPlansResponseAllOf {
	return v.value
}

func (v *NullableGetPlansResponseAllOf) Set(val *GetPlansResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPlansResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPlansResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPlansResponseAllOf(val *GetPlansResponseAllOf) *NullableGetPlansResponseAllOf {
	return &NullableGetPlansResponseAllOf{value: val, isSet: true}
}

func (v NullableGetPlansResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPlansResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

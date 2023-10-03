package conekta

import (
	"encoding/json"
)

// checks if the GetEventsResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventsResponseAllOf{}

// GetEventsResponseAllOf struct for GetEventsResponseAllOf
type GetEventsResponseAllOf struct {
	Data []EventResponse `json:"data,omitempty"`
}

// NewGetEventsResponseAllOf instantiates a new GetEventsResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventsResponseAllOf() *GetEventsResponseAllOf {
	this := GetEventsResponseAllOf{}
	return &this
}

// NewGetEventsResponseAllOfWithDefaults instantiates a new GetEventsResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventsResponseAllOfWithDefaults() *GetEventsResponseAllOf {
	this := GetEventsResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetEventsResponseAllOf) GetData() []EventResponse {
	if o == nil || IsNil(o.Data) {
		var ret []EventResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventsResponseAllOf) GetDataOk() ([]EventResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetEventsResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EventResponse and assigns it to the Data field.
func (o *GetEventsResponseAllOf) SetData(v []EventResponse) {
	o.Data = v
}

func (o GetEventsResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventsResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetEventsResponseAllOf struct {
	value *GetEventsResponseAllOf
	isSet bool
}

func (v NullableGetEventsResponseAllOf) Get() *GetEventsResponseAllOf {
	return v.value
}

func (v *NullableGetEventsResponseAllOf) Set(val *GetEventsResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventsResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventsResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventsResponseAllOf(val *GetEventsResponseAllOf) *NullableGetEventsResponseAllOf {
	return &NullableGetEventsResponseAllOf{value: val, isSet: true}
}

func (v NullableGetEventsResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventsResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the GetOrderDiscountLinesResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderDiscountLinesResponseAllOf{}

// GetOrderDiscountLinesResponseAllOf struct for GetOrderDiscountLinesResponseAllOf
type GetOrderDiscountLinesResponseAllOf struct {
	Data []DiscountLinesResponse `json:"data,omitempty"`
}

// NewGetOrderDiscountLinesResponseAllOf instantiates a new GetOrderDiscountLinesResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderDiscountLinesResponseAllOf() *GetOrderDiscountLinesResponseAllOf {
	this := GetOrderDiscountLinesResponseAllOf{}
	return &this
}

// NewGetOrderDiscountLinesResponseAllOfWithDefaults instantiates a new GetOrderDiscountLinesResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderDiscountLinesResponseAllOfWithDefaults() *GetOrderDiscountLinesResponseAllOf {
	this := GetOrderDiscountLinesResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetOrderDiscountLinesResponseAllOf) GetData() []DiscountLinesResponse {
	if o == nil || IsNil(o.Data) {
		var ret []DiscountLinesResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDiscountLinesResponseAllOf) GetDataOk() ([]DiscountLinesResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetOrderDiscountLinesResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DiscountLinesResponse and assigns it to the Data field.
func (o *GetOrderDiscountLinesResponseAllOf) SetData(v []DiscountLinesResponse) {
	o.Data = v
}

func (o GetOrderDiscountLinesResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderDiscountLinesResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetOrderDiscountLinesResponseAllOf struct {
	value *GetOrderDiscountLinesResponseAllOf
	isSet bool
}

func (v NullableGetOrderDiscountLinesResponseAllOf) Get() *GetOrderDiscountLinesResponseAllOf {
	return v.value
}

func (v *NullableGetOrderDiscountLinesResponseAllOf) Set(val *GetOrderDiscountLinesResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderDiscountLinesResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderDiscountLinesResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderDiscountLinesResponseAllOf(val *GetOrderDiscountLinesResponseAllOf) *NullableGetOrderDiscountLinesResponseAllOf {
	return &NullableGetOrderDiscountLinesResponseAllOf{value: val, isSet: true}
}

func (v NullableGetOrderDiscountLinesResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderDiscountLinesResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

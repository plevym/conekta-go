package conekta

import (
	"encoding/json"
)

// checks if the CheckoutsResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutsResponseAllOf{}

// CheckoutsResponseAllOf struct for CheckoutsResponseAllOf
type CheckoutsResponseAllOf struct {
	Data []CheckoutResponse `json:"data,omitempty"`
}

// NewCheckoutsResponseAllOf instantiates a new CheckoutsResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutsResponseAllOf() *CheckoutsResponseAllOf {
	this := CheckoutsResponseAllOf{}
	return &this
}

// NewCheckoutsResponseAllOfWithDefaults instantiates a new CheckoutsResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutsResponseAllOfWithDefaults() *CheckoutsResponseAllOf {
	this := CheckoutsResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckoutsResponseAllOf) GetData() []CheckoutResponse {
	if o == nil || IsNil(o.Data) {
		var ret []CheckoutResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutsResponseAllOf) GetDataOk() ([]CheckoutResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckoutsResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CheckoutResponse and assigns it to the Data field.
func (o *CheckoutsResponseAllOf) SetData(v []CheckoutResponse) {
	o.Data = v
}

func (o CheckoutsResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutsResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCheckoutsResponseAllOf struct {
	value *CheckoutsResponseAllOf
	isSet bool
}

func (v NullableCheckoutsResponseAllOf) Get() *CheckoutsResponseAllOf {
	return v.value
}

func (v *NullableCheckoutsResponseAllOf) Set(val *CheckoutsResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutsResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutsResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutsResponseAllOf(val *CheckoutsResponseAllOf) *NullableCheckoutsResponseAllOf {
	return &NullableCheckoutsResponseAllOf{value: val, isSet: true}
}

func (v NullableCheckoutsResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutsResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the ChargeResponseRefundsAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeResponseRefundsAllOf{}

// ChargeResponseRefundsAllOf struct for ChargeResponseRefundsAllOf
type ChargeResponseRefundsAllOf struct {
	// refunds
	Data []ChargeResponseRefundsData `json:"data,omitempty"`
}

// NewChargeResponseRefundsAllOf instantiates a new ChargeResponseRefundsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeResponseRefundsAllOf() *ChargeResponseRefundsAllOf {
	this := ChargeResponseRefundsAllOf{}
	return &this
}

// NewChargeResponseRefundsAllOfWithDefaults instantiates a new ChargeResponseRefundsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeResponseRefundsAllOfWithDefaults() *ChargeResponseRefundsAllOf {
	this := ChargeResponseRefundsAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ChargeResponseRefundsAllOf) GetData() []ChargeResponseRefundsData {
	if o == nil || IsNil(o.Data) {
		var ret []ChargeResponseRefundsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponseRefundsAllOf) GetDataOk() ([]ChargeResponseRefundsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ChargeResponseRefundsAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ChargeResponseRefundsData and assigns it to the Data field.
func (o *ChargeResponseRefundsAllOf) SetData(v []ChargeResponseRefundsData) {
	o.Data = v
}

func (o ChargeResponseRefundsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeResponseRefundsAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableChargeResponseRefundsAllOf struct {
	value *ChargeResponseRefundsAllOf
	isSet bool
}

func (v NullableChargeResponseRefundsAllOf) Get() *ChargeResponseRefundsAllOf {
	return v.value
}

func (v *NullableChargeResponseRefundsAllOf) Set(val *ChargeResponseRefundsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeResponseRefundsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeResponseRefundsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeResponseRefundsAllOf(val *ChargeResponseRefundsAllOf) *NullableChargeResponseRefundsAllOf {
	return &NullableChargeResponseRefundsAllOf{value: val, isSet: true}
}

func (v NullableChargeResponseRefundsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeResponseRefundsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

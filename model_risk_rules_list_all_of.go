package conekta

import (
	"encoding/json"
)

// checks if the RiskRulesListAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskRulesListAllOf{}

// RiskRulesListAllOf struct for RiskRulesListAllOf
type RiskRulesListAllOf struct {
	Data []RiskRulesListAllOfData `json:"data,omitempty"`
}

// NewRiskRulesListAllOf instantiates a new RiskRulesListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRulesListAllOf() *RiskRulesListAllOf {
	this := RiskRulesListAllOf{}
	return &this
}

// NewRiskRulesListAllOfWithDefaults instantiates a new RiskRulesListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRulesListAllOfWithDefaults() *RiskRulesListAllOf {
	this := RiskRulesListAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RiskRulesListAllOf) GetData() []RiskRulesListAllOfData {
	if o == nil || IsNil(o.Data) {
		var ret []RiskRulesListAllOfData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesListAllOf) GetDataOk() ([]RiskRulesListAllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RiskRulesListAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RiskRulesListAllOfData and assigns it to the Data field.
func (o *RiskRulesListAllOf) SetData(v []RiskRulesListAllOfData) {
	o.Data = v
}

func (o RiskRulesListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskRulesListAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRiskRulesListAllOf struct {
	value *RiskRulesListAllOf
	isSet bool
}

func (v NullableRiskRulesListAllOf) Get() *RiskRulesListAllOf {
	return v.value
}

func (v *NullableRiskRulesListAllOf) Set(val *RiskRulesListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRulesListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRulesListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRulesListAllOf(val *RiskRulesListAllOf) *NullableRiskRulesListAllOf {
	return &NullableRiskRulesListAllOf{value: val, isSet: true}
}

func (v NullableRiskRulesListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRulesListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the RiskRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskRules{}

// RiskRules struct for RiskRules
type RiskRules struct {
	Data []RiskRulesData `json:"data,omitempty"`
}

// NewRiskRules instantiates a new RiskRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRules() *RiskRules {
	this := RiskRules{}
	return &this
}

// NewRiskRulesWithDefaults instantiates a new RiskRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRulesWithDefaults() *RiskRules {
	this := RiskRules{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RiskRules) GetData() []RiskRulesData {
	if o == nil || IsNil(o.Data) {
		var ret []RiskRulesData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRules) GetDataOk() ([]RiskRulesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RiskRules) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RiskRulesData and assigns it to the Data field.
func (o *RiskRules) SetData(v []RiskRulesData) {
	o.Data = v
}

func (o RiskRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRiskRules struct {
	value *RiskRules
	isSet bool
}

func (v NullableRiskRules) Get() *RiskRules {
	return v.value
}

func (v *NullableRiskRules) Set(val *RiskRules) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRules) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRules(val *RiskRules) *NullableRiskRules {
	return &NullableRiskRules{value: val, isSet: true}
}

func (v NullableRiskRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

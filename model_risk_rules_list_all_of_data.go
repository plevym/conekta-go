package conekta

import (
	"encoding/json"
)

// checks if the RiskRulesListAllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskRulesListAllOfData{}

// RiskRulesListAllOfData struct for RiskRulesListAllOfData
type RiskRulesListAllOfData struct {
	// rule id
	Id *string `json:"id,omitempty"`
	// field to be used for the rule
	Field *string `json:"field,omitempty"`
	// rule creation date
	CreatedAt *string `json:"created_at,omitempty"`
	// value to be used for the rule
	Value *string `json:"value,omitempty"`
	// if the rule is global
	IsGlobal *bool `json:"is_global,omitempty"`
	// if the rule is test
	IsTest *bool `json:"is_test,omitempty"`
	// description of the rule
	Description *string `json:"description,omitempty"`
}

// NewRiskRulesListAllOfData instantiates a new RiskRulesListAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRulesListAllOfData() *RiskRulesListAllOfData {
	this := RiskRulesListAllOfData{}
	return &this
}

// NewRiskRulesListAllOfDataWithDefaults instantiates a new RiskRulesListAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRulesListAllOfDataWithDefaults() *RiskRulesListAllOfData {
	this := RiskRulesListAllOfData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskRulesListAllOfData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesListAllOfData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskRulesListAllOfData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskRulesListAllOfData) SetId(v string) {
	o.Id = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *RiskRulesListAllOfData) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesListAllOfData) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *RiskRulesListAllOfData) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *RiskRulesListAllOfData) SetField(v string) {
	o.Field = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskRulesListAllOfData) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesListAllOfData) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskRulesListAllOfData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RiskRulesListAllOfData) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskRulesListAllOfData) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesListAllOfData) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskRulesListAllOfData) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RiskRulesListAllOfData) SetValue(v string) {
	o.Value = &v
}

// GetIsGlobal returns the IsGlobal field value if set, zero value otherwise.
func (o *RiskRulesListAllOfData) GetIsGlobal() bool {
	if o == nil || IsNil(o.IsGlobal) {
		var ret bool
		return ret
	}
	return *o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesListAllOfData) GetIsGlobalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobal) {
		return nil, false
	}
	return o.IsGlobal, true
}

// HasIsGlobal returns a boolean if a field has been set.
func (o *RiskRulesListAllOfData) HasIsGlobal() bool {
	if o != nil && !IsNil(o.IsGlobal) {
		return true
	}

	return false
}

// SetIsGlobal gets a reference to the given bool and assigns it to the IsGlobal field.
func (o *RiskRulesListAllOfData) SetIsGlobal(v bool) {
	o.IsGlobal = &v
}

// GetIsTest returns the IsTest field value if set, zero value otherwise.
func (o *RiskRulesListAllOfData) GetIsTest() bool {
	if o == nil || IsNil(o.IsTest) {
		var ret bool
		return ret
	}
	return *o.IsTest
}

// GetIsTestOk returns a tuple with the IsTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesListAllOfData) GetIsTestOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTest) {
		return nil, false
	}
	return o.IsTest, true
}

// HasIsTest returns a boolean if a field has been set.
func (o *RiskRulesListAllOfData) HasIsTest() bool {
	if o != nil && !IsNil(o.IsTest) {
		return true
	}

	return false
}

// SetIsTest gets a reference to the given bool and assigns it to the IsTest field.
func (o *RiskRulesListAllOfData) SetIsTest(v bool) {
	o.IsTest = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskRulesListAllOfData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesListAllOfData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskRulesListAllOfData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskRulesListAllOfData) SetDescription(v string) {
	o.Description = &v
}

func (o RiskRulesListAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskRulesListAllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.IsGlobal) {
		toSerialize["is_global"] = o.IsGlobal
	}
	if !IsNil(o.IsTest) {
		toSerialize["is_test"] = o.IsTest
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableRiskRulesListAllOfData struct {
	value *RiskRulesListAllOfData
	isSet bool
}

func (v NullableRiskRulesListAllOfData) Get() *RiskRulesListAllOfData {
	return v.value
}

func (v *NullableRiskRulesListAllOfData) Set(val *RiskRulesListAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRulesListAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRulesListAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRulesListAllOfData(val *RiskRulesListAllOfData) *NullableRiskRulesListAllOfData {
	return &NullableRiskRulesListAllOfData{value: val, isSet: true}
}

func (v NullableRiskRulesListAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRulesListAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

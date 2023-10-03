package conekta

import (
	"encoding/json"
)

// checks if the RiskRulesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskRulesData{}

// RiskRulesData struct for RiskRulesData
type RiskRulesData struct {
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

// NewRiskRulesData instantiates a new RiskRulesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRulesData() *RiskRulesData {
	this := RiskRulesData{}
	return &this
}

// NewRiskRulesDataWithDefaults instantiates a new RiskRulesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRulesDataWithDefaults() *RiskRulesData {
	this := RiskRulesData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskRulesData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskRulesData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskRulesData) SetId(v string) {
	o.Id = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *RiskRulesData) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesData) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *RiskRulesData) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *RiskRulesData) SetField(v string) {
	o.Field = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskRulesData) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesData) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskRulesData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RiskRulesData) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskRulesData) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesData) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskRulesData) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RiskRulesData) SetValue(v string) {
	o.Value = &v
}

// GetIsGlobal returns the IsGlobal field value if set, zero value otherwise.
func (o *RiskRulesData) GetIsGlobal() bool {
	if o == nil || IsNil(o.IsGlobal) {
		var ret bool
		return ret
	}
	return *o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesData) GetIsGlobalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobal) {
		return nil, false
	}
	return o.IsGlobal, true
}

// HasIsGlobal returns a boolean if a field has been set.
func (o *RiskRulesData) HasIsGlobal() bool {
	if o != nil && !IsNil(o.IsGlobal) {
		return true
	}

	return false
}

// SetIsGlobal gets a reference to the given bool and assigns it to the IsGlobal field.
func (o *RiskRulesData) SetIsGlobal(v bool) {
	o.IsGlobal = &v
}

// GetIsTest returns the IsTest field value if set, zero value otherwise.
func (o *RiskRulesData) GetIsTest() bool {
	if o == nil || IsNil(o.IsTest) {
		var ret bool
		return ret
	}
	return *o.IsTest
}

// GetIsTestOk returns a tuple with the IsTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesData) GetIsTestOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTest) {
		return nil, false
	}
	return o.IsTest, true
}

// HasIsTest returns a boolean if a field has been set.
func (o *RiskRulesData) HasIsTest() bool {
	if o != nil && !IsNil(o.IsTest) {
		return true
	}

	return false
}

// SetIsTest gets a reference to the given bool and assigns it to the IsTest field.
func (o *RiskRulesData) SetIsTest(v bool) {
	o.IsTest = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskRulesData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskRulesData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskRulesData) SetDescription(v string) {
	o.Description = &v
}

func (o RiskRulesData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskRulesData) ToMap() (map[string]interface{}, error) {
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

type NullableRiskRulesData struct {
	value *RiskRulesData
	isSet bool
}

func (v NullableRiskRulesData) Get() *RiskRulesData {
	return v.value
}

func (v *NullableRiskRulesData) Set(val *RiskRulesData) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRulesData) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRulesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRulesData(val *RiskRulesData) *NullableRiskRulesData {
	return &NullableRiskRulesData{value: val, isSet: true}
}

func (v NullableRiskRulesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRulesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

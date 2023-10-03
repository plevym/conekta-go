package conekta

import (
	"encoding/json"
)

// checks if the RiskRulesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskRulesList{}

// RiskRulesList struct for RiskRulesList
type RiskRulesList struct {
	// Indicates if there are more pages to be requested
	HasMore bool `json:"has_more"`
	// Object type, in this case is list
	Object string `json:"object"`
	// URL of the next page.
	NextPageUrl NullableString `json:"next_page_url,omitempty"`
	// Url of the previous page.
	PreviousPageUrl NullableString  `json:"previous_page_url,omitempty"`
	Data            []RiskRulesData `json:"data,omitempty"`
}

// NewRiskRulesList instantiates a new RiskRulesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRulesList(hasMore bool, object string) *RiskRulesList {
	this := RiskRulesList{}
	this.HasMore = hasMore
	this.Object = object
	return &this
}

// NewRiskRulesListWithDefaults instantiates a new RiskRulesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRulesListWithDefaults() *RiskRulesList {
	this := RiskRulesList{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *RiskRulesList) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *RiskRulesList) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *RiskRulesList) SetHasMore(v bool) {
	o.HasMore = v
}

// GetObject returns the Object field value
func (o *RiskRulesList) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *RiskRulesList) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *RiskRulesList) SetObject(v string) {
	o.Object = v
}

// GetNextPageUrl returns the NextPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskRulesList) GetNextPageUrl() string {
	if o == nil || IsNil(o.NextPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageUrl.Get()
}

// GetNextPageUrlOk returns a tuple with the NextPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskRulesList) GetNextPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageUrl.Get(), o.NextPageUrl.IsSet()
}

// HasNextPageUrl returns a boolean if a field has been set.
func (o *RiskRulesList) HasNextPageUrl() bool {
	if o != nil && o.NextPageUrl.IsSet() {
		return true
	}

	return false
}

// SetNextPageUrl gets a reference to the given NullableString and assigns it to the NextPageUrl field.
func (o *RiskRulesList) SetNextPageUrl(v string) {
	o.NextPageUrl.Set(&v)
}

// SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil
func (o *RiskRulesList) SetNextPageUrlNil() {
	o.NextPageUrl.Set(nil)
}

// UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
func (o *RiskRulesList) UnsetNextPageUrl() {
	o.NextPageUrl.Unset()
}

// GetPreviousPageUrl returns the PreviousPageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskRulesList) GetPreviousPageUrl() string {
	if o == nil || IsNil(o.PreviousPageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousPageUrl.Get()
}

// GetPreviousPageUrlOk returns a tuple with the PreviousPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskRulesList) GetPreviousPageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousPageUrl.Get(), o.PreviousPageUrl.IsSet()
}

// HasPreviousPageUrl returns a boolean if a field has been set.
func (o *RiskRulesList) HasPreviousPageUrl() bool {
	if o != nil && o.PreviousPageUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviousPageUrl gets a reference to the given NullableString and assigns it to the PreviousPageUrl field.
func (o *RiskRulesList) SetPreviousPageUrl(v string) {
	o.PreviousPageUrl.Set(&v)
}

// SetPreviousPageUrlNil sets the value for PreviousPageUrl to be an explicit nil
func (o *RiskRulesList) SetPreviousPageUrlNil() {
	o.PreviousPageUrl.Set(nil)
}

// UnsetPreviousPageUrl ensures that no value is present for PreviousPageUrl, not even an explicit nil
func (o *RiskRulesList) UnsetPreviousPageUrl() {
	o.PreviousPageUrl.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RiskRulesList) GetData() []RiskRulesData {
	if o == nil || IsNil(o.Data) {
		var ret []RiskRulesData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRulesList) GetDataOk() ([]RiskRulesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RiskRulesList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RiskRulesData and assigns it to the Data field.
func (o *RiskRulesList) SetData(v []RiskRulesData) {
	o.Data = v
}

func (o RiskRulesList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskRulesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_more"] = o.HasMore
	toSerialize["object"] = o.Object
	if o.NextPageUrl.IsSet() {
		toSerialize["next_page_url"] = o.NextPageUrl.Get()
	}
	if o.PreviousPageUrl.IsSet() {
		toSerialize["previous_page_url"] = o.PreviousPageUrl.Get()
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRiskRulesList struct {
	value *RiskRulesList
	isSet bool
}

func (v NullableRiskRulesList) Get() *RiskRulesList {
	return v.value
}

func (v *NullableRiskRulesList) Set(val *RiskRulesList) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRulesList) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRulesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRulesList(val *RiskRulesList) *NullableRiskRulesList {
	return &NullableRiskRulesList{value: val, isSet: true}
}

func (v NullableRiskRulesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRulesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the GetCompaniesResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCompaniesResponseAllOf{}

// GetCompaniesResponseAllOf struct for GetCompaniesResponseAllOf
type GetCompaniesResponseAllOf struct {
	Data []CompanyResponse `json:"data,omitempty"`
}

// NewGetCompaniesResponseAllOf instantiates a new GetCompaniesResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCompaniesResponseAllOf() *GetCompaniesResponseAllOf {
	this := GetCompaniesResponseAllOf{}
	return &this
}

// NewGetCompaniesResponseAllOfWithDefaults instantiates a new GetCompaniesResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCompaniesResponseAllOfWithDefaults() *GetCompaniesResponseAllOf {
	this := GetCompaniesResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCompaniesResponseAllOf) GetData() []CompanyResponse {
	if o == nil || IsNil(o.Data) {
		var ret []CompanyResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCompaniesResponseAllOf) GetDataOk() ([]CompanyResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCompaniesResponseAllOf) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CompanyResponse and assigns it to the Data field.
func (o *GetCompaniesResponseAllOf) SetData(v []CompanyResponse) {
	o.Data = v
}

func (o GetCompaniesResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCompaniesResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetCompaniesResponseAllOf struct {
	value *GetCompaniesResponseAllOf
	isSet bool
}

func (v NullableGetCompaniesResponseAllOf) Get() *GetCompaniesResponseAllOf {
	return v.value
}

func (v *NullableGetCompaniesResponseAllOf) Set(val *GetCompaniesResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCompaniesResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCompaniesResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCompaniesResponseAllOf(val *GetCompaniesResponseAllOf) *NullableGetCompaniesResponseAllOf {
	return &NullableGetCompaniesResponseAllOf{value: val, isSet: true}
}

func (v NullableGetCompaniesResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCompaniesResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

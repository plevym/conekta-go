package conekta

import (
	"encoding/json"
)

// checks if the Details type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Details{}

// Details struct for Details
type Details struct {
	Details []DetailsError `json:"details,omitempty"`
}

// NewDetails instantiates a new Details object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetails() *Details {
	this := Details{}
	return &this
}

// NewDetailsWithDefaults instantiates a new Details object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailsWithDefaults() *Details {
	this := Details{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Details) GetDetails() []DetailsError {
	if o == nil || IsNil(o.Details) {
		var ret []DetailsError
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Details) GetDetailsOk() ([]DetailsError, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Details) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []DetailsError and assigns it to the Details field.
func (o *Details) SetDetails(v []DetailsError) {
	o.Details = v
}

func (o Details) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Details) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableDetails struct {
	value *Details
	isSet bool
}

func (v NullableDetails) Get() *Details {
	return v.value
}

func (v *NullableDetails) Set(val *Details) {
	v.value = val
	v.isSet = true
}

func (v NullableDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetails(val *Details) *NullableDetails {
	return &NullableDetails{value: val, isSet: true}
}

func (v NullableDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

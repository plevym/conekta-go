package conekta

import (
	"encoding/json"
)

// checks if the DeleteApiKeysResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteApiKeysResponseAllOf{}

// DeleteApiKeysResponseAllOf struct for DeleteApiKeysResponseAllOf
type DeleteApiKeysResponseAllOf struct {
	Deleted *bool `json:"deleted,omitempty"`
}

// NewDeleteApiKeysResponseAllOf instantiates a new DeleteApiKeysResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteApiKeysResponseAllOf() *DeleteApiKeysResponseAllOf {
	this := DeleteApiKeysResponseAllOf{}
	return &this
}

// NewDeleteApiKeysResponseAllOfWithDefaults instantiates a new DeleteApiKeysResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteApiKeysResponseAllOfWithDefaults() *DeleteApiKeysResponseAllOf {
	this := DeleteApiKeysResponseAllOf{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *DeleteApiKeysResponseAllOf) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteApiKeysResponseAllOf) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *DeleteApiKeysResponseAllOf) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *DeleteApiKeysResponseAllOf) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o DeleteApiKeysResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteApiKeysResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

type NullableDeleteApiKeysResponseAllOf struct {
	value *DeleteApiKeysResponseAllOf
	isSet bool
}

func (v NullableDeleteApiKeysResponseAllOf) Get() *DeleteApiKeysResponseAllOf {
	return v.value
}

func (v *NullableDeleteApiKeysResponseAllOf) Set(val *DeleteApiKeysResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteApiKeysResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteApiKeysResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteApiKeysResponseAllOf(val *DeleteApiKeysResponseAllOf) *NullableDeleteApiKeysResponseAllOf {
	return &NullableDeleteApiKeysResponseAllOf{value: val, isSet: true}
}

func (v NullableDeleteApiKeysResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteApiKeysResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

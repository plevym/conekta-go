package conekta

import (
	"encoding/json"
)

// checks if the ProductDataResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDataResponseAllOf{}

// ProductDataResponseAllOf struct for ProductDataResponseAllOf
type ProductDataResponseAllOf struct {
	Id       *string `json:"id,omitempty"`
	Object   *string `json:"object,omitempty"`
	ParentId *string `json:"parent_id,omitempty"`
}

// NewProductDataResponseAllOf instantiates a new ProductDataResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDataResponseAllOf() *ProductDataResponseAllOf {
	this := ProductDataResponseAllOf{}
	return &this
}

// NewProductDataResponseAllOfWithDefaults instantiates a new ProductDataResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDataResponseAllOfWithDefaults() *ProductDataResponseAllOf {
	this := ProductDataResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductDataResponseAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductDataResponseAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductDataResponseAllOf) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ProductDataResponseAllOf) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponseAllOf) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ProductDataResponseAllOf) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ProductDataResponseAllOf) SetObject(v string) {
	o.Object = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ProductDataResponseAllOf) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponseAllOf) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ProductDataResponseAllOf) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ProductDataResponseAllOf) SetParentId(v string) {
	o.ParentId = &v
}

func (o ProductDataResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDataResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	return toSerialize, nil
}

type NullableProductDataResponseAllOf struct {
	value *ProductDataResponseAllOf
	isSet bool
}

func (v NullableProductDataResponseAllOf) Get() *ProductDataResponseAllOf {
	return v.value
}

func (v *NullableProductDataResponseAllOf) Set(val *ProductDataResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDataResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDataResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDataResponseAllOf(val *ProductDataResponseAllOf) *NullableProductDataResponseAllOf {
	return &NullableProductDataResponseAllOf{value: val, isSet: true}
}

func (v NullableProductDataResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDataResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

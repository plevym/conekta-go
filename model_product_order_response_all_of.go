package conekta

import (
	"encoding/json"
)

// checks if the ProductOrderResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductOrderResponseAllOf{}

// ProductOrderResponseAllOf struct for ProductOrderResponseAllOf
type ProductOrderResponseAllOf struct {
	Id       *string `json:"id,omitempty"`
	Object   *string `json:"object,omitempty"`
	ParentId *string `json:"parent_id,omitempty"`
}

// NewProductOrderResponseAllOf instantiates a new ProductOrderResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductOrderResponseAllOf() *ProductOrderResponseAllOf {
	this := ProductOrderResponseAllOf{}
	return &this
}

// NewProductOrderResponseAllOfWithDefaults instantiates a new ProductOrderResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductOrderResponseAllOfWithDefaults() *ProductOrderResponseAllOf {
	this := ProductOrderResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductOrderResponseAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductOrderResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductOrderResponseAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductOrderResponseAllOf) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ProductOrderResponseAllOf) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductOrderResponseAllOf) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ProductOrderResponseAllOf) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ProductOrderResponseAllOf) SetObject(v string) {
	o.Object = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ProductOrderResponseAllOf) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductOrderResponseAllOf) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ProductOrderResponseAllOf) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ProductOrderResponseAllOf) SetParentId(v string) {
	o.ParentId = &v
}

func (o ProductOrderResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductOrderResponseAllOf) ToMap() (map[string]interface{}, error) {
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

type NullableProductOrderResponseAllOf struct {
	value *ProductOrderResponseAllOf
	isSet bool
}

func (v NullableProductOrderResponseAllOf) Get() *ProductOrderResponseAllOf {
	return v.value
}

func (v *NullableProductOrderResponseAllOf) Set(val *ProductOrderResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProductOrderResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProductOrderResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductOrderResponseAllOf(val *ProductOrderResponseAllOf) *NullableProductOrderResponseAllOf {
	return &NullableProductOrderResponseAllOf{value: val, isSet: true}
}

func (v NullableProductOrderResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductOrderResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

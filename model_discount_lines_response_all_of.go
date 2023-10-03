package conekta

import (
	"encoding/json"
)

// checks if the DiscountLinesResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountLinesResponseAllOf{}

// DiscountLinesResponseAllOf struct for DiscountLinesResponseAllOf
type DiscountLinesResponseAllOf struct {
	// The discount line id
	Id string `json:"id"`
	// The object name
	Object string `json:"object"`
	// The order id
	ParentId string `json:"parent_id"`
}

// NewDiscountLinesResponseAllOf instantiates a new DiscountLinesResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountLinesResponseAllOf(id string, object string, parentId string) *DiscountLinesResponseAllOf {
	this := DiscountLinesResponseAllOf{}
	this.Id = id
	this.Object = object
	this.ParentId = parentId
	return &this
}

// NewDiscountLinesResponseAllOfWithDefaults instantiates a new DiscountLinesResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountLinesResponseAllOfWithDefaults() *DiscountLinesResponseAllOf {
	this := DiscountLinesResponseAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *DiscountLinesResponseAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscountLinesResponseAllOf) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *DiscountLinesResponseAllOf) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesResponseAllOf) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *DiscountLinesResponseAllOf) SetObject(v string) {
	o.Object = v
}

// GetParentId returns the ParentId field value
func (o *DiscountLinesResponseAllOf) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *DiscountLinesResponseAllOf) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *DiscountLinesResponseAllOf) SetParentId(v string) {
	o.ParentId = v
}

func (o DiscountLinesResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountLinesResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["parent_id"] = o.ParentId
	return toSerialize, nil
}

type NullableDiscountLinesResponseAllOf struct {
	value *DiscountLinesResponseAllOf
	isSet bool
}

func (v NullableDiscountLinesResponseAllOf) Get() *DiscountLinesResponseAllOf {
	return v.value
}

func (v *NullableDiscountLinesResponseAllOf) Set(val *DiscountLinesResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountLinesResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountLinesResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountLinesResponseAllOf(val *DiscountLinesResponseAllOf) *NullableDiscountLinesResponseAllOf {
	return &NullableDiscountLinesResponseAllOf{value: val, isSet: true}
}

func (v NullableDiscountLinesResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountLinesResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

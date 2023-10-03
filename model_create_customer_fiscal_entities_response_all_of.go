package conekta

import (
	"encoding/json"
)

// checks if the CreateCustomerFiscalEntitiesResponseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomerFiscalEntitiesResponseAllOf{}

// CreateCustomerFiscalEntitiesResponseAllOf struct for CreateCustomerFiscalEntitiesResponseAllOf
type CreateCustomerFiscalEntitiesResponseAllOf struct {
	Id        string  `json:"id"`
	Object    string  `json:"object"`
	CreatedAt int64   `json:"created_at"`
	ParentId  *string `json:"parent_id,omitempty"`
	Default   *bool   `json:"default,omitempty"`
}

// NewCreateCustomerFiscalEntitiesResponseAllOf instantiates a new CreateCustomerFiscalEntitiesResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomerFiscalEntitiesResponseAllOf(id string, object string, createdAt int64) *CreateCustomerFiscalEntitiesResponseAllOf {
	this := CreateCustomerFiscalEntitiesResponseAllOf{}
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	return &this
}

// NewCreateCustomerFiscalEntitiesResponseAllOfWithDefaults instantiates a new CreateCustomerFiscalEntitiesResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomerFiscalEntitiesResponseAllOfWithDefaults() *CreateCustomerFiscalEntitiesResponseAllOf {
	this := CreateCustomerFiscalEntitiesResponseAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateCustomerFiscalEntitiesResponseAllOf) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CreateCustomerFiscalEntitiesResponseAllOf) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreateCustomerFiscalEntitiesResponseAllOf) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) SetParentId(v string) {
	o.ParentId = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CreateCustomerFiscalEntitiesResponseAllOf) SetDefault(v bool) {
	o.Default = &v
}

func (o CreateCustomerFiscalEntitiesResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomerFiscalEntitiesResponseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

type NullableCreateCustomerFiscalEntitiesResponseAllOf struct {
	value *CreateCustomerFiscalEntitiesResponseAllOf
	isSet bool
}

func (v NullableCreateCustomerFiscalEntitiesResponseAllOf) Get() *CreateCustomerFiscalEntitiesResponseAllOf {
	return v.value
}

func (v *NullableCreateCustomerFiscalEntitiesResponseAllOf) Set(val *CreateCustomerFiscalEntitiesResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerFiscalEntitiesResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerFiscalEntitiesResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerFiscalEntitiesResponseAllOf(val *CreateCustomerFiscalEntitiesResponseAllOf) *NullableCreateCustomerFiscalEntitiesResponseAllOf {
	return &NullableCreateCustomerFiscalEntitiesResponseAllOf{value: val, isSet: true}
}

func (v NullableCreateCustomerFiscalEntitiesResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerFiscalEntitiesResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

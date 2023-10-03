package conekta

import (
	"encoding/json"
)

// checks if the UpdateOrderTaxResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrderTaxResponse{}

// UpdateOrderTaxResponse create new taxes for an existing order response
type UpdateOrderTaxResponse struct {
	// The amount to be collected for tax in cents
	Amount int64 `json:"amount"`
	// description or tax's name
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Id          string                 `json:"id"`
	Object      *string                `json:"object,omitempty"`
	ParentId    *string                `json:"parent_id,omitempty"`
}

// NewUpdateOrderTaxResponse instantiates a new UpdateOrderTaxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrderTaxResponse(amount int64, description string, id string) *UpdateOrderTaxResponse {
	this := UpdateOrderTaxResponse{}
	this.Amount = amount
	this.Description = description
	this.Id = id
	return &this
}

// NewUpdateOrderTaxResponseWithDefaults instantiates a new UpdateOrderTaxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrderTaxResponseWithDefaults() *UpdateOrderTaxResponse {
	this := UpdateOrderTaxResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *UpdateOrderTaxResponse) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *UpdateOrderTaxResponse) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *UpdateOrderTaxResponse) SetAmount(v int64) {
	o.Amount = v
}

// GetDescription returns the Description field value
func (o *UpdateOrderTaxResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpdateOrderTaxResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpdateOrderTaxResponse) SetDescription(v string) {
	o.Description = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateOrderTaxResponse) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderTaxResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateOrderTaxResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UpdateOrderTaxResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetId returns the Id field value
func (o *UpdateOrderTaxResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateOrderTaxResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateOrderTaxResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *UpdateOrderTaxResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderTaxResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *UpdateOrderTaxResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *UpdateOrderTaxResponse) SetObject(v string) {
	o.Object = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *UpdateOrderTaxResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrderTaxResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *UpdateOrderTaxResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *UpdateOrderTaxResponse) SetParentId(v string) {
	o.ParentId = &v
}

func (o UpdateOrderTaxResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrderTaxResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["description"] = o.Description
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	return toSerialize, nil
}

type NullableUpdateOrderTaxResponse struct {
	value *UpdateOrderTaxResponse
	isSet bool
}

func (v NullableUpdateOrderTaxResponse) Get() *UpdateOrderTaxResponse {
	return v.value
}

func (v *NullableUpdateOrderTaxResponse) Set(val *UpdateOrderTaxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrderTaxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrderTaxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrderTaxResponse(val *UpdateOrderTaxResponse) *NullableUpdateOrderTaxResponse {
	return &NullableUpdateOrderTaxResponse{value: val, isSet: true}
}

func (v NullableUpdateOrderTaxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrderTaxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

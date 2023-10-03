package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodResponse{}

// PaymentMethodResponse struct for PaymentMethodResponse
type PaymentMethodResponse struct {
	Type      string  `json:"type"`
	Id        string  `json:"id"`
	Object    string  `json:"object"`
	CreatedAt int64   `json:"created_at"`
	ParentId  *string `json:"parent_id,omitempty"`
}

// NewPaymentMethodResponse instantiates a new PaymentMethodResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodResponse(type_ string, id string, object string, createdAt int64) *PaymentMethodResponse {
	this := PaymentMethodResponse{}
	this.Type = type_
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	return &this
}

// NewPaymentMethodResponseWithDefaults instantiates a new PaymentMethodResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodResponseWithDefaults() *PaymentMethodResponse {
	this := PaymentMethodResponse{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentMethodResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodResponse) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PaymentMethodResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentMethodResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *PaymentMethodResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentMethodResponse) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentMethodResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentMethodResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *PaymentMethodResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *PaymentMethodResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *PaymentMethodResponse) SetParentId(v string) {
	o.ParentId = &v
}

func (o PaymentMethodResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	return toSerialize, nil
}

type NullablePaymentMethodResponse struct {
	value *PaymentMethodResponse
	isSet bool
}

func (v NullablePaymentMethodResponse) Get() *PaymentMethodResponse {
	return v.value
}

func (v *NullablePaymentMethodResponse) Set(val *PaymentMethodResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodResponse(val *PaymentMethodResponse) *NullablePaymentMethodResponse {
	return &NullablePaymentMethodResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

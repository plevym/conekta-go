package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodSpeiRecurrent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodSpeiRecurrent{}

// PaymentMethodSpeiRecurrent struct for PaymentMethodSpeiRecurrent
type PaymentMethodSpeiRecurrent struct {
	Type      string  `json:"type"`
	Id        string  `json:"id"`
	Object    string  `json:"object"`
	CreatedAt int64   `json:"created_at"`
	ParentId  *string `json:"parent_id,omitempty"`
	Reference *string `json:"reference,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
}

// NewPaymentMethodSpeiRecurrent instantiates a new PaymentMethodSpeiRecurrent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodSpeiRecurrent(type_ string, id string, object string, createdAt int64) *PaymentMethodSpeiRecurrent {
	this := PaymentMethodSpeiRecurrent{}
	this.Type = type_
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	return &this
}

// NewPaymentMethodSpeiRecurrentWithDefaults instantiates a new PaymentMethodSpeiRecurrent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodSpeiRecurrentWithDefaults() *PaymentMethodSpeiRecurrent {
	this := PaymentMethodSpeiRecurrent{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentMethodSpeiRecurrent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodSpeiRecurrent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodSpeiRecurrent) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PaymentMethodSpeiRecurrent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodSpeiRecurrent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentMethodSpeiRecurrent) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *PaymentMethodSpeiRecurrent) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodSpeiRecurrent) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentMethodSpeiRecurrent) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentMethodSpeiRecurrent) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodSpeiRecurrent) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentMethodSpeiRecurrent) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *PaymentMethodSpeiRecurrent) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSpeiRecurrent) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *PaymentMethodSpeiRecurrent) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *PaymentMethodSpeiRecurrent) SetParentId(v string) {
	o.ParentId = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentMethodSpeiRecurrent) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSpeiRecurrent) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentMethodSpeiRecurrent) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentMethodSpeiRecurrent) SetReference(v string) {
	o.Reference = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PaymentMethodSpeiRecurrent) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSpeiRecurrent) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentMethodSpeiRecurrent) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *PaymentMethodSpeiRecurrent) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o PaymentMethodSpeiRecurrent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodSpeiRecurrent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullablePaymentMethodSpeiRecurrent struct {
	value *PaymentMethodSpeiRecurrent
	isSet bool
}

func (v NullablePaymentMethodSpeiRecurrent) Get() *PaymentMethodSpeiRecurrent {
	return v.value
}

func (v *NullablePaymentMethodSpeiRecurrent) Set(val *PaymentMethodSpeiRecurrent) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodSpeiRecurrent) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodSpeiRecurrent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodSpeiRecurrent(val *PaymentMethodSpeiRecurrent) *NullablePaymentMethodSpeiRecurrent {
	return &NullablePaymentMethodSpeiRecurrent{value: val, isSet: true}
}

func (v NullablePaymentMethodSpeiRecurrent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodSpeiRecurrent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

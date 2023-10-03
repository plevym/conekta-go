package conekta

import (
	"encoding/json"
)

// checks if the ChargeResponseChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeResponseChannel{}

// ChargeResponseChannel struct for ChargeResponseChannel
type ChargeResponseChannel struct {
	Segment             *string `json:"segment,omitempty"`
	CheckoutRequestId   *string `json:"checkout_request_id,omitempty"`
	CheckoutRequestType *string `json:"checkout_request_type,omitempty"`
	Id                  *string `json:"id,omitempty"`
}

// NewChargeResponseChannel instantiates a new ChargeResponseChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeResponseChannel() *ChargeResponseChannel {
	this := ChargeResponseChannel{}
	return &this
}

// NewChargeResponseChannelWithDefaults instantiates a new ChargeResponseChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeResponseChannelWithDefaults() *ChargeResponseChannel {
	this := ChargeResponseChannel{}
	return &this
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *ChargeResponseChannel) GetSegment() string {
	if o == nil || IsNil(o.Segment) {
		var ret string
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponseChannel) GetSegmentOk() (*string, bool) {
	if o == nil || IsNil(o.Segment) {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *ChargeResponseChannel) HasSegment() bool {
	if o != nil && !IsNil(o.Segment) {
		return true
	}

	return false
}

// SetSegment gets a reference to the given string and assigns it to the Segment field.
func (o *ChargeResponseChannel) SetSegment(v string) {
	o.Segment = &v
}

// GetCheckoutRequestId returns the CheckoutRequestId field value if set, zero value otherwise.
func (o *ChargeResponseChannel) GetCheckoutRequestId() string {
	if o == nil || IsNil(o.CheckoutRequestId) {
		var ret string
		return ret
	}
	return *o.CheckoutRequestId
}

// GetCheckoutRequestIdOk returns a tuple with the CheckoutRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponseChannel) GetCheckoutRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.CheckoutRequestId) {
		return nil, false
	}
	return o.CheckoutRequestId, true
}

// HasCheckoutRequestId returns a boolean if a field has been set.
func (o *ChargeResponseChannel) HasCheckoutRequestId() bool {
	if o != nil && !IsNil(o.CheckoutRequestId) {
		return true
	}

	return false
}

// SetCheckoutRequestId gets a reference to the given string and assigns it to the CheckoutRequestId field.
func (o *ChargeResponseChannel) SetCheckoutRequestId(v string) {
	o.CheckoutRequestId = &v
}

// GetCheckoutRequestType returns the CheckoutRequestType field value if set, zero value otherwise.
func (o *ChargeResponseChannel) GetCheckoutRequestType() string {
	if o == nil || IsNil(o.CheckoutRequestType) {
		var ret string
		return ret
	}
	return *o.CheckoutRequestType
}

// GetCheckoutRequestTypeOk returns a tuple with the CheckoutRequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponseChannel) GetCheckoutRequestTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CheckoutRequestType) {
		return nil, false
	}
	return o.CheckoutRequestType, true
}

// HasCheckoutRequestType returns a boolean if a field has been set.
func (o *ChargeResponseChannel) HasCheckoutRequestType() bool {
	if o != nil && !IsNil(o.CheckoutRequestType) {
		return true
	}

	return false
}

// SetCheckoutRequestType gets a reference to the given string and assigns it to the CheckoutRequestType field.
func (o *ChargeResponseChannel) SetCheckoutRequestType(v string) {
	o.CheckoutRequestType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChargeResponseChannel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeResponseChannel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChargeResponseChannel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChargeResponseChannel) SetId(v string) {
	o.Id = &v
}

func (o ChargeResponseChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeResponseChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Segment) {
		toSerialize["segment"] = o.Segment
	}
	if !IsNil(o.CheckoutRequestId) {
		toSerialize["checkout_request_id"] = o.CheckoutRequestId
	}
	if !IsNil(o.CheckoutRequestType) {
		toSerialize["checkout_request_type"] = o.CheckoutRequestType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableChargeResponseChannel struct {
	value *ChargeResponseChannel
	isSet bool
}

func (v NullableChargeResponseChannel) Get() *ChargeResponseChannel {
	return v.value
}

func (v *NullableChargeResponseChannel) Set(val *ChargeResponseChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeResponseChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeResponseChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeResponseChannel(val *ChargeResponseChannel) *NullableChargeResponseChannel {
	return &NullableChargeResponseChannel{value: val, isSet: true}
}

func (v NullableChargeResponseChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeResponseChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

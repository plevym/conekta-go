package conekta

import (
	"encoding/json"
)

// checks if the ShippingOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingOrderResponse{}

// ShippingOrderResponse struct for ShippingOrderResponse
type ShippingOrderResponse struct {
	// Shipping amount in cents
	Amount int64 `json:"amount"`
	// Carrier name for the shipment
	Carrier *string `json:"carrier,omitempty"`
	// Tracking number can be used to track the shipment
	TrackingNumber *string `json:"tracking_number,omitempty"`
	// Method of shipment
	Method *string `json:"method,omitempty"`
	// Hash where the user can send additional information for each 'shipping'.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Id       *string                `json:"id,omitempty"`
	Object   *string                `json:"object,omitempty"`
	ParentId *string                `json:"parent_id,omitempty"`
}

// NewShippingOrderResponse instantiates a new ShippingOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingOrderResponse(amount int64) *ShippingOrderResponse {
	this := ShippingOrderResponse{}
	this.Amount = amount
	return &this
}

// NewShippingOrderResponseWithDefaults instantiates a new ShippingOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingOrderResponseWithDefaults() *ShippingOrderResponse {
	this := ShippingOrderResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ShippingOrderResponse) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ShippingOrderResponse) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ShippingOrderResponse) SetAmount(v int64) {
	o.Amount = v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *ShippingOrderResponse) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOrderResponse) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *ShippingOrderResponse) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *ShippingOrderResponse) SetCarrier(v string) {
	o.Carrier = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *ShippingOrderResponse) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOrderResponse) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *ShippingOrderResponse) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *ShippingOrderResponse) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ShippingOrderResponse) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOrderResponse) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ShippingOrderResponse) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ShippingOrderResponse) SetMethod(v string) {
	o.Method = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShippingOrderResponse) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOrderResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShippingOrderResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ShippingOrderResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShippingOrderResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOrderResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShippingOrderResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ShippingOrderResponse) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ShippingOrderResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOrderResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ShippingOrderResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ShippingOrderResponse) SetObject(v string) {
	o.Object = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ShippingOrderResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOrderResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ShippingOrderResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ShippingOrderResponse) SetParentId(v string) {
	o.ParentId = &v
}

func (o ShippingOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.TrackingNumber) {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
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

type NullableShippingOrderResponse struct {
	value *ShippingOrderResponse
	isSet bool
}

func (v NullableShippingOrderResponse) Get() *ShippingOrderResponse {
	return v.value
}

func (v *NullableShippingOrderResponse) Set(val *ShippingOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingOrderResponse(val *ShippingOrderResponse) *NullableShippingOrderResponse {
	return &NullableShippingOrderResponse{value: val, isSet: true}
}

func (v NullableShippingOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

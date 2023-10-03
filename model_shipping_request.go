package conekta

import (
	"encoding/json"
)

// checks if the ShippingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingRequest{}

// ShippingRequest struct for ShippingRequest
type ShippingRequest struct {
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
}

// NewShippingRequest instantiates a new ShippingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingRequest(amount int64) *ShippingRequest {
	this := ShippingRequest{}
	this.Amount = amount
	return &this
}

// NewShippingRequestWithDefaults instantiates a new ShippingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingRequestWithDefaults() *ShippingRequest {
	this := ShippingRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ShippingRequest) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ShippingRequest) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ShippingRequest) SetAmount(v int64) {
	o.Amount = v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *ShippingRequest) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingRequest) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *ShippingRequest) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *ShippingRequest) SetCarrier(v string) {
	o.Carrier = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *ShippingRequest) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingRequest) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *ShippingRequest) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *ShippingRequest) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ShippingRequest) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingRequest) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ShippingRequest) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ShippingRequest) SetMethod(v string) {
	o.Method = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShippingRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShippingRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ShippingRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ShippingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingRequest) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableShippingRequest struct {
	value *ShippingRequest
	isSet bool
}

func (v NullableShippingRequest) Get() *ShippingRequest {
	return v.value
}

func (v *NullableShippingRequest) Set(val *ShippingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingRequest(val *ShippingRequest) *NullableShippingRequest {
	return &NullableShippingRequest{value: val, isSet: true}
}

func (v NullableShippingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

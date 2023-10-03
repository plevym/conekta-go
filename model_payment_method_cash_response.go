package conekta

import (
	"encoding/json"
)

// checks if the PaymentMethodCashResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCashResponse{}

// PaymentMethodCashResponse struct for PaymentMethodCashResponse
type PaymentMethodCashResponse struct {
	Type       string  `json:"type"`
	Id         string  `json:"id"`
	Object     string  `json:"object"`
	CreatedAt  int64   `json:"created_at"`
	ParentId   *string `json:"parent_id,omitempty"`
	Reference  *string `json:"reference,omitempty"`
	Barcode    *string `json:"barcode,omitempty"`
	BarcodeUrl *string `json:"barcode_url,omitempty"`
	ExpiresAt  *int64  `json:"expires_at,omitempty"`
	Provider   *string `json:"provider,omitempty"`
}

// NewPaymentMethodCashResponse instantiates a new PaymentMethodCashResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCashResponse(type_ string, id string, object string, createdAt int64) *PaymentMethodCashResponse {
	this := PaymentMethodCashResponse{}
	this.Type = type_
	this.Id = id
	this.Object = object
	this.CreatedAt = createdAt
	return &this
}

// NewPaymentMethodCashResponseWithDefaults instantiates a new PaymentMethodCashResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCashResponseWithDefaults() *PaymentMethodCashResponse {
	this := PaymentMethodCashResponse{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentMethodCashResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodCashResponse) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PaymentMethodCashResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentMethodCashResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *PaymentMethodCashResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentMethodCashResponse) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentMethodCashResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentMethodCashResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *PaymentMethodCashResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *PaymentMethodCashResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *PaymentMethodCashResponse) SetParentId(v string) {
	o.ParentId = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentMethodCashResponse) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentMethodCashResponse) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentMethodCashResponse) SetReference(v string) {
	o.Reference = &v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *PaymentMethodCashResponse) GetBarcode() string {
	if o == nil || IsNil(o.Barcode) {
		var ret string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetBarcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Barcode) {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *PaymentMethodCashResponse) HasBarcode() bool {
	if o != nil && !IsNil(o.Barcode) {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given string and assigns it to the Barcode field.
func (o *PaymentMethodCashResponse) SetBarcode(v string) {
	o.Barcode = &v
}

// GetBarcodeUrl returns the BarcodeUrl field value if set, zero value otherwise.
func (o *PaymentMethodCashResponse) GetBarcodeUrl() string {
	if o == nil || IsNil(o.BarcodeUrl) {
		var ret string
		return ret
	}
	return *o.BarcodeUrl
}

// GetBarcodeUrlOk returns a tuple with the BarcodeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetBarcodeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BarcodeUrl) {
		return nil, false
	}
	return o.BarcodeUrl, true
}

// HasBarcodeUrl returns a boolean if a field has been set.
func (o *PaymentMethodCashResponse) HasBarcodeUrl() bool {
	if o != nil && !IsNil(o.BarcodeUrl) {
		return true
	}

	return false
}

// SetBarcodeUrl gets a reference to the given string and assigns it to the BarcodeUrl field.
func (o *PaymentMethodCashResponse) SetBarcodeUrl(v string) {
	o.BarcodeUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PaymentMethodCashResponse) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentMethodCashResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *PaymentMethodCashResponse) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PaymentMethodCashResponse) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCashResponse) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PaymentMethodCashResponse) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *PaymentMethodCashResponse) SetProvider(v string) {
	o.Provider = &v
}

func (o PaymentMethodCashResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCashResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Barcode) {
		toSerialize["barcode"] = o.Barcode
	}
	if !IsNil(o.BarcodeUrl) {
		toSerialize["barcode_url"] = o.BarcodeUrl
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullablePaymentMethodCashResponse struct {
	value *PaymentMethodCashResponse
	isSet bool
}

func (v NullablePaymentMethodCashResponse) Get() *PaymentMethodCashResponse {
	return v.value
}

func (v *NullablePaymentMethodCashResponse) Set(val *PaymentMethodCashResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCashResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCashResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCashResponse(val *PaymentMethodCashResponse) *NullablePaymentMethodCashResponse {
	return &NullablePaymentMethodCashResponse{value: val, isSet: true}
}

func (v NullablePaymentMethodCashResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCashResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

package conekta

import (
	"encoding/json"
)

// checks if the ProductDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDataResponse{}

// ProductDataResponse struct for ProductDataResponse
type ProductDataResponse struct {
	AntifraudInfo map[string]interface{} `json:"antifraud_info,omitempty"`
	// The brand of the item.
	Brand *string `json:"brand,omitempty"`
	// Short description of the item
	Description *string `json:"description,omitempty"`
	// It is a key/value hash that can hold custom fields. Maximum 100 elements and allows special characters.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The name of the item. It will be displayed in the order.
	Name string `json:"name"`
	// The quantity of the item in the order.
	Quantity int32 `json:"quantity"`
	// The stock keeping unit for the item. It is used to identify the item in the order.
	Sku *string `json:"sku,omitempty"`
	// List of tags for the item. It is used to identify the item in the order.
	Tags []string `json:"tags,omitempty"`
	// The price of the item in cents.
	UnitPrice int32   `json:"unit_price"`
	Id        *string `json:"id,omitempty"`
	Object    *string `json:"object,omitempty"`
	ParentId  *string `json:"parent_id,omitempty"`
}

// NewProductDataResponse instantiates a new ProductDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDataResponse(name string, quantity int32, unitPrice int32) *ProductDataResponse {
	this := ProductDataResponse{}
	this.Name = name
	this.Quantity = quantity
	this.UnitPrice = unitPrice
	return &this
}

// NewProductDataResponseWithDefaults instantiates a new ProductDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDataResponseWithDefaults() *ProductDataResponse {
	this := ProductDataResponse{}
	return &this
}

// GetAntifraudInfo returns the AntifraudInfo field value if set, zero value otherwise.
func (o *ProductDataResponse) GetAntifraudInfo() map[string]interface{} {
	if o == nil || IsNil(o.AntifraudInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.AntifraudInfo
}

// GetAntifraudInfoOk returns a tuple with the AntifraudInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetAntifraudInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AntifraudInfo) {
		return map[string]interface{}{}, false
	}
	return o.AntifraudInfo, true
}

// HasAntifraudInfo returns a boolean if a field has been set.
func (o *ProductDataResponse) HasAntifraudInfo() bool {
	if o != nil && !IsNil(o.AntifraudInfo) {
		return true
	}

	return false
}

// SetAntifraudInfo gets a reference to the given map[string]interface{} and assigns it to the AntifraudInfo field.
func (o *ProductDataResponse) SetAntifraudInfo(v map[string]interface{}) {
	o.AntifraudInfo = v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *ProductDataResponse) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *ProductDataResponse) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *ProductDataResponse) SetBrand(v string) {
	o.Brand = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductDataResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductDataResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductDataResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ProductDataResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ProductDataResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ProductDataResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *ProductDataResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductDataResponse) SetName(v string) {
	o.Name = v
}

// GetQuantity returns the Quantity field value
func (o *ProductDataResponse) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ProductDataResponse) SetQuantity(v int32) {
	o.Quantity = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductDataResponse) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductDataResponse) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductDataResponse) SetSku(v string) {
	o.Sku = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProductDataResponse) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProductDataResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ProductDataResponse) SetTags(v []string) {
	o.Tags = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *ProductDataResponse) GetUnitPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetUnitPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *ProductDataResponse) SetUnitPrice(v int32) {
	o.UnitPrice = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductDataResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductDataResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductDataResponse) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ProductDataResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ProductDataResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ProductDataResponse) SetObject(v string) {
	o.Object = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ProductDataResponse) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataResponse) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ProductDataResponse) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ProductDataResponse) SetParentId(v string) {
	o.ParentId = &v
}

func (o ProductDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AntifraudInfo) {
		toSerialize["antifraud_info"] = o.AntifraudInfo
	}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["unit_price"] = o.UnitPrice
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

type NullableProductDataResponse struct {
	value *ProductDataResponse
	isSet bool
}

func (v NullableProductDataResponse) Get() *ProductDataResponse {
	return v.value
}

func (v *NullableProductDataResponse) Set(val *ProductDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDataResponse(val *ProductDataResponse) *NullableProductDataResponse {
	return &NullableProductDataResponse{value: val, isSet: true}
}

func (v NullableProductDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

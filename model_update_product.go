package conekta

import (
	"encoding/json"
)

// checks if the UpdateProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProduct{}

// UpdateProduct struct for UpdateProduct
type UpdateProduct struct {
	AntifraudInfo map[string]map[string]interface{} `json:"antifraud_info,omitempty"`
	Description   *string                           `json:"description,omitempty"`
	Sku           *string                           `json:"sku,omitempty"`
	Name          *string                           `json:"name,omitempty"`
	UnitPrice     *int64                            `json:"unit_price,omitempty"`
	Quantity      *int32                            `json:"quantity,omitempty"`
	Tags          []string                          `json:"tags,omitempty"`
	Brand         *string                           `json:"brand,omitempty"`
	Metadata      *map[string]string                `json:"metadata,omitempty"`
}

// NewUpdateProduct instantiates a new UpdateProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProduct() *UpdateProduct {
	this := UpdateProduct{}
	return &this
}

// NewUpdateProductWithDefaults instantiates a new UpdateProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductWithDefaults() *UpdateProduct {
	this := UpdateProduct{}
	return &this
}

// GetAntifraudInfo returns the AntifraudInfo field value if set, zero value otherwise.
func (o *UpdateProduct) GetAntifraudInfo() map[string]map[string]interface{} {
	if o == nil || IsNil(o.AntifraudInfo) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.AntifraudInfo
}

// GetAntifraudInfoOk returns a tuple with the AntifraudInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProduct) GetAntifraudInfoOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.AntifraudInfo) {
		return map[string]map[string]interface{}{}, false
	}
	return o.AntifraudInfo, true
}

// HasAntifraudInfo returns a boolean if a field has been set.
func (o *UpdateProduct) HasAntifraudInfo() bool {
	if o != nil && !IsNil(o.AntifraudInfo) {
		return true
	}

	return false
}

// SetAntifraudInfo gets a reference to the given map[string]map[string]interface{} and assigns it to the AntifraudInfo field.
func (o *UpdateProduct) SetAntifraudInfo(v map[string]map[string]interface{}) {
	o.AntifraudInfo = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateProduct) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProduct) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateProduct) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateProduct) SetDescription(v string) {
	o.Description = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *UpdateProduct) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProduct) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *UpdateProduct) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *UpdateProduct) SetSku(v string) {
	o.Sku = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateProduct) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProduct) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateProduct) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateProduct) SetName(v string) {
	o.Name = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *UpdateProduct) GetUnitPrice() int64 {
	if o == nil || IsNil(o.UnitPrice) {
		var ret int64
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProduct) GetUnitPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *UpdateProduct) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given int64 and assigns it to the UnitPrice field.
func (o *UpdateProduct) SetUnitPrice(v int64) {
	o.UnitPrice = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UpdateProduct) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProduct) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UpdateProduct) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *UpdateProduct) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateProduct) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProduct) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateProduct) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateProduct) SetTags(v []string) {
	o.Tags = v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *UpdateProduct) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProduct) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *UpdateProduct) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *UpdateProduct) SetBrand(v string) {
	o.Brand = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateProduct) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProduct) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateProduct) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UpdateProduct) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o UpdateProduct) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AntifraudInfo) {
		toSerialize["antifraud_info"] = o.AntifraudInfo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unit_price"] = o.UnitPrice
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableUpdateProduct struct {
	value *UpdateProduct
	isSet bool
}

func (v NullableUpdateProduct) Get() *UpdateProduct {
	return v.value
}

func (v *NullableUpdateProduct) Set(val *UpdateProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProduct(val *UpdateProduct) *NullableUpdateProduct {
	return &NullableUpdateProduct{value: val, isSet: true}
}

func (v NullableUpdateProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

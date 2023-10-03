package conekta

import (
	"encoding/json"
)

// checks if the Product type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Product{}

// Product struct for Product
type Product struct {
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
	UnitPrice int32 `json:"unit_price"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(name string, quantity int32, unitPrice int32) *Product {
	this := Product{}
	this.Name = name
	this.Quantity = quantity
	this.UnitPrice = unitPrice
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetAntifraudInfo returns the AntifraudInfo field value if set, zero value otherwise.
func (o *Product) GetAntifraudInfo() map[string]interface{} {
	if o == nil || IsNil(o.AntifraudInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.AntifraudInfo
}

// GetAntifraudInfoOk returns a tuple with the AntifraudInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetAntifraudInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AntifraudInfo) {
		return map[string]interface{}{}, false
	}
	return o.AntifraudInfo, true
}

// HasAntifraudInfo returns a boolean if a field has been set.
func (o *Product) HasAntifraudInfo() bool {
	if o != nil && !IsNil(o.AntifraudInfo) {
		return true
	}

	return false
}

// SetAntifraudInfo gets a reference to the given map[string]interface{} and assigns it to the AntifraudInfo field.
func (o *Product) SetAntifraudInfo(v map[string]interface{}) {
	o.AntifraudInfo = v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *Product) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *Product) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *Product) SetBrand(v string) {
	o.Brand = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Product) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Product) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Product) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Product) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Product) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Product) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *Product) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Product) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Product) SetName(v string) {
	o.Name = v
}

// GetQuantity returns the Quantity field value
func (o *Product) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *Product) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *Product) SetQuantity(v int32) {
	o.Quantity = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *Product) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *Product) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *Product) SetSku(v string) {
	o.Sku = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Product) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Product) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Product) SetTags(v []string) {
	o.Tags = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *Product) GetUnitPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *Product) GetUnitPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *Product) SetUnitPrice(v int32) {
	o.UnitPrice = v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Product) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

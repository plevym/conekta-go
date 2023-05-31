/*
Conekta API

Conekta sdk

API version: 2.1.0
Contact: engineering@conekta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conekta

import (
	"encoding/json"
)

// checks if the ApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyResponse{}

// ApiKeyResponse api keys model
type ApiKeyResponse struct {
	// Indicates if the api key is active
	Active *bool `json:"active,omitempty"`
	// Unix timestamp in seconds with the creation date of the api key
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Detail of the use that will be given to the api key
	Description *string `json:"description,omitempty"`
	// Unique identifier of the api key
	Id *string `json:"id,omitempty"`
	// Indicates if the api key is in live mode
	Livemode *bool `json:"livemode,omitempty"`
	// Object name, value is api_key
	Object *string `json:"object,omitempty"`
	// The first few characters of the authentication_token
	Prefix *string `json:"prefix,omitempty"`
	// Indicates the user account private=owner or public=public
	Role *string `json:"role,omitempty"`
}

// NewApiKeyResponse instantiates a new ApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyResponse() *ApiKeyResponse {
	this := ApiKeyResponse{}
	return &this
}

// NewApiKeyResponseWithDefaults instantiates a new ApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyResponseWithDefaults() *ApiKeyResponse {
	this := ApiKeyResponse{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApiKeyResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApiKeyResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApiKeyResponse) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiKeyResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiKeyResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ApiKeyResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiKeyResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiKeyResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiKeyResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiKeyResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiKeyResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiKeyResponse) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *ApiKeyResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *ApiKeyResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *ApiKeyResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ApiKeyResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ApiKeyResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ApiKeyResponse) SetObject(v string) {
	o.Object = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ApiKeyResponse) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponse) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ApiKeyResponse) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ApiKeyResponse) SetPrefix(v string) {
	o.Prefix = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ApiKeyResponse) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponse) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ApiKeyResponse) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ApiKeyResponse) SetRole(v string) {
	o.Role = &v
}

func (o ApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableApiKeyResponse struct {
	value *ApiKeyResponse
	isSet bool
}

func (v NullableApiKeyResponse) Get() *ApiKeyResponse {
	return v.value
}

func (v *NullableApiKeyResponse) Set(val *ApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyResponse(val *ApiKeyResponse) *NullableApiKeyResponse {
	return &NullableApiKeyResponse{value: val, isSet: true}
}

func (v NullableApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


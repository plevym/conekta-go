package conekta

import (
	"encoding/json"
)

// checks if the ApiKeyResponseOnDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyResponseOnDelete{}

// ApiKeyResponseOnDelete api keys model
type ApiKeyResponseOnDelete struct {
	// Indicates if the api key is active
	Active *bool `json:"active,omitempty"`
	// Unix timestamp in seconds of when the api key was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// A name or brief explanation of what this api key is used for
	Description *string `json:"description,omitempty"`
	// Indicates if the api key is in production
	Livemode *bool `json:"livemode,omitempty"`
	// The first few characters of the authentication_token
	Prefix *string `json:"prefix,omitempty"`
	// Unique identifier of the api key
	Id *string `json:"id,omitempty"`
	// Object name, value is 'api_key'
	Object *string `json:"object,omitempty"`
	// Indicates if the api key was deleted
	Deleted *bool `json:"deleted,omitempty"`
	// Indicates if the api key is private or public
	Role *string `json:"role,omitempty"`
}

// NewApiKeyResponseOnDelete instantiates a new ApiKeyResponseOnDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyResponseOnDelete() *ApiKeyResponseOnDelete {
	this := ApiKeyResponseOnDelete{}
	return &this
}

// NewApiKeyResponseOnDeleteWithDefaults instantiates a new ApiKeyResponseOnDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyResponseOnDeleteWithDefaults() *ApiKeyResponseOnDelete {
	this := ApiKeyResponseOnDelete{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApiKeyResponseOnDelete) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseOnDelete) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApiKeyResponseOnDelete) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApiKeyResponseOnDelete) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiKeyResponseOnDelete) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseOnDelete) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiKeyResponseOnDelete) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ApiKeyResponseOnDelete) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiKeyResponseOnDelete) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseOnDelete) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiKeyResponseOnDelete) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiKeyResponseOnDelete) SetDescription(v string) {
	o.Description = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *ApiKeyResponseOnDelete) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseOnDelete) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *ApiKeyResponseOnDelete) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *ApiKeyResponseOnDelete) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ApiKeyResponseOnDelete) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseOnDelete) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ApiKeyResponseOnDelete) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ApiKeyResponseOnDelete) SetPrefix(v string) {
	o.Prefix = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiKeyResponseOnDelete) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseOnDelete) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiKeyResponseOnDelete) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiKeyResponseOnDelete) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ApiKeyResponseOnDelete) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseOnDelete) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ApiKeyResponseOnDelete) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ApiKeyResponseOnDelete) SetObject(v string) {
	o.Object = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *ApiKeyResponseOnDelete) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseOnDelete) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *ApiKeyResponseOnDelete) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *ApiKeyResponseOnDelete) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ApiKeyResponseOnDelete) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyResponseOnDelete) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ApiKeyResponseOnDelete) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ApiKeyResponseOnDelete) SetRole(v string) {
	o.Role = &v
}

func (o ApiKeyResponseOnDelete) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyResponseOnDelete) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableApiKeyResponseOnDelete struct {
	value *ApiKeyResponseOnDelete
	isSet bool
}

func (v NullableApiKeyResponseOnDelete) Get() *ApiKeyResponseOnDelete {
	return v.value
}

func (v *NullableApiKeyResponseOnDelete) Set(val *ApiKeyResponseOnDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyResponseOnDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyResponseOnDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyResponseOnDelete(val *ApiKeyResponseOnDelete) *NullableApiKeyResponseOnDelete {
	return &NullableApiKeyResponseOnDelete{value: val, isSet: true}
}

func (v NullableApiKeyResponseOnDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyResponseOnDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

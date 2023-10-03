package conekta

import (
	"encoding/json"
)

// checks if the ApiKeyCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyCreateResponse{}

// ApiKeyCreateResponse struct for ApiKeyCreateResponse
type ApiKeyCreateResponse struct {
	// It is occupied as a user when authenticated with basic authentication, with a blank password. This value will only appear once, in the request to create a new key. Copy and save it in a safe place.
	AuthenticationToken *string `json:"authentication_token,omitempty"`
	// Indicates if the api key is active
	Active *bool `json:"active,omitempty"`
	// Unix timestamp in seconds of when the api key was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Unix timestamp in seconds of when the api key was last updated
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// Unix timestamp in seconds of when the api key was deleted
	DeactivatedAt NullableInt64 `json:"deactivated_at,omitempty"`
	// A name or brief explanation of what this api key is used for
	Description *string `json:"description,omitempty"`
	// Unique identifier of the api key
	Id *string `json:"id,omitempty"`
	// Indicates if the api key is in production
	Livemode *bool `json:"livemode,omitempty"`
	// Indicates if the api key was deleted
	Deleted *bool `json:"deleted,omitempty"`
	// Object name, value is 'api_key'
	Object *string `json:"object,omitempty"`
	// The first few characters of the authentication_token
	Prefix *string `json:"prefix,omitempty"`
	// Indicates if the api key is private or public
	Role *string `json:"role,omitempty"`
}

// NewApiKeyCreateResponse instantiates a new ApiKeyCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyCreateResponse() *ApiKeyCreateResponse {
	this := ApiKeyCreateResponse{}
	return &this
}

// NewApiKeyCreateResponseWithDefaults instantiates a new ApiKeyCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyCreateResponseWithDefaults() *ApiKeyCreateResponse {
	this := ApiKeyCreateResponse{}
	return &this
}

// GetAuthenticationToken returns the AuthenticationToken field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetAuthenticationToken() string {
	if o == nil || IsNil(o.AuthenticationToken) {
		var ret string
		return ret
	}
	return *o.AuthenticationToken
}

// GetAuthenticationTokenOk returns a tuple with the AuthenticationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetAuthenticationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationToken) {
		return nil, false
	}
	return o.AuthenticationToken, true
}

// HasAuthenticationToken returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasAuthenticationToken() bool {
	if o != nil && !IsNil(o.AuthenticationToken) {
		return true
	}

	return false
}

// SetAuthenticationToken gets a reference to the given string and assigns it to the AuthenticationToken field.
func (o *ApiKeyCreateResponse) SetAuthenticationToken(v string) {
	o.AuthenticationToken = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApiKeyCreateResponse) SetActive(v bool) {
	o.Active = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ApiKeyCreateResponse) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *ApiKeyCreateResponse) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetDeactivatedAt returns the DeactivatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyCreateResponse) GetDeactivatedAt() int64 {
	if o == nil || IsNil(o.DeactivatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.DeactivatedAt.Get()
}

// GetDeactivatedAtOk returns a tuple with the DeactivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyCreateResponse) GetDeactivatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeactivatedAt.Get(), o.DeactivatedAt.IsSet()
}

// HasDeactivatedAt returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasDeactivatedAt() bool {
	if o != nil && o.DeactivatedAt.IsSet() {
		return true
	}

	return false
}

// SetDeactivatedAt gets a reference to the given NullableInt64 and assigns it to the DeactivatedAt field.
func (o *ApiKeyCreateResponse) SetDeactivatedAt(v int64) {
	o.DeactivatedAt.Set(&v)
}

// SetDeactivatedAtNil sets the value for DeactivatedAt to be an explicit nil
func (o *ApiKeyCreateResponse) SetDeactivatedAtNil() {
	o.DeactivatedAt.Set(nil)
}

// UnsetDeactivatedAt ensures that no value is present for DeactivatedAt, not even an explicit nil
func (o *ApiKeyCreateResponse) UnsetDeactivatedAt() {
	o.DeactivatedAt.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiKeyCreateResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiKeyCreateResponse) SetId(v string) {
	o.Id = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *ApiKeyCreateResponse) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *ApiKeyCreateResponse) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ApiKeyCreateResponse) SetObject(v string) {
	o.Object = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ApiKeyCreateResponse) SetPrefix(v string) {
	o.Prefix = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ApiKeyCreateResponse) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyCreateResponse) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ApiKeyCreateResponse) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ApiKeyCreateResponse) SetRole(v string) {
	o.Role = &v
}

func (o ApiKeyCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationToken) {
		toSerialize["authentication_token"] = o.AuthenticationToken
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.DeactivatedAt.IsSet() {
		toSerialize["deactivated_at"] = o.DeactivatedAt.Get()
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
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
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

type NullableApiKeyCreateResponse struct {
	value *ApiKeyCreateResponse
	isSet bool
}

func (v NullableApiKeyCreateResponse) Get() *ApiKeyCreateResponse {
	return v.value
}

func (v *NullableApiKeyCreateResponse) Set(val *ApiKeyCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyCreateResponse(val *ApiKeyCreateResponse) *NullableApiKeyCreateResponse {
	return &NullableApiKeyCreateResponse{value: val, isSet: true}
}

func (v NullableApiKeyCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

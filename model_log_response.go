package conekta

import (
	"encoding/json"
)

// checks if the LogResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogResponse{}

// LogResponse log model
type LogResponse struct {
	CreatedAt       int64                  `json:"created_at"`
	Id              string                 `json:"id"`
	IpAddress       *string                `json:"ip_address,omitempty"`
	Livemode        bool                   `json:"livemode"`
	LoggableId      NullableString         `json:"loggable_id,omitempty"`
	LoggableType    NullableString         `json:"loggable_type,omitempty"`
	Method          *string                `json:"method,omitempty"`
	OauthTokenId    NullableString         `json:"oauth_token_id,omitempty"`
	QueryString     map[string]interface{} `json:"query_string,omitempty"`
	Related         *string                `json:"related,omitempty"`
	RequestBody     map[string]interface{} `json:"request_body,omitempty"`
	RequestHeaders  *map[string]string     `json:"request_headers,omitempty"`
	ResponseBody    map[string]interface{} `json:"response_body,omitempty"`
	ResponseHeaders *map[string]string     `json:"response_headers,omitempty"`
	SearchableTags  []string               `json:"searchable_tags,omitempty"`
	Status          *string                `json:"status,omitempty"`
	UpdatedAt       *string                `json:"updated_at,omitempty"`
	Url             *string                `json:"url,omitempty"`
	UserAccountId   *string                `json:"user_account_id,omitempty"`
	Version         *string                `json:"version,omitempty"`
}

// NewLogResponse instantiates a new LogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogResponse(createdAt int64, id string, livemode bool) *LogResponse {
	this := LogResponse{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Livemode = livemode
	return &this
}

// NewLogResponseWithDefaults instantiates a new LogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogResponseWithDefaults() *LogResponse {
	this := LogResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *LogResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LogResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LogResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *LogResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogResponse) SetId(v string) {
	o.Id = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *LogResponse) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *LogResponse) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *LogResponse) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLivemode returns the Livemode field value
func (o *LogResponse) GetLivemode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value
// and a boolean to check if the value has been set.
func (o *LogResponse) GetLivemodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Livemode, true
}

// SetLivemode sets field value
func (o *LogResponse) SetLivemode(v bool) {
	o.Livemode = v
}

// GetLoggableId returns the LoggableId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResponse) GetLoggableId() string {
	if o == nil || IsNil(o.LoggableId.Get()) {
		var ret string
		return ret
	}
	return *o.LoggableId.Get()
}

// GetLoggableIdOk returns a tuple with the LoggableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResponse) GetLoggableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoggableId.Get(), o.LoggableId.IsSet()
}

// HasLoggableId returns a boolean if a field has been set.
func (o *LogResponse) HasLoggableId() bool {
	if o != nil && o.LoggableId.IsSet() {
		return true
	}

	return false
}

// SetLoggableId gets a reference to the given NullableString and assigns it to the LoggableId field.
func (o *LogResponse) SetLoggableId(v string) {
	o.LoggableId.Set(&v)
}

// SetLoggableIdNil sets the value for LoggableId to be an explicit nil
func (o *LogResponse) SetLoggableIdNil() {
	o.LoggableId.Set(nil)
}

// UnsetLoggableId ensures that no value is present for LoggableId, not even an explicit nil
func (o *LogResponse) UnsetLoggableId() {
	o.LoggableId.Unset()
}

// GetLoggableType returns the LoggableType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResponse) GetLoggableType() string {
	if o == nil || IsNil(o.LoggableType.Get()) {
		var ret string
		return ret
	}
	return *o.LoggableType.Get()
}

// GetLoggableTypeOk returns a tuple with the LoggableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResponse) GetLoggableTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoggableType.Get(), o.LoggableType.IsSet()
}

// HasLoggableType returns a boolean if a field has been set.
func (o *LogResponse) HasLoggableType() bool {
	if o != nil && o.LoggableType.IsSet() {
		return true
	}

	return false
}

// SetLoggableType gets a reference to the given NullableString and assigns it to the LoggableType field.
func (o *LogResponse) SetLoggableType(v string) {
	o.LoggableType.Set(&v)
}

// SetLoggableTypeNil sets the value for LoggableType to be an explicit nil
func (o *LogResponse) SetLoggableTypeNil() {
	o.LoggableType.Set(nil)
}

// UnsetLoggableType ensures that no value is present for LoggableType, not even an explicit nil
func (o *LogResponse) UnsetLoggableType() {
	o.LoggableType.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LogResponse) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LogResponse) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LogResponse) SetMethod(v string) {
	o.Method = &v
}

// GetOauthTokenId returns the OauthTokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResponse) GetOauthTokenId() string {
	if o == nil || IsNil(o.OauthTokenId.Get()) {
		var ret string
		return ret
	}
	return *o.OauthTokenId.Get()
}

// GetOauthTokenIdOk returns a tuple with the OauthTokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResponse) GetOauthTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OauthTokenId.Get(), o.OauthTokenId.IsSet()
}

// HasOauthTokenId returns a boolean if a field has been set.
func (o *LogResponse) HasOauthTokenId() bool {
	if o != nil && o.OauthTokenId.IsSet() {
		return true
	}

	return false
}

// SetOauthTokenId gets a reference to the given NullableString and assigns it to the OauthTokenId field.
func (o *LogResponse) SetOauthTokenId(v string) {
	o.OauthTokenId.Set(&v)
}

// SetOauthTokenIdNil sets the value for OauthTokenId to be an explicit nil
func (o *LogResponse) SetOauthTokenIdNil() {
	o.OauthTokenId.Set(nil)
}

// UnsetOauthTokenId ensures that no value is present for OauthTokenId, not even an explicit nil
func (o *LogResponse) UnsetOauthTokenId() {
	o.OauthTokenId.Unset()
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *LogResponse) GetQueryString() map[string]interface{} {
	if o == nil || IsNil(o.QueryString) {
		var ret map[string]interface{}
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetQueryStringOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.QueryString) {
		return map[string]interface{}{}, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *LogResponse) HasQueryString() bool {
	if o != nil && !IsNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given map[string]interface{} and assigns it to the QueryString field.
func (o *LogResponse) SetQueryString(v map[string]interface{}) {
	o.QueryString = v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *LogResponse) GetRelated() string {
	if o == nil || IsNil(o.Related) {
		var ret string
		return ret
	}
	return *o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetRelatedOk() (*string, bool) {
	if o == nil || IsNil(o.Related) {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *LogResponse) HasRelated() bool {
	if o != nil && !IsNil(o.Related) {
		return true
	}

	return false
}

// SetRelated gets a reference to the given string and assigns it to the Related field.
func (o *LogResponse) SetRelated(v string) {
	o.Related = &v
}

// GetRequestBody returns the RequestBody field value if set, zero value otherwise.
func (o *LogResponse) GetRequestBody() map[string]interface{} {
	if o == nil || IsNil(o.RequestBody) {
		var ret map[string]interface{}
		return ret
	}
	return o.RequestBody
}

// GetRequestBodyOk returns a tuple with the RequestBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetRequestBodyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RequestBody) {
		return map[string]interface{}{}, false
	}
	return o.RequestBody, true
}

// HasRequestBody returns a boolean if a field has been set.
func (o *LogResponse) HasRequestBody() bool {
	if o != nil && !IsNil(o.RequestBody) {
		return true
	}

	return false
}

// SetRequestBody gets a reference to the given map[string]interface{} and assigns it to the RequestBody field.
func (o *LogResponse) SetRequestBody(v map[string]interface{}) {
	o.RequestBody = v
}

// GetRequestHeaders returns the RequestHeaders field value if set, zero value otherwise.
func (o *LogResponse) GetRequestHeaders() map[string]string {
	if o == nil || IsNil(o.RequestHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.RequestHeaders
}

// GetRequestHeadersOk returns a tuple with the RequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetRequestHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.RequestHeaders) {
		return nil, false
	}
	return o.RequestHeaders, true
}

// HasRequestHeaders returns a boolean if a field has been set.
func (o *LogResponse) HasRequestHeaders() bool {
	if o != nil && !IsNil(o.RequestHeaders) {
		return true
	}

	return false
}

// SetRequestHeaders gets a reference to the given map[string]string and assigns it to the RequestHeaders field.
func (o *LogResponse) SetRequestHeaders(v map[string]string) {
	o.RequestHeaders = &v
}

// GetResponseBody returns the ResponseBody field value if set, zero value otherwise.
func (o *LogResponse) GetResponseBody() map[string]interface{} {
	if o == nil || IsNil(o.ResponseBody) {
		var ret map[string]interface{}
		return ret
	}
	return o.ResponseBody
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetResponseBodyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ResponseBody) {
		return map[string]interface{}{}, false
	}
	return o.ResponseBody, true
}

// HasResponseBody returns a boolean if a field has been set.
func (o *LogResponse) HasResponseBody() bool {
	if o != nil && !IsNil(o.ResponseBody) {
		return true
	}

	return false
}

// SetResponseBody gets a reference to the given map[string]interface{} and assigns it to the ResponseBody field.
func (o *LogResponse) SetResponseBody(v map[string]interface{}) {
	o.ResponseBody = v
}

// GetResponseHeaders returns the ResponseHeaders field value if set, zero value otherwise.
func (o *LogResponse) GetResponseHeaders() map[string]string {
	if o == nil || IsNil(o.ResponseHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.ResponseHeaders
}

// GetResponseHeadersOk returns a tuple with the ResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetResponseHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ResponseHeaders) {
		return nil, false
	}
	return o.ResponseHeaders, true
}

// HasResponseHeaders returns a boolean if a field has been set.
func (o *LogResponse) HasResponseHeaders() bool {
	if o != nil && !IsNil(o.ResponseHeaders) {
		return true
	}

	return false
}

// SetResponseHeaders gets a reference to the given map[string]string and assigns it to the ResponseHeaders field.
func (o *LogResponse) SetResponseHeaders(v map[string]string) {
	o.ResponseHeaders = &v
}

// GetSearchableTags returns the SearchableTags field value if set, zero value otherwise.
func (o *LogResponse) GetSearchableTags() []string {
	if o == nil || IsNil(o.SearchableTags) {
		var ret []string
		return ret
	}
	return o.SearchableTags
}

// GetSearchableTagsOk returns a tuple with the SearchableTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetSearchableTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.SearchableTags) {
		return nil, false
	}
	return o.SearchableTags, true
}

// HasSearchableTags returns a boolean if a field has been set.
func (o *LogResponse) HasSearchableTags() bool {
	if o != nil && !IsNil(o.SearchableTags) {
		return true
	}

	return false
}

// SetSearchableTags gets a reference to the given []string and assigns it to the SearchableTags field.
func (o *LogResponse) SetSearchableTags(v []string) {
	o.SearchableTags = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LogResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LogResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LogResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LogResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LogResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *LogResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LogResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LogResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LogResponse) SetUrl(v string) {
	o.Url = &v
}

// GetUserAccountId returns the UserAccountId field value if set, zero value otherwise.
func (o *LogResponse) GetUserAccountId() string {
	if o == nil || IsNil(o.UserAccountId) {
		var ret string
		return ret
	}
	return *o.UserAccountId
}

// GetUserAccountIdOk returns a tuple with the UserAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetUserAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserAccountId) {
		return nil, false
	}
	return o.UserAccountId, true
}

// HasUserAccountId returns a boolean if a field has been set.
func (o *LogResponse) HasUserAccountId() bool {
	if o != nil && !IsNil(o.UserAccountId) {
		return true
	}

	return false
}

// SetUserAccountId gets a reference to the given string and assigns it to the UserAccountId field.
func (o *LogResponse) SetUserAccountId(v string) {
	o.UserAccountId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LogResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LogResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LogResponse) SetVersion(v string) {
	o.Version = &v
}

func (o LogResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	toSerialize["livemode"] = o.Livemode
	if o.LoggableId.IsSet() {
		toSerialize["loggable_id"] = o.LoggableId.Get()
	}
	if o.LoggableType.IsSet() {
		toSerialize["loggable_type"] = o.LoggableType.Get()
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if o.OauthTokenId.IsSet() {
		toSerialize["oauth_token_id"] = o.OauthTokenId.Get()
	}
	if !IsNil(o.QueryString) {
		toSerialize["query_string"] = o.QueryString
	}
	if !IsNil(o.Related) {
		toSerialize["related"] = o.Related
	}
	if !IsNil(o.RequestBody) {
		toSerialize["request_body"] = o.RequestBody
	}
	if !IsNil(o.RequestHeaders) {
		toSerialize["request_headers"] = o.RequestHeaders
	}
	if !IsNil(o.ResponseBody) {
		toSerialize["response_body"] = o.ResponseBody
	}
	if !IsNil(o.ResponseHeaders) {
		toSerialize["response_headers"] = o.ResponseHeaders
	}
	if !IsNil(o.SearchableTags) {
		toSerialize["searchable_tags"] = o.SearchableTags
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UserAccountId) {
		toSerialize["user_account_id"] = o.UserAccountId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableLogResponse struct {
	value *LogResponse
	isSet bool
}

func (v NullableLogResponse) Get() *LogResponse {
	return v.value
}

func (v *NullableLogResponse) Set(val *LogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogResponse(val *LogResponse) *NullableLogResponse {
	return &NullableLogResponse{value: val, isSet: true}
}

func (v NullableLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

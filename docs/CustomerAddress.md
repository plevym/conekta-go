# CustomerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street1** | **string** |  | 
**Street2** | Pointer to **string** |  | [optional] 
**PostalCode** | **string** |  | 
**City** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** | this field follows the [ISO 3166-1 alpha-2 standard](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) | [optional] 
**Residential** | Pointer to **bool** |  | [optional] 
**ExternalNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerAddress

`func NewCustomerAddress(street1 string, postalCode string, city string, ) *CustomerAddress`

NewCustomerAddress instantiates a new CustomerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressWithDefaults

`func NewCustomerAddressWithDefaults() *CustomerAddress`

NewCustomerAddressWithDefaults instantiates a new CustomerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet1

`func (o *CustomerAddress) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *CustomerAddress) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *CustomerAddress) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.


### GetStreet2

`func (o *CustomerAddress) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *CustomerAddress) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *CustomerAddress) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *CustomerAddress) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetPostalCode

`func (o *CustomerAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CustomerAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CustomerAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCity

`func (o *CustomerAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *CustomerAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *CustomerAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomerAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetResidential

`func (o *CustomerAddress) GetResidential() bool`

GetResidential returns the Residential field if non-nil, zero value otherwise.

### GetResidentialOk

`func (o *CustomerAddress) GetResidentialOk() (*bool, bool)`

GetResidentialOk returns a tuple with the Residential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidential

`func (o *CustomerAddress) SetResidential(v bool)`

SetResidential sets Residential field to given value.

### HasResidential

`func (o *CustomerAddress) HasResidential() bool`

HasResidential returns a boolean if a field has been set.

### GetExternalNumber

`func (o *CustomerAddress) GetExternalNumber() string`

GetExternalNumber returns the ExternalNumber field if non-nil, zero value otherwise.

### GetExternalNumberOk

`func (o *CustomerAddress) GetExternalNumberOk() (*string, bool)`

GetExternalNumberOk returns a tuple with the ExternalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNumber

`func (o *CustomerAddress) SetExternalNumber(v string)`

SetExternalNumber sets ExternalNumber field to given value.

### HasExternalNumber

`func (o *CustomerAddress) HasExternalNumber() bool`

HasExternalNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



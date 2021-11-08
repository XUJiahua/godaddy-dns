# ConsentRedemption

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **int32** | Price for the domain renewal (if domain renewal required for redemption). Please use GET /v2/customers/{customerId}/domains/{domain} to retrieve the redemption price and currency for the domain | [default to null]
**Fee** | **int32** | Fee charged for the domain redemption. Please use GET /v2/customers/{customerId}/domains/{domain} to retrieve the redemption fee and currency for the domain | [default to null]
**Currency** | **string** | Currency in which the &#x60;price&#x60; and &#x60;fee&#x60; are listed | [default to USD]
**AgreedBy** | **string** | Originating client IP address of the end-user&#x27;s computer when they consented to these legal agreements | [default to null]
**AgreedAt** | **string** | Timestamp indicating when the end-user consented to these legal agreements | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


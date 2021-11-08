# {{classname}}

All URIs are relative to *//api.ote-godaddy.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsForwardsDelete**](DomainsApi.md#DomainsForwardsDelete) | **Delete** /v2/customers/{customerId}/domains/forwards/{fqdn} | Submit a forwarding cancellation request for the given fqdn
[**DomainsForwardsGet**](DomainsApi.md#DomainsForwardsGet) | **Get** /v2/customers/{customerId}/domains/forwards/{fqdn} | Retrieve the forwarding information for the given fqdn
[**DomainsForwardsPost**](DomainsApi.md#DomainsForwardsPost) | **Post** /v2/customers/{customerId}/domains/forwards/{fqdn} | Create a new forwarding configuration for the given FQDN
[**DomainsForwardsPut**](DomainsApi.md#DomainsForwardsPut) | **Put** /v2/customers/{customerId}/domains/forwards/{fqdn} | Modify the forwarding information for the given fqdn
[**V2CustomersCustomerIdDomainsDomainGet**](DomainsApi.md#V2CustomersCustomerIdDomainsDomainGet) | **Get** /v2/customers/{customerId}/domains/{domain} | Retrieve details for the specified Domain
[**V2CustomersCustomerIdDomainsDomainRedeemPost**](DomainsApi.md#V2CustomersCustomerIdDomainsDomainRedeemPost) | **Post** /v2/customers/{customerId}/domains/{domain}/redeem | Purchase a restore for the given domain to bring it out of redemption
[**V2CustomersCustomerIdDomainsDomainTransferOutPost**](DomainsApi.md#V2CustomersCustomerIdDomainsDomainTransferOutPost) | **Post** /v2/customers/{customerId}/domains/{domain}/transferOut | Initiate transfer out to another registrar for a .uk domain.

# **DomainsForwardsDelete**
> DomainsForwardsDelete(ctx, customerId, fqdn)
Submit a forwarding cancellation request for the given fqdn

<strong>Notes:</strong><ul><li>**shopperId** is **not the same** as **customerId**.  **shopperId** is a number of max length 10 digits (*ex:* 1234567890) whereas **customerId** is a UUIDv4 (*ex:* 295e3bc3-b3b9-4d95-aae5-ede41a994d13)</li></ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **fqdn** | **string**| The fully qualified domain name whose forwarding details are to be deleted. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainsForwardsGet**
> []DomainForwarding DomainsForwardsGet(ctx, customerId, fqdn, optional)
Retrieve the forwarding information for the given fqdn

<strong>Notes:</strong><ul><li>**shopperId** is **not the same** as **customerId**.  **shopperId** is a number of max length 10 digits (*ex:* 1234567890) whereas **customerId** is a UUIDv4 (*ex:* 295e3bc3-b3b9-4d95-aae5-ede41a994d13)</li></ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **fqdn** | **string**| The fully qualified domain name whose forwarding details are to be retrieved. | 
 **optional** | ***DomainsApiDomainsForwardsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainsApiDomainsForwardsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeSubs** | **optional.Bool**| Optionally include all sub domains if the fqdn specified is a domain and not a sub domain. | 

### Return type

[**[]DomainForwarding**](DomainForwarding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainsForwardsPost**
> DomainsForwardsPost(ctx, body, customerId, fqdn)
Create a new forwarding configuration for the given FQDN

<strong>Notes:</strong><ul><li>**shopperId** is **not the same** as **customerId**.  **shopperId** is a number of max length 10 digits (*ex:* 1234567890) whereas **customerId** is a UUIDv4 (*ex:* 295e3bc3-b3b9-4d95-aae5-ede41a994d13)</li></ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainForwardingCreate**](DomainForwardingCreate.md)| Domain forwarding rule to create for the specified fqdn | 
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your own customer id. | 
  **fqdn** | **string**| The fully qualified domain name whose forwarding details are to be modified. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainsForwardsPut**
> DomainsForwardsPut(ctx, body, customerId, fqdn)
Modify the forwarding information for the given fqdn

<strong>Notes:</strong><ul><li>**shopperId** is **not the same** as **customerId**.  **shopperId** is a number of max length 10 digits (*ex:* 1234567890) whereas **customerId** is a UUIDv4 (*ex:* 295e3bc3-b3b9-4d95-aae5-ede41a994d13)</li></ul>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainForwardingCreate**](DomainForwardingCreate.md)| Domain forwarding rule to create or replace on the fqdn | 
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **fqdn** | **string**| The fully qualified domain name whose forwarding details are to be modified. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2CustomersCustomerIdDomainsDomainGet**
> DomainDetailV2 V2CustomersCustomerIdDomainsDomainGet(ctx, customerId, domain, optional)
Retrieve details for the specified Domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **domain** | **string**| Domain name whose details are to be retrieved | 
 **optional** | ***DomainsApiV2CustomersCustomerIdDomainsDomainGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainsApiV2CustomersCustomerIdDomainsDomainGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| A client provided identifier for tracking this request. | 
 **includes** | [**optional.Interface of []string**](string.md)| Optional details to be included in the response | 

### Return type

[**DomainDetailV2**](DomainDetailV2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2CustomersCustomerIdDomainsDomainRedeemPost**
> V2CustomersCustomerIdDomainsDomainRedeemPost(ctx, customerId, domain, optional)
Purchase a restore for the given domain to bring it out of redemption

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **domain** | **string**| Domain to request redeem for | 
 **optional** | ***DomainsApiV2CustomersCustomerIdDomainsDomainRedeemPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainsApiV2CustomersCustomerIdDomainsDomainRedeemPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DomainRedeemV2**](DomainRedeemV2.md)| Options for redeeming existing Domain | 
 **xRequestId** | **optional.**| A client provided identifier for tracking this request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2CustomersCustomerIdDomainsDomainTransferOutPost**
> V2CustomersCustomerIdDomainsDomainTransferOutPost(ctx, customerId, domain, registrar, optional)
Initiate transfer out to another registrar for a .uk domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **domain** | **string**| Domain to initiate the transfer out for | 
  **registrar** | **string**| Registrar tag to push transfer to | 
 **optional** | ***DomainsApiV2CustomersCustomerIdDomainsDomainTransferOutPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DomainsApiV2CustomersCustomerIdDomainsDomainTransferOutPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestId** | **optional.String**| A client provided identifier for tracking this request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


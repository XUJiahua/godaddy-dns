# {{classname}}

All URIs are relative to *//api.ote-godaddy.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2CustomersCustomerIdDomainsDomainActionsGet**](ActionsApi.md#V2CustomersCustomerIdDomainsDomainActionsGet) | **Get** /v2/customers/{customerId}/domains/{domain}/actions | Retrieves a list of the most recent actions for the specified domain
[**V2CustomersCustomerIdDomainsDomainActionsTypeDelete**](ActionsApi.md#V2CustomersCustomerIdDomainsDomainActionsTypeDelete) | **Delete** /v2/customers/{customerId}/domains/{domain}/actions/{type} | Cancel the most recent user action for the specified domain
[**V2CustomersCustomerIdDomainsDomainActionsTypeGet**](ActionsApi.md#V2CustomersCustomerIdDomainsDomainActionsTypeGet) | **Get** /v2/customers/{customerId}/domains/{domain}/actions/{type} | Retrieves the most recent action for the specified domain

# **V2CustomersCustomerIdDomainsDomainActionsGet**
> []Action V2CustomersCustomerIdDomainsDomainActionsGet(ctx, customerId, domain, optional)
Retrieves a list of the most recent actions for the specified domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **domain** | **string**| Domain whose actions are to be retrieved | 
 **optional** | ***ActionsApiV2CustomersCustomerIdDomainsDomainActionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionsApiV2CustomersCustomerIdDomainsDomainActionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| A client provided identifier for tracking this request. | 

### Return type

[**[]Action**](Action.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2CustomersCustomerIdDomainsDomainActionsTypeDelete**
> V2CustomersCustomerIdDomainsDomainActionsTypeDelete(ctx, customerId, domain, type_, optional)
Cancel the most recent user action for the specified domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **domain** | **string**| Domain whose action is to be cancelled | 
  **type_** | **string**| The type of action to cancel | 
 **optional** | ***ActionsApiV2CustomersCustomerIdDomainsDomainActionsTypeDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionsApiV2CustomersCustomerIdDomainsDomainActionsTypeDeleteOpts struct
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

# **V2CustomersCustomerIdDomainsDomainActionsTypeGet**
> Action V2CustomersCustomerIdDomainsDomainActionsTypeGet(ctx, customerId, domain, type_, optional)
Retrieves the most recent action for the specified domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **domain** | **string**| Domain whose action is to be retrieved | 
  **type_** | **string**| The type of action to retrieve | 
 **optional** | ***ActionsApiV2CustomersCustomerIdDomainsDomainActionsTypeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionsApiV2CustomersCustomerIdDomainsDomainActionsTypeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestId** | **optional.String**| A client provided identifier for tracking this request. | 

### Return type

[**Action**](Action.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


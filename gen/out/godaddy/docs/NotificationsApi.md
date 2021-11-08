# {{classname}}

All URIs are relative to *//api.ote-godaddy.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2CustomersCustomerIdDomainsNotificationsGet**](NotificationsApi.md#V2CustomersCustomerIdDomainsNotificationsGet) | **Get** /v2/customers/{customerId}/domains/notifications | Retrieve the next domain notification
[**V2CustomersCustomerIdDomainsNotificationsNotificationIdAcknowledgePost**](NotificationsApi.md#V2CustomersCustomerIdDomainsNotificationsNotificationIdAcknowledgePost) | **Post** /v2/customers/{customerId}/domains/notifications/{notificationId}/acknowledge | Acknowledge a domain notification
[**V2CustomersCustomerIdDomainsNotificationsOptInGet**](NotificationsApi.md#V2CustomersCustomerIdDomainsNotificationsOptInGet) | **Get** /v2/customers/{customerId}/domains/notifications/optIn | Retrieve a list of notification types that are opted in
[**V2CustomersCustomerIdDomainsNotificationsOptInPut**](NotificationsApi.md#V2CustomersCustomerIdDomainsNotificationsOptInPut) | **Put** /v2/customers/{customerId}/domains/notifications/optIn | Opt in to recieve notifications for the submitted notification types
[**V2CustomersCustomerIdDomainsNotificationsSchemasTypeGet**](NotificationsApi.md#V2CustomersCustomerIdDomainsNotificationsSchemasTypeGet) | **Get** /v2/customers/{customerId}/domains/notifications/schemas/{type} | Retrieve the schema for the notification data for the specified notification type

# **V2CustomersCustomerIdDomainsNotificationsGet**
> DomainNotification V2CustomersCustomerIdDomainsNotificationsGet(ctx, customerId, optional)
Retrieve the next domain notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
 **optional** | ***NotificationsApiV2CustomersCustomerIdDomainsNotificationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiV2CustomersCustomerIdDomainsNotificationsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| A client provided identifier for tracking this request. | 

### Return type

[**DomainNotification**](DomainNotification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2CustomersCustomerIdDomainsNotificationsNotificationIdAcknowledgePost**
> V2CustomersCustomerIdDomainsNotificationsNotificationIdAcknowledgePost(ctx, customerId, notificationId, optional)
Acknowledge a domain notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **notificationId** | **string**| The notification ID to acknowledge | 
 **optional** | ***NotificationsApiV2CustomersCustomerIdDomainsNotificationsNotificationIdAcknowledgePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiV2CustomersCustomerIdDomainsNotificationsNotificationIdAcknowledgePostOpts struct
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

# **V2CustomersCustomerIdDomainsNotificationsOptInGet**
> []DomainNotification V2CustomersCustomerIdDomainsNotificationsOptInGet(ctx, customerId, optional)
Retrieve a list of notification types that are opted in

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
 **optional** | ***NotificationsApiV2CustomersCustomerIdDomainsNotificationsOptInGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiV2CustomersCustomerIdDomainsNotificationsOptInGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| A client provided identifier for tracking this request. | 

### Return type

[**[]DomainNotification**](DomainNotification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2CustomersCustomerIdDomainsNotificationsOptInPut**
> V2CustomersCustomerIdDomainsNotificationsOptInPut(ctx, customerId, types, optional)
Opt in to recieve notifications for the submitted notification types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **types** | [**[]string**](string.md)| The notification types that should be opted in | 
 **optional** | ***NotificationsApiV2CustomersCustomerIdDomainsNotificationsOptInPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiV2CustomersCustomerIdDomainsNotificationsOptInPutOpts struct
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

# **V2CustomersCustomerIdDomainsNotificationsSchemasTypeGet**
> JsonSchema V2CustomersCustomerIdDomainsNotificationsSchemasTypeGet(ctx, customerId, type_, optional)
Retrieve the schema for the notification data for the specified notification type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#x27;re operating on behalf of; otherwise use your shopper id. | 
  **type_** | **string**| The notification type whose schema should be retrieved | 
 **optional** | ***NotificationsApiV2CustomersCustomerIdDomainsNotificationsSchemasTypeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiV2CustomersCustomerIdDomainsNotificationsSchemasTypeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| A client provided identifier for tracking this request. | 

### Return type

[**JsonSchema**](JsonSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


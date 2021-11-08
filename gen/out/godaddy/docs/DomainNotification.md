# DomainNotification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationId** | **string** | The notification ID to be used in POST /v2/customers/{customerId}/domains/notifications to acknowledge the notification | 
**Type_** | **string** | The type of action the notification relates to | [default to null]
**Resource** | **string** | The resource the notification pertains to. | 
**ResourceType** | **string** | The type of resource the notification relates to | [default to null]
**Status** | **string** | The resulting status of the action. | [default to null]
**AddedAt** | **string** | The date the notification was added | 
**RequestId** | **string** | A client provided identifier (via X-Request-Id header) indicating the request this notification is for | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | The notification data for the given type as specifed by GET /v2/customers/{customerId}/domains/notifications/schema | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


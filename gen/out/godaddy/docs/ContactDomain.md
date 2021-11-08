# ContactDomain

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | **string** | Unique identifier for this Contact | [optional] [default to null]
**Encoding** | **string** | The encoding of the contact data&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;ASCII&lt;/strong&gt; - Data contains only ASCII characters that are not region or language specific&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;UTF-8&lt;/strong&gt; - Data contains characters that are specific to a region or language&lt;/li&gt;&lt;/ul&gt; | [optional] [default to ENCODING.ASCII]
**NameFirst** | **string** |  | [default to null]
**NameMiddle** | **string** |  | [optional] [default to null]
**NameLast** | **string** |  | [default to null]
**Organization** | **string** |  | [optional] [default to null]
**JobTitle** | **string** |  | [optional] [default to null]
**Email** | **string** |  | [default to null]
**Phone** | **string** |  | [default to null]
**Fax** | **string** |  | [optional] [default to null]
**AddressMailing** | [***Address**](Address.md) |  | [default to null]
**ExposeWhois** | **bool** | Whether or not the contact details should be shown in the WHOIS | [default to null]
**Metadata** | [***interface{}**](interface{}.md) | The contact eligibility data fields as specified by GET /v2/customers/{customerId}/domains/contacts/schema/{tld} | [optional] [default to null]
**Tlds** | **[]string** | The tlds that this contact can be assigned to | [optional] [default to null]
**CreatedAt** | **string** | Timestamp indicating when the contact was created | [optional] [default to null]
**ModifiedAt** | **string** | Timestamp indicating when the contact was last modified | [optional] [default to null]
**Deleted** | **bool** | Flag indicating if the contact has been logically deleted in the system | [optional] [default to null]
**Revision** | **int32** | The current revision number of the contact. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


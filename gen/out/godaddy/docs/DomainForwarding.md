# DomainForwarding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | **string** | The fqdn (domain or sub domain) to forward (ex somedomain.com or sub.somedomain.com) | [default to null]
**Type_** | **string** | The type of fowarding to implement&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;MASKED&lt;/strong&gt; - Prevents the forwarded domain or subdomain URL from displaying in the browser&#x27;s address bar.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;REDIRECT_PERMANENT*&lt;/strong&gt; - Redirects to the url you specified in the forwardTo field using a &#x60;301 Moved Permanently&#x60; HTTP response. The HTTP 301 response code tells user-agents (including search engines) that the location has permanently moved.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;REDIRECT_TEMPORARY&lt;/strong&gt; - Redirects to the url you specified in the forwardTo field using a &#x60;302 Found&#x60; HTTP response. The HTTP 302 response code tells user-agents (including search engines) that the location has temporarily moved.&lt;/li&gt;&lt;/ul&gt; | [default to TYPE_.REDIRECT_PERMANENT]
**Url** | **string** | Forwards http(s) traffic to this destination url (ex. http://www.somedomain.com/) | [default to null]
**Mask** | [***DomainForwardingMask**](DomainForwardingMask.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


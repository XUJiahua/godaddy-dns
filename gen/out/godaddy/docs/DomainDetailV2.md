# DomainDetailV2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Unique identifier for this Domain | [default to null]
**Domain** | **string** | Name of the domain | [default to null]
**SubaccountId** | **string** | Reseller subaccount shopperid who can manage the domain | [optional] [default to null]
**Status** | **string** | The current status of the domain&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;ACTIVE&lt;/strong&gt; - Domain has been registered and is active.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;CANCELLED&lt;/strong&gt; - Domain has been cancelled by the user or system, and is not be reclaimable.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;DELETED_REDEEMABLE&lt;/strong&gt; - Domain is deleted but is redeemable.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;EXPIRED&lt;/strong&gt; - Domain has expired.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;FAILED&lt;/strong&gt; - Domain registration or transfer error.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;LOCKED_REGISTRAR&lt;/strong&gt; - Domain is locked at the registrar - this is usually the result of a spam, abuse, etc.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PARKED&lt;/strong&gt; - Domain has been parked.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;HELD_REGISTRAR&lt;/strong&gt; - Domain is held at the registrar and cannot be transferred or modified - this is usually the result of a dispute.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;OWNERSHIP_CHANGED&lt;/strong&gt; - Domain has been moved to another account.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PENDING_TRANSFER&lt;/strong&gt; - Domain transfer has been requested and is pending the transfer process.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PENDING_REGISTRATION&lt;/strong&gt; - Domain is pending setup at the registry.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;REPOSSESSED&lt;/strong&gt; - Domain has been confiscated - this is usually the result of a chargeback, fraud, abuse, etc.).&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;SUSPENDED&lt;/strong&gt; - Domain is in violation and has been suspended.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;TRANSFERRED&lt;/strong&gt; - Domain has been transferred to another registrar.&lt;/li&gt;&lt;/ul&gt; | [default to null]
**ExpiresAt** | **string** | Date and time when this domain will expire | [optional] [default to null]
**ExpirationProtected** | **bool** | Whether or not the domain is protected from expiration | [default to null]
**HoldRegistrar** | **bool** | Whether or not the domain is on-hold by the registrar | [default to null]
**Locked** | **bool** | Whether or not the domain is locked to prevent transfers | [default to null]
**Privacy** | **bool** | Whether or not the domain has privacy protection | [default to null]
**RenewAuto** | **bool** | Whether or not the domain is configured to automatically renew | [default to null]
**RenewDeadline** | **string** | Date the domain must renew on | [default to null]
**TransferProtected** | **bool** | Whether or not the domain is protected from transfer | [default to null]
**CreatedAt** | **string** | Date and time when this domain was created | [default to null]
**DeletedAt** | **string** | Date and time when this domain was deleted | [optional] [default to null]
**ModifiedAt** | **string** | Date and time when this domain was last modified | [optional] [default to null]
**TransferAwayEligibleAt** | **string** | Date and time when this domain is eligible to transfer | [optional] [default to null]
**AuthCode** | **string** | Authorization code for transferring the Domain | [default to null]
**NameServers** | **[]string** | Fully-qualified domain names for DNS servers | [default to null]
**Hostnames** | **[]string** | Hostnames owned by the domain | [optional] [default to null]
**Renewal** | [***RenewalDetails**](RenewalDetails.md) |  | [optional] [default to null]
**Verifications** | [***VerificationsDomainV2**](VerificationsDomainV2.md) |  | [optional] [default to null]
**Contacts** | [***DomainContactsV2**](DomainContactsV2.md) |  | [default to null]
**Actions** | [**[]Action**](Action.md) | List of current actions in progress for this domain | [optional] [default to null]
**DnssecRecords** | [**[]DomainDnssec**](DomainDnssec.md) | List of active DNSSEC records for this domain | [optional] [default to null]
**RegistryStatusCodes** | **[]string** | The current registry status codes of the domain&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;ADD_PERIOD&lt;/strong&gt; - This grace period is provided after the initial registration of a domain name.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;AUTO_RENEW_PERIOD&lt;/strong&gt; - This grace period is provided after a domain name registration period expires and is extended (renewed) automatically by the registry.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;CLIENT_DELETE_PROHIBITED&lt;/strong&gt; - This status code tells your domain&#x27;s registry to reject requests to delete the domain.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;CLIENT_HOLD&lt;/strong&gt; - This status code tells your domain&#x27;s registry to not activate your domain in the DNS and as a consequence, it will not resolve.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;CLIENT_RENEW_PROHIBITED&lt;/strong&gt; - This status code tells your domain&#x27;s registry to reject requests to renew your domain.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;CLIENT_TRANSFER_PROHIBITED&lt;/strong&gt; - This status code tells your domain&#x27;s registry to reject requests to transfer the domain from your current registrar to another.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;CLIENT_UPDATE_PROHIBITED&lt;/strong&gt; - This status code tells your domain&#x27;s registry to reject requests to update the domain.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;INACTIVE&lt;/strong&gt; - This status code indicates that delegation information (name servers) has not been associated with your domain.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;OK&lt;/strong&gt; - This is the standard status for a domain, meaning it has no pending operations or prohibitions.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PENDING_CREATE&lt;/strong&gt; - This status code indicates that a request to create your domain has been received and is being processed.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PENDING_DELETE&lt;/strong&gt; - This status code indicates that the domain is either in a redemption period if combined with either REDEMPTION_PERIOD or PENDING_RESTORE, if not combined with these, then indicates that the redemption period for the domain has ended and domain will be be purged and dropped from the registry database.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PENDING_RENEW&lt;/strong&gt; - This status code indicates that a request to renew your domain has been received and is being processed.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PENDING_RESTORE&lt;/strong&gt; - This status code indicates that your registrar has asked the registry to restore your domain that was in REDEMPTION_PERIOD status&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PENDING_TRANSFER&lt;/strong&gt; - This status code indicates that a request to transfer your domain to a new registrar has been received and is being processed.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;PENDING_UPDATE&lt;/strong&gt; - This status code indicates that a request to update your domain has been received and is being processed.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;REDEMPTION_PERIOD&lt;/strong&gt; - This status code indicates that your registrar has asked the registry to delete your domain.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;RENEW_PERIOD&lt;/strong&gt; - This grace period is provided after a domain name registration period is explicitly extended (renewed) by the registrar.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;SERVER_DELETE_PROHIBITED&lt;/strong&gt; - This status code prevents your domain from being deleted. &lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;SERVER_HOLD&lt;/strong&gt; - This status code is set by your domain&#x27;s Registry Operator. Your domain is not activated in the DNS.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;SERVER_RENEW_PROHIBITED&lt;/strong&gt; - This status code indicates your domain&#x27;s Registry Operator will not allow your registrar to renew your domain.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;SERVER_TRANSFER_PROHIBITED&lt;/strong&gt; - This status code prevents your domain from being transferred from your current registrar to another. &lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;SERVER_UPDATE_PROHIBITED&lt;/strong&gt; - This status code locks your domain preventing it from being updated.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#x27;margin-left: 12px;&#x27;&gt;TRANSFER_PERIOD&lt;/strong&gt; - This grace period is provided after the successful transfer of a domain name from one registrar to another. &lt;/li&gt;&lt;/ul&gt; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

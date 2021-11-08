/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type DomainDetailV2 struct {
	// Unique identifier for this Domain
	DomainId string `json:"domainId"`
	// Name of the domain
	Domain string `json:"domain"`
	// Reseller subaccount shopperid who can manage the domain
	SubaccountId string `json:"subaccountId,omitempty"`
	// The current status of the domain<br/><ul><li><strong style='margin-left: 12px;'>ACTIVE</strong> - Domain has been registered and is active.</li><li><strong style='margin-left: 12px;'>CANCELLED</strong> - Domain has been cancelled by the user or system, and is not be reclaimable.</li><li><strong style='margin-left: 12px;'>DELETED_REDEEMABLE</strong> - Domain is deleted but is redeemable.</li><li><strong style='margin-left: 12px;'>EXPIRED</strong> - Domain has expired.</li><li><strong style='margin-left: 12px;'>FAILED</strong> - Domain registration or transfer error.</li><li><strong style='margin-left: 12px;'>LOCKED_REGISTRAR</strong> - Domain is locked at the registrar - this is usually the result of a spam, abuse, etc.</li><li><strong style='margin-left: 12px;'>PARKED</strong> - Domain has been parked.</li><li><strong style='margin-left: 12px;'>HELD_REGISTRAR</strong> - Domain is held at the registrar and cannot be transferred or modified - this is usually the result of a dispute.</li><li><strong style='margin-left: 12px;'>OWNERSHIP_CHANGED</strong> - Domain has been moved to another account.</li><li><strong style='margin-left: 12px;'>PENDING_TRANSFER</strong> - Domain transfer has been requested and is pending the transfer process.</li><li><strong style='margin-left: 12px;'>PENDING_REGISTRATION</strong> - Domain is pending setup at the registry.</li><li><strong style='margin-left: 12px;'>REPOSSESSED</strong> - Domain has been confiscated - this is usually the result of a chargeback, fraud, abuse, etc.).</li><li><strong style='margin-left: 12px;'>SUSPENDED</strong> - Domain is in violation and has been suspended.</li><li><strong style='margin-left: 12px;'>TRANSFERRED</strong> - Domain has been transferred to another registrar.</li></ul>
	Status string `json:"status"`
	// Date and time when this domain will expire
	ExpiresAt string `json:"expiresAt,omitempty"`
	// Whether or not the domain is protected from expiration
	ExpirationProtected bool `json:"expirationProtected"`
	// Whether or not the domain is on-hold by the registrar
	HoldRegistrar bool `json:"holdRegistrar"`
	// Whether or not the domain is locked to prevent transfers
	Locked bool `json:"locked"`
	// Whether or not the domain has privacy protection
	Privacy bool `json:"privacy"`
	// Whether or not the domain is configured to automatically renew
	RenewAuto bool `json:"renewAuto"`
	// Date the domain must renew on
	RenewDeadline string `json:"renewDeadline"`
	// Whether or not the domain is protected from transfer
	TransferProtected bool `json:"transferProtected"`
	// Date and time when this domain was created
	CreatedAt string `json:"createdAt"`
	// Date and time when this domain was deleted
	DeletedAt string `json:"deletedAt,omitempty"`
	// Date and time when this domain was last modified
	ModifiedAt string `json:"modifiedAt,omitempty"`
	// Date and time when this domain is eligible to transfer
	TransferAwayEligibleAt string `json:"transferAwayEligibleAt,omitempty"`
	// Authorization code for transferring the Domain
	AuthCode string `json:"authCode"`
	// Fully-qualified domain names for DNS servers
	NameServers []string `json:"nameServers"`
	// Hostnames owned by the domain
	Hostnames     []string               `json:"hostnames,omitempty"`
	Renewal       *RenewalDetails        `json:"renewal,omitempty"`
	Verifications *VerificationsDomainV2 `json:"verifications,omitempty"`
	Contacts      *DomainContactsV2      `json:"contacts"`
	// List of current actions in progress for this domain
	Actions []Action `json:"actions,omitempty"`
	// List of active DNSSEC records for this domain
	DnssecRecords []DomainDnssec `json:"dnssecRecords,omitempty"`
	// The current registry status codes of the domain<br/><ul><li><strong style='margin-left: 12px;'>ADD_PERIOD</strong> - This grace period is provided after the initial registration of a domain name.</li><li><strong style='margin-left: 12px;'>AUTO_RENEW_PERIOD</strong> - This grace period is provided after a domain name registration period expires and is extended (renewed) automatically by the registry.</li><li><strong style='margin-left: 12px;'>CLIENT_DELETE_PROHIBITED</strong> - This status code tells your domain's registry to reject requests to delete the domain.</li><li><strong style='margin-left: 12px;'>CLIENT_HOLD</strong> - This status code tells your domain's registry to not activate your domain in the DNS and as a consequence, it will not resolve.</li><li><strong style='margin-left: 12px;'>CLIENT_RENEW_PROHIBITED</strong> - This status code tells your domain's registry to reject requests to renew your domain.</li><li><strong style='margin-left: 12px;'>CLIENT_TRANSFER_PROHIBITED</strong> - This status code tells your domain's registry to reject requests to transfer the domain from your current registrar to another.</li><li><strong style='margin-left: 12px;'>CLIENT_UPDATE_PROHIBITED</strong> - This status code tells your domain's registry to reject requests to update the domain.</li><li><strong style='margin-left: 12px;'>INACTIVE</strong> - This status code indicates that delegation information (name servers) has not been associated with your domain.</li><li><strong style='margin-left: 12px;'>OK</strong> - This is the standard status for a domain, meaning it has no pending operations or prohibitions.</li><li><strong style='margin-left: 12px;'>PENDING_CREATE</strong> - This status code indicates that a request to create your domain has been received and is being processed.</li><li><strong style='margin-left: 12px;'>PENDING_DELETE</strong> - This status code indicates that the domain is either in a redemption period if combined with either REDEMPTION_PERIOD or PENDING_RESTORE, if not combined with these, then indicates that the redemption period for the domain has ended and domain will be be purged and dropped from the registry database.</li><li><strong style='margin-left: 12px;'>PENDING_RENEW</strong> - This status code indicates that a request to renew your domain has been received and is being processed.</li><li><strong style='margin-left: 12px;'>PENDING_RESTORE</strong> - This status code indicates that your registrar has asked the registry to restore your domain that was in REDEMPTION_PERIOD status</li><li><strong style='margin-left: 12px;'>PENDING_TRANSFER</strong> - This status code indicates that a request to transfer your domain to a new registrar has been received and is being processed.</li><li><strong style='margin-left: 12px;'>PENDING_UPDATE</strong> - This status code indicates that a request to update your domain has been received and is being processed.</li><li><strong style='margin-left: 12px;'>REDEMPTION_PERIOD</strong> - This status code indicates that your registrar has asked the registry to delete your domain.</li><li><strong style='margin-left: 12px;'>RENEW_PERIOD</strong> - This grace period is provided after a domain name registration period is explicitly extended (renewed) by the registrar.</li><li><strong style='margin-left: 12px;'>SERVER_DELETE_PROHIBITED</strong> - This status code prevents your domain from being deleted. </li><li><strong style='margin-left: 12px;'>SERVER_HOLD</strong> - This status code is set by your domain's Registry Operator. Your domain is not activated in the DNS.</li><li><strong style='margin-left: 12px;'>SERVER_RENEW_PROHIBITED</strong> - This status code indicates your domain's Registry Operator will not allow your registrar to renew your domain.</li><li><strong style='margin-left: 12px;'>SERVER_TRANSFER_PROHIBITED</strong> - This status code prevents your domain from being transferred from your current registrar to another. </li><li><strong style='margin-left: 12px;'>SERVER_UPDATE_PROHIBITED</strong> - This status code locks your domain preventing it from being updated.</li><li><strong style='margin-left: 12px;'>TRANSFER_PERIOD</strong> - This grace period is provided after the successful transfer of a domain name from one registrar to another. </li></ul>
	RegistryStatusCodes []string `json:"registryStatusCodes,omitempty"`
}

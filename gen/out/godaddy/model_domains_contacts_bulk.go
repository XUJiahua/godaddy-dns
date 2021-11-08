/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type DomainsContactsBulk struct {
	ContactAdmin      *Contact `json:"contactAdmin,omitempty"`
	ContactBilling    *Contact `json:"contactBilling,omitempty"`
	ContactPresence   *Contact `json:"contactPresence,omitempty"`
	ContactRegistrant *Contact `json:"contactRegistrant,omitempty"`
	ContactTech       *Contact `json:"contactTech,omitempty"`
	// An array of domain names to be validated against. Alternatively, you can specify the extracted tlds. However, full domain names are required if the tld is `uk`
	Domains []string `json:"domains"`
	// Canadian Presence Requirement (CA)
	EntityType string `json:"entityType,omitempty"`
}

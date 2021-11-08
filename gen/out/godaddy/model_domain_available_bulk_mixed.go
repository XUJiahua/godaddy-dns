/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type DomainAvailableBulkMixed struct {
	// Domain available response array
	Domains []DomainAvailableResponse `json:"domains"`
	// Errors encountered while performing a domain available check
	Errors []DomainAvailableError `json:"errors,omitempty"`
}

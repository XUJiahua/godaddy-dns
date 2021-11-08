/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type DomainAvailableResponse struct {
	// Whether or not the domain name is available
	Available bool `json:"available"`
	// Currency in which the `price` is listed. Only returned if tld is offered
	Currency string `json:"currency,omitempty"`
	// Whether or not the `available` answer has been definitively verified with the registry
	Definitive bool `json:"definitive"`
	// Domain name
	Domain string `json:"domain"`
	// Number of years included in the price. Only returned if tld is offered
	Period int32 `json:"period,omitempty"`
	// Price of the domain excluding taxes or fees. Only returned if tld is offered
	Price int32 `json:"price,omitempty"`
}

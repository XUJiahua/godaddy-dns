/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type Consent struct {
	// Timestamp indicating when the end-user consented to these legal agreements
	AgreedAt string `json:"agreedAt"`
	// Originating client IP address of the end-user's computer when they consented to these legal agreements
	AgreedBy string `json:"agreedBy"`
	// Unique identifiers of the legal agreements to which the end-user has agreed, as returned from the/domains/agreements endpoint
	AgreementKeys []string `json:"agreementKeys"`
}

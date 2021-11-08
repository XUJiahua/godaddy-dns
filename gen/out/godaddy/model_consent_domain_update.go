/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type ConsentDomainUpdate struct {
	// Timestamp indicating when the end-user consented to these agreements
	AgreedAt string `json:"agreedAt"`
	// Originating client IP address of the end-user's computer when they consented to the agreements
	AgreedBy string `json:"agreedBy"`
	// Unique identifiers of the agreements to which the end-user has agreed, as required by the elements being updated<br/><ul><li><strong style='margin-left: 12px;'>EXPOSE_WHOIS</strong> - Required when the exposeWhois field is updated to true</li></ul>
	AgreementKeys []string `json:"agreementKeys"`
}

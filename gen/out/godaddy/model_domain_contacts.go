/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type DomainContacts struct {
	ContactAdmin      *Contact `json:"contactAdmin,omitempty"`
	ContactBilling    *Contact `json:"contactBilling,omitempty"`
	ContactRegistrant *Contact `json:"contactRegistrant"`
	ContactTech       *Contact `json:"contactTech,omitempty"`
}
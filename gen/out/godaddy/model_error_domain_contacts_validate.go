/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type ErrorDomainContactsValidate struct {
	// Short identifier for the error, suitable for indicating the specific error within client code
	Code string `json:"code"`
	// List of the specific fields, and the errors found with their contents
	Fields []ErrorFieldDomainContactsValidate `json:"fields,omitempty"`
	// Human-readable, English description of the error
	Message string `json:"message,omitempty"`
	// Stack trace indicating where the error occurred.<br/>NOTE: This attribute <strong>MAY</strong> be included for Development and Test environments. However, it <strong>MUST NOT</strong> be exposed from OTE nor Production systems
	Stack []string `json:"stack,omitempty"`
}

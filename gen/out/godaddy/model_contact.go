/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type Contact struct {
	AddressMailing *Address `json:"addressMailing"`
	Email          string   `json:"email"`
	Fax            string   `json:"fax,omitempty"`
	JobTitle       string   `json:"jobTitle,omitempty"`
	NameFirst      string   `json:"nameFirst"`
	NameLast       string   `json:"nameLast"`
	NameMiddle     string   `json:"nameMiddle,omitempty"`
	Organization   string   `json:"organization,omitempty"`
	Phone          string   `json:"phone"`
}
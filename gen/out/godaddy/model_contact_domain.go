/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type ContactDomain struct {
	// Unique identifier for this Contact
	ContactId string `json:"contactId,omitempty"`
	// The encoding of the contact data<br/><ul><li><strong style='margin-left: 12px;'>ASCII</strong> - Data contains only ASCII characters that are not region or language specific</li><li><strong style='margin-left: 12px;'>UTF-8</strong> - Data contains characters that are specific to a region or language</li></ul>
	Encoding       string   `json:"encoding,omitempty"`
	NameFirst      string   `json:"nameFirst"`
	NameMiddle     string   `json:"nameMiddle,omitempty"`
	NameLast       string   `json:"nameLast"`
	Organization   string   `json:"organization,omitempty"`
	JobTitle       string   `json:"jobTitle,omitempty"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Fax            string   `json:"fax,omitempty"`
	AddressMailing *Address `json:"addressMailing"`
	// Whether or not the contact details should be shown in the WHOIS
	ExposeWhois bool `json:"exposeWhois"`
	// The contact eligibility data fields as specified by GET /v2/customers/{customerId}/domains/contacts/schema/{tld}
	Metadata *interface{} `json:"metadata,omitempty"`
	// The tlds that this contact can be assigned to
	Tlds []string `json:"tlds,omitempty"`
	// Timestamp indicating when the contact was created
	CreatedAt string `json:"_createdAt,omitempty"`
	// Timestamp indicating when the contact was last modified
	ModifiedAt string `json:"_modifiedAt,omitempty"`
	// Flag indicating if the contact has been logically deleted in the system
	Deleted bool `json:"_deleted,omitempty"`
	// The current revision number of the contact.
	Revision int32 `json:"_revision,omitempty"`
}

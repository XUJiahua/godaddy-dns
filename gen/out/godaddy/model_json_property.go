/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package godaddy

type JsonProperty struct {
	DefaultValue string       `json:"defaultValue,omitempty"`
	Format       string       `json:"format,omitempty"`
	Items        *interface{} `json:"items,omitempty"`
	MaxItems     int32        `json:"maxItems,omitempty"`
	Maximum      int32        `json:"maximum,omitempty"`
	MinItems     int32        `json:"minItems,omitempty"`
	Minimum      int32        `json:"minimum,omitempty"`
	Pattern      string       `json:"pattern,omitempty"`
	Required     bool         `json:"required"`
	Type_        string       `json:"type"`
}

/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RestListDocstoreRequest struct {

	StoreID string `json:"StoreID,omitempty"`

	Query *DocstoreDocumentQuery `json:"Query,omitempty"`

	CountOnly bool `json:"CountOnly,omitempty"`
}

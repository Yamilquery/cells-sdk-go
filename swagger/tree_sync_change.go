/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TreeSyncChange struct {

	Seq string `json:"seq,omitempty"`

	NodeId string `json:"nodeId,omitempty"`

	Type_ *TreeSyncChangeType `json:"type,omitempty"`

	Source string `json:"source,omitempty"`

	Target string `json:"target,omitempty"`

	Node *TreeSyncChangeNode `json:"node,omitempty"`
}

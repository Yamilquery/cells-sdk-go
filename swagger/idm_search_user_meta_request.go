/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IdmSearchUserMetaRequest struct {

	MetaUuids []string `json:"MetaUuids,omitempty"`

	NodeUuids []string `json:"NodeUuids,omitempty"`

	Namespace string `json:"Namespace,omitempty"`

	ResourceSubjectOwner string `json:"ResourceSubjectOwner,omitempty"`

	ResourceQuery *ServiceResourcePolicyQuery `json:"ResourceQuery,omitempty"`
}

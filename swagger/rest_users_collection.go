/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RestUsersCollection struct {

	Groups []IdmUser `json:"Groups,omitempty"`

	Users []IdmUser `json:"Users,omitempty"`

	Total int32 `json:"Total,omitempty"`
}

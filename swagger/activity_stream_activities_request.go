/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ActivityStreamActivitiesRequest struct {

	Context *ActivityStreamContext `json:"Context,omitempty"`

	ContextData string `json:"ContextData,omitempty"`

	StreamFilter string `json:"StreamFilter,omitempty"`

	BoxName string `json:"BoxName,omitempty"`

	UnreadCountOnly bool `json:"UnreadCountOnly,omitempty"`

	Offset string `json:"Offset,omitempty"`

	Limit string `json:"Limit,omitempty"`

	AsDigest bool `json:"AsDigest,omitempty"`

	PointOfView *ActivitySummaryPointOfView `json:"PointOfView,omitempty"`

	Language string `json:"Language,omitempty"`
}

/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NodeChangeEventEventType string

// List of NodeChangeEventEventType
const (
	CREATE NodeChangeEventEventType = "CREATE"
	READ NodeChangeEventEventType = "READ"
	UPDATE_PATH NodeChangeEventEventType = "UPDATE_PATH"
	UPDATE_CONTENT NodeChangeEventEventType = "UPDATE_CONTENT"
	UPDATE_META NodeChangeEventEventType = "UPDATE_META"
	DELETE NodeChangeEventEventType = "DELETE"
)

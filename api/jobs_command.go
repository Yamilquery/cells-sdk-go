/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type JobsCommand string

// List of jobsCommand
const (
	NONE        JobsCommand = "None"
	PAUSE       JobsCommand = "Pause"
	RESUME      JobsCommand = "Resume"
	JC_STOP     JobsCommand = "Stop"
	JC_DELETE   JobsCommand = "Delete"
	JC_RUN_ONCE JobsCommand = "RunOnce"
	INACTIVE    JobsCommand = "Inactive"
	ACTIVE      JobsCommand = "Active"
)

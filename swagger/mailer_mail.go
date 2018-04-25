/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MailerMail struct {

	From *MailerUser `json:"From,omitempty"`

	To []MailerUser `json:"To,omitempty"`

	Cc []MailerUser `json:"Cc,omitempty"`

	DateSent string `json:"DateSent,omitempty"`

	Subject string `json:"Subject,omitempty"`

	ContentPlain string `json:"ContentPlain,omitempty"`

	ContentHtml string `json:"ContentHtml,omitempty"`

	ContentMarkdown string `json:"ContentMarkdown,omitempty"`

	Attachments []string `json:"Attachments,omitempty"`

	ThreadUuid string `json:"ThreadUuid,omitempty"`

	ThreadIndex string `json:"ThreadIndex,omitempty"`

	TemplateId string `json:"TemplateId,omitempty"`

	TemplateData map[string]string `json:"TemplateData,omitempty"`

	Retries int32 `json:"Retries,omitempty"`

	SendErrors []string `json:"sendErrors,omitempty"`
}

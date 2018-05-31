// Code generated by go-swagger; DO NOT EDIT.

package enterprise_log_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// AuditReader is a Reader for the Audit structure.
type AuditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAuditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuditOK creates a AuditOK with default headers values
func NewAuditOK() *AuditOK {
	return &AuditOK{}
}

/*AuditOK handles this case with default header values.

AuditOK audit o k
*/
type AuditOK struct {
	Payload *models.RestLogMessageCollection
}

func (o *AuditOK) Error() string {
	return fmt.Sprintf("[POST /log/audit][%d] auditOK  %+v", 200, o.Payload)
}

func (o *AuditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestLogMessageCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

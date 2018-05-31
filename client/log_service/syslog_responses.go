// Code generated by go-swagger; DO NOT EDIT.

package log_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// SyslogReader is a Reader for the Syslog structure.
type SyslogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyslogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSyslogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSyslogOK creates a SyslogOK with default headers values
func NewSyslogOK() *SyslogOK {
	return &SyslogOK{}
}

/*SyslogOK handles this case with default header values.

SyslogOK syslog o k
*/
type SyslogOK struct {
	Payload *models.RestLogMessageCollection
}

func (o *SyslogOK) Error() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogOK  %+v", 200, o.Payload)
}

func (o *SyslogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestLogMessageCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

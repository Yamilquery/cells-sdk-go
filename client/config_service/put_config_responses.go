// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// PutConfigReader is a Reader for the PutConfig structure.
type PutConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutConfigOK creates a PutConfigOK with default headers values
func NewPutConfigOK() *PutConfigOK {
	return &PutConfigOK{}
}

/*PutConfigOK handles this case with default header values.

PutConfigOK put config o k
*/
type PutConfigOK struct {
	Payload *models.RestConfiguration
}

func (o *PutConfigOK) Error() string {
	return fmt.Sprintf("[PUT /config/{FullPath}][%d] putConfigOK  %+v", 200, o.Payload)
}

func (o *PutConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

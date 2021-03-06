// Code generated by go-swagger; DO NOT EDIT.

package user_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// PutRolesReader is a Reader for the PutRoles structure.
type PutRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutRolesOK creates a PutRolesOK with default headers values
func NewPutRolesOK() *PutRolesOK {
	return &PutRolesOK{}
}

/*PutRolesOK handles this case with default header values.

PutRolesOK put roles o k
*/
type PutRolesOK struct {
	Payload *models.IdmUser
}

func (o *PutRolesOK) Error() string {
	return fmt.Sprintf("[PUT /user/roles/{Login}][%d] putRolesOK  %+v", 200, o.Payload)
}

func (o *PutRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

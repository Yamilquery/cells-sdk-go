// Code generated by go-swagger; DO NOT EDIT.

package acl_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// SearchAclsReader is a Reader for the SearchAcls structure.
type SearchAclsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAclsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchAclsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchAclsOK creates a SearchAclsOK with default headers values
func NewSearchAclsOK() *SearchAclsOK {
	return &SearchAclsOK{}
}

/*SearchAclsOK handles this case with default header values.

SearchAclsOK search acls o k
*/
type SearchAclsOK struct {
	Payload *models.RestACLCollection
}

func (o *SearchAclsOK) Error() string {
	return fmt.Sprintf("[POST /acl][%d] searchAclsOK  %+v", 200, o.Payload)
}

func (o *SearchAclsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestACLCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

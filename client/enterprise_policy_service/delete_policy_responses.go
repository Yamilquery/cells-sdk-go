// Code generated by go-swagger; DO NOT EDIT.

package enterprise_policy_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// DeletePolicyReader is a Reader for the DeletePolicy structure.
type DeletePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePolicyOK creates a DeletePolicyOK with default headers values
func NewDeletePolicyOK() *DeletePolicyOK {
	return &DeletePolicyOK{}
}

/*DeletePolicyOK handles this case with default header values.

DeletePolicyOK delete policy o k
*/
type DeletePolicyOK struct {
	Payload *models.RestDeleteResponse
}

func (o *DeletePolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /policy/{Uuid}][%d] deletePolicyOK  %+v", 200, o.Payload)
}

func (o *DeletePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

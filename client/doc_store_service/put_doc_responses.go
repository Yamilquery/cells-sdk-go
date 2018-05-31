// Code generated by go-swagger; DO NOT EDIT.

package doc_store_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// PutDocReader is a Reader for the PutDoc structure.
type PutDocReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDocReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutDocOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDocOK creates a PutDocOK with default headers values
func NewPutDocOK() *PutDocOK {
	return &PutDocOK{}
}

/*PutDocOK handles this case with default header values.

PutDocOK put doc o k
*/
type PutDocOK struct {
	Payload *models.DocstorePutDocumentResponse
}

func (o *PutDocOK) Error() string {
	return fmt.Sprintf("[PUT /docstore/{StoreID}/{DocumentID}][%d] putDocOK  %+v", 200, o.Payload)
}

func (o *PutDocOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocstorePutDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

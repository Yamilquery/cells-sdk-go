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

// ListPeerFoldersReader is a Reader for the ListPeerFolders structure.
type ListPeerFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPeerFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListPeerFoldersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPeerFoldersOK creates a ListPeerFoldersOK with default headers values
func NewListPeerFoldersOK() *ListPeerFoldersOK {
	return &ListPeerFoldersOK{}
}

/*ListPeerFoldersOK handles this case with default header values.

ListPeerFoldersOK list peer folders o k
*/
type ListPeerFoldersOK struct {
	Payload *models.RestNodesCollection
}

func (o *ListPeerFoldersOK) Error() string {
	return fmt.Sprintf("[POST /config/peers/{PeerAddress}][%d] listPeerFoldersOK  %+v", 200, o.Payload)
}

func (o *ListPeerFoldersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestNodesCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

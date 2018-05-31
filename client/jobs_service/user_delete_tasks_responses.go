// Code generated by go-swagger; DO NOT EDIT.

package jobs_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// UserDeleteTasksReader is a Reader for the UserDeleteTasks structure.
type UserDeleteTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserDeleteTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserDeleteTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserDeleteTasksOK creates a UserDeleteTasksOK with default headers values
func NewUserDeleteTasksOK() *UserDeleteTasksOK {
	return &UserDeleteTasksOK{}
}

/*UserDeleteTasksOK handles this case with default header values.

UserDeleteTasksOK user delete tasks o k
*/
type UserDeleteTasksOK struct {
	Payload *models.JobsDeleteTasksResponse
}

func (o *UserDeleteTasksOK) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksOK  %+v", 200, o.Payload)
}

func (o *UserDeleteTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobsDeleteTasksResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

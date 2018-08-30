// Code generated by go-swagger; DO NOT EDIT.

package scheduler_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// NewPutJobParams creates a new PutJobParams object
// with the default values initialized.
func NewPutJobParams() *PutJobParams {
	var ()
	return &PutJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutJobParamsWithTimeout creates a new PutJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutJobParamsWithTimeout(timeout time.Duration) *PutJobParams {
	var ()
	return &PutJobParams{

		timeout: timeout,
	}
}

// NewPutJobParamsWithContext creates a new PutJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutJobParamsWithContext(ctx context.Context) *PutJobParams {
	var ()
	return &PutJobParams{

		Context: ctx,
	}
}

// NewPutJobParamsWithHTTPClient creates a new PutJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutJobParamsWithHTTPClient(client *http.Client) *PutJobParams {
	var ()
	return &PutJobParams{
		HTTPClient: client,
	}
}

/*PutJobParams contains all the parameters to send to the API endpoint
for the put job operation typically these are written to a http.Request
*/
type PutJobParams struct {

	/*Body*/
	Body *models.JobsPutJobRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put job params
func (o *PutJobParams) WithTimeout(timeout time.Duration) *PutJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put job params
func (o *PutJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put job params
func (o *PutJobParams) WithContext(ctx context.Context) *PutJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put job params
func (o *PutJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put job params
func (o *PutJobParams) WithHTTPClient(client *http.Client) *PutJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put job params
func (o *PutJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put job params
func (o *PutJobParams) WithBody(body *models.JobsPutJobRequest) *PutJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put job params
func (o *PutJobParams) SetBody(body *models.JobsPutJobRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
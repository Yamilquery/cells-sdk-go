// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

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

// NewFrontLogParams creates a new FrontLogParams object
// with the default values initialized.
func NewFrontLogParams() *FrontLogParams {
	var ()
	return &FrontLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFrontLogParamsWithTimeout creates a new FrontLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFrontLogParamsWithTimeout(timeout time.Duration) *FrontLogParams {
	var ()
	return &FrontLogParams{

		timeout: timeout,
	}
}

// NewFrontLogParamsWithContext creates a new FrontLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewFrontLogParamsWithContext(ctx context.Context) *FrontLogParams {
	var ()
	return &FrontLogParams{

		Context: ctx,
	}
}

// NewFrontLogParamsWithHTTPClient creates a new FrontLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFrontLogParamsWithHTTPClient(client *http.Client) *FrontLogParams {
	var ()
	return &FrontLogParams{
		HTTPClient: client,
	}
}

/*FrontLogParams contains all the parameters to send to the API endpoint
for the front log operation typically these are written to a http.Request
*/
type FrontLogParams struct {

	/*Body*/
	Body *models.RestFrontLogMessage

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the front log params
func (o *FrontLogParams) WithTimeout(timeout time.Duration) *FrontLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the front log params
func (o *FrontLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the front log params
func (o *FrontLogParams) WithContext(ctx context.Context) *FrontLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the front log params
func (o *FrontLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the front log params
func (o *FrontLogParams) WithHTTPClient(client *http.Client) *FrontLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the front log params
func (o *FrontLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the front log params
func (o *FrontLogParams) WithBody(body *models.RestFrontLogMessage) *FrontLogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the front log params
func (o *FrontLogParams) SetBody(body *models.RestFrontLogMessage) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FrontLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

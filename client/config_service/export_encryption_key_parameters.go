// Code generated by go-swagger; DO NOT EDIT.

package config_service

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

// NewExportEncryptionKeyParams creates a new ExportEncryptionKeyParams object
// with the default values initialized.
func NewExportEncryptionKeyParams() *ExportEncryptionKeyParams {
	var ()
	return &ExportEncryptionKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExportEncryptionKeyParamsWithTimeout creates a new ExportEncryptionKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportEncryptionKeyParamsWithTimeout(timeout time.Duration) *ExportEncryptionKeyParams {
	var ()
	return &ExportEncryptionKeyParams{

		timeout: timeout,
	}
}

// NewExportEncryptionKeyParamsWithContext creates a new ExportEncryptionKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportEncryptionKeyParamsWithContext(ctx context.Context) *ExportEncryptionKeyParams {
	var ()
	return &ExportEncryptionKeyParams{

		Context: ctx,
	}
}

// NewExportEncryptionKeyParamsWithHTTPClient creates a new ExportEncryptionKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportEncryptionKeyParamsWithHTTPClient(client *http.Client) *ExportEncryptionKeyParams {
	var ()
	return &ExportEncryptionKeyParams{
		HTTPClient: client,
	}
}

/*ExportEncryptionKeyParams contains all the parameters to send to the API endpoint
for the export encryption key operation typically these are written to a http.Request
*/
type ExportEncryptionKeyParams struct {

	/*Body*/
	Body *models.EncryptionAdminExportKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the export encryption key params
func (o *ExportEncryptionKeyParams) WithTimeout(timeout time.Duration) *ExportEncryptionKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export encryption key params
func (o *ExportEncryptionKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export encryption key params
func (o *ExportEncryptionKeyParams) WithContext(ctx context.Context) *ExportEncryptionKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export encryption key params
func (o *ExportEncryptionKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export encryption key params
func (o *ExportEncryptionKeyParams) WithHTTPClient(client *http.Client) *ExportEncryptionKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export encryption key params
func (o *ExportEncryptionKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the export encryption key params
func (o *ExportEncryptionKeyParams) WithBody(body *models.EncryptionAdminExportKeyRequest) *ExportEncryptionKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the export encryption key params
func (o *ExportEncryptionKeyParams) SetBody(body *models.EncryptionAdminExportKeyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ExportEncryptionKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

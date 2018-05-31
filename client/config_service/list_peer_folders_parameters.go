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

// NewListPeerFoldersParams creates a new ListPeerFoldersParams object
// with the default values initialized.
func NewListPeerFoldersParams() *ListPeerFoldersParams {
	var ()
	return &ListPeerFoldersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPeerFoldersParamsWithTimeout creates a new ListPeerFoldersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPeerFoldersParamsWithTimeout(timeout time.Duration) *ListPeerFoldersParams {
	var ()
	return &ListPeerFoldersParams{

		timeout: timeout,
	}
}

// NewListPeerFoldersParamsWithContext creates a new ListPeerFoldersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPeerFoldersParamsWithContext(ctx context.Context) *ListPeerFoldersParams {
	var ()
	return &ListPeerFoldersParams{

		Context: ctx,
	}
}

// NewListPeerFoldersParamsWithHTTPClient creates a new ListPeerFoldersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPeerFoldersParamsWithHTTPClient(client *http.Client) *ListPeerFoldersParams {
	var ()
	return &ListPeerFoldersParams{
		HTTPClient: client,
	}
}

/*ListPeerFoldersParams contains all the parameters to send to the API endpoint
for the list peer folders operation typically these are written to a http.Request
*/
type ListPeerFoldersParams struct {

	/*PeerAddress*/
	PeerAddress string
	/*Body*/
	Body *models.RestListPeerFoldersRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list peer folders params
func (o *ListPeerFoldersParams) WithTimeout(timeout time.Duration) *ListPeerFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list peer folders params
func (o *ListPeerFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list peer folders params
func (o *ListPeerFoldersParams) WithContext(ctx context.Context) *ListPeerFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list peer folders params
func (o *ListPeerFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list peer folders params
func (o *ListPeerFoldersParams) WithHTTPClient(client *http.Client) *ListPeerFoldersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list peer folders params
func (o *ListPeerFoldersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeerAddress adds the peerAddress to the list peer folders params
func (o *ListPeerFoldersParams) WithPeerAddress(peerAddress string) *ListPeerFoldersParams {
	o.SetPeerAddress(peerAddress)
	return o
}

// SetPeerAddress adds the peerAddress to the list peer folders params
func (o *ListPeerFoldersParams) SetPeerAddress(peerAddress string) {
	o.PeerAddress = peerAddress
}

// WithBody adds the body to the list peer folders params
func (o *ListPeerFoldersParams) WithBody(body *models.RestListPeerFoldersRequest) *ListPeerFoldersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list peer folders params
func (o *ListPeerFoldersParams) SetBody(body *models.RestListPeerFoldersRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListPeerFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param PeerAddress
	if err := r.SetPathParam("PeerAddress", o.PeerAddress); err != nil {
		return err
	}

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

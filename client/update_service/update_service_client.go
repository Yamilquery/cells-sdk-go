// Code generated by go-swagger; DO NOT EDIT.

package update_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new update service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for update service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ApplyUpdate applies an update to a given version
*/
func (a *Client) ApplyUpdate(params *ApplyUpdateParams) (*ApplyUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplyUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApplyUpdate",
		Method:             "GET",
		PathPattern:        "/update/{TargetVersion}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ApplyUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ApplyUpdateOK), nil

}

/*
UpdateRequired checks the remote server to see if there are available binaries
*/
func (a *Client) UpdateRequired(params *UpdateRequiredParams) (*UpdateRequiredOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRequiredParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateRequired",
		Method:             "GET",
		PathPattern:        "/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &UpdateRequiredReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateRequiredOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

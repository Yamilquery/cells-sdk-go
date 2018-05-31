// Code generated by go-swagger; DO NOT EDIT.

package user_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
BindUser binds a user with her login and password
*/
func (a *Client) BindUser(params *BindUserParams) (*BindUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBindUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BindUser",
		Method:             "POST",
		PathPattern:        "/user/{Login}/bind",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &BindUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BindUserOK), nil

}

/*
DeleteUser deletes a user
*/
func (a *Client) DeleteUser(params *DeleteUserParams) (*DeleteUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUser",
		Method:             "DELETE",
		PathPattern:        "/user/{Login}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserOK), nil

}

/*
GetUser gets a user by login
*/
func (a *Client) GetUser(params *GetUserParams) (*GetUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUser",
		Method:             "GET",
		PathPattern:        "/user/{Login}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &GetUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserOK), nil

}

/*
PutRoles justs save a user roles without other datas
*/
func (a *Client) PutRoles(params *PutRolesParams) (*PutRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutRoles",
		Method:             "PUT",
		PathPattern:        "/user/roles/{Login}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PutRolesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutRolesOK), nil

}

/*
PutUser creates or update a user
*/
func (a *Client) PutUser(params *PutUserParams) (*PutUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutUser",
		Method:             "PUT",
		PathPattern:        "/user/{Login}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PutUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUserOK), nil

}

/*
SearchUsers lists search users
*/
func (a *Client) SearchUsers(params *SearchUsersParams) (*SearchUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SearchUsers",
		Method:             "POST",
		PathPattern:        "/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &SearchUsersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchUsersOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

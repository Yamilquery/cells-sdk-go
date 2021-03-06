// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new config service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for config service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ConfigFormsDiscovery publishes forms definition for building screens in frontend
*/
func (a *Client) ConfigFormsDiscovery(params *ConfigFormsDiscoveryParams) (*ConfigFormsDiscoveryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigFormsDiscoveryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ConfigFormsDiscovery",
		Method:             "GET",
		PathPattern:        "/config/discovery/forms/{ServiceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ConfigFormsDiscoveryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ConfigFormsDiscoveryOK), nil

}

/*
ControlService nots implemented start stop a service
*/
func (a *Client) ControlService(params *ControlServiceParams) (*ControlServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewControlServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ControlService",
		Method:             "POST",
		PathPattern:        "/config/ctl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ControlServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ControlServiceOK), nil

}

/*
CreateEncryptionKey creates a new master key
*/
func (a *Client) CreateEncryptionKey(params *CreateEncryptionKeyParams) (*CreateEncryptionKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEncryptionKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateEncryptionKey",
		Method:             "POST",
		PathPattern:        "/config/encryption/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &CreateEncryptionKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateEncryptionKeyOK), nil

}

/*
DeleteDataSource deletes a datasource
*/
func (a *Client) DeleteDataSource(params *DeleteDataSourceParams) (*DeleteDataSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDataSourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteDataSource",
		Method:             "DELETE",
		PathPattern:        "/config/datasource/{Name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &DeleteDataSourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDataSourceOK), nil

}

/*
DeleteEncryptionKey deletes an existing master key
*/
func (a *Client) DeleteEncryptionKey(params *DeleteEncryptionKeyParams) (*DeleteEncryptionKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEncryptionKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteEncryptionKey",
		Method:             "POST",
		PathPattern:        "/config/encryption/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &DeleteEncryptionKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEncryptionKeyOK), nil

}

/*
EndpointsDiscovery publishes available endpoints
*/
func (a *Client) EndpointsDiscovery(params *EndpointsDiscoveryParams) (*EndpointsDiscoveryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEndpointsDiscoveryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EndpointsDiscovery",
		Method:             "GET",
		PathPattern:        "/config/discovery",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &EndpointsDiscoveryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EndpointsDiscoveryOK), nil

}

/*
ExportEncryptionKey exports a master key for backup purpose protected with a password
*/
func (a *Client) ExportEncryptionKey(params *ExportEncryptionKeyParams) (*ExportEncryptionKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportEncryptionKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ExportEncryptionKey",
		Method:             "POST",
		PathPattern:        "/config/encryption/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ExportEncryptionKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExportEncryptionKeyOK), nil

}

/*
GetConfig generics config get using a full path in the config tree
*/
func (a *Client) GetConfig(params *GetConfigParams) (*GetConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetConfig",
		Method:             "GET",
		PathPattern:        "/config/{FullPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &GetConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigOK), nil

}

/*
GetDataSource loads datasource information
*/
func (a *Client) GetDataSource(params *GetDataSourceParams) (*GetDataSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataSourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDataSource",
		Method:             "GET",
		PathPattern:        "/config/datasource/{Name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &GetDataSourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDataSourceOK), nil

}

/*
GetVersioningPolicy loads a given versioning policy
*/
func (a *Client) GetVersioningPolicy(params *GetVersioningPolicyParams) (*GetVersioningPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersioningPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVersioningPolicy",
		Method:             "GET",
		PathPattern:        "/config/versioning/{Uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &GetVersioningPolicyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVersioningPolicyOK), nil

}

/*
ImportEncryptionKey imports a previously exported master key requires the password created at export time
*/
func (a *Client) ImportEncryptionKey(params *ImportEncryptionKeyParams) (*ImportEncryptionKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportEncryptionKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ImportEncryptionKey",
		Method:             "PUT",
		PathPattern:        "/config/encryption/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ImportEncryptionKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ImportEncryptionKeyOK), nil

}

/*
ListDataSources lists all defined datasources
*/
func (a *Client) ListDataSources(params *ListDataSourcesParams) (*ListDataSourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDataSourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListDataSources",
		Method:             "GET",
		PathPattern:        "/config/datasource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListDataSourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDataSourcesOK), nil

}

/*
ListEncryptionKeys lists registered master keys
*/
func (a *Client) ListEncryptionKeys(params *ListEncryptionKeysParams) (*ListEncryptionKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEncryptionKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListEncryptionKeys",
		Method:             "POST",
		PathPattern:        "/config/encryption/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListEncryptionKeysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListEncryptionKeysOK), nil

}

/*
ListPeerFolders lists folders on a peer starting from root
*/
func (a *Client) ListPeerFolders(params *ListPeerFoldersParams) (*ListPeerFoldersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPeerFoldersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPeerFolders",
		Method:             "POST",
		PathPattern:        "/config/peers/{PeerAddress}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListPeerFoldersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPeerFoldersOK), nil

}

/*
ListPeersAddresses lists all detected peers servers on which the app is running
*/
func (a *Client) ListPeersAddresses(params *ListPeersAddressesParams) (*ListPeersAddressesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPeersAddressesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPeersAddresses",
		Method:             "GET",
		PathPattern:        "/config/peers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListPeersAddressesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPeersAddressesOK), nil

}

/*
ListServices lists all services and their status
*/
func (a *Client) ListServices(params *ListServicesParams) (*ListServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListServices",
		Method:             "GET",
		PathPattern:        "/config/ctl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListServicesOK), nil

}

/*
ListVersioningPolicies lists all defined versioning policies
*/
func (a *Client) ListVersioningPolicies(params *ListVersioningPoliciesParams) (*ListVersioningPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersioningPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVersioningPolicies",
		Method:             "GET",
		PathPattern:        "/config/versioning",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListVersioningPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListVersioningPoliciesOK), nil

}

/*
ListVirtualNodes lists all defined virtual nodes
*/
func (a *Client) ListVirtualNodes(params *ListVirtualNodesParams) (*ListVirtualNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVirtualNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListVirtualNodes",
		Method:             "GET",
		PathPattern:        "/config/virtualnodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListVirtualNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListVirtualNodesOK), nil

}

/*
OpenAPIDiscovery publishes available r e s t apis
*/
func (a *Client) OpenAPIDiscovery(params *OpenAPIDiscoveryParams) (*OpenAPIDiscoveryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpenAPIDiscoveryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "OpenApiDiscovery",
		Method:             "GET",
		PathPattern:        "/config/discovery/openapi",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &OpenAPIDiscoveryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OpenAPIDiscoveryOK), nil

}

/*
PutConfig generics config put using a full path in the config tree
*/
func (a *Client) PutConfig(params *PutConfigParams) (*PutConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutConfig",
		Method:             "PUT",
		PathPattern:        "/config/{FullPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PutConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutConfigOK), nil

}

/*
PutDataSource creates or update a datasource
*/
func (a *Client) PutDataSource(params *PutDataSourceParams) (*PutDataSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDataSourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutDataSource",
		Method:             "POST",
		PathPattern:        "/config/datasource/{Name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PutDataSourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutDataSourceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

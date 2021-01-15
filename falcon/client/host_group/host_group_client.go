// Code generated by go-swagger; DO NOT EDIT.

package host_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new host group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for host group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateHostGroups(params *CreateHostGroupsParams) (*CreateHostGroupsCreated, error)

	DeleteHostGroups(params *DeleteHostGroupsParams) (*DeleteHostGroupsOK, error)

	GetHostGroups(params *GetHostGroupsParams) (*GetHostGroupsOK, error)

	PerformGroupAction(params *PerformGroupActionParams) (*PerformGroupActionOK, error)

	QueryCombinedGroupMembers(params *QueryCombinedGroupMembersParams) (*QueryCombinedGroupMembersOK, error)

	QueryCombinedHostGroups(params *QueryCombinedHostGroupsParams) (*QueryCombinedHostGroupsOK, error)

	QueryGroupMembers(params *QueryGroupMembersParams) (*QueryGroupMembersOK, error)

	QueryHostGroups(params *QueryHostGroupsParams) (*QueryHostGroupsOK, error)

	UpdateHostGroups(params *UpdateHostGroupsParams) (*UpdateHostGroupsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateHostGroups creates host groups by specifying details about the group to create
*/
func (a *Client) CreateHostGroups(params *CreateHostGroupsParams) (*CreateHostGroupsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateHostGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createHostGroups",
		Method:             "POST",
		PathPattern:        "/devices/entities/host-groups/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateHostGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateHostGroupsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createHostGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteHostGroups deletes a set of host groups by specifying their i ds
*/
func (a *Client) DeleteHostGroups(params *DeleteHostGroupsParams) (*DeleteHostGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHostGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteHostGroups",
		Method:             "DELETE",
		PathPattern:        "/devices/entities/host-groups/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteHostGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteHostGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteHostGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetHostGroups retrieves a set of host groups by specifying their i ds
*/
func (a *Client) GetHostGroups(params *GetHostGroupsParams) (*GetHostGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHostGroups",
		Method:             "GET",
		PathPattern:        "/devices/entities/host-groups/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHostGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHostGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHostGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PerformGroupAction performs the specified action on the host groups specified in the request
*/
func (a *Client) PerformGroupAction(params *PerformGroupActionParams) (*PerformGroupActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformGroupActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "performGroupAction",
		Method:             "POST",
		PathPattern:        "/devices/entities/host-group-actions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PerformGroupActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PerformGroupActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PerformGroupActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryCombinedGroupMembers searches for members of a host group in your environment by providing an f q l filter and paging details returns a set of host details which match the filter criteria
*/
func (a *Client) QueryCombinedGroupMembers(params *QueryCombinedGroupMembersParams) (*QueryCombinedGroupMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedGroupMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryCombinedGroupMembers",
		Method:             "GET",
		PathPattern:        "/devices/combined/host-group-members/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedGroupMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryCombinedGroupMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedGroupMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryCombinedHostGroups searches for host groups in your environment by providing an f q l filter and paging details returns a set of host groups which match the filter criteria
*/
func (a *Client) QueryCombinedHostGroups(params *QueryCombinedHostGroupsParams) (*QueryCombinedHostGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryCombinedHostGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryCombinedHostGroups",
		Method:             "GET",
		PathPattern:        "/devices/combined/host-groups/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryCombinedHostGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryCombinedHostGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryCombinedHostGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryGroupMembers searches for members of a host group in your environment by providing an f q l filter and paging details returns a set of agent i ds which match the filter criteria
*/
func (a *Client) QueryGroupMembers(params *QueryGroupMembersParams) (*QueryGroupMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryGroupMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryGroupMembers",
		Method:             "GET",
		PathPattern:        "/devices/queries/host-group-members/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryGroupMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryGroupMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryGroupMembersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  QueryHostGroups searches for host groups in your environment by providing an f q l filter and paging details returns a set of host group i ds which match the filter criteria
*/
func (a *Client) QueryHostGroups(params *QueryHostGroupsParams) (*QueryHostGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryHostGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryHostGroups",
		Method:             "GET",
		PathPattern:        "/devices/queries/host-groups/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryHostGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryHostGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QueryHostGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateHostGroups updates host groups by specifying the ID of the group and details to update
*/
func (a *Client) UpdateHostGroups(params *UpdateHostGroupsParams) (*UpdateHostGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateHostGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateHostGroups",
		Method:             "PATCH",
		PathPattern:        "/devices/entities/host-groups/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateHostGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateHostGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateHostGroupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
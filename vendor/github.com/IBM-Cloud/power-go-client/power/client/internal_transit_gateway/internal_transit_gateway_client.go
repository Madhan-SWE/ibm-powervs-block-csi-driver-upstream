// Code generated by go-swagger; DO NOT EDIT.

package internal_transit_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new internal transit gateway API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for internal transit gateway API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	InternalV1TransitgatewayGet(params *InternalV1TransitgatewayGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1TransitgatewayGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  InternalV1TransitgatewayGet gets the cloud instance transit gateway information
*/
func (a *Client) InternalV1TransitgatewayGet(params *InternalV1TransitgatewayGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1TransitgatewayGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1TransitgatewayGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.transitgateway.get",
		Method:             "GET",
		PathPattern:        "/internal/v1/transit-gateway/{powervs_service_crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1TransitgatewayGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InternalV1TransitgatewayGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.transitgateway.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

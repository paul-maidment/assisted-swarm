// Code generated by go-swagger; DO NOT EDIT.

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteAgentParams creates a new DeleteAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAgentParams() *DeleteAgentParams {
	return &DeleteAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAgentParamsWithTimeout creates a new DeleteAgentParams object
// with the ability to set a timeout on a request.
func NewDeleteAgentParamsWithTimeout(timeout time.Duration) *DeleteAgentParams {
	return &DeleteAgentParams{
		timeout: timeout,
	}
}

// NewDeleteAgentParamsWithContext creates a new DeleteAgentParams object
// with the ability to set a context for a request.
func NewDeleteAgentParamsWithContext(ctx context.Context) *DeleteAgentParams {
	return &DeleteAgentParams{
		Context: ctx,
	}
}

// NewDeleteAgentParamsWithHTTPClient creates a new DeleteAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAgentParamsWithHTTPClient(client *http.Client) *DeleteAgentParams {
	return &DeleteAgentParams{
		HTTPClient: client,
	}
}

/* DeleteAgentParams contains all the parameters to send to the API endpoint
   for the delete agent operation.

   Typically these are written to a http.Request.
*/
type DeleteAgentParams struct {

	// AgentID.
	AgentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentParams) WithDefaults() *DeleteAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete agent params
func (o *DeleteAgentParams) WithTimeout(timeout time.Duration) *DeleteAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete agent params
func (o *DeleteAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete agent params
func (o *DeleteAgentParams) WithContext(ctx context.Context) *DeleteAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete agent params
func (o *DeleteAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete agent params
func (o *DeleteAgentParams) WithHTTPClient(client *http.Client) *DeleteAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete agent params
func (o *DeleteAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentID adds the agentID to the delete agent params
func (o *DeleteAgentParams) WithAgentID(agentID int64) *DeleteAgentParams {
	o.SetAgentID(agentID)
	return o
}

// SetAgentID adds the agentId to the delete agent params
func (o *DeleteAgentParams) SetAgentID(agentID int64) {
	o.AgentID = agentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agent_id
	if err := r.SetPathParam("agent_id", swag.FormatInt64(o.AgentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

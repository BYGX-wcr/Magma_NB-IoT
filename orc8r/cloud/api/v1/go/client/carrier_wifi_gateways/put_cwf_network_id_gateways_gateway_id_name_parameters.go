// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutCwfNetworkIDGatewaysGatewayIDNameParams creates a new PutCwfNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized.
func NewPutCwfNetworkIDGatewaysGatewayIDNameParams() *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDNameParamsWithTimeout creates a new PutCwfNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCwfNetworkIDGatewaysGatewayIDNameParamsWithTimeout(timeout time.Duration) *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDNameParams{

		timeout: timeout,
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDNameParamsWithContext creates a new PutCwfNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCwfNetworkIDGatewaysGatewayIDNameParamsWithContext(ctx context.Context) *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDNameParams{

		Context: ctx,
	}
}

// NewPutCwfNetworkIDGatewaysGatewayIDNameParamsWithHTTPClient creates a new PutCwfNetworkIDGatewaysGatewayIDNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCwfNetworkIDGatewaysGatewayIDNameParamsWithHTTPClient(client *http.Client) *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	var ()
	return &PutCwfNetworkIDGatewaysGatewayIDNameParams{
		HTTPClient: client,
	}
}

/*PutCwfNetworkIDGatewaysGatewayIDNameParams contains all the parameters to send to the API endpoint
for the put cwf network ID gateways gateway ID name operation typically these are written to a http.Request
*/
type PutCwfNetworkIDGatewaysGatewayIDNameParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*Name*/
	Name models.GatewayName
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) WithTimeout(timeout time.Duration) *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) WithContext(ctx context.Context) *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) WithHTTPClient(client *http.Client) *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) WithGatewayID(gatewayID string) *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithName adds the name to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) WithName(name models.GatewayName) *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) SetName(name models.GatewayName) {
	o.Name = name
}

// WithNetworkID adds the networkID to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) WithNetworkID(networkID string) *PutCwfNetworkIDGatewaysGatewayIDNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put cwf network ID gateways gateway ID name params
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCwfNetworkIDGatewaysGatewayIDNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Name); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

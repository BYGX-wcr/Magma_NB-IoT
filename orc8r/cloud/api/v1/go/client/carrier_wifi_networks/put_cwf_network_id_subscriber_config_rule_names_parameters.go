// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

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

// NewPutCwfNetworkIDSubscriberConfigRuleNamesParams creates a new PutCwfNetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized.
func NewPutCwfNetworkIDSubscriberConfigRuleNamesParams() *PutCwfNetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &PutCwfNetworkIDSubscriberConfigRuleNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCwfNetworkIDSubscriberConfigRuleNamesParamsWithTimeout creates a new PutCwfNetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCwfNetworkIDSubscriberConfigRuleNamesParamsWithTimeout(timeout time.Duration) *PutCwfNetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &PutCwfNetworkIDSubscriberConfigRuleNamesParams{

		timeout: timeout,
	}
}

// NewPutCwfNetworkIDSubscriberConfigRuleNamesParamsWithContext creates a new PutCwfNetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCwfNetworkIDSubscriberConfigRuleNamesParamsWithContext(ctx context.Context) *PutCwfNetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &PutCwfNetworkIDSubscriberConfigRuleNamesParams{

		Context: ctx,
	}
}

// NewPutCwfNetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient creates a new PutCwfNetworkIDSubscriberConfigRuleNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutCwfNetworkIDSubscriberConfigRuleNamesParamsWithHTTPClient(client *http.Client) *PutCwfNetworkIDSubscriberConfigRuleNamesParams {
	var ()
	return &PutCwfNetworkIDSubscriberConfigRuleNamesParams{
		HTTPClient: client,
	}
}

/*PutCwfNetworkIDSubscriberConfigRuleNamesParams contains all the parameters to send to the API endpoint
for the put cwf network ID subscriber config rule names operation typically these are written to a http.Request
*/
type PutCwfNetworkIDSubscriberConfigRuleNamesParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Record
	  Subscriber Config for the Network

	*/
	Record models.RuleNames

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) WithTimeout(timeout time.Duration) *PutCwfNetworkIDSubscriberConfigRuleNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) WithContext(ctx context.Context) *PutCwfNetworkIDSubscriberConfigRuleNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) WithHTTPClient(client *http.Client) *PutCwfNetworkIDSubscriberConfigRuleNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) WithNetworkID(networkID string) *PutCwfNetworkIDSubscriberConfigRuleNamesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecord adds the record to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) WithRecord(record models.RuleNames) *PutCwfNetworkIDSubscriberConfigRuleNamesParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the put cwf network ID subscriber config rule names params
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) SetRecord(record models.RuleNames) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PutCwfNetworkIDSubscriberConfigRuleNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Record != nil {
		if err := r.SetBodyParam(o.Record); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

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

// NewPostFegParams creates a new PostFegParams object
// with the default values initialized.
func NewPostFegParams() *PostFegParams {
	var ()
	return &PostFegParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFegParamsWithTimeout creates a new PostFegParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFegParamsWithTimeout(timeout time.Duration) *PostFegParams {
	var ()
	return &PostFegParams{

		timeout: timeout,
	}
}

// NewPostFegParamsWithContext creates a new PostFegParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFegParamsWithContext(ctx context.Context) *PostFegParams {
	var ()
	return &PostFegParams{

		Context: ctx,
	}
}

// NewPostFegParamsWithHTTPClient creates a new PostFegParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFegParamsWithHTTPClient(client *http.Client) *PostFegParams {
	var ()
	return &PostFegParams{
		HTTPClient: client,
	}
}

/*PostFegParams contains all the parameters to send to the API endpoint
for the post feg operation typically these are written to a http.Request
*/
type PostFegParams struct {

	/*FegNetwork
	  Configuration of the network to create

	*/
	FegNetwork *models.FegNetwork

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post feg params
func (o *PostFegParams) WithTimeout(timeout time.Duration) *PostFegParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post feg params
func (o *PostFegParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post feg params
func (o *PostFegParams) WithContext(ctx context.Context) *PostFegParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post feg params
func (o *PostFegParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post feg params
func (o *PostFegParams) WithHTTPClient(client *http.Client) *PostFegParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post feg params
func (o *PostFegParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFegNetwork adds the fegNetwork to the post feg params
func (o *PostFegParams) WithFegNetwork(fegNetwork *models.FegNetwork) *PostFegParams {
	o.SetFegNetwork(fegNetwork)
	return o
}

// SetFegNetwork adds the fegNetwork to the post feg params
func (o *PostFegParams) SetFegNetwork(fegNetwork *models.FegNetwork) {
	o.FegNetwork = fegNetwork
}

// WriteToRequest writes these params to a swagger request
func (o *PostFegParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FegNetwork != nil {
		if err := r.SetBodyParam(o.FegNetwork); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

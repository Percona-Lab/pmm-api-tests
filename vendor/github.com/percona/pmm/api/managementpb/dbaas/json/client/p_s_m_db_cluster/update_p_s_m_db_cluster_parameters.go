// Code generated by go-swagger; DO NOT EDIT.

package p_s_m_db_cluster

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
)

// NewUpdatePSMDBClusterParams creates a new UpdatePSMDBClusterParams object
// with the default values initialized.
func NewUpdatePSMDBClusterParams() *UpdatePSMDBClusterParams {
	var ()
	return &UpdatePSMDBClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePSMDBClusterParamsWithTimeout creates a new UpdatePSMDBClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePSMDBClusterParamsWithTimeout(timeout time.Duration) *UpdatePSMDBClusterParams {
	var ()
	return &UpdatePSMDBClusterParams{

		timeout: timeout,
	}
}

// NewUpdatePSMDBClusterParamsWithContext creates a new UpdatePSMDBClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePSMDBClusterParamsWithContext(ctx context.Context) *UpdatePSMDBClusterParams {
	var ()
	return &UpdatePSMDBClusterParams{

		Context: ctx,
	}
}

// NewUpdatePSMDBClusterParamsWithHTTPClient creates a new UpdatePSMDBClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePSMDBClusterParamsWithHTTPClient(client *http.Client) *UpdatePSMDBClusterParams {
	var ()
	return &UpdatePSMDBClusterParams{
		HTTPClient: client,
	}
}

/*UpdatePSMDBClusterParams contains all the parameters to send to the API endpoint
for the update p s m DB cluster operation typically these are written to a http.Request
*/
type UpdatePSMDBClusterParams struct {

	/*Body*/
	Body UpdatePSMDBClusterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update p s m DB cluster params
func (o *UpdatePSMDBClusterParams) WithTimeout(timeout time.Duration) *UpdatePSMDBClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update p s m DB cluster params
func (o *UpdatePSMDBClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update p s m DB cluster params
func (o *UpdatePSMDBClusterParams) WithContext(ctx context.Context) *UpdatePSMDBClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update p s m DB cluster params
func (o *UpdatePSMDBClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update p s m DB cluster params
func (o *UpdatePSMDBClusterParams) WithHTTPClient(client *http.Client) *UpdatePSMDBClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update p s m DB cluster params
func (o *UpdatePSMDBClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update p s m DB cluster params
func (o *UpdatePSMDBClusterParams) WithBody(body UpdatePSMDBClusterBody) *UpdatePSMDBClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update p s m DB cluster params
func (o *UpdatePSMDBClusterParams) SetBody(body UpdatePSMDBClusterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePSMDBClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

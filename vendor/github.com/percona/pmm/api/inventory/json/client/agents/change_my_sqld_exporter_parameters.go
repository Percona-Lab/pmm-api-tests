// Code generated by go-swagger; DO NOT EDIT.

package agents

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
)

// NewChangeMySqldExporterParams creates a new ChangeMySqldExporterParams object
// with the default values initialized.
func NewChangeMySqldExporterParams() *ChangeMySqldExporterParams {
	var ()
	return &ChangeMySqldExporterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeMySqldExporterParamsWithTimeout creates a new ChangeMySqldExporterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeMySqldExporterParamsWithTimeout(timeout time.Duration) *ChangeMySqldExporterParams {
	var ()
	return &ChangeMySqldExporterParams{

		timeout: timeout,
	}
}

// NewChangeMySqldExporterParamsWithContext creates a new ChangeMySqldExporterParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeMySqldExporterParamsWithContext(ctx context.Context) *ChangeMySqldExporterParams {
	var ()
	return &ChangeMySqldExporterParams{

		Context: ctx,
	}
}

// NewChangeMySqldExporterParamsWithHTTPClient creates a new ChangeMySqldExporterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeMySqldExporterParamsWithHTTPClient(client *http.Client) *ChangeMySqldExporterParams {
	var ()
	return &ChangeMySqldExporterParams{
		HTTPClient: client,
	}
}

/*ChangeMySqldExporterParams contains all the parameters to send to the API endpoint
for the change my sqld exporter operation typically these are written to a http.Request
*/
type ChangeMySqldExporterParams struct {

	/*Body*/
	Body ChangeMySqldExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change my sqld exporter params
func (o *ChangeMySqldExporterParams) WithTimeout(timeout time.Duration) *ChangeMySqldExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change my sqld exporter params
func (o *ChangeMySqldExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change my sqld exporter params
func (o *ChangeMySqldExporterParams) WithContext(ctx context.Context) *ChangeMySqldExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change my sqld exporter params
func (o *ChangeMySqldExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change my sqld exporter params
func (o *ChangeMySqldExporterParams) WithHTTPClient(client *http.Client) *ChangeMySqldExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change my sqld exporter params
func (o *ChangeMySqldExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change my sqld exporter params
func (o *ChangeMySqldExporterParams) WithBody(body ChangeMySqldExporterBody) *ChangeMySqldExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change my sqld exporter params
func (o *ChangeMySqldExporterParams) SetBody(body ChangeMySqldExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeMySqldExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

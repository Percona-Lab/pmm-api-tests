// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewStartSecurityChecksParams creates a new StartSecurityChecksParams object
// with the default values initialized.
func NewStartSecurityChecksParams() *StartSecurityChecksParams {
	var ()
	return &StartSecurityChecksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartSecurityChecksParamsWithTimeout creates a new StartSecurityChecksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartSecurityChecksParamsWithTimeout(timeout time.Duration) *StartSecurityChecksParams {
	var ()
	return &StartSecurityChecksParams{

		timeout: timeout,
	}
}

// NewStartSecurityChecksParamsWithContext creates a new StartSecurityChecksParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartSecurityChecksParamsWithContext(ctx context.Context) *StartSecurityChecksParams {
	var ()
	return &StartSecurityChecksParams{

		Context: ctx,
	}
}

// NewStartSecurityChecksParamsWithHTTPClient creates a new StartSecurityChecksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartSecurityChecksParamsWithHTTPClient(client *http.Client) *StartSecurityChecksParams {
	var ()
	return &StartSecurityChecksParams{
		HTTPClient: client,
	}
}

/*StartSecurityChecksParams contains all the parameters to send to the API endpoint
for the start security checks operation typically these are written to a http.Request
*/
type StartSecurityChecksParams struct {

	/*Body*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start security checks params
func (o *StartSecurityChecksParams) WithTimeout(timeout time.Duration) *StartSecurityChecksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start security checks params
func (o *StartSecurityChecksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start security checks params
func (o *StartSecurityChecksParams) WithContext(ctx context.Context) *StartSecurityChecksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start security checks params
func (o *StartSecurityChecksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start security checks params
func (o *StartSecurityChecksParams) WithHTTPClient(client *http.Client) *StartSecurityChecksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start security checks params
func (o *StartSecurityChecksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start security checks params
func (o *StartSecurityChecksParams) WithBody(body interface{}) *StartSecurityChecksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start security checks params
func (o *StartSecurityChecksParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StartSecurityChecksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

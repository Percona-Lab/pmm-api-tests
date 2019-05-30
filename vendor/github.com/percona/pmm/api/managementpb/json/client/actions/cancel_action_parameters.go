// Code generated by go-swagger; DO NOT EDIT.

package actions

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

// NewCancelActionParams creates a new CancelActionParams object
// with the default values initialized.
func NewCancelActionParams() *CancelActionParams {
	var ()
	return &CancelActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelActionParamsWithTimeout creates a new CancelActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelActionParamsWithTimeout(timeout time.Duration) *CancelActionParams {
	var ()
	return &CancelActionParams{

		timeout: timeout,
	}
}

// NewCancelActionParamsWithContext creates a new CancelActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelActionParamsWithContext(ctx context.Context) *CancelActionParams {
	var ()
	return &CancelActionParams{

		Context: ctx,
	}
}

// NewCancelActionParamsWithHTTPClient creates a new CancelActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelActionParamsWithHTTPClient(client *http.Client) *CancelActionParams {
	var ()
	return &CancelActionParams{
		HTTPClient: client,
	}
}

/*CancelActionParams contains all the parameters to send to the API endpoint
for the cancel action operation typically these are written to a http.Request
*/
type CancelActionParams struct {

	/*Body*/
	Body CancelActionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel action params
func (o *CancelActionParams) WithTimeout(timeout time.Duration) *CancelActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel action params
func (o *CancelActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel action params
func (o *CancelActionParams) WithContext(ctx context.Context) *CancelActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel action params
func (o *CancelActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel action params
func (o *CancelActionParams) WithHTTPClient(client *http.Client) *CancelActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel action params
func (o *CancelActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cancel action params
func (o *CancelActionParams) WithBody(body CancelActionBody) *CancelActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cancel action params
func (o *CancelActionParams) SetBody(body CancelActionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CancelActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

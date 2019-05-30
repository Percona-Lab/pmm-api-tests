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

// NewStartMySQLExplainJSONActionParams creates a new StartMySQLExplainJSONActionParams object
// with the default values initialized.
func NewStartMySQLExplainJSONActionParams() *StartMySQLExplainJSONActionParams {
	var ()
	return &StartMySQLExplainJSONActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartMySQLExplainJSONActionParamsWithTimeout creates a new StartMySQLExplainJSONActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartMySQLExplainJSONActionParamsWithTimeout(timeout time.Duration) *StartMySQLExplainJSONActionParams {
	var ()
	return &StartMySQLExplainJSONActionParams{

		timeout: timeout,
	}
}

// NewStartMySQLExplainJSONActionParamsWithContext creates a new StartMySQLExplainJSONActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartMySQLExplainJSONActionParamsWithContext(ctx context.Context) *StartMySQLExplainJSONActionParams {
	var ()
	return &StartMySQLExplainJSONActionParams{

		Context: ctx,
	}
}

// NewStartMySQLExplainJSONActionParamsWithHTTPClient creates a new StartMySQLExplainJSONActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartMySQLExplainJSONActionParamsWithHTTPClient(client *http.Client) *StartMySQLExplainJSONActionParams {
	var ()
	return &StartMySQLExplainJSONActionParams{
		HTTPClient: client,
	}
}

/*StartMySQLExplainJSONActionParams contains all the parameters to send to the API endpoint
for the start my SQL explain JSON action operation typically these are written to a http.Request
*/
type StartMySQLExplainJSONActionParams struct {

	/*Body*/
	Body StartMySQLExplainJSONActionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start my SQL explain JSON action params
func (o *StartMySQLExplainJSONActionParams) WithTimeout(timeout time.Duration) *StartMySQLExplainJSONActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start my SQL explain JSON action params
func (o *StartMySQLExplainJSONActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start my SQL explain JSON action params
func (o *StartMySQLExplainJSONActionParams) WithContext(ctx context.Context) *StartMySQLExplainJSONActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start my SQL explain JSON action params
func (o *StartMySQLExplainJSONActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start my SQL explain JSON action params
func (o *StartMySQLExplainJSONActionParams) WithHTTPClient(client *http.Client) *StartMySQLExplainJSONActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start my SQL explain JSON action params
func (o *StartMySQLExplainJSONActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start my SQL explain JSON action params
func (o *StartMySQLExplainJSONActionParams) WithBody(body StartMySQLExplainJSONActionBody) *StartMySQLExplainJSONActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start my SQL explain JSON action params
func (o *StartMySQLExplainJSONActionParams) SetBody(body StartMySQLExplainJSONActionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StartMySQLExplainJSONActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

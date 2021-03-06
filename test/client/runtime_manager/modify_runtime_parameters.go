// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewModifyRuntimeParams creates a new ModifyRuntimeParams object
// with the default values initialized.
func NewModifyRuntimeParams() *ModifyRuntimeParams {
	var ()
	return &ModifyRuntimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyRuntimeParamsWithTimeout creates a new ModifyRuntimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyRuntimeParamsWithTimeout(timeout time.Duration) *ModifyRuntimeParams {
	var ()
	return &ModifyRuntimeParams{

		timeout: timeout,
	}
}

// NewModifyRuntimeParamsWithContext creates a new ModifyRuntimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyRuntimeParamsWithContext(ctx context.Context) *ModifyRuntimeParams {
	var ()
	return &ModifyRuntimeParams{

		Context: ctx,
	}
}

// NewModifyRuntimeParamsWithHTTPClient creates a new ModifyRuntimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyRuntimeParamsWithHTTPClient(client *http.Client) *ModifyRuntimeParams {
	var ()
	return &ModifyRuntimeParams{
		HTTPClient: client,
	}
}

/*ModifyRuntimeParams contains all the parameters to send to the API endpoint
for the modify runtime operation typically these are written to a http.Request
*/
type ModifyRuntimeParams struct {

	/*Body*/
	Body *models.OpenpitrixModifyRuntimeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify runtime params
func (o *ModifyRuntimeParams) WithTimeout(timeout time.Duration) *ModifyRuntimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify runtime params
func (o *ModifyRuntimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify runtime params
func (o *ModifyRuntimeParams) WithContext(ctx context.Context) *ModifyRuntimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify runtime params
func (o *ModifyRuntimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify runtime params
func (o *ModifyRuntimeParams) WithHTTPClient(client *http.Client) *ModifyRuntimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify runtime params
func (o *ModifyRuntimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify runtime params
func (o *ModifyRuntimeParams) WithBody(body *models.OpenpitrixModifyRuntimeRequest) *ModifyRuntimeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify runtime params
func (o *ModifyRuntimeParams) SetBody(body *models.OpenpitrixModifyRuntimeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyRuntimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

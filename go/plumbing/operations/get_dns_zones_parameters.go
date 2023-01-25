// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetDNSZonesParams creates a new GetDNSZonesParams object
// with the default values initialized.
func NewGetDNSZonesParams() *GetDNSZonesParams {
	var ()
	return &GetDNSZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDNSZonesParamsWithTimeout creates a new GetDNSZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDNSZonesParamsWithTimeout(timeout time.Duration) *GetDNSZonesParams {
	var ()
	return &GetDNSZonesParams{

		timeout: timeout,
	}
}

// NewGetDNSZonesParamsWithContext creates a new GetDNSZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDNSZonesParamsWithContext(ctx context.Context) *GetDNSZonesParams {
	var ()
	return &GetDNSZonesParams{

		Context: ctx,
	}
}

// NewGetDNSZonesParamsWithHTTPClient creates a new GetDNSZonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDNSZonesParamsWithHTTPClient(client *http.Client) *GetDNSZonesParams {
	var ()
	return &GetDNSZonesParams{
		HTTPClient: client,
	}
}

/*
GetDNSZonesParams contains all the parameters to send to the API endpoint
for the get Dns zones operation typically these are written to a http.Request
*/
type GetDNSZonesParams struct {

	/*AccountSlug*/
	AccountSlug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Dns zones params
func (o *GetDNSZonesParams) WithTimeout(timeout time.Duration) *GetDNSZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Dns zones params
func (o *GetDNSZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Dns zones params
func (o *GetDNSZonesParams) WithContext(ctx context.Context) *GetDNSZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Dns zones params
func (o *GetDNSZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Dns zones params
func (o *GetDNSZonesParams) WithHTTPClient(client *http.Client) *GetDNSZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Dns zones params
func (o *GetDNSZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountSlug adds the accountSlug to the get Dns zones params
func (o *GetDNSZonesParams) WithAccountSlug(accountSlug *string) *GetDNSZonesParams {
	o.SetAccountSlug(accountSlug)
	return o
}

// SetAccountSlug adds the accountSlug to the get Dns zones params
func (o *GetDNSZonesParams) SetAccountSlug(accountSlug *string) {
	o.AccountSlug = accountSlug
}

// WriteToRequest writes these params to a swagger request
func (o *GetDNSZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountSlug != nil {

		// query param account_slug
		var qrAccountSlug string
		if o.AccountSlug != nil {
			qrAccountSlug = *o.AccountSlug
		}
		qAccountSlug := qrAccountSlug
		if qAccountSlug != "" {
			if err := r.SetQueryParam("account_slug", qAccountSlug); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

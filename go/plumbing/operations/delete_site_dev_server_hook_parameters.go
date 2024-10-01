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

// NewDeleteSiteDevServerHookParams creates a new DeleteSiteDevServerHookParams object
// with the default values initialized.
func NewDeleteSiteDevServerHookParams() *DeleteSiteDevServerHookParams {
	var ()
	return &DeleteSiteDevServerHookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSiteDevServerHookParamsWithTimeout creates a new DeleteSiteDevServerHookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSiteDevServerHookParamsWithTimeout(timeout time.Duration) *DeleteSiteDevServerHookParams {
	var ()
	return &DeleteSiteDevServerHookParams{

		timeout: timeout,
	}
}

// NewDeleteSiteDevServerHookParamsWithContext creates a new DeleteSiteDevServerHookParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSiteDevServerHookParamsWithContext(ctx context.Context) *DeleteSiteDevServerHookParams {
	var ()
	return &DeleteSiteDevServerHookParams{

		Context: ctx,
	}
}

// NewDeleteSiteDevServerHookParamsWithHTTPClient creates a new DeleteSiteDevServerHookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSiteDevServerHookParamsWithHTTPClient(client *http.Client) *DeleteSiteDevServerHookParams {
	var ()
	return &DeleteSiteDevServerHookParams{
		HTTPClient: client,
	}
}

/*
DeleteSiteDevServerHookParams contains all the parameters to send to the API endpoint
for the delete site dev server hook operation typically these are written to a http.Request
*/
type DeleteSiteDevServerHookParams struct {

	/*ID*/
	ID string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) WithTimeout(timeout time.Duration) *DeleteSiteDevServerHookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) WithContext(ctx context.Context) *DeleteSiteDevServerHookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) WithHTTPClient(client *http.Client) *DeleteSiteDevServerHookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) WithID(id string) *DeleteSiteDevServerHookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) SetID(id string) {
	o.ID = id
}

// WithSiteID adds the siteID to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) WithSiteID(siteID string) *DeleteSiteDevServerHookParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete site dev server hook params
func (o *DeleteSiteDevServerHookParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSiteDevServerHookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
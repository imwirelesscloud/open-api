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
	"github.com/go-openapi/swag"

	"github.com/netlify/open-api/v2/go/models"
)

// NewCreateSiteDeployParams creates a new CreateSiteDeployParams object
// with the default values initialized.
func NewCreateSiteDeployParams() *CreateSiteDeployParams {
	var ()
	return &CreateSiteDeployParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSiteDeployParamsWithTimeout creates a new CreateSiteDeployParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSiteDeployParamsWithTimeout(timeout time.Duration) *CreateSiteDeployParams {
	var ()
	return &CreateSiteDeployParams{

		timeout: timeout,
	}
}

// NewCreateSiteDeployParamsWithContext creates a new CreateSiteDeployParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSiteDeployParamsWithContext(ctx context.Context) *CreateSiteDeployParams {
	var ()
	return &CreateSiteDeployParams{

		Context: ctx,
	}
}

// NewCreateSiteDeployParamsWithHTTPClient creates a new CreateSiteDeployParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSiteDeployParamsWithHTTPClient(client *http.Client) *CreateSiteDeployParams {
	var ()
	return &CreateSiteDeployParams{
		HTTPClient: client,
	}
}

/*CreateSiteDeployParams contains all the parameters to send to the API endpoint
for the create site deploy operation typically these are written to a http.Request
*/
type CreateSiteDeployParams struct {

	/*Branch*/
	Branch *string
	/*Deploy*/
	Deploy *models.DeployFiles
	/*DeployPreviews*/
	DeployPreviews *bool
	/*LatestPublished*/
	LatestPublished *bool
	/*Production*/
	Production *bool
	/*SiteID*/
	SiteID string
	/*State*/
	State *string
	/*Title*/
	Title *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create site deploy params
func (o *CreateSiteDeployParams) WithTimeout(timeout time.Duration) *CreateSiteDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create site deploy params
func (o *CreateSiteDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create site deploy params
func (o *CreateSiteDeployParams) WithContext(ctx context.Context) *CreateSiteDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create site deploy params
func (o *CreateSiteDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create site deploy params
func (o *CreateSiteDeployParams) WithHTTPClient(client *http.Client) *CreateSiteDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create site deploy params
func (o *CreateSiteDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranch adds the branch to the create site deploy params
func (o *CreateSiteDeployParams) WithBranch(branch *string) *CreateSiteDeployParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the create site deploy params
func (o *CreateSiteDeployParams) SetBranch(branch *string) {
	o.Branch = branch
}

// WithDeploy adds the deploy to the create site deploy params
func (o *CreateSiteDeployParams) WithDeploy(deploy *models.DeployFiles) *CreateSiteDeployParams {
	o.SetDeploy(deploy)
	return o
}

// SetDeploy adds the deploy to the create site deploy params
func (o *CreateSiteDeployParams) SetDeploy(deploy *models.DeployFiles) {
	o.Deploy = deploy
}

// WithDeployPreviews adds the deployPreviews to the create site deploy params
func (o *CreateSiteDeployParams) WithDeployPreviews(deployPreviews *bool) *CreateSiteDeployParams {
	o.SetDeployPreviews(deployPreviews)
	return o
}

// SetDeployPreviews adds the deployPreviews to the create site deploy params
func (o *CreateSiteDeployParams) SetDeployPreviews(deployPreviews *bool) {
	o.DeployPreviews = deployPreviews
}

// WithLatestPublished adds the latestPublished to the create site deploy params
func (o *CreateSiteDeployParams) WithLatestPublished(latestPublished *bool) *CreateSiteDeployParams {
	o.SetLatestPublished(latestPublished)
	return o
}

// SetLatestPublished adds the latestPublished to the create site deploy params
func (o *CreateSiteDeployParams) SetLatestPublished(latestPublished *bool) {
	o.LatestPublished = latestPublished
}

// WithProduction adds the production to the create site deploy params
func (o *CreateSiteDeployParams) WithProduction(production *bool) *CreateSiteDeployParams {
	o.SetProduction(production)
	return o
}

// SetProduction adds the production to the create site deploy params
func (o *CreateSiteDeployParams) SetProduction(production *bool) {
	o.Production = production
}

// WithSiteID adds the siteID to the create site deploy params
func (o *CreateSiteDeployParams) WithSiteID(siteID string) *CreateSiteDeployParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the create site deploy params
func (o *CreateSiteDeployParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithState adds the state to the create site deploy params
func (o *CreateSiteDeployParams) WithState(state *string) *CreateSiteDeployParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the create site deploy params
func (o *CreateSiteDeployParams) SetState(state *string) {
	o.State = state
}

// WithTitle adds the title to the create site deploy params
func (o *CreateSiteDeployParams) WithTitle(title *string) *CreateSiteDeployParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the create site deploy params
func (o *CreateSiteDeployParams) SetTitle(title *string) {
	o.Title = title
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSiteDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Branch != nil {

		// query param branch
		var qrBranch string
		if o.Branch != nil {
			qrBranch = *o.Branch
		}
		qBranch := qrBranch
		if qBranch != "" {
			if err := r.SetQueryParam("branch", qBranch); err != nil {
				return err
			}
		}

	}

	if o.Deploy != nil {
		if err := r.SetBodyParam(o.Deploy); err != nil {
			return err
		}
	}

	if o.DeployPreviews != nil {

		// query param deploy-previews
		var qrDeployPreviews bool
		if o.DeployPreviews != nil {
			qrDeployPreviews = *o.DeployPreviews
		}
		qDeployPreviews := swag.FormatBool(qrDeployPreviews)
		if qDeployPreviews != "" {
			if err := r.SetQueryParam("deploy-previews", qDeployPreviews); err != nil {
				return err
			}
		}

	}

	if o.LatestPublished != nil {

		// query param latest-published
		var qrLatestPublished bool
		if o.LatestPublished != nil {
			qrLatestPublished = *o.LatestPublished
		}
		qLatestPublished := swag.FormatBool(qrLatestPublished)
		if qLatestPublished != "" {
			if err := r.SetQueryParam("latest-published", qLatestPublished); err != nil {
				return err
			}
		}

	}

	if o.Production != nil {

		// query param production
		var qrProduction bool
		if o.Production != nil {
			qrProduction = *o.Production
		}
		qProduction := swag.FormatBool(qrProduction)
		if qProduction != "" {
			if err := r.SetQueryParam("production", qProduction); err != nil {
				return err
			}
		}

	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if o.Title != nil {

		// query param title
		var qrTitle string
		if o.Title != nil {
			qrTitle = *o.Title
		}
		qTitle := qrTitle
		if qTitle != "" {
			if err := r.SetQueryParam("title", qTitle); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

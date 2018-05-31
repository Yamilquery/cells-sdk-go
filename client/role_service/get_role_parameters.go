// Code generated by go-swagger; DO NOT EDIT.

package role_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRoleParams creates a new GetRoleParams object
// with the default values initialized.
func NewGetRoleParams() *GetRoleParams {
	var ()
	return &GetRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoleParamsWithTimeout creates a new GetRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoleParamsWithTimeout(timeout time.Duration) *GetRoleParams {
	var ()
	return &GetRoleParams{

		timeout: timeout,
	}
}

// NewGetRoleParamsWithContext creates a new GetRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoleParamsWithContext(ctx context.Context) *GetRoleParams {
	var ()
	return &GetRoleParams{

		Context: ctx,
	}
}

// NewGetRoleParamsWithHTTPClient creates a new GetRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoleParamsWithHTTPClient(client *http.Client) *GetRoleParams {
	var ()
	return &GetRoleParams{
		HTTPClient: client,
	}
}

/*GetRoleParams contains all the parameters to send to the API endpoint
for the get role operation typically these are written to a http.Request
*/
type GetRoleParams struct {

	/*AutoApplies*/
	AutoApplies []string
	/*GroupRole*/
	GroupRole *bool
	/*IsTeam*/
	IsTeam *bool
	/*Label*/
	Label *string
	/*LastUpdated*/
	LastUpdated *int32
	/*PoliciesContextEditable*/
	PoliciesContextEditable *bool
	/*UserRole*/
	UserRole *bool
	/*UUID*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get role params
func (o *GetRoleParams) WithTimeout(timeout time.Duration) *GetRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role params
func (o *GetRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role params
func (o *GetRoleParams) WithContext(ctx context.Context) *GetRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role params
func (o *GetRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role params
func (o *GetRoleParams) WithHTTPClient(client *http.Client) *GetRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role params
func (o *GetRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAutoApplies adds the autoApplies to the get role params
func (o *GetRoleParams) WithAutoApplies(autoApplies []string) *GetRoleParams {
	o.SetAutoApplies(autoApplies)
	return o
}

// SetAutoApplies adds the autoApplies to the get role params
func (o *GetRoleParams) SetAutoApplies(autoApplies []string) {
	o.AutoApplies = autoApplies
}

// WithGroupRole adds the groupRole to the get role params
func (o *GetRoleParams) WithGroupRole(groupRole *bool) *GetRoleParams {
	o.SetGroupRole(groupRole)
	return o
}

// SetGroupRole adds the groupRole to the get role params
func (o *GetRoleParams) SetGroupRole(groupRole *bool) {
	o.GroupRole = groupRole
}

// WithIsTeam adds the isTeam to the get role params
func (o *GetRoleParams) WithIsTeam(isTeam *bool) *GetRoleParams {
	o.SetIsTeam(isTeam)
	return o
}

// SetIsTeam adds the isTeam to the get role params
func (o *GetRoleParams) SetIsTeam(isTeam *bool) {
	o.IsTeam = isTeam
}

// WithLabel adds the label to the get role params
func (o *GetRoleParams) WithLabel(label *string) *GetRoleParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the get role params
func (o *GetRoleParams) SetLabel(label *string) {
	o.Label = label
}

// WithLastUpdated adds the lastUpdated to the get role params
func (o *GetRoleParams) WithLastUpdated(lastUpdated *int32) *GetRoleParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the get role params
func (o *GetRoleParams) SetLastUpdated(lastUpdated *int32) {
	o.LastUpdated = lastUpdated
}

// WithPoliciesContextEditable adds the policiesContextEditable to the get role params
func (o *GetRoleParams) WithPoliciesContextEditable(policiesContextEditable *bool) *GetRoleParams {
	o.SetPoliciesContextEditable(policiesContextEditable)
	return o
}

// SetPoliciesContextEditable adds the policiesContextEditable to the get role params
func (o *GetRoleParams) SetPoliciesContextEditable(policiesContextEditable *bool) {
	o.PoliciesContextEditable = policiesContextEditable
}

// WithUserRole adds the userRole to the get role params
func (o *GetRoleParams) WithUserRole(userRole *bool) *GetRoleParams {
	o.SetUserRole(userRole)
	return o
}

// SetUserRole adds the userRole to the get role params
func (o *GetRoleParams) SetUserRole(userRole *bool) {
	o.UserRole = userRole
}

// WithUUID adds the uuid to the get role params
func (o *GetRoleParams) WithUUID(uuid string) *GetRoleParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get role params
func (o *GetRoleParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAutoApplies := o.AutoApplies

	joinedAutoApplies := swag.JoinByFormat(valuesAutoApplies, "")
	// query array param AutoApplies
	if err := r.SetQueryParam("AutoApplies", joinedAutoApplies...); err != nil {
		return err
	}

	if o.GroupRole != nil {

		// query param GroupRole
		var qrGroupRole bool
		if o.GroupRole != nil {
			qrGroupRole = *o.GroupRole
		}
		qGroupRole := swag.FormatBool(qrGroupRole)
		if qGroupRole != "" {
			if err := r.SetQueryParam("GroupRole", qGroupRole); err != nil {
				return err
			}
		}

	}

	if o.IsTeam != nil {

		// query param IsTeam
		var qrIsTeam bool
		if o.IsTeam != nil {
			qrIsTeam = *o.IsTeam
		}
		qIsTeam := swag.FormatBool(qrIsTeam)
		if qIsTeam != "" {
			if err := r.SetQueryParam("IsTeam", qIsTeam); err != nil {
				return err
			}
		}

	}

	if o.Label != nil {

		// query param Label
		var qrLabel string
		if o.Label != nil {
			qrLabel = *o.Label
		}
		qLabel := qrLabel
		if qLabel != "" {
			if err := r.SetQueryParam("Label", qLabel); err != nil {
				return err
			}
		}

	}

	if o.LastUpdated != nil {

		// query param LastUpdated
		var qrLastUpdated int32
		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := swag.FormatInt32(qrLastUpdated)
		if qLastUpdated != "" {
			if err := r.SetQueryParam("LastUpdated", qLastUpdated); err != nil {
				return err
			}
		}

	}

	if o.PoliciesContextEditable != nil {

		// query param PoliciesContextEditable
		var qrPoliciesContextEditable bool
		if o.PoliciesContextEditable != nil {
			qrPoliciesContextEditable = *o.PoliciesContextEditable
		}
		qPoliciesContextEditable := swag.FormatBool(qrPoliciesContextEditable)
		if qPoliciesContextEditable != "" {
			if err := r.SetQueryParam("PoliciesContextEditable", qPoliciesContextEditable); err != nil {
				return err
			}
		}

	}

	if o.UserRole != nil {

		// query param UserRole
		var qrUserRole bool
		if o.UserRole != nil {
			qrUserRole = *o.UserRole
		}
		qUserRole := swag.FormatBool(qrUserRole)
		if qUserRole != "" {
			if err := r.SetQueryParam("UserRole", qUserRole); err != nil {
				return err
			}
		}

	}

	// path param Uuid
	if err := r.SetPathParam("Uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

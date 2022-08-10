// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudPvminstancesCapturePostParams creates a new PcloudPvminstancesCapturePostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesCapturePostParams() *PcloudPvminstancesCapturePostParams {
	return &PcloudPvminstancesCapturePostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesCapturePostParamsWithTimeout creates a new PcloudPvminstancesCapturePostParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesCapturePostParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesCapturePostParams {
	return &PcloudPvminstancesCapturePostParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesCapturePostParamsWithContext creates a new PcloudPvminstancesCapturePostParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesCapturePostParamsWithContext(ctx context.Context) *PcloudPvminstancesCapturePostParams {
	return &PcloudPvminstancesCapturePostParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesCapturePostParamsWithHTTPClient creates a new PcloudPvminstancesCapturePostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesCapturePostParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesCapturePostParams {
	return &PcloudPvminstancesCapturePostParams{
		HTTPClient: client,
	}
}

/* PcloudPvminstancesCapturePostParams contains all the parameters to send to the API endpoint
   for the pcloud pvminstances capture post operation.

   Typically these are written to a http.Request.
*/
type PcloudPvminstancesCapturePostParams struct {

	/* Body.

	   Parameters for the capture PVMInstance
	*/
	Body *models.PVMInstanceCapture

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances capture post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesCapturePostParams) WithDefaults() *PcloudPvminstancesCapturePostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances capture post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesCapturePostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesCapturePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) WithContext(ctx context.Context) *PcloudPvminstancesCapturePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesCapturePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) WithBody(body *models.PVMInstanceCapture) *PcloudPvminstancesCapturePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) SetBody(body *models.PVMInstanceCapture) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesCapturePostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesCapturePostParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances capture post params
func (o *PcloudPvminstancesCapturePostParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesCapturePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

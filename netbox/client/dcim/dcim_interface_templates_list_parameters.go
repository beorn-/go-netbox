// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dcim

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
)

// NewDcimInterfaceTemplatesListParams creates a new DcimInterfaceTemplatesListParams object
// with the default values initialized.
func NewDcimInterfaceTemplatesListParams() *DcimInterfaceTemplatesListParams {
	var ()
	return &DcimInterfaceTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesListParamsWithTimeout creates a new DcimInterfaceTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfaceTemplatesListParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesListParams {
	var ()
	return &DcimInterfaceTemplatesListParams{

		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesListParamsWithContext creates a new DcimInterfaceTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfaceTemplatesListParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesListParams {
	var ()
	return &DcimInterfaceTemplatesListParams{

		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesListParamsWithHTTPClient creates a new DcimInterfaceTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfaceTemplatesListParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesListParams {
	var ()
	return &DcimInterfaceTemplatesListParams{
		HTTPClient: client,
	}
}

/*DcimInterfaceTemplatesListParams contains all the parameters to send to the API endpoint
for the dcim interface templates list operation typically these are written to a http.Request
*/
type DcimInterfaceTemplatesListParams struct {

	/*DevicetypeID*/
	DevicetypeID *string
	/*FormFactor*/
	FormFactor *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*MgmtOnly*/
	MgmtOnly *string
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevicetypeID adds the devicetypeID to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimInterfaceTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithFormFactor adds the formFactor to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) WithFormFactor(formFactor *string) *DcimInterfaceTemplatesListParams {
	o.SetFormFactor(formFactor)
	return o
}

// SetFormFactor adds the formFactor to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) SetFormFactor(formFactor *string) {
	o.FormFactor = formFactor
}

// WithLimit adds the limit to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) WithLimit(limit *int64) *DcimInterfaceTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMgmtOnly adds the mgmtOnly to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) WithMgmtOnly(mgmtOnly *string) *DcimInterfaceTemplatesListParams {
	o.SetMgmtOnly(mgmtOnly)
	return o
}

// SetMgmtOnly adds the mgmtOnly to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) SetMgmtOnly(mgmtOnly *string) {
	o.MgmtOnly = mgmtOnly
}

// WithName adds the name to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) WithName(name *string) *DcimInterfaceTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) WithOffset(offset *int64) *DcimInterfaceTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim interface templates list params
func (o *DcimInterfaceTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string
		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {
			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
				return err
			}
		}

	}

	if o.FormFactor != nil {

		// query param form_factor
		var qrFormFactor string
		if o.FormFactor != nil {
			qrFormFactor = *o.FormFactor
		}
		qFormFactor := qrFormFactor
		if qFormFactor != "" {
			if err := r.SetQueryParam("form_factor", qFormFactor); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.MgmtOnly != nil {

		// query param mgmt_only
		var qrMgmtOnly string
		if o.MgmtOnly != nil {
			qrMgmtOnly = *o.MgmtOnly
		}
		qMgmtOnly := qrMgmtOnly
		if qMgmtOnly != "" {
			if err := r.SetQueryParam("mgmt_only", qMgmtOnly); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

	"github.com/beorn-/go-netbox/netbox/models"
)

// NewDcimInterfaceTemplatesPartialUpdateParams creates a new DcimInterfaceTemplatesPartialUpdateParams object
// with the default values initialized.
func NewDcimInterfaceTemplatesPartialUpdateParams() *DcimInterfaceTemplatesPartialUpdateParams {
	var ()
	return &DcimInterfaceTemplatesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesPartialUpdateParamsWithTimeout creates a new DcimInterfaceTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfaceTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesPartialUpdateParams {
	var ()
	return &DcimInterfaceTemplatesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesPartialUpdateParamsWithContext creates a new DcimInterfaceTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfaceTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesPartialUpdateParams {
	var ()
	return &DcimInterfaceTemplatesPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimInterfaceTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfaceTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesPartialUpdateParams {
	var ()
	return &DcimInterfaceTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimInterfaceTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim interface templates partial update operation typically these are written to a http.Request
*/
type DcimInterfaceTemplatesPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableInterfaceTemplate
	/*ID
	  A unique integer value identifying this interface template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithData(data *models.WritableInterfaceTemplate) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetData(data *models.WritableInterfaceTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithID(id int64) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/beorn-/go-netbox/netbox/models"
)

// NewDcimInterfaceConnectionsPartialUpdateParams creates a new DcimInterfaceConnectionsPartialUpdateParams object
// with the default values initialized.
func NewDcimInterfaceConnectionsPartialUpdateParams() *DcimInterfaceConnectionsPartialUpdateParams {
	var ()
	return &DcimInterfaceConnectionsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceConnectionsPartialUpdateParamsWithTimeout creates a new DcimInterfaceConnectionsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfaceConnectionsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimInterfaceConnectionsPartialUpdateParams {
	var ()
	return &DcimInterfaceConnectionsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimInterfaceConnectionsPartialUpdateParamsWithContext creates a new DcimInterfaceConnectionsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfaceConnectionsPartialUpdateParamsWithContext(ctx context.Context) *DcimInterfaceConnectionsPartialUpdateParams {
	var ()
	return &DcimInterfaceConnectionsPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimInterfaceConnectionsPartialUpdateParamsWithHTTPClient creates a new DcimInterfaceConnectionsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfaceConnectionsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimInterfaceConnectionsPartialUpdateParams {
	var ()
	return &DcimInterfaceConnectionsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimInterfaceConnectionsPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim interface connections partial update operation typically these are written to a http.Request
*/
type DcimInterfaceConnectionsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableInterfaceConnection
	/*ID
	  A unique integer value identifying this interface connection.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimInterfaceConnectionsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) WithContext(ctx context.Context) *DcimInterfaceConnectionsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimInterfaceConnectionsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) WithData(data *models.WritableInterfaceConnection) *DcimInterfaceConnectionsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) SetData(data *models.WritableInterfaceConnection) {
	o.Data = data
}

// WithID adds the id to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) WithID(id int64) *DcimInterfaceConnectionsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interface connections partial update params
func (o *DcimInterfaceConnectionsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceConnectionsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

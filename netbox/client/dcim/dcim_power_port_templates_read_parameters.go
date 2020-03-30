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

// NewDcimPowerPortTemplatesReadParams creates a new DcimPowerPortTemplatesReadParams object
// with the default values initialized.
func NewDcimPowerPortTemplatesReadParams() *DcimPowerPortTemplatesReadParams {
	var ()
	return &DcimPowerPortTemplatesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesReadParamsWithTimeout creates a new DcimPowerPortTemplatesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPortTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesReadParams {
	var ()
	return &DcimPowerPortTemplatesReadParams{

		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesReadParamsWithContext creates a new DcimPowerPortTemplatesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPortTemplatesReadParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesReadParams {
	var ()
	return &DcimPowerPortTemplatesReadParams{

		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesReadParamsWithHTTPClient creates a new DcimPowerPortTemplatesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPortTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesReadParams {
	var ()
	return &DcimPowerPortTemplatesReadParams{
		HTTPClient: client,
	}
}

/*DcimPowerPortTemplatesReadParams contains all the parameters to send to the API endpoint
for the dcim power port templates read operation typically these are written to a http.Request
*/
type DcimPowerPortTemplatesReadParams struct {

	/*ID
	  A unique integer value identifying this power port template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) WithID(id int64) *DcimPowerPortTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power port templates read params
func (o *DcimPowerPortTemplatesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

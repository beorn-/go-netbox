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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/beorn-/go-netbox/netbox/models"
)

// NewDcimPowerPortTemplatesCreateParams creates a new DcimPowerPortTemplatesCreateParams object
// with the default values initialized.
func NewDcimPowerPortTemplatesCreateParams() *DcimPowerPortTemplatesCreateParams {
	var ()
	return &DcimPowerPortTemplatesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesCreateParamsWithTimeout creates a new DcimPowerPortTemplatesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPortTemplatesCreateParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesCreateParams {
	var ()
	return &DcimPowerPortTemplatesCreateParams{

		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesCreateParamsWithContext creates a new DcimPowerPortTemplatesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPortTemplatesCreateParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesCreateParams {
	var ()
	return &DcimPowerPortTemplatesCreateParams{

		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesCreateParamsWithHTTPClient creates a new DcimPowerPortTemplatesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPortTemplatesCreateParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesCreateParams {
	var ()
	return &DcimPowerPortTemplatesCreateParams{
		HTTPClient: client,
	}
}

/*DcimPowerPortTemplatesCreateParams contains all the parameters to send to the API endpoint
for the dcim power port templates create operation typically these are written to a http.Request
*/
type DcimPowerPortTemplatesCreateParams struct {

	/*Data*/
	Data *models.WritablePowerPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power port templates create params
func (o *DcimPowerPortTemplatesCreateParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates create params
func (o *DcimPowerPortTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates create params
func (o *DcimPowerPortTemplatesCreateParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates create params
func (o *DcimPowerPortTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates create params
func (o *DcimPowerPortTemplatesCreateParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates create params
func (o *DcimPowerPortTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power port templates create params
func (o *DcimPowerPortTemplatesCreateParams) WithData(data *models.WritablePowerPortTemplate) *DcimPowerPortTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power port templates create params
func (o *DcimPowerPortTemplatesCreateParams) SetData(data *models.WritablePowerPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

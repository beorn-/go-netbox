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

	"github.com/beorn-/go-netbox/netbox/models"
)

// NewDcimSitesCreateParams creates a new DcimSitesCreateParams object
// with the default values initialized.
func NewDcimSitesCreateParams() *DcimSitesCreateParams {
	var ()
	return &DcimSitesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesCreateParamsWithTimeout creates a new DcimSitesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimSitesCreateParamsWithTimeout(timeout time.Duration) *DcimSitesCreateParams {
	var ()
	return &DcimSitesCreateParams{

		timeout: timeout,
	}
}

// NewDcimSitesCreateParamsWithContext creates a new DcimSitesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimSitesCreateParamsWithContext(ctx context.Context) *DcimSitesCreateParams {
	var ()
	return &DcimSitesCreateParams{

		Context: ctx,
	}
}

// NewDcimSitesCreateParamsWithHTTPClient creates a new DcimSitesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimSitesCreateParamsWithHTTPClient(client *http.Client) *DcimSitesCreateParams {
	var ()
	return &DcimSitesCreateParams{
		HTTPClient: client,
	}
}

/*DcimSitesCreateParams contains all the parameters to send to the API endpoint
for the dcim sites create operation typically these are written to a http.Request
*/
type DcimSitesCreateParams struct {

	/*Data*/
	Data *models.WritableSite

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim sites create params
func (o *DcimSitesCreateParams) WithTimeout(timeout time.Duration) *DcimSitesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites create params
func (o *DcimSitesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites create params
func (o *DcimSitesCreateParams) WithContext(ctx context.Context) *DcimSitesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites create params
func (o *DcimSitesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites create params
func (o *DcimSitesCreateParams) WithHTTPClient(client *http.Client) *DcimSitesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites create params
func (o *DcimSitesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim sites create params
func (o *DcimSitesCreateParams) WithData(data *models.WritableSite) *DcimSitesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim sites create params
func (o *DcimSitesCreateParams) SetData(data *models.WritableSite) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

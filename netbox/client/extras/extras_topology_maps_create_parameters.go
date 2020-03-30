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

package extras

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

// NewExtrasTopologyMapsCreateParams creates a new ExtrasTopologyMapsCreateParams object
// with the default values initialized.
func NewExtrasTopologyMapsCreateParams() *ExtrasTopologyMapsCreateParams {
	var ()
	return &ExtrasTopologyMapsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTopologyMapsCreateParamsWithTimeout creates a new ExtrasTopologyMapsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTopologyMapsCreateParamsWithTimeout(timeout time.Duration) *ExtrasTopologyMapsCreateParams {
	var ()
	return &ExtrasTopologyMapsCreateParams{

		timeout: timeout,
	}
}

// NewExtrasTopologyMapsCreateParamsWithContext creates a new ExtrasTopologyMapsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTopologyMapsCreateParamsWithContext(ctx context.Context) *ExtrasTopologyMapsCreateParams {
	var ()
	return &ExtrasTopologyMapsCreateParams{

		Context: ctx,
	}
}

// NewExtrasTopologyMapsCreateParamsWithHTTPClient creates a new ExtrasTopologyMapsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTopologyMapsCreateParamsWithHTTPClient(client *http.Client) *ExtrasTopologyMapsCreateParams {
	var ()
	return &ExtrasTopologyMapsCreateParams{
		HTTPClient: client,
	}
}

/*ExtrasTopologyMapsCreateParams contains all the parameters to send to the API endpoint
for the extras topology maps create operation typically these are written to a http.Request
*/
type ExtrasTopologyMapsCreateParams struct {

	/*Data*/
	Data *models.WritableTopologyMap

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) WithTimeout(timeout time.Duration) *ExtrasTopologyMapsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) WithContext(ctx context.Context) *ExtrasTopologyMapsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) WithHTTPClient(client *http.Client) *ExtrasTopologyMapsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) WithData(data *models.WritableTopologyMap) *ExtrasTopologyMapsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) SetData(data *models.WritableTopologyMap) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTopologyMapsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

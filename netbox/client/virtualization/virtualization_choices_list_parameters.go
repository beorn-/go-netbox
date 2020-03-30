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

package virtualization

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
)

// NewVirtualizationChoicesListParams creates a new VirtualizationChoicesListParams object
// with the default values initialized.
func NewVirtualizationChoicesListParams() *VirtualizationChoicesListParams {

	return &VirtualizationChoicesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationChoicesListParamsWithTimeout creates a new VirtualizationChoicesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationChoicesListParamsWithTimeout(timeout time.Duration) *VirtualizationChoicesListParams {

	return &VirtualizationChoicesListParams{

		timeout: timeout,
	}
}

// NewVirtualizationChoicesListParamsWithContext creates a new VirtualizationChoicesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationChoicesListParamsWithContext(ctx context.Context) *VirtualizationChoicesListParams {

	return &VirtualizationChoicesListParams{

		Context: ctx,
	}
}

// NewVirtualizationChoicesListParamsWithHTTPClient creates a new VirtualizationChoicesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationChoicesListParamsWithHTTPClient(client *http.Client) *VirtualizationChoicesListParams {

	return &VirtualizationChoicesListParams{
		HTTPClient: client,
	}
}

/*VirtualizationChoicesListParams contains all the parameters to send to the API endpoint
for the virtualization choices list operation typically these are written to a http.Request
*/
type VirtualizationChoicesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization choices list params
func (o *VirtualizationChoicesListParams) WithTimeout(timeout time.Duration) *VirtualizationChoicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization choices list params
func (o *VirtualizationChoicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization choices list params
func (o *VirtualizationChoicesListParams) WithContext(ctx context.Context) *VirtualizationChoicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization choices list params
func (o *VirtualizationChoicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization choices list params
func (o *VirtualizationChoicesListParams) WithHTTPClient(client *http.Client) *VirtualizationChoicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization choices list params
func (o *VirtualizationChoicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationChoicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

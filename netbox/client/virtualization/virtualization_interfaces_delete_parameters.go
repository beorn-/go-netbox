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
	"github.com/go-openapi/swag"
)

// NewVirtualizationInterfacesDeleteParams creates a new VirtualizationInterfacesDeleteParams object
// with the default values initialized.
func NewVirtualizationInterfacesDeleteParams() *VirtualizationInterfacesDeleteParams {
	var ()
	return &VirtualizationInterfacesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesDeleteParamsWithTimeout creates a new VirtualizationInterfacesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationInterfacesDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesDeleteParams {
	var ()
	return &VirtualizationInterfacesDeleteParams{

		timeout: timeout,
	}
}

// NewVirtualizationInterfacesDeleteParamsWithContext creates a new VirtualizationInterfacesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationInterfacesDeleteParamsWithContext(ctx context.Context) *VirtualizationInterfacesDeleteParams {
	var ()
	return &VirtualizationInterfacesDeleteParams{

		Context: ctx,
	}
}

// NewVirtualizationInterfacesDeleteParamsWithHTTPClient creates a new VirtualizationInterfacesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationInterfacesDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesDeleteParams {
	var ()
	return &VirtualizationInterfacesDeleteParams{
		HTTPClient: client,
	}
}

/*VirtualizationInterfacesDeleteParams contains all the parameters to send to the API endpoint
for the virtualization interfaces delete operation typically these are written to a http.Request
*/
type VirtualizationInterfacesDeleteParams struct {

	/*ID
	  A unique integer value identifying this interface.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) WithContext(ctx context.Context) *VirtualizationInterfacesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) WithID(id int64) *VirtualizationInterfacesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

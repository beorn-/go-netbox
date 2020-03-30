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

package ipam

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

// NewIpamAggregatesListParams creates a new IpamAggregatesListParams object
// with the default values initialized.
func NewIpamAggregatesListParams() *IpamAggregatesListParams {
	var ()
	return &IpamAggregatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAggregatesListParamsWithTimeout creates a new IpamAggregatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamAggregatesListParamsWithTimeout(timeout time.Duration) *IpamAggregatesListParams {
	var ()
	return &IpamAggregatesListParams{

		timeout: timeout,
	}
}

// NewIpamAggregatesListParamsWithContext creates a new IpamAggregatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamAggregatesListParamsWithContext(ctx context.Context) *IpamAggregatesListParams {
	var ()
	return &IpamAggregatesListParams{

		Context: ctx,
	}
}

// NewIpamAggregatesListParamsWithHTTPClient creates a new IpamAggregatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamAggregatesListParamsWithHTTPClient(client *http.Client) *IpamAggregatesListParams {
	var ()
	return &IpamAggregatesListParams{
		HTTPClient: client,
	}
}

/*IpamAggregatesListParams contains all the parameters to send to the API endpoint
for the ipam aggregates list operation typically these are written to a http.Request
*/
type IpamAggregatesListParams struct {

	/*DateAdded*/
	DateAdded *string
	/*Family*/
	Family *string
	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Rir*/
	Rir *string
	/*RirID*/
	RirID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithTimeout(timeout time.Duration) *IpamAggregatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithContext(ctx context.Context) *IpamAggregatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithHTTPClient(client *http.Client) *IpamAggregatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDateAdded adds the dateAdded to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithDateAdded(dateAdded *string) *IpamAggregatesListParams {
	o.SetDateAdded(dateAdded)
	return o
}

// SetDateAdded adds the dateAdded to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetDateAdded(dateAdded *string) {
	o.DateAdded = dateAdded
}

// WithFamily adds the family to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithFamily(family *string) *IpamAggregatesListParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetFamily(family *string) {
	o.Family = family
}

// WithIDIn adds the iDIn to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithIDIn(iDIn *string) *IpamAggregatesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithLimit adds the limit to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithLimit(limit *int64) *IpamAggregatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithOffset(offset *int64) *IpamAggregatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithQ(q *string) *IpamAggregatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRir adds the rir to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithRir(rir *string) *IpamAggregatesListParams {
	o.SetRir(rir)
	return o
}

// SetRir adds the rir to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetRir(rir *string) {
	o.Rir = rir
}

// WithRirID adds the rirID to the ipam aggregates list params
func (o *IpamAggregatesListParams) WithRirID(rirID *string) *IpamAggregatesListParams {
	o.SetRirID(rirID)
	return o
}

// SetRirID adds the rirId to the ipam aggregates list params
func (o *IpamAggregatesListParams) SetRirID(rirID *string) {
	o.RirID = rirID
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAggregatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DateAdded != nil {

		// query param date_added
		var qrDateAdded string
		if o.DateAdded != nil {
			qrDateAdded = *o.DateAdded
		}
		qDateAdded := qrDateAdded
		if qDateAdded != "" {
			if err := r.SetQueryParam("date_added", qDateAdded); err != nil {
				return err
			}
		}

	}

	if o.Family != nil {

		// query param family
		var qrFamily string
		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := qrFamily
		if qFamily != "" {
			if err := r.SetQueryParam("family", qFamily); err != nil {
				return err
			}
		}

	}

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn string
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := qrIDIn
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
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

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Rir != nil {

		// query param rir
		var qrRir string
		if o.Rir != nil {
			qrRir = *o.Rir
		}
		qRir := qrRir
		if qRir != "" {
			if err := r.SetQueryParam("rir", qRir); err != nil {
				return err
			}
		}

	}

	if o.RirID != nil {

		// query param rir_id
		var qrRirID string
		if o.RirID != nil {
			qrRirID = *o.RirID
		}
		qRirID := qrRirID
		if qRirID != "" {
			if err := r.SetQueryParam("rir_id", qRirID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

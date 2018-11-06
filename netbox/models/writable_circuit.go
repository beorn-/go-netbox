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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableCircuit writable circuit
// swagger:model WritableCircuit
type WritableCircuit struct {

	// Circuit ID
	// Required: true
	// Max Length: 50
	Cid *string `json:"cid"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Commit rate (Kbps)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	CommitRate *int64 `json:"commit_rate,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Date installed
	// Format: date
	InstallDate strfmt.Date `json:"install_date,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Provider
	// Required: true
	Provider *int64 `json:"provider"`

	// Status
	// Enum: [2 3 1 4 0 5]
	Status int64 `json:"status,omitempty"`

	// Tenant
	Tenant int64 `json:"tenant,omitempty"`

	// Type
	// Required: true
	Type *int64 `json:"type"`
}

// Validate validates this writable circuit
func (m *WritableCircuit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCircuit) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	if err := validate.MaxLength("cid", "body", string(*m.Cid), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateCommitRate(formats strfmt.Registry) error {

	if swag.IsZero(m.CommitRate) { // not required
		return nil
	}

	if err := validate.MinimumInt("commit_rate", "body", int64(*m.CommitRate), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("commit_rate", "body", int64(*m.CommitRate), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateInstallDate(formats strfmt.Registry) error {

	if swag.IsZero(m.InstallDate) { // not required
		return nil
	}

	if err := validate.FormatOf("install_date", "body", "date", m.InstallDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	return nil
}

var writableCircuitTypeStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[2,3,1,4,0,5]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCircuitTypeStatusPropEnum = append(writableCircuitTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableCircuit) validateStatusEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableCircuitTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableCircuit) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuit) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableCircuit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableCircuit) UnmarshalBinary(b []byte) error {
	var res WritableCircuit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

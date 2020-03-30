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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableInventoryItem writable inventory item
//
// swagger:model WritableInventoryItem
type WritableInventoryItem struct {

	// Asset tag
	//
	// A unique tag used to identify this item
	// Max Length: 50
	AssetTag string `json:"asset_tag,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// Discovered
	Discovered bool `json:"discovered,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Manufacturer
	Manufacturer int64 `json:"manufacturer,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	Name *string `json:"name"`

	// Parent
	Parent int64 `json:"parent,omitempty"`

	// Part ID
	// Max Length: 50
	PartID string `json:"part_id,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`
}

// Validate validates this writable inventory item
func (m *WritableInventoryItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableInventoryItem) validateAssetTag(formats strfmt.Registry) error {

	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", string(m.AssetTag), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableInventoryItem) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableInventoryItem) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *WritableInventoryItem) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableInventoryItem) validatePartID(formats strfmt.Registry) error {

	if swag.IsZero(m.PartID) { // not required
		return nil
	}

	if err := validate.MaxLength("part_id", "body", string(m.PartID), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableInventoryItem) validateSerial(formats strfmt.Registry) error {

	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", string(m.Serial), 50); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableInventoryItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableInventoryItem) UnmarshalBinary(b []byte) error {
	var res WritableInventoryItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

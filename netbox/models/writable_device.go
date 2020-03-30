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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableDevice writable device
//
// swagger:model WritableDevice
type WritableDevice struct {

	// Asset tag
	//
	// A unique tag used to identify this device
	// Max Length: 50
	AssetTag string `json:"asset_tag,omitempty"`

	// Cluster
	Cluster int64 `json:"cluster,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Device role
	// Required: true
	DeviceRole *int64 `json:"device_role"`

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// Rack face
	// Enum: [0 1]
	Face int64 `json:"face,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Max Length: 64
	Name string `json:"name,omitempty"`

	// Platform
	Platform int64 `json:"platform,omitempty"`

	// Position (U)
	//
	// The lowest-numbered unit occupied by the device
	// Maximum: 32767
	// Minimum: 1
	Position int64 `json:"position,omitempty"`

	// Primary IPv4
	PrimaryIp4 int64 `json:"primary_ip4,omitempty"`

	// Primary IPv6
	PrimaryIp6 int64 `json:"primary_ip6,omitempty"`

	// Rack
	Rack int64 `json:"rack,omitempty"`

	// Serial number
	// Max Length: 50
	Serial string `json:"serial,omitempty"`

	// Site
	// Required: true
	Site *int64 `json:"site"`

	// Status
	// Enum: [1 0 2 3 4 5]
	Status int64 `json:"status,omitempty"`

	// Tenant
	Tenant int64 `json:"tenant,omitempty"`

	// Vc position
	// Maximum: 255
	// Minimum: 0
	VcPosition *int64 `json:"vc_position,omitempty"`

	// Vc priority
	// Maximum: 255
	// Minimum: 0
	VcPriority *int64 `json:"vc_priority,omitempty"`

	// Virtual chassis
	VirtualChassis int64 `json:"virtual_chassis,omitempty"`
}

// Validate validates this writable device
func (m *WritableDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcPriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableDevice) validateAssetTag(formats strfmt.Registry) error {

	if swag.IsZero(m.AssetTag) { // not required
		return nil
	}

	if err := validate.MaxLength("asset_tag", "body", string(m.AssetTag), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validateDeviceRole(formats strfmt.Registry) error {

	if err := validate.Required("device_role", "body", m.DeviceRole); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

var writableDeviceTypeFacePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceTypeFacePropEnum = append(writableDeviceTypeFacePropEnum, v)
	}
}

// prop value enum
func (m *WritableDevice) validateFaceEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableDeviceTypeFacePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableDevice) validateFace(formats strfmt.Registry) error {

	if swag.IsZero(m.Face) { // not required
		return nil
	}

	// value enum
	if err := m.validateFaceEnum("face", "body", m.Face); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", int64(m.Position), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("position", "body", int64(m.Position), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validateSerial(formats strfmt.Registry) error {

	if swag.IsZero(m.Serial) { // not required
		return nil
	}

	if err := validate.MaxLength("serial", "body", string(m.Serial), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	return nil
}

var writableDeviceTypeStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,0,2,3,4,5]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableDeviceTypeStatusPropEnum = append(writableDeviceTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableDevice) validateStatusEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableDeviceTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableDevice) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validateVcPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.VcPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_position", "body", int64(*m.VcPosition), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_position", "body", int64(*m.VcPosition), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableDevice) validateVcPriority(formats strfmt.Registry) error {

	if swag.IsZero(m.VcPriority) { // not required
		return nil
	}

	if err := validate.MinimumInt("vc_priority", "body", int64(*m.VcPriority), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vc_priority", "body", int64(*m.VcPriority), 255, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableDevice) UnmarshalBinary(b []byte) error {
	var res WritableDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

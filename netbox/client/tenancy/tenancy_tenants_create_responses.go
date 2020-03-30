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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/beorn-/go-netbox/netbox/models"
)

// TenancyTenantsCreateReader is a Reader for the TenancyTenantsCreate structure.
type TenancyTenantsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyTenantsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTenancyTenantsCreateCreated creates a TenancyTenantsCreateCreated with default headers values
func NewTenancyTenantsCreateCreated() *TenancyTenantsCreateCreated {
	return &TenancyTenantsCreateCreated{}
}

/*TenancyTenantsCreateCreated handles this case with default header values.

TenancyTenantsCreateCreated tenancy tenants create created
*/
type TenancyTenantsCreateCreated struct {
	Payload *models.WritableTenant
}

func (o *TenancyTenantsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /tenancy/tenants/][%d] tenancyTenantsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyTenantsCreateCreated) GetPayload() *models.WritableTenant {
	return o.Payload
}

func (o *TenancyTenantsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableTenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/beorn-/go-netbox/netbox/models"
)

// ExtrasTopologyMapsPartialUpdateReader is a Reader for the ExtrasTopologyMapsPartialUpdate structure.
type ExtrasTopologyMapsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTopologyMapsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTopologyMapsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasTopologyMapsPartialUpdateOK creates a ExtrasTopologyMapsPartialUpdateOK with default headers values
func NewExtrasTopologyMapsPartialUpdateOK() *ExtrasTopologyMapsPartialUpdateOK {
	return &ExtrasTopologyMapsPartialUpdateOK{}
}

/*ExtrasTopologyMapsPartialUpdateOK handles this case with default header values.

ExtrasTopologyMapsPartialUpdateOK extras topology maps partial update o k
*/
type ExtrasTopologyMapsPartialUpdateOK struct {
	Payload *models.WritableTopologyMap
}

func (o *ExtrasTopologyMapsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/topology-maps/{id}/][%d] extrasTopologyMapsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasTopologyMapsPartialUpdateOK) GetPayload() *models.WritableTopologyMap {
	return o.Payload
}

func (o *ExtrasTopologyMapsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableTopologyMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

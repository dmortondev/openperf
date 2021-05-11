// Code generated by go-swagger; DO NOT EDIT.

package packet_generators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spirent/openperf/api/client/golang/models"
)

// ListPacketGeneratorsReader is a Reader for the ListPacketGenerators structure.
type ListPacketGeneratorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPacketGeneratorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPacketGeneratorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPacketGeneratorsOK creates a ListPacketGeneratorsOK with default headers values
func NewListPacketGeneratorsOK() *ListPacketGeneratorsOK {
	return &ListPacketGeneratorsOK{}
}

/* ListPacketGeneratorsOK describes a response with status code 200, with default header values.

Success
*/
type ListPacketGeneratorsOK struct {
	Payload []*models.PacketGenerator
}

func (o *ListPacketGeneratorsOK) Error() string {
	return fmt.Sprintf("[GET /packet/generators][%d] listPacketGeneratorsOK  %+v", 200, o.Payload)
}
func (o *ListPacketGeneratorsOK) GetPayload() []*models.PacketGenerator {
	return o.Payload
}

func (o *ListPacketGeneratorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

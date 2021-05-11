// Code generated by go-swagger; DO NOT EDIT.

package packet_captures

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spirent/openperf/api/client/golang/models"
)

// GetPacketCaptureResultReader is a Reader for the GetPacketCaptureResult structure.
type GetPacketCaptureResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPacketCaptureResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPacketCaptureResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPacketCaptureResultOK creates a GetPacketCaptureResultOK with default headers values
func NewGetPacketCaptureResultOK() *GetPacketCaptureResultOK {
	return &GetPacketCaptureResultOK{}
}

/* GetPacketCaptureResultOK describes a response with status code 200, with default header values.

Success
*/
type GetPacketCaptureResultOK struct {
	Payload *models.PacketCaptureResult
}

func (o *GetPacketCaptureResultOK) Error() string {
	return fmt.Sprintf("[GET /packet/capture-results/{id}][%d] getPacketCaptureResultOK  %+v", 200, o.Payload)
}
func (o *GetPacketCaptureResultOK) GetPayload() *models.PacketCaptureResult {
	return o.Payload
}

func (o *GetPacketCaptureResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PacketCaptureResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

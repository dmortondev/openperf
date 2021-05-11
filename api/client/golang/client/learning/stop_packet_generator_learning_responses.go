// Code generated by go-swagger; DO NOT EDIT.

package learning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StopPacketGeneratorLearningReader is a Reader for the StopPacketGeneratorLearning structure.
type StopPacketGeneratorLearningReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPacketGeneratorLearningReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStopPacketGeneratorLearningNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopPacketGeneratorLearningNoContent creates a StopPacketGeneratorLearningNoContent with default headers values
func NewStopPacketGeneratorLearningNoContent() *StopPacketGeneratorLearningNoContent {
	return &StopPacketGeneratorLearningNoContent{}
}

/* StopPacketGeneratorLearningNoContent describes a response with status code 204, with default header values.

No Content
*/
type StopPacketGeneratorLearningNoContent struct {
}

func (o *StopPacketGeneratorLearningNoContent) Error() string {
	return fmt.Sprintf("[POST /packet/generators/{id}/learning/stop][%d] stopPacketGeneratorLearningNoContent ", 204)
}

func (o *StopPacketGeneratorLearningNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

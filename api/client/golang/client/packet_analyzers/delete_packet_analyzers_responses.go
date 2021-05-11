// Code generated by go-swagger; DO NOT EDIT.

package packet_analyzers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePacketAnalyzersReader is a Reader for the DeletePacketAnalyzers structure.
type DeletePacketAnalyzersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePacketAnalyzersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePacketAnalyzersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePacketAnalyzersNoContent creates a DeletePacketAnalyzersNoContent with default headers values
func NewDeletePacketAnalyzersNoContent() *DeletePacketAnalyzersNoContent {
	return &DeletePacketAnalyzersNoContent{}
}

/* DeletePacketAnalyzersNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeletePacketAnalyzersNoContent struct {
}

func (o *DeletePacketAnalyzersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /packet/analyzers][%d] deletePacketAnalyzersNoContent ", 204)
}

func (o *DeletePacketAnalyzersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

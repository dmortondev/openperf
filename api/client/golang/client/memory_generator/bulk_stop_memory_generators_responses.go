// Code generated by go-swagger; DO NOT EDIT.

package memory_generator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BulkStopMemoryGeneratorsReader is a Reader for the BulkStopMemoryGenerators structure.
type BulkStopMemoryGeneratorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkStopMemoryGeneratorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewBulkStopMemoryGeneratorsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBulkStopMemoryGeneratorsNoContent creates a BulkStopMemoryGeneratorsNoContent with default headers values
func NewBulkStopMemoryGeneratorsNoContent() *BulkStopMemoryGeneratorsNoContent {
	return &BulkStopMemoryGeneratorsNoContent{}
}

/* BulkStopMemoryGeneratorsNoContent describes a response with status code 204, with default header values.

No Content
*/
type BulkStopMemoryGeneratorsNoContent struct {
}

func (o *BulkStopMemoryGeneratorsNoContent) Error() string {
	return fmt.Sprintf("[POST /memory-generators/x/bulk-stop][%d] bulkStopMemoryGeneratorsNoContent ", 204)
}

func (o *BulkStopMemoryGeneratorsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*BulkStopMemoryGeneratorsBody BulkStopMemoryGeneratorsRequest
//
// Parameters for the bulk stop operation
swagger:model BulkStopMemoryGeneratorsBody
*/
type BulkStopMemoryGeneratorsBody struct {

	// List of memory generator identifiers
	// Required: true
	// Min Items: 1
	Ids []string `json:"ids"`
}

// Validate validates this bulk stop memory generators body
func (o *BulkStopMemoryGeneratorsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BulkStopMemoryGeneratorsBody) validateIds(formats strfmt.Registry) error {

	if err := validate.Required("bulk_stop"+"."+"ids", "body", o.Ids); err != nil {
		return err
	}

	iIdsSize := int64(len(o.Ids))

	if err := validate.MinItems("bulk_stop"+"."+"ids", "body", iIdsSize, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this bulk stop memory generators body based on context it is used
func (o *BulkStopMemoryGeneratorsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BulkStopMemoryGeneratorsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BulkStopMemoryGeneratorsBody) UnmarshalBinary(b []byte) error {
	var res BulkStopMemoryGeneratorsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

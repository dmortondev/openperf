// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BulkStartMemoryGeneratorsRequest Parameters for the bulk start operation
//
// swagger:model BulkStartMemoryGeneratorsRequest
type BulkStartMemoryGeneratorsRequest struct {

	// dynamic results
	DynamicResults *DynamicResultsConfig `json:"dynamic_results,omitempty"`

	// List of memory generator identifiers
	// Required: true
	// Min Items: 1
	Ids []string `json:"ids"`
}

// Validate validates this bulk start memory generators request
func (m *BulkStartMemoryGeneratorsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDynamicResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkStartMemoryGeneratorsRequest) validateDynamicResults(formats strfmt.Registry) error {
	if swag.IsZero(m.DynamicResults) { // not required
		return nil
	}

	if m.DynamicResults != nil {
		if err := m.DynamicResults.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic_results")
			}
			return err
		}
	}

	return nil
}

func (m *BulkStartMemoryGeneratorsRequest) validateIds(formats strfmt.Registry) error {

	if err := validate.Required("ids", "body", m.Ids); err != nil {
		return err
	}

	iIdsSize := int64(len(m.Ids))

	if err := validate.MinItems("ids", "body", iIdsSize, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bulk start memory generators request based on the context it is used
func (m *BulkStartMemoryGeneratorsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDynamicResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkStartMemoryGeneratorsRequest) contextValidateDynamicResults(ctx context.Context, formats strfmt.Registry) error {

	if m.DynamicResults != nil {
		if err := m.DynamicResults.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic_results")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkStartMemoryGeneratorsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkStartMemoryGeneratorsRequest) UnmarshalBinary(b []byte) error {
	var res BulkStartMemoryGeneratorsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

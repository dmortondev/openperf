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

// NetworkGeneratorStats Network generator statistics
//
// swagger:model NetworkGeneratorStats
type NetworkGeneratorStats struct {

	// The actual number of bytes read or written
	// Required: true
	// Minimum: 0
	BytesActual *int64 `json:"bytes_actual"`

	// The intended number of bytes read or written
	// Required: true
	// Minimum: 0
	BytesTarget *int64 `json:"bytes_target"`

	// The number of io_errors observed during reading or writing
	// Required: true
	// Minimum: 0
	IoErrors *int64 `json:"io_errors"`

	// The maximum observed latency value (in nanoseconds)
	// Minimum: 0
	LatencyMax *int64 `json:"latency_max,omitempty"`

	// The minimum observed latency value (in nanoseconds)
	// Minimum: 0
	LatencyMin *int64 `json:"latency_min,omitempty"`

	// The total amount of time required to perform all operations (in nanoseconds)
	// Required: true
	// Minimum: 0
	LatencyTotal *int64 `json:"latency_total"`

	// The actual number of operations performed
	// Required: true
	// Minimum: 0
	OpsActual *int64 `json:"ops_actual"`

	// The intended number of operations performed
	// Required: true
	// Minimum: 0
	OpsTarget *int64 `json:"ops_target"`
}

// Validate validates this network generator stats
func (m *NetworkGeneratorStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytesActual(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBytesTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencyMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencyMin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencyTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpsActual(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpsTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkGeneratorStats) validateBytesActual(formats strfmt.Registry) error {

	if err := validate.Required("bytes_actual", "body", m.BytesActual); err != nil {
		return err
	}

	if err := validate.MinimumInt("bytes_actual", "body", *m.BytesActual, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkGeneratorStats) validateBytesTarget(formats strfmt.Registry) error {

	if err := validate.Required("bytes_target", "body", m.BytesTarget); err != nil {
		return err
	}

	if err := validate.MinimumInt("bytes_target", "body", *m.BytesTarget, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkGeneratorStats) validateIoErrors(formats strfmt.Registry) error {

	if err := validate.Required("io_errors", "body", m.IoErrors); err != nil {
		return err
	}

	if err := validate.MinimumInt("io_errors", "body", *m.IoErrors, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkGeneratorStats) validateLatencyMax(formats strfmt.Registry) error {
	if swag.IsZero(m.LatencyMax) { // not required
		return nil
	}

	if err := validate.MinimumInt("latency_max", "body", *m.LatencyMax, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkGeneratorStats) validateLatencyMin(formats strfmt.Registry) error {
	if swag.IsZero(m.LatencyMin) { // not required
		return nil
	}

	if err := validate.MinimumInt("latency_min", "body", *m.LatencyMin, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkGeneratorStats) validateLatencyTotal(formats strfmt.Registry) error {

	if err := validate.Required("latency_total", "body", m.LatencyTotal); err != nil {
		return err
	}

	if err := validate.MinimumInt("latency_total", "body", *m.LatencyTotal, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkGeneratorStats) validateOpsActual(formats strfmt.Registry) error {

	if err := validate.Required("ops_actual", "body", m.OpsActual); err != nil {
		return err
	}

	if err := validate.MinimumInt("ops_actual", "body", *m.OpsActual, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkGeneratorStats) validateOpsTarget(formats strfmt.Registry) error {

	if err := validate.Required("ops_target", "body", m.OpsTarget); err != nil {
		return err
	}

	if err := validate.MinimumInt("ops_target", "body", *m.OpsTarget, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network generator stats based on context it is used
func (m *NetworkGeneratorStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkGeneratorStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkGeneratorStats) UnmarshalBinary(b []byte) error {
	var res NetworkGeneratorStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PacketCaptureResult Packet capture results
//
// swagger:model PacketCaptureResult
type PacketCaptureResult struct {

	// Indicates whether this object is currently capturing packets or not.
	//
	Active bool `json:"active,omitempty"`

	// Number of bytes captured
	// Required: true
	// Minimum: 0
	Bytes *int64 `json:"bytes"`

	// Unique capture identifier that generated this result
	// Required: true
	CaptureID *string `json:"capture_id"`

	// Unique capture result identifier
	// Required: true
	ID *string `json:"id"`

	// Number of packets captured
	// Required: true
	// Minimum: 0
	Packets *int64 `json:"packets"`

	// Capture state
	// Required: true
	// Enum: [stopped armed running]
	State *string `json:"state"`
}

// Validate validates this packet capture result
func (m *PacketCaptureResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaptureID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PacketCaptureResult) validateBytes(formats strfmt.Registry) error {

	if err := validate.Required("bytes", "body", m.Bytes); err != nil {
		return err
	}

	if err := validate.MinimumInt("bytes", "body", *m.Bytes, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PacketCaptureResult) validateCaptureID(formats strfmt.Registry) error {

	if err := validate.Required("capture_id", "body", m.CaptureID); err != nil {
		return err
	}

	return nil
}

func (m *PacketCaptureResult) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PacketCaptureResult) validatePackets(formats strfmt.Registry) error {

	if err := validate.Required("packets", "body", m.Packets); err != nil {
		return err
	}

	if err := validate.MinimumInt("packets", "body", *m.Packets, 0, false); err != nil {
		return err
	}

	return nil
}

var packetCaptureResultTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["stopped","armed","running"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packetCaptureResultTypeStatePropEnum = append(packetCaptureResultTypeStatePropEnum, v)
	}
}

const (

	// PacketCaptureResultStateStopped captures enum value "stopped"
	PacketCaptureResultStateStopped string = "stopped"

	// PacketCaptureResultStateArmed captures enum value "armed"
	PacketCaptureResultStateArmed string = "armed"

	// PacketCaptureResultStateRunning captures enum value "running"
	PacketCaptureResultStateRunning string = "running"
)

// prop value enum
func (m *PacketCaptureResult) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, packetCaptureResultTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PacketCaptureResult) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this packet capture result based on context it is used
func (m *PacketCaptureResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PacketCaptureResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PacketCaptureResult) UnmarshalBinary(b []byte) error {
	var res PacketCaptureResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

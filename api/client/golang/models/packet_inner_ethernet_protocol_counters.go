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

// PacketInnerEthernetProtocolCounters Inner layer 2 protocol counters
//
// swagger:model PacketInnerEthernetProtocolCounters
type PacketInnerEthernetProtocolCounters struct {

	// Number of IPv6/IPv6 frames
	// Required: true
	// Minimum: 0
	IP *int64 `json:"ip"`

	// Number of Queue-in-Queue frames
	// Required: true
	// Minimum: 0
	Qinq *int64 `json:"qinq"`

	// Number of Virtual LAN frames
	// Required: true
	// Minimum: 0
	Vlan *int64 `json:"vlan"`
}

// Validate validates this packet inner ethernet protocol counters
func (m *PacketInnerEthernetProtocolCounters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQinq(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PacketInnerEthernetProtocolCounters) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	if err := validate.MinimumInt("ip", "body", *m.IP, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PacketInnerEthernetProtocolCounters) validateQinq(formats strfmt.Registry) error {

	if err := validate.Required("qinq", "body", m.Qinq); err != nil {
		return err
	}

	if err := validate.MinimumInt("qinq", "body", *m.Qinq, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PacketInnerEthernetProtocolCounters) validateVlan(formats strfmt.Registry) error {

	if err := validate.Required("vlan", "body", m.Vlan); err != nil {
		return err
	}

	if err := validate.MinimumInt("vlan", "body", *m.Vlan, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this packet inner ethernet protocol counters based on context it is used
func (m *PacketInnerEthernetProtocolCounters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PacketInnerEthernetProtocolCounters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PacketInnerEthernetProtocolCounters) UnmarshalBinary(b []byte) error {
	var res PacketInnerEthernetProtocolCounters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

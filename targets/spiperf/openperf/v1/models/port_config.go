// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PortConfig Port configuration
// swagger:model PortConfig
type PortConfig struct {

	// bond
	Bond *PortConfigBond `json:"bond,omitempty"`

	// dpdk
	Dpdk *PortConfigDpdk `json:"dpdk,omitempty"`
}

// Validate validates this port config
func (m *PortConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDpdk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortConfig) validateBond(formats strfmt.Registry) error {

	if swag.IsZero(m.Bond) { // not required
		return nil
	}

	if m.Bond != nil {
		if err := m.Bond.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bond")
			}
			return err
		}
	}

	return nil
}

func (m *PortConfig) validateDpdk(formats strfmt.Registry) error {

	if swag.IsZero(m.Dpdk) { // not required
		return nil
	}

	if m.Dpdk != nil {
		if err := m.Dpdk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dpdk")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortConfig) UnmarshalBinary(b []byte) error {
	var res PortConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortConfigBond Bond-specific configuration information
// swagger:model PortConfigBond
type PortConfigBond struct {

	// Port bonding mode
	// Required: true
	// Enum: [lag_802_3_ad]
	Mode *string `json:"mode"`

	// Unique identifiers of bonded ports
	// Required: true
	// Min Items: 1
	Ports []string `json:"ports"`
}

// Validate validates this port config bond
func (m *PortConfigBond) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var portConfigBondTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["lag_802_3_ad"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portConfigBondTypeModePropEnum = append(portConfigBondTypeModePropEnum, v)
	}
}

const (

	// PortConfigBondModeLag8023Ad captures enum value "lag_802_3_ad"
	PortConfigBondModeLag8023Ad string = "lag_802_3_ad"
)

// prop value enum
func (m *PortConfigBond) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, portConfigBondTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PortConfigBond) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("bond"+"."+"mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("bond"+"."+"mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *PortConfigBond) validatePorts(formats strfmt.Registry) error {

	if err := validate.Required("bond"+"."+"ports", "body", m.Ports); err != nil {
		return err
	}

	iPortsSize := int64(len(m.Ports))

	if err := validate.MinItems("bond"+"."+"ports", "body", iPortsSize, 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortConfigBond) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortConfigBond) UnmarshalBinary(b []byte) error {
	var res PortConfigBond
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortConfigDpdk DPDK-specific configuration information
// swagger:model PortConfigDpdk
type PortConfigDpdk struct {

	// Enable link auto-negotiation
	// Required: true
	AutoNegotiation *bool `json:"auto_negotiation"`

	// Manually-configured port duplex
	// Enum: [full half]
	Duplex string `json:"duplex,omitempty"`

	// Manually-configured port speed (in Mbps)
	Speed int64 `json:"speed,omitempty"`
}

// Validate validates this port config dpdk
func (m *PortConfigDpdk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoNegotiation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuplex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortConfigDpdk) validateAutoNegotiation(formats strfmt.Registry) error {

	if err := validate.Required("dpdk"+"."+"auto_negotiation", "body", m.AutoNegotiation); err != nil {
		return err
	}

	return nil
}

var portConfigDpdkTypeDuplexPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full","half"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portConfigDpdkTypeDuplexPropEnum = append(portConfigDpdkTypeDuplexPropEnum, v)
	}
}

const (

	// PortConfigDpdkDuplexFull captures enum value "full"
	PortConfigDpdkDuplexFull string = "full"

	// PortConfigDpdkDuplexHalf captures enum value "half"
	PortConfigDpdkDuplexHalf string = "half"
)

// prop value enum
func (m *PortConfigDpdk) validateDuplexEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, portConfigDpdkTypeDuplexPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PortConfigDpdk) validateDuplex(formats strfmt.Registry) error {

	if swag.IsZero(m.Duplex) { // not required
		return nil
	}

	// value enum
	if err := m.validateDuplexEnum("dpdk"+"."+"duplex", "body", m.Duplex); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortConfigDpdk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortConfigDpdk) UnmarshalBinary(b []byte) error {
	var res PortConfigDpdk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// HardwarePlatform Hardware platform detailing its limits and statistics
// swagger:model HardwarePlatform
type HardwarePlatform struct {

	// Description
	Description string `json:"description,omitempty"`

	// The DataCenter list of servers and their available resources
	HostsResources []*HostResources `json:"hostsResources"`

	// Configured Memory GB
	Memory float64 `json:"memory,omitempty"`

	// Processor to Memory (GB) Ratio
	ProcessorMemoryRatio float64 `json:"processorMemoryRatio,omitempty"`

	// Configured Processors
	Processors float64 `json:"processors,omitempty"`

	// Allowable granularity for shared processors
	SharedProcessorStep float64 `json:"sharedProcessorStep,omitempty"`

	// Short code for hardware
	Type string `json:"type,omitempty"`
}

// Validate validates this hardware platform
func (m *HardwarePlatform) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostsResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HardwarePlatform) validateHostsResources(formats strfmt.Registry) error {

	if swag.IsZero(m.HostsResources) { // not required
		return nil
	}

	for i := 0; i < len(m.HostsResources); i++ {
		if swag.IsZero(m.HostsResources[i]) { // not required
			continue
		}

		if m.HostsResources[i] != nil {
			if err := m.HostsResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostsResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HardwarePlatform) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HardwarePlatform) UnmarshalBinary(b []byte) error {
	var res HardwarePlatform
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
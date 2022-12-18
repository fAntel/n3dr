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

// RoutingRuleXO routing rule x o
//
// swagger:model RoutingRuleXO
type RoutingRuleXO struct {

	// description
	Description string `json:"description,omitempty"`

	// Regular expressions used to identify request paths that are allowed or blocked (depending on mode)
	Matchers []string `json:"matchers"`

	// Determines what should be done with requests when their path matches any of the matchers
	// Enum: [BLOCK ALLOW]
	Mode string `json:"mode,omitempty"`

	// name
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this routing rule x o
func (m *RoutingRuleXO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var routingRuleXOTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BLOCK","ALLOW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routingRuleXOTypeModePropEnum = append(routingRuleXOTypeModePropEnum, v)
	}
}

const (

	// RoutingRuleXOModeBLOCK captures enum value "BLOCK"
	RoutingRuleXOModeBLOCK string = "BLOCK"

	// RoutingRuleXOModeALLOW captures enum value "ALLOW"
	RoutingRuleXOModeALLOW string = "ALLOW"
)

// prop value enum
func (m *RoutingRuleXO) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routingRuleXOTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RoutingRuleXO) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *RoutingRuleXO) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this routing rule x o based on context it is used
func (m *RoutingRuleXO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoutingRuleXO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutingRuleXO) UnmarshalBinary(b []byte) error {
	var res RoutingRuleXO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
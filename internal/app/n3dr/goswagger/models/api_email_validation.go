// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIEmailValidation Api email validation
//
// swagger:model ApiEmailValidation
type APIEmailValidation struct {

	// reason
	Reason string `json:"reason,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this Api email validation
func (m *APIEmailValidation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api email validation based on context it is used
func (m *APIEmailValidation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIEmailValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIEmailValidation) UnmarshalBinary(b []byte) error {
	var res APIEmailValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
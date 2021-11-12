// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository10Reader is a Reader for the CreateRepository10 structure.
type CreateRepository10Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository10Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository10Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository10Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository10Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository10Created creates a CreateRepository10Created with default headers values
func NewCreateRepository10Created() *CreateRepository10Created {
	return &CreateRepository10Created{}
}

/* CreateRepository10Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository10Created struct {
}

func (o *CreateRepository10Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/proxy][%d] createRepository10Created ", 201)
}

func (o *CreateRepository10Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository10Unauthorized creates a CreateRepository10Unauthorized with default headers values
func NewCreateRepository10Unauthorized() *CreateRepository10Unauthorized {
	return &CreateRepository10Unauthorized{}
}

/* CreateRepository10Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository10Unauthorized struct {
}

func (o *CreateRepository10Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/proxy][%d] createRepository10Unauthorized ", 401)
}

func (o *CreateRepository10Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository10Forbidden creates a CreateRepository10Forbidden with default headers values
func NewCreateRepository10Forbidden() *CreateRepository10Forbidden {
	return &CreateRepository10Forbidden{}
}

/* CreateRepository10Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository10Forbidden struct {
}

func (o *CreateRepository10Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/npm/proxy][%d] createRepository10Forbidden ", 403)
}

func (o *CreateRepository10Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

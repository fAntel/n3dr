// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateRepository25Reader is a Reader for the CreateRepository25 structure.
type CreateRepository25Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRepository25Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRepository25Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRepository25Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRepository25Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRepository25Created creates a CreateRepository25Created with default headers values
func NewCreateRepository25Created() *CreateRepository25Created {
	return &CreateRepository25Created{}
}

/*
CreateRepository25Created describes a response with status code 201, with default header values.

Repository created
*/
type CreateRepository25Created struct {
}

// IsSuccess returns true when this create repository25 created response has a 2xx status code
func (o *CreateRepository25Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create repository25 created response has a 3xx status code
func (o *CreateRepository25Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository25 created response has a 4xx status code
func (o *CreateRepository25Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create repository25 created response has a 5xx status code
func (o *CreateRepository25Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository25 created response a status code equal to that given
func (o *CreateRepository25Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create repository25 created response
func (o *CreateRepository25Created) Code() int {
	return 201
}

func (o *CreateRepository25Created) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/gitlfs/hosted][%d] createRepository25Created ", 201)
}

func (o *CreateRepository25Created) String() string {
	return fmt.Sprintf("[POST /v1/repositories/gitlfs/hosted][%d] createRepository25Created ", 201)
}

func (o *CreateRepository25Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository25Unauthorized creates a CreateRepository25Unauthorized with default headers values
func NewCreateRepository25Unauthorized() *CreateRepository25Unauthorized {
	return &CreateRepository25Unauthorized{}
}

/*
CreateRepository25Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type CreateRepository25Unauthorized struct {
}

// IsSuccess returns true when this create repository25 unauthorized response has a 2xx status code
func (o *CreateRepository25Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository25 unauthorized response has a 3xx status code
func (o *CreateRepository25Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository25 unauthorized response has a 4xx status code
func (o *CreateRepository25Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository25 unauthorized response has a 5xx status code
func (o *CreateRepository25Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository25 unauthorized response a status code equal to that given
func (o *CreateRepository25Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create repository25 unauthorized response
func (o *CreateRepository25Unauthorized) Code() int {
	return 401
}

func (o *CreateRepository25Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/gitlfs/hosted][%d] createRepository25Unauthorized ", 401)
}

func (o *CreateRepository25Unauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/gitlfs/hosted][%d] createRepository25Unauthorized ", 401)
}

func (o *CreateRepository25Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRepository25Forbidden creates a CreateRepository25Forbidden with default headers values
func NewCreateRepository25Forbidden() *CreateRepository25Forbidden {
	return &CreateRepository25Forbidden{}
}

/*
CreateRepository25Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type CreateRepository25Forbidden struct {
}

// IsSuccess returns true when this create repository25 forbidden response has a 2xx status code
func (o *CreateRepository25Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create repository25 forbidden response has a 3xx status code
func (o *CreateRepository25Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create repository25 forbidden response has a 4xx status code
func (o *CreateRepository25Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create repository25 forbidden response has a 5xx status code
func (o *CreateRepository25Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create repository25 forbidden response a status code equal to that given
func (o *CreateRepository25Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create repository25 forbidden response
func (o *CreateRepository25Forbidden) Code() int {
	return 403
}

func (o *CreateRepository25Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/gitlfs/hosted][%d] createRepository25Forbidden ", 403)
}

func (o *CreateRepository25Forbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/gitlfs/hosted][%d] createRepository25Forbidden ", 403)
}

func (o *CreateRepository25Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

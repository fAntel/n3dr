// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRepository20Reader is a Reader for the UpdateRepository20 structure.
type UpdateRepository20Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepository20Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRepository20NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateRepository20Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepository20Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepository20NoContent creates a UpdateRepository20NoContent with default headers values
func NewUpdateRepository20NoContent() *UpdateRepository20NoContent {
	return &UpdateRepository20NoContent{}
}

/* UpdateRepository20NoContent describes a response with status code 204, with default header values.

Repository updated
*/
type UpdateRepository20NoContent struct {
}

func (o *UpdateRepository20NoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] updateRepository20NoContent ", 204)
}

func (o *UpdateRepository20NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository20Unauthorized creates a UpdateRepository20Unauthorized with default headers values
func NewUpdateRepository20Unauthorized() *UpdateRepository20Unauthorized {
	return &UpdateRepository20Unauthorized{}
}

/* UpdateRepository20Unauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type UpdateRepository20Unauthorized struct {
}

func (o *UpdateRepository20Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] updateRepository20Unauthorized ", 401)
}

func (o *UpdateRepository20Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRepository20Forbidden creates a UpdateRepository20Forbidden with default headers values
func NewUpdateRepository20Forbidden() *UpdateRepository20Forbidden {
	return &UpdateRepository20Forbidden{}
}

/* UpdateRepository20Forbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type UpdateRepository20Forbidden struct {
}

func (o *UpdateRepository20Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/repositories/yum/group/{repositoryName}][%d] updateRepository20Forbidden ", 403)
}

func (o *UpdateRepository20Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

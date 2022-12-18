// Code generated by go-swagger; DO NOT EDIT.

package security_management_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewUpdateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserBadRequest creates a UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {
	return &UpdateUserBadRequest{}
}

/* UpdateUserBadRequest describes a response with status code 400, with default header values.

Password was not supplied in the body of the request
*/
type UpdateUserBadRequest struct {
}

func (o *UpdateUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/security/users/{userId}][%d] updateUserBadRequest ", 400)
}

func (o *UpdateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserForbidden creates a UpdateUserForbidden with default headers values
func NewUpdateUserForbidden() *UpdateUserForbidden {
	return &UpdateUserForbidden{}
}

/* UpdateUserForbidden describes a response with status code 403, with default header values.

The user does not have permission to perform the operation.
*/
type UpdateUserForbidden struct {
}

func (o *UpdateUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/security/users/{userId}][%d] updateUserForbidden ", 403)
}

func (o *UpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserNotFound creates a UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {
	return &UpdateUserNotFound{}
}

/* UpdateUserNotFound describes a response with status code 404, with default header values.

User or user source not found in the system.
*/
type UpdateUserNotFound struct {
}

func (o *UpdateUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/security/users/{userId}][%d] updateUserNotFound ", 404)
}

func (o *UpdateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
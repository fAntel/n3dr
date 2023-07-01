// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IsWritableReader is a Reader for the IsWritable structure.
type IsWritableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsWritableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsWritableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewIsWritableServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIsWritableOK creates a IsWritableOK with default headers values
func NewIsWritableOK() *IsWritableOK {
	return &IsWritableOK{}
}

/*
IsWritableOK describes a response with status code 200, with default header values.

Available to service requests
*/
type IsWritableOK struct {
}

// IsSuccess returns true when this is writable o k response has a 2xx status code
func (o *IsWritableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this is writable o k response has a 3xx status code
func (o *IsWritableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is writable o k response has a 4xx status code
func (o *IsWritableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this is writable o k response has a 5xx status code
func (o *IsWritableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this is writable o k response a status code equal to that given
func (o *IsWritableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the is writable o k response
func (o *IsWritableOK) Code() int {
	return 200
}

func (o *IsWritableOK) Error() string {
	return fmt.Sprintf("[GET /v1/status/writable][%d] isWritableOK ", 200)
}

func (o *IsWritableOK) String() string {
	return fmt.Sprintf("[GET /v1/status/writable][%d] isWritableOK ", 200)
}

func (o *IsWritableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsWritableServiceUnavailable creates a IsWritableServiceUnavailable with default headers values
func NewIsWritableServiceUnavailable() *IsWritableServiceUnavailable {
	return &IsWritableServiceUnavailable{}
}

/*
IsWritableServiceUnavailable describes a response with status code 503, with default header values.

Unavailable to service requests
*/
type IsWritableServiceUnavailable struct {
}

// IsSuccess returns true when this is writable service unavailable response has a 2xx status code
func (o *IsWritableServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this is writable service unavailable response has a 3xx status code
func (o *IsWritableServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this is writable service unavailable response has a 4xx status code
func (o *IsWritableServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this is writable service unavailable response has a 5xx status code
func (o *IsWritableServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this is writable service unavailable response a status code equal to that given
func (o *IsWritableServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the is writable service unavailable response
func (o *IsWritableServiceUnavailable) Code() int {
	return 503
}

func (o *IsWritableServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/status/writable][%d] isWritableServiceUnavailable ", 503)
}

func (o *IsWritableServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/status/writable][%d] isWritableServiceUnavailable ", 503)
}

func (o *IsWritableServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

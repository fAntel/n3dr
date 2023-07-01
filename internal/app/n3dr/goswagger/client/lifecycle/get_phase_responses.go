// Code generated by go-swagger; DO NOT EDIT.

package lifecycle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetPhaseReader is a Reader for the GetPhase structure.
type GetPhaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPhaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPhaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPhaseOK creates a GetPhaseOK with default headers values
func NewGetPhaseOK() *GetPhaseOK {
	return &GetPhaseOK{}
}

/*
GetPhaseOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPhaseOK struct {
	Payload string
}

// IsSuccess returns true when this get phase o k response has a 2xx status code
func (o *GetPhaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get phase o k response has a 3xx status code
func (o *GetPhaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get phase o k response has a 4xx status code
func (o *GetPhaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get phase o k response has a 5xx status code
func (o *GetPhaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get phase o k response a status code equal to that given
func (o *GetPhaseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get phase o k response
func (o *GetPhaseOK) Code() int {
	return 200
}

func (o *GetPhaseOK) Error() string {
	return fmt.Sprintf("[GET /v1/lifecycle/phase][%d] getPhaseOK  %+v", 200, o.Payload)
}

func (o *GetPhaseOK) String() string {
	return fmt.Sprintf("[GET /v1/lifecycle/phase][%d] getPhaseOK  %+v", 200, o.Payload)
}

func (o *GetPhaseOK) GetPayload() string {
	return o.Payload
}

func (o *GetPhaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

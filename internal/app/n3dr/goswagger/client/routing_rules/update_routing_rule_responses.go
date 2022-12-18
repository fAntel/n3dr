// Code generated by go-swagger; DO NOT EDIT.

package routing_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateRoutingRuleReader is a Reader for the UpdateRoutingRule structure.
type UpdateRoutingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoutingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRoutingRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRoutingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRoutingRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRoutingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRoutingRuleNoContent creates a UpdateRoutingRuleNoContent with default headers values
func NewUpdateRoutingRuleNoContent() *UpdateRoutingRuleNoContent {
	return &UpdateRoutingRuleNoContent{}
}

/* UpdateRoutingRuleNoContent describes a response with status code 204, with default header values.

Routing rule was successfully updated
*/
type UpdateRoutingRuleNoContent struct {
}

func (o *UpdateRoutingRuleNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/routing-rules/{name}][%d] updateRoutingRuleNoContent ", 204)
}

func (o *UpdateRoutingRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRoutingRuleBadRequest creates a UpdateRoutingRuleBadRequest with default headers values
func NewUpdateRoutingRuleBadRequest() *UpdateRoutingRuleBadRequest {
	return &UpdateRoutingRuleBadRequest{}
}

/* UpdateRoutingRuleBadRequest describes a response with status code 400, with default header values.

Another routing rule with the same name already exists or required parameters missing
*/
type UpdateRoutingRuleBadRequest struct {
}

func (o *UpdateRoutingRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/routing-rules/{name}][%d] updateRoutingRuleBadRequest ", 400)
}

func (o *UpdateRoutingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRoutingRuleForbidden creates a UpdateRoutingRuleForbidden with default headers values
func NewUpdateRoutingRuleForbidden() *UpdateRoutingRuleForbidden {
	return &UpdateRoutingRuleForbidden{}
}

/* UpdateRoutingRuleForbidden describes a response with status code 403, with default header values.

Insufficient permissions to edit routing rules
*/
type UpdateRoutingRuleForbidden struct {
}

func (o *UpdateRoutingRuleForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/routing-rules/{name}][%d] updateRoutingRuleForbidden ", 403)
}

func (o *UpdateRoutingRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRoutingRuleNotFound creates a UpdateRoutingRuleNotFound with default headers values
func NewUpdateRoutingRuleNotFound() *UpdateRoutingRuleNotFound {
	return &UpdateRoutingRuleNotFound{}
}

/* UpdateRoutingRuleNotFound describes a response with status code 404, with default header values.

Routing rule not found
*/
type UpdateRoutingRuleNotFound struct {
}

func (o *UpdateRoutingRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/routing-rules/{name}][%d] updateRoutingRuleNotFound ", 404)
}

func (o *UpdateRoutingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
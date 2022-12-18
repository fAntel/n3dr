// Code generated by go-swagger; DO NOT EDIT.

package routing_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// GetRoutingRuleReader is a Reader for the GetRoutingRule structure.
type GetRoutingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRoutingRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingRuleOK creates a GetRoutingRuleOK with default headers values
func NewGetRoutingRuleOK() *GetRoutingRuleOK {
	return &GetRoutingRuleOK{}
}

/* GetRoutingRuleOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingRuleOK struct {
	Payload *models.RoutingRuleXO
}

func (o *GetRoutingRuleOK) Error() string {
	return fmt.Sprintf("[GET /v1/routing-rules/{name}][%d] getRoutingRuleOK  %+v", 200, o.Payload)
}
func (o *GetRoutingRuleOK) GetPayload() *models.RoutingRuleXO {
	return o.Payload
}

func (o *GetRoutingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoutingRuleXO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingRuleForbidden creates a GetRoutingRuleForbidden with default headers values
func NewGetRoutingRuleForbidden() *GetRoutingRuleForbidden {
	return &GetRoutingRuleForbidden{}
}

/* GetRoutingRuleForbidden describes a response with status code 403, with default header values.

Insufficient permissions to read routing rules
*/
type GetRoutingRuleForbidden struct {
}

func (o *GetRoutingRuleForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/routing-rules/{name}][%d] getRoutingRuleForbidden ", 403)
}

func (o *GetRoutingRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoutingRuleNotFound creates a GetRoutingRuleNotFound with default headers values
func NewGetRoutingRuleNotFound() *GetRoutingRuleNotFound {
	return &GetRoutingRuleNotFound{}
}

/* GetRoutingRuleNotFound describes a response with status code 404, with default header values.

Routing rule not found
*/
type GetRoutingRuleNotFound struct {
}

func (o *GetRoutingRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/routing-rules/{name}][%d] getRoutingRuleNotFound ", 404)
}

func (o *GetRoutingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
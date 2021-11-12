// Code generated by go-swagger; DO NOT EDIT.

package security_certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/go-swagger/models"
)

// AddCertificateReader is a Reader for the AddCertificate structure.
type AddCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddCertificateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAddCertificateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddCertificateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddCertificateCreated creates a AddCertificateCreated with default headers values
func NewAddCertificateCreated() *AddCertificateCreated {
	return &AddCertificateCreated{}
}

/* AddCertificateCreated describes a response with status code 201, with default header values.

The certificate was successfully added.
*/
type AddCertificateCreated struct {
	Payload *models.APICertificate
}

func (o *AddCertificateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/security/ssl/truststore][%d] addCertificateCreated  %+v", 201, o.Payload)
}
func (o *AddCertificateCreated) GetPayload() *models.APICertificate {
	return o.Payload
}

func (o *AddCertificateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCertificateForbidden creates a AddCertificateForbidden with default headers values
func NewAddCertificateForbidden() *AddCertificateForbidden {
	return &AddCertificateForbidden{}
}

/* AddCertificateForbidden describes a response with status code 403, with default header values.

Insufficient permissions to add certificate to the trust store.
*/
type AddCertificateForbidden struct {
}

func (o *AddCertificateForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/security/ssl/truststore][%d] addCertificateForbidden ", 403)
}

func (o *AddCertificateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddCertificateConflict creates a AddCertificateConflict with default headers values
func NewAddCertificateConflict() *AddCertificateConflict {
	return &AddCertificateConflict{}
}

/* AddCertificateConflict describes a response with status code 409, with default header values.

The certificate already exists in the system.
*/
type AddCertificateConflict struct {
}

func (o *AddCertificateConflict) Error() string {
	return fmt.Sprintf("[POST /v1/security/ssl/truststore][%d] addCertificateConflict ", 409)
}

func (o *AddCertificateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

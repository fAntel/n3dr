// Code generated by go-swagger; DO NOT EDIT.

package blob_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/app/n3dr/goswagger/models"
)

// GetFileBlobStoreConfigurationReader is a Reader for the GetFileBlobStoreConfiguration structure.
type GetFileBlobStoreConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileBlobStoreConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileBlobStoreConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFileBlobStoreConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFileBlobStoreConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFileBlobStoreConfigurationOK creates a GetFileBlobStoreConfigurationOK with default headers values
func NewGetFileBlobStoreConfigurationOK() *GetFileBlobStoreConfigurationOK {
	return &GetFileBlobStoreConfigurationOK{}
}

/*
GetFileBlobStoreConfigurationOK describes a response with status code 200, with default header values.

Success
*/
type GetFileBlobStoreConfigurationOK struct {
	Payload *models.FileBlobStoreAPIModel
}

// IsSuccess returns true when this get file blob store configuration o k response has a 2xx status code
func (o *GetFileBlobStoreConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file blob store configuration o k response has a 3xx status code
func (o *GetFileBlobStoreConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file blob store configuration o k response has a 4xx status code
func (o *GetFileBlobStoreConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file blob store configuration o k response has a 5xx status code
func (o *GetFileBlobStoreConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file blob store configuration o k response a status code equal to that given
func (o *GetFileBlobStoreConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file blob store configuration o k response
func (o *GetFileBlobStoreConfigurationOK) Code() int {
	return 200
}

func (o *GetFileBlobStoreConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/file/{name}][%d] getFileBlobStoreConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFileBlobStoreConfigurationOK) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/file/{name}][%d] getFileBlobStoreConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFileBlobStoreConfigurationOK) GetPayload() *models.FileBlobStoreAPIModel {
	return o.Payload
}

func (o *GetFileBlobStoreConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileBlobStoreAPIModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileBlobStoreConfigurationForbidden creates a GetFileBlobStoreConfigurationForbidden with default headers values
func NewGetFileBlobStoreConfigurationForbidden() *GetFileBlobStoreConfigurationForbidden {
	return &GetFileBlobStoreConfigurationForbidden{}
}

/*
GetFileBlobStoreConfigurationForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type GetFileBlobStoreConfigurationForbidden struct {
}

// IsSuccess returns true when this get file blob store configuration forbidden response has a 2xx status code
func (o *GetFileBlobStoreConfigurationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file blob store configuration forbidden response has a 3xx status code
func (o *GetFileBlobStoreConfigurationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file blob store configuration forbidden response has a 4xx status code
func (o *GetFileBlobStoreConfigurationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get file blob store configuration forbidden response has a 5xx status code
func (o *GetFileBlobStoreConfigurationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get file blob store configuration forbidden response a status code equal to that given
func (o *GetFileBlobStoreConfigurationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get file blob store configuration forbidden response
func (o *GetFileBlobStoreConfigurationForbidden) Code() int {
	return 403
}

func (o *GetFileBlobStoreConfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/file/{name}][%d] getFileBlobStoreConfigurationForbidden ", 403)
}

func (o *GetFileBlobStoreConfigurationForbidden) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/file/{name}][%d] getFileBlobStoreConfigurationForbidden ", 403)
}

func (o *GetFileBlobStoreConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFileBlobStoreConfigurationNotFound creates a GetFileBlobStoreConfigurationNotFound with default headers values
func NewGetFileBlobStoreConfigurationNotFound() *GetFileBlobStoreConfigurationNotFound {
	return &GetFileBlobStoreConfigurationNotFound{}
}

/*
GetFileBlobStoreConfigurationNotFound describes a response with status code 404, with default header values.

Blob store not found
*/
type GetFileBlobStoreConfigurationNotFound struct {
}

// IsSuccess returns true when this get file blob store configuration not found response has a 2xx status code
func (o *GetFileBlobStoreConfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file blob store configuration not found response has a 3xx status code
func (o *GetFileBlobStoreConfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file blob store configuration not found response has a 4xx status code
func (o *GetFileBlobStoreConfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get file blob store configuration not found response has a 5xx status code
func (o *GetFileBlobStoreConfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get file blob store configuration not found response a status code equal to that given
func (o *GetFileBlobStoreConfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get file blob store configuration not found response
func (o *GetFileBlobStoreConfigurationNotFound) Code() int {
	return 404
}

func (o *GetFileBlobStoreConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/blobstores/file/{name}][%d] getFileBlobStoreConfigurationNotFound ", 404)
}

func (o *GetFileBlobStoreConfigurationNotFound) String() string {
	return fmt.Sprintf("[GET /v1/blobstores/file/{name}][%d] getFileBlobStoreConfigurationNotFound ", 404)
}

func (o *GetFileBlobStoreConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

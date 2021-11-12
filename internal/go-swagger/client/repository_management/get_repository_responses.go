// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/030/n3dr/internal/go-swagger/models"
)

// GetRepositoryReader is a Reader for the GetRepository structure.
type GetRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRepositoryOK creates a GetRepositoryOK with default headers values
func NewGetRepositoryOK() *GetRepositoryOK {
	return &GetRepositoryOK{}
}

/* GetRepositoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRepositoryOK struct {
	Payload *models.SimpleAPIGroupRepository
}

func (o *GetRepositoryOK) Error() string {
	return fmt.Sprintf("[GET /v1/repositories/maven/group/{repositoryName}][%d] getRepositoryOK  %+v", 200, o.Payload)
}
func (o *GetRepositoryOK) GetPayload() *models.SimpleAPIGroupRepository {
	return o.Payload
}

func (o *GetRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleAPIGroupRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

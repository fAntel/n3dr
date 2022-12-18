// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAssetReader is a Reader for the DeleteAsset structure.
type DeleteAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAssetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteAssetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAssetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteAssetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAssetNoContent creates a DeleteAssetNoContent with default headers values
func NewDeleteAssetNoContent() *DeleteAssetNoContent {
	return &DeleteAssetNoContent{}
}

/* DeleteAssetNoContent describes a response with status code 204, with default header values.

Asset was successfully deleted
*/
type DeleteAssetNoContent struct {
}

func (o *DeleteAssetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/assets/{id}][%d] deleteAssetNoContent ", 204)
}

func (o *DeleteAssetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAssetForbidden creates a DeleteAssetForbidden with default headers values
func NewDeleteAssetForbidden() *DeleteAssetForbidden {
	return &DeleteAssetForbidden{}
}

/* DeleteAssetForbidden describes a response with status code 403, with default header values.

Insufficient permissions to delete asset
*/
type DeleteAssetForbidden struct {
}

func (o *DeleteAssetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/assets/{id}][%d] deleteAssetForbidden ", 403)
}

func (o *DeleteAssetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAssetNotFound creates a DeleteAssetNotFound with default headers values
func NewDeleteAssetNotFound() *DeleteAssetNotFound {
	return &DeleteAssetNotFound{}
}

/* DeleteAssetNotFound describes a response with status code 404, with default header values.

Asset not found
*/
type DeleteAssetNotFound struct {
}

func (o *DeleteAssetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/assets/{id}][%d] deleteAssetNotFound ", 404)
}

func (o *DeleteAssetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAssetUnprocessableEntity creates a DeleteAssetUnprocessableEntity with default headers values
func NewDeleteAssetUnprocessableEntity() *DeleteAssetUnprocessableEntity {
	return &DeleteAssetUnprocessableEntity{}
}

/* DeleteAssetUnprocessableEntity describes a response with status code 422, with default header values.

Malformed ID
*/
type DeleteAssetUnprocessableEntity struct {
}

func (o *DeleteAssetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /v1/assets/{id}][%d] deleteAssetUnprocessableEntity ", 422)
}

func (o *DeleteAssetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
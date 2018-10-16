// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionUpdateReader is a Reader for the WeaviateActionUpdate structure.
type WeaviateActionUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewWeaviateActionUpdateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateActionUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateActionUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateActionUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionUpdateAccepted creates a WeaviateActionUpdateAccepted with default headers values
func NewWeaviateActionUpdateAccepted() *WeaviateActionUpdateAccepted {
	return &WeaviateActionUpdateAccepted{}
}

/*WeaviateActionUpdateAccepted handles this case with default header values.

Successfully received.
*/
type WeaviateActionUpdateAccepted struct {
	Payload *models.ActionGetResponse
}

func (o *WeaviateActionUpdateAccepted) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}][%d] weaviateActionUpdateAccepted  %+v", 202, o.Payload)
}

func (o *WeaviateActionUpdateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionUpdateUnauthorized creates a WeaviateActionUpdateUnauthorized with default headers values
func NewWeaviateActionUpdateUnauthorized() *WeaviateActionUpdateUnauthorized {
	return &WeaviateActionUpdateUnauthorized{}
}

/*WeaviateActionUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionUpdateUnauthorized struct {
}

func (o *WeaviateActionUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}][%d] weaviateActionUpdateUnauthorized ", 401)
}

func (o *WeaviateActionUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionUpdateForbidden creates a WeaviateActionUpdateForbidden with default headers values
func NewWeaviateActionUpdateForbidden() *WeaviateActionUpdateForbidden {
	return &WeaviateActionUpdateForbidden{}
}

/*WeaviateActionUpdateForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateActionUpdateForbidden struct {
}

func (o *WeaviateActionUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}][%d] weaviateActionUpdateForbidden ", 403)
}

func (o *WeaviateActionUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionUpdateNotFound creates a WeaviateActionUpdateNotFound with default headers values
func NewWeaviateActionUpdateNotFound() *WeaviateActionUpdateNotFound {
	return &WeaviateActionUpdateNotFound{}
}

/*WeaviateActionUpdateNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateActionUpdateNotFound struct {
}

func (o *WeaviateActionUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}][%d] weaviateActionUpdateNotFound ", 404)
}

func (o *WeaviateActionUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionUpdateUnprocessableEntity creates a WeaviateActionUpdateUnprocessableEntity with default headers values
func NewWeaviateActionUpdateUnprocessableEntity() *WeaviateActionUpdateUnprocessableEntity {
	return &WeaviateActionUpdateUnprocessableEntity{}
}

/*WeaviateActionUpdateUnprocessableEntity handles this case with default header values.

Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateActionUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}][%d] weaviateActionUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateActionUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionUpdateInternalServerError creates a WeaviateActionUpdateInternalServerError with default headers values
func NewWeaviateActionUpdateInternalServerError() *WeaviateActionUpdateInternalServerError {
	return &WeaviateActionUpdateInternalServerError{}
}

/*WeaviateActionUpdateInternalServerError handles this case with default header values.

Internal server error; see the ErrorResponse in the response body for the reason.
*/
type WeaviateActionUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /actions/{actionId}][%d] weaviateActionUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateActionUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

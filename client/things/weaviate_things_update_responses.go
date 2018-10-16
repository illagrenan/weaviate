// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsUpdateReader is a Reader for the WeaviateThingsUpdate structure.
type WeaviateThingsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewWeaviateThingsUpdateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateThingsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateThingsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateThingsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsUpdateAccepted creates a WeaviateThingsUpdateAccepted with default headers values
func NewWeaviateThingsUpdateAccepted() *WeaviateThingsUpdateAccepted {
	return &WeaviateThingsUpdateAccepted{}
}

/*WeaviateThingsUpdateAccepted handles this case with default header values.

Successfully received.
*/
type WeaviateThingsUpdateAccepted struct {
	Payload *models.ThingGetResponse
}

func (o *WeaviateThingsUpdateAccepted) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}][%d] weaviateThingsUpdateAccepted  %+v", 202, o.Payload)
}

func (o *WeaviateThingsUpdateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThingGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsUpdateUnauthorized creates a WeaviateThingsUpdateUnauthorized with default headers values
func NewWeaviateThingsUpdateUnauthorized() *WeaviateThingsUpdateUnauthorized {
	return &WeaviateThingsUpdateUnauthorized{}
}

/*WeaviateThingsUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsUpdateUnauthorized struct {
}

func (o *WeaviateThingsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}][%d] weaviateThingsUpdateUnauthorized ", 401)
}

func (o *WeaviateThingsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsUpdateForbidden creates a WeaviateThingsUpdateForbidden with default headers values
func NewWeaviateThingsUpdateForbidden() *WeaviateThingsUpdateForbidden {
	return &WeaviateThingsUpdateForbidden{}
}

/*WeaviateThingsUpdateForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateThingsUpdateForbidden struct {
}

func (o *WeaviateThingsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}][%d] weaviateThingsUpdateForbidden ", 403)
}

func (o *WeaviateThingsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsUpdateNotFound creates a WeaviateThingsUpdateNotFound with default headers values
func NewWeaviateThingsUpdateNotFound() *WeaviateThingsUpdateNotFound {
	return &WeaviateThingsUpdateNotFound{}
}

/*WeaviateThingsUpdateNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateThingsUpdateNotFound struct {
}

func (o *WeaviateThingsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}][%d] weaviateThingsUpdateNotFound ", 404)
}

func (o *WeaviateThingsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsUpdateUnprocessableEntity creates a WeaviateThingsUpdateUnprocessableEntity with default headers values
func NewWeaviateThingsUpdateUnprocessableEntity() *WeaviateThingsUpdateUnprocessableEntity {
	return &WeaviateThingsUpdateUnprocessableEntity{}
}

/*WeaviateThingsUpdateUnprocessableEntity handles this case with default header values.

Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateThingsUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}][%d] weaviateThingsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateThingsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsUpdateInternalServerError creates a WeaviateThingsUpdateInternalServerError with default headers values
func NewWeaviateThingsUpdateInternalServerError() *WeaviateThingsUpdateInternalServerError {
	return &WeaviateThingsUpdateInternalServerError{}
}

/*WeaviateThingsUpdateInternalServerError handles this case with default header values.

Internal server error; see the ErrorResponse in the response body for the reason.
*/
type WeaviateThingsUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}][%d] weaviateThingsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateThingsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

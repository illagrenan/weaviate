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

// WeaviateThingHistoryGetReader is a Reader for the WeaviateThingHistoryGet structure.
type WeaviateThingHistoryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingHistoryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateThingHistoryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingHistoryGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingHistoryGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateThingHistoryGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateThingHistoryGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewWeaviateThingHistoryGetNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingHistoryGetOK creates a WeaviateThingHistoryGetOK with default headers values
func NewWeaviateThingHistoryGetOK() *WeaviateThingHistoryGetOK {
	return &WeaviateThingHistoryGetOK{}
}

/*WeaviateThingHistoryGetOK handles this case with default header values.

Successful response.
*/
type WeaviateThingHistoryGetOK struct {
	Payload *models.ThingGetHistoryResponse
}

func (o *WeaviateThingHistoryGetOK) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/history][%d] weaviateThingHistoryGetOK  %+v", 200, o.Payload)
}

func (o *WeaviateThingHistoryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThingGetHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingHistoryGetUnauthorized creates a WeaviateThingHistoryGetUnauthorized with default headers values
func NewWeaviateThingHistoryGetUnauthorized() *WeaviateThingHistoryGetUnauthorized {
	return &WeaviateThingHistoryGetUnauthorized{}
}

/*WeaviateThingHistoryGetUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingHistoryGetUnauthorized struct {
}

func (o *WeaviateThingHistoryGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/history][%d] weaviateThingHistoryGetUnauthorized ", 401)
}

func (o *WeaviateThingHistoryGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingHistoryGetForbidden creates a WeaviateThingHistoryGetForbidden with default headers values
func NewWeaviateThingHistoryGetForbidden() *WeaviateThingHistoryGetForbidden {
	return &WeaviateThingHistoryGetForbidden{}
}

/*WeaviateThingHistoryGetForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateThingHistoryGetForbidden struct {
}

func (o *WeaviateThingHistoryGetForbidden) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/history][%d] weaviateThingHistoryGetForbidden ", 403)
}

func (o *WeaviateThingHistoryGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingHistoryGetNotFound creates a WeaviateThingHistoryGetNotFound with default headers values
func NewWeaviateThingHistoryGetNotFound() *WeaviateThingHistoryGetNotFound {
	return &WeaviateThingHistoryGetNotFound{}
}

/*WeaviateThingHistoryGetNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateThingHistoryGetNotFound struct {
}

func (o *WeaviateThingHistoryGetNotFound) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/history][%d] weaviateThingHistoryGetNotFound ", 404)
}

func (o *WeaviateThingHistoryGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingHistoryGetInternalServerError creates a WeaviateThingHistoryGetInternalServerError with default headers values
func NewWeaviateThingHistoryGetInternalServerError() *WeaviateThingHistoryGetInternalServerError {
	return &WeaviateThingHistoryGetInternalServerError{}
}

/*WeaviateThingHistoryGetInternalServerError handles this case with default header values.

Internal server error; see the ErrorResponse in the response body for the reason.
*/
type WeaviateThingHistoryGetInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingHistoryGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/history][%d] weaviateThingHistoryGetInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateThingHistoryGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingHistoryGetNotImplemented creates a WeaviateThingHistoryGetNotImplemented with default headers values
func NewWeaviateThingHistoryGetNotImplemented() *WeaviateThingHistoryGetNotImplemented {
	return &WeaviateThingHistoryGetNotImplemented{}
}

/*WeaviateThingHistoryGetNotImplemented handles this case with default header values.

Not (yet) implemented.
*/
type WeaviateThingHistoryGetNotImplemented struct {
}

func (o *WeaviateThingHistoryGetNotImplemented) Error() string {
	return fmt.Sprintf("[GET /things/{thingId}/history][%d] weaviateThingHistoryGetNotImplemented ", 501)
}

func (o *WeaviateThingHistoryGetNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

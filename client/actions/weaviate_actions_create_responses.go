// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsCreateReader is a Reader for the WeaviateActionsCreate structure.
type WeaviateActionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateActionsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewWeaviateActionsCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateActionsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateActionsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsCreateOK creates a WeaviateActionsCreateOK with default headers values
func NewWeaviateActionsCreateOK() *WeaviateActionsCreateOK {
	return &WeaviateActionsCreateOK{}
}

/*WeaviateActionsCreateOK handles this case with default header values.

Action created
*/
type WeaviateActionsCreateOK struct {
	Payload *models.ActionGetResponse
}

func (o *WeaviateActionsCreateOK) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateOK  %+v", 200, o.Payload)
}

func (o *WeaviateActionsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsCreateAccepted creates a WeaviateActionsCreateAccepted with default headers values
func NewWeaviateActionsCreateAccepted() *WeaviateActionsCreateAccepted {
	return &WeaviateActionsCreateAccepted{}
}

/*WeaviateActionsCreateAccepted handles this case with default header values.

Successfully received. No guarantees are made that the Action is persisted.
*/
type WeaviateActionsCreateAccepted struct {
	Payload *models.ActionGetResponse
}

func (o *WeaviateActionsCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateAccepted  %+v", 202, o.Payload)
}

func (o *WeaviateActionsCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsCreateUnauthorized creates a WeaviateActionsCreateUnauthorized with default headers values
func NewWeaviateActionsCreateUnauthorized() *WeaviateActionsCreateUnauthorized {
	return &WeaviateActionsCreateUnauthorized{}
}

/*WeaviateActionsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsCreateUnauthorized struct {
}

func (o *WeaviateActionsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateUnauthorized ", 401)
}

func (o *WeaviateActionsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsCreateForbidden creates a WeaviateActionsCreateForbidden with default headers values
func NewWeaviateActionsCreateForbidden() *WeaviateActionsCreateForbidden {
	return &WeaviateActionsCreateForbidden{}
}

/*WeaviateActionsCreateForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateActionsCreateForbidden struct {
}

func (o *WeaviateActionsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateForbidden ", 403)
}

func (o *WeaviateActionsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsCreateUnprocessableEntity creates a WeaviateActionsCreateUnprocessableEntity with default headers values
func NewWeaviateActionsCreateUnprocessableEntity() *WeaviateActionsCreateUnprocessableEntity {
	return &WeaviateActionsCreateUnprocessableEntity{}
}

/*WeaviateActionsCreateUnprocessableEntity handles this case with default header values.

Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateActionsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateActionsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsCreateInternalServerError creates a WeaviateActionsCreateInternalServerError with default headers values
func NewWeaviateActionsCreateInternalServerError() *WeaviateActionsCreateInternalServerError {
	return &WeaviateActionsCreateInternalServerError{}
}

/*WeaviateActionsCreateInternalServerError handles this case with default header values.

Internal server error; see the ErrorResponse in the response body for the reason.
*/
type WeaviateActionsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /actions][%d] weaviateActionsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateActionsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WeaviateActionsCreateBody weaviate actions create body
swagger:model WeaviateActionsCreateBody
*/
type WeaviateActionsCreateBody struct {

	// action
	Action *models.ActionCreate `json:"action,omitempty"`

	// If `async` is true, return a 202 with the new ID of the Action. You will receive this response before the data is persisted. If `async` is false, you will receive confirmation after the value is persisted. The value of `async` defaults to false.
	Async bool `json:"async,omitempty"`
}

// Validate validates this weaviate actions create body
func (o *WeaviateActionsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateActionsCreateBody) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(o.Action) { // not required
		return nil
	}

	if o.Action != nil {
		if err := o.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "action")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *WeaviateActionsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WeaviateActionsCreateBody) UnmarshalBinary(b []byte) error {
	var res WeaviateActionsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

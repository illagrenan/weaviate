/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsPatchAcceptedCode is the HTTP code returned for type WeaviateThingsPatchAccepted
const WeaviateThingsPatchAcceptedCode int = 202

/*WeaviateThingsPatchAccepted Successfully received.

swagger:response weaviateThingsPatchAccepted
*/
type WeaviateThingsPatchAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ThingGetResponse `json:"body,omitempty"`
}

// NewWeaviateThingsPatchAccepted creates WeaviateThingsPatchAccepted with default headers values
func NewWeaviateThingsPatchAccepted() *WeaviateThingsPatchAccepted {

	return &WeaviateThingsPatchAccepted{}
}

// WithPayload adds the payload to the weaviate things patch accepted response
func (o *WeaviateThingsPatchAccepted) WithPayload(payload *models.ThingGetResponse) *WeaviateThingsPatchAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things patch accepted response
func (o *WeaviateThingsPatchAccepted) SetPayload(payload *models.ThingGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsPatchAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsPatchBadRequestCode is the HTTP code returned for type WeaviateThingsPatchBadRequest
const WeaviateThingsPatchBadRequestCode int = 400

/*WeaviateThingsPatchBadRequest The patch-JSON is malformed.

swagger:response weaviateThingsPatchBadRequest
*/
type WeaviateThingsPatchBadRequest struct {
}

// NewWeaviateThingsPatchBadRequest creates WeaviateThingsPatchBadRequest with default headers values
func NewWeaviateThingsPatchBadRequest() *WeaviateThingsPatchBadRequest {

	return &WeaviateThingsPatchBadRequest{}
}

// WriteResponse to the client
func (o *WeaviateThingsPatchBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// WeaviateThingsPatchUnauthorizedCode is the HTTP code returned for type WeaviateThingsPatchUnauthorized
const WeaviateThingsPatchUnauthorizedCode int = 401

/*WeaviateThingsPatchUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingsPatchUnauthorized
*/
type WeaviateThingsPatchUnauthorized struct {
}

// NewWeaviateThingsPatchUnauthorized creates WeaviateThingsPatchUnauthorized with default headers values
func NewWeaviateThingsPatchUnauthorized() *WeaviateThingsPatchUnauthorized {

	return &WeaviateThingsPatchUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingsPatchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateThingsPatchForbiddenCode is the HTTP code returned for type WeaviateThingsPatchForbidden
const WeaviateThingsPatchForbiddenCode int = 403

/*WeaviateThingsPatchForbidden The used API-key has insufficient permissions.

swagger:response weaviateThingsPatchForbidden
*/
type WeaviateThingsPatchForbidden struct {
}

// NewWeaviateThingsPatchForbidden creates WeaviateThingsPatchForbidden with default headers values
func NewWeaviateThingsPatchForbidden() *WeaviateThingsPatchForbidden {

	return &WeaviateThingsPatchForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingsPatchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateThingsPatchNotFoundCode is the HTTP code returned for type WeaviateThingsPatchNotFound
const WeaviateThingsPatchNotFoundCode int = 404

/*WeaviateThingsPatchNotFound Successful query result but no resource was found.

swagger:response weaviateThingsPatchNotFound
*/
type WeaviateThingsPatchNotFound struct {
}

// NewWeaviateThingsPatchNotFound creates WeaviateThingsPatchNotFound with default headers values
func NewWeaviateThingsPatchNotFound() *WeaviateThingsPatchNotFound {

	return &WeaviateThingsPatchNotFound{}
}

// WriteResponse to the client
func (o *WeaviateThingsPatchNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// WeaviateThingsPatchUnprocessableEntityCode is the HTTP code returned for type WeaviateThingsPatchUnprocessableEntity
const WeaviateThingsPatchUnprocessableEntityCode int = 422

/*WeaviateThingsPatchUnprocessableEntity The patch-JSON is valid but unprocessable.

swagger:response weaviateThingsPatchUnprocessableEntity
*/
type WeaviateThingsPatchUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsPatchUnprocessableEntity creates WeaviateThingsPatchUnprocessableEntity with default headers values
func NewWeaviateThingsPatchUnprocessableEntity() *WeaviateThingsPatchUnprocessableEntity {

	return &WeaviateThingsPatchUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate things patch unprocessable entity response
func (o *WeaviateThingsPatchUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateThingsPatchUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things patch unprocessable entity response
func (o *WeaviateThingsPatchUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsPatchUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsPatchInternalServerErrorCode is the HTTP code returned for type WeaviateThingsPatchInternalServerError
const WeaviateThingsPatchInternalServerErrorCode int = 500

/*WeaviateThingsPatchInternalServerError Internal server error; see the ErrorResponse in the response body for the reason.

swagger:response weaviateThingsPatchInternalServerError
*/
type WeaviateThingsPatchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsPatchInternalServerError creates WeaviateThingsPatchInternalServerError with default headers values
func NewWeaviateThingsPatchInternalServerError() *WeaviateThingsPatchInternalServerError {

	return &WeaviateThingsPatchInternalServerError{}
}

// WithPayload adds the payload to the weaviate things patch internal server error response
func (o *WeaviateThingsPatchInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateThingsPatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things patch internal server error response
func (o *WeaviateThingsPatchInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsPatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsGetOKCode is the HTTP code returned for type WeaviateActionsGetOK
const WeaviateActionsGetOKCode int = 200

/*WeaviateActionsGetOK Successful response.

swagger:response weaviateActionsGetOK
*/
type WeaviateActionsGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.ActionGetResponse `json:"body,omitempty"`
}

// NewWeaviateActionsGetOK creates WeaviateActionsGetOK with default headers values
func NewWeaviateActionsGetOK() *WeaviateActionsGetOK {

	return &WeaviateActionsGetOK{}
}

// WithPayload adds the payload to the weaviate actions get o k response
func (o *WeaviateActionsGetOK) WithPayload(payload *models.ActionGetResponse) *WeaviateActionsGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions get o k response
func (o *WeaviateActionsGetOK) SetPayload(payload *models.ActionGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsGetUnauthorizedCode is the HTTP code returned for type WeaviateActionsGetUnauthorized
const WeaviateActionsGetUnauthorizedCode int = 401

/*WeaviateActionsGetUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsGetUnauthorized
*/
type WeaviateActionsGetUnauthorized struct {
}

// NewWeaviateActionsGetUnauthorized creates WeaviateActionsGetUnauthorized with default headers values
func NewWeaviateActionsGetUnauthorized() *WeaviateActionsGetUnauthorized {

	return &WeaviateActionsGetUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionsGetForbiddenCode is the HTTP code returned for type WeaviateActionsGetForbidden
const WeaviateActionsGetForbiddenCode int = 403

/*WeaviateActionsGetForbidden The used API-key has insufficient permissions.

swagger:response weaviateActionsGetForbidden
*/
type WeaviateActionsGetForbidden struct {
}

// NewWeaviateActionsGetForbidden creates WeaviateActionsGetForbidden with default headers values
func NewWeaviateActionsGetForbidden() *WeaviateActionsGetForbidden {

	return &WeaviateActionsGetForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionsGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateActionsGetNotFoundCode is the HTTP code returned for type WeaviateActionsGetNotFound
const WeaviateActionsGetNotFoundCode int = 404

/*WeaviateActionsGetNotFound Successful query result but no resource was found.

swagger:response weaviateActionsGetNotFound
*/
type WeaviateActionsGetNotFound struct {
}

// NewWeaviateActionsGetNotFound creates WeaviateActionsGetNotFound with default headers values
func NewWeaviateActionsGetNotFound() *WeaviateActionsGetNotFound {

	return &WeaviateActionsGetNotFound{}
}

// WriteResponse to the client
func (o *WeaviateActionsGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// WeaviateActionsGetInternalServerErrorCode is the HTTP code returned for type WeaviateActionsGetInternalServerError
const WeaviateActionsGetInternalServerErrorCode int = 500

/*WeaviateActionsGetInternalServerError Internal server error; see the ErrorResponse in the response body for the reason.

swagger:response weaviateActionsGetInternalServerError
*/
type WeaviateActionsGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsGetInternalServerError creates WeaviateActionsGetInternalServerError with default headers values
func NewWeaviateActionsGetInternalServerError() *WeaviateActionsGetInternalServerError {

	return &WeaviateActionsGetInternalServerError{}
}

// WithPayload adds the payload to the weaviate actions get internal server error response
func (o *WeaviateActionsGetInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateActionsGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions get internal server error response
func (o *WeaviateActionsGetInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

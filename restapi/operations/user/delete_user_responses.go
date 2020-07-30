// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ntq.project/project.test/models"
)

// DeleteUserOKCode is the HTTP code returned for type DeleteUserOK
const DeleteUserOKCode int = 200

/*DeleteUserOK OK

swagger:response deleteUserOK
*/
type DeleteUserOK struct {

	/*
	  In: Body
	*/
	Payload []*models.CommonUser `json:"body,omitempty"`
}

// NewDeleteUserOK creates DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {

	return &DeleteUserOK{}
}

// WithPayload adds the payload to the delete user o k response
func (o *DeleteUserOK) WithPayload(payload []*models.CommonUser) *DeleteUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user o k response
func (o *DeleteUserOK) SetPayload(payload []*models.CommonUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.CommonUser, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteUserBadRequestCode is the HTTP code returned for type DeleteUserBadRequest
const DeleteUserBadRequestCode int = 400

/*DeleteUserBadRequest Bad Request

swagger:response deleteUserBadRequest
*/
type DeleteUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewDeleteUserBadRequest creates DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {

	return &DeleteUserBadRequest{}
}

// WithPayload adds the payload to the delete user bad request response
func (o *DeleteUserBadRequest) WithPayload(payload *models.ErrorResult) *DeleteUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user bad request response
func (o *DeleteUserBadRequest) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
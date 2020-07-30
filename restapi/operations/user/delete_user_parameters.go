// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewDeleteUserParams creates a new DeleteUserParams object
// no default values defined in spec.
func NewDeleteUserParams() DeleteUserParams {

	return DeleteUserParams{}
}

// DeleteUserParams contains all the bound params for the delete user operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteUser
type DeleteUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Id of users to delete
	  Required: true
	  In: body
	*/
	Ids []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteUserParams() beforehand.
func (o *DeleteUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("ids", "body", ""))
			} else {
				res = append(res, errors.NewParseError("ids", "body", "", err))
			}
		} else {
			// no validation required on inline body
			o.Ids = body
		}
	} else {
		res = append(res, errors.Required("ids", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package qq

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPostSendPrivateMsgParams creates a new PostSendPrivateMsgParams object
//
// There are no default values defined in the spec.
func NewPostSendPrivateMsgParams() PostSendPrivateMsgParams {

	return PostSendPrivateMsgParams{}
}

// PostSendPrivateMsgParams contains all the bound params for the post send private msg operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostSendPrivateMsg
type PostSendPrivateMsgParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*message
	  Required: true
	  In: query
	*/
	Message string
	/*QQ number
	  Required: true
	  In: query
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostSendPrivateMsgParams() beforehand.
func (o *PostSendPrivateMsgParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qMessage, qhkMessage, _ := qs.GetOK("message")
	if err := o.bindMessage(qMessage, qhkMessage, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserID, qhkUserID, _ := qs.GetOK("user_id")
	if err := o.bindUserID(qUserID, qhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMessage binds and validates parameter Message from query.
func (o *PostSendPrivateMsgParams) bindMessage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("message", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("message", "query", raw); err != nil {
		return err
	}
	o.Message = raw

	return nil
}

// bindUserID binds and validates parameter UserID from query.
func (o *PostSendPrivateMsgParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("user_id", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("user_id", "query", raw); err != nil {
		return err
	}
	o.UserID = raw

	return nil
}

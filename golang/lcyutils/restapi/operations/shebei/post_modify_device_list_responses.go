// Code generated by go-swagger; DO NOT EDIT.

package shebei

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.corp.bcollie.net/whistle/lcyutils/models"
)

// PostModifyDeviceListOKCode is the HTTP code returned for type PostModifyDeviceListOK
const PostModifyDeviceListOKCode int = 200

/*PostModifyDeviceListOK token check no problem

swagger:response postModifyDeviceListOK
*/
type PostModifyDeviceListOK struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostModifyDeviceListOK creates PostModifyDeviceListOK with default headers values
func NewPostModifyDeviceListOK() *PostModifyDeviceListOK {

	return &PostModifyDeviceListOK{}
}

// WithPayload adds the payload to the post modify device list o k response
func (o *PostModifyDeviceListOK) WithPayload(payload *models.APIResponse) *PostModifyDeviceListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post modify device list o k response
func (o *PostModifyDeviceListOK) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostModifyDeviceListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostModifyDeviceListInternalServerErrorCode is the HTTP code returned for type PostModifyDeviceListInternalServerError
const PostModifyDeviceListInternalServerErrorCode int = 500

/*PostModifyDeviceListInternalServerError internal error

swagger:response postModifyDeviceListInternalServerError
*/
type PostModifyDeviceListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostModifyDeviceListInternalServerError creates PostModifyDeviceListInternalServerError with default headers values
func NewPostModifyDeviceListInternalServerError() *PostModifyDeviceListInternalServerError {

	return &PostModifyDeviceListInternalServerError{}
}

// WithPayload adds the payload to the post modify device list internal server error response
func (o *PostModifyDeviceListInternalServerError) WithPayload(payload *models.APIResponse) *PostModifyDeviceListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post modify device list internal server error response
func (o *PostModifyDeviceListInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostModifyDeviceListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

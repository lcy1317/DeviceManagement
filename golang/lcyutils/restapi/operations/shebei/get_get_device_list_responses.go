// Code generated by go-swagger; DO NOT EDIT.

package shebei

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"code.corp.bcollie.net/whistle/lcyutils/models"
)

// GetGetDeviceListOKCode is the HTTP code returned for type GetGetDeviceListOK
const GetGetDeviceListOKCode int = 200

/*GetGetDeviceListOK token check no problem

swagger:response getGetDeviceListOK
*/
type GetGetDeviceListOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeviceListInfo `json:"body,omitempty"`
}

// NewGetGetDeviceListOK creates GetGetDeviceListOK with default headers values
func NewGetGetDeviceListOK() *GetGetDeviceListOK {

	return &GetGetDeviceListOK{}
}

// WithPayload adds the payload to the get get device list o k response
func (o *GetGetDeviceListOK) WithPayload(payload *models.DeviceListInfo) *GetGetDeviceListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get get device list o k response
func (o *GetGetDeviceListOK) SetPayload(payload *models.DeviceListInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGetDeviceListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetGetDeviceListInternalServerErrorCode is the HTTP code returned for type GetGetDeviceListInternalServerError
const GetGetDeviceListInternalServerErrorCode int = 500

/*GetGetDeviceListInternalServerError internal error

swagger:response getGetDeviceListInternalServerError
*/
type GetGetDeviceListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetGetDeviceListInternalServerError creates GetGetDeviceListInternalServerError with default headers values
func NewGetGetDeviceListInternalServerError() *GetGetDeviceListInternalServerError {

	return &GetGetDeviceListInternalServerError{}
}

// WithPayload adds the payload to the get get device list internal server error response
func (o *GetGetDeviceListInternalServerError) WithPayload(payload *models.APIResponse) *GetGetDeviceListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get get device list internal server error response
func (o *GetGetDeviceListInternalServerError) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGetDeviceListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
// Code generated by go-swagger; DO NOT EDIT.

package shebei

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetGetDeviceHandlerFunc turns a function with the right signature into a get get device handler
type GetGetDeviceHandlerFunc func(GetGetDeviceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGetDeviceHandlerFunc) Handle(params GetGetDeviceParams) middleware.Responder {
	return fn(params)
}

// GetGetDeviceHandler interface for that can handle valid get get device params
type GetGetDeviceHandler interface {
	Handle(GetGetDeviceParams) middleware.Responder
}

// NewGetGetDevice creates a new http.Handler for the get get device operation
func NewGetGetDevice(ctx *middleware.Context, handler GetGetDeviceHandler) *GetGetDevice {
	return &GetGetDevice{Context: ctx, Handler: handler}
}

/* GetGetDevice swagger:route GET /getDevice shebei getGetDevice

device record

*/
type GetGetDevice struct {
	Context *middleware.Context
	Handler GetGetDeviceHandler
}

func (o *GetGetDevice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetGetDeviceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

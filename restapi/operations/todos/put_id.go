// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutIDHandlerFunc turns a function with the right signature into a put ID handler
type PutIDHandlerFunc func(PutIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutIDHandlerFunc) Handle(params PutIDParams) middleware.Responder {
	return fn(params)
}

// PutIDHandler interface for that can handle valid put ID params
type PutIDHandler interface {
	Handle(PutIDParams) middleware.Responder
}

// NewPutID creates a new http.Handler for the put ID operation
func NewPutID(ctx *middleware.Context, handler PutIDHandler) *PutID {
	return &PutID{Context: ctx, Handler: handler}
}

/*PutID swagger:route PUT /{id} todos putId

PutID put ID API

*/
type PutID struct {
	Context *middleware.Context
	Handler PutIDHandler
}

func (o *PutID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
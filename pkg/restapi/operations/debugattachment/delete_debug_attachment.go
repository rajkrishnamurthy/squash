// Code generated by go-swagger; DO NOT EDIT.

package debugattachment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteDebugAttachmentHandlerFunc turns a function with the right signature into a delete debug attachment handler
type DeleteDebugAttachmentHandlerFunc func(DeleteDebugAttachmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDebugAttachmentHandlerFunc) Handle(params DeleteDebugAttachmentParams) middleware.Responder {
	return fn(params)
}

// DeleteDebugAttachmentHandler interface for that can handle valid delete debug attachment params
type DeleteDebugAttachmentHandler interface {
	Handle(DeleteDebugAttachmentParams) middleware.Responder
}

// NewDeleteDebugAttachment creates a new http.Handler for the delete debug attachment operation
func NewDeleteDebugAttachment(ctx *middleware.Context, handler DeleteDebugAttachmentHandler) *DeleteDebugAttachment {
	return &DeleteDebugAttachment{Context: ctx, Handler: handler}
}

/*DeleteDebugAttachment swagger:route DELETE /debugattachment/{debugAttachmentId} debugattachment deleteDebugAttachment

Delete a debug attachment

Delete a debug attachment. be careful not to delete on during attaching phase.

*/
type DeleteDebugAttachment struct {
	Context *middleware.Context
	Handler DeleteDebugAttachmentHandler
}

func (o *DeleteDebugAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteDebugAttachmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
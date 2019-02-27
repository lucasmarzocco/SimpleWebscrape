// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostGetJobsHandlerFunc turns a function with the right signature into a post get jobs handler
type PostGetJobsHandlerFunc func(PostGetJobsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostGetJobsHandlerFunc) Handle(params PostGetJobsParams) middleware.Responder {
	return fn(params)
}

// PostGetJobsHandler interface for that can handle valid post get jobs params
type PostGetJobsHandler interface {
	Handle(PostGetJobsParams) middleware.Responder
}

// NewPostGetJobs creates a new http.Handler for the post get jobs operation
func NewPostGetJobs(ctx *middleware.Context, handler PostGetJobsHandler) *PostGetJobs {
	return &PostGetJobs{Context: ctx, Handler: handler}
}

/*PostGetJobs swagger:route POST /get_jobs postGetJobs

Obtain information about jobs from Indeed

Obtain information about jobs from Indeed

*/
type PostGetJobs struct {
	Context *middleware.Context
	Handler PostGetJobsHandler
}

func (o *PostGetJobs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostGetJobsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

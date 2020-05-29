// Code generated by goa v3.1.3, DO NOT EDIT.
//
// resource HTTP server
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	"context"
	"net/http"
	"regexp"

	resource "github.com/tektoncd/hub/api/gen/resource"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the resource service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	All    http.Handler
	CORS   http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the resource service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *resource.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"All", "GET", "/resource"},
			{"CORS", "OPTIONS", "/resource"},
		},
		All:  NewAllHandler(e.All, mux, decoder, encoder, errhandler, formatter),
		CORS: NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "resource" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.All = m(s.All)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the resource endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAllHandler(mux, h.All)
	MountCORSHandler(mux, h.CORS)
}

// MountAllHandler configures the mux to serve the "resource" service "All"
// endpoint.
func MountAllHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleResourceOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/resource", f)
}

// NewAllHandler creates a HTTP handler which loads the HTTP request and calls
// the "resource" service "All" endpoint.
func NewAllHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeAllResponse(encoder)
		encodeError    = EncodeAllError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "All")
		ctx = context.WithValue(ctx, goa.ServiceKey, "resource")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service resource.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleResourceOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/resource", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleResourceOrigin applies the CORS response headers corresponding to the
// origin for the service resource.
func handleResourceOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*localhost")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "X-Time, X-Api-Version")
			w.Header().Set("Access-Control-Max-Age", "100")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
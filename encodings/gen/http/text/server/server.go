// Code generated by goa v3.0.2, DO NOT EDIT.
//
// text HTTP server
//
// Command:
// $ goa gen goa.design/examples/encodings/design -o
// $(GOPATH)/src/goa.design/examples/encodings

package server

import (
	"context"
	"net/http"

	text "goa.design/examples/encodings/gen/text"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the text service endpoint HTTP handlers.
type Server struct {
	Mounts            []*MountPoint
	Concatstrings     http.Handler
	Concatbytes       http.Handler
	Concatstringfield http.Handler
	Concatbytesfield  http.Handler
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

// New instantiates HTTP handlers for all the text service endpoints.
func New(
	e *text.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Concatstrings", "GET", "/concatstrings/{a}/{b}"},
			{"Concatbytes", "GET", "/concatbytes/{a}/{b}"},
			{"Concatstringfield", "GET", "/concatstringfield/{a}/{b}"},
			{"Concatbytesfield", "GET", "/concatbytesfield/{a}/{b}"},
			{"../../gen/http/openapi.json", "GET", "/swagger.json"},
		},
		Concatstrings:     NewConcatstringsHandler(e.Concatstrings, mux, dec, enc, eh),
		Concatbytes:       NewConcatbytesHandler(e.Concatbytes, mux, dec, enc, eh),
		Concatstringfield: NewConcatstringfieldHandler(e.Concatstringfield, mux, dec, enc, eh),
		Concatbytesfield:  NewConcatbytesfieldHandler(e.Concatbytesfield, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "text" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Concatstrings = m(s.Concatstrings)
	s.Concatbytes = m(s.Concatbytes)
	s.Concatstringfield = m(s.Concatstringfield)
	s.Concatbytesfield = m(s.Concatbytesfield)
}

// Mount configures the mux to serve the text endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountConcatstringsHandler(mux, h.Concatstrings)
	MountConcatbytesHandler(mux, h.Concatbytes)
	MountConcatstringfieldHandler(mux, h.Concatstringfield)
	MountConcatbytesfieldHandler(mux, h.Concatbytesfield)
	MountGenHTTPOpenapiJSON(mux, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../gen/http/openapi.json")
	}))
}

// MountConcatstringsHandler configures the mux to serve the "text" service
// "concatstrings" endpoint.
func MountConcatstringsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/concatstrings/{a}/{b}", f)
}

// NewConcatstringsHandler creates a HTTP handler which loads the HTTP request
// and calls the "text" service "concatstrings" endpoint.
func NewConcatstringsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeConcatstringsRequest(mux, dec)
		encodeResponse = EncodeConcatstringsResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "concatstrings")
		ctx = context.WithValue(ctx, goa.ServiceKey, "text")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountConcatbytesHandler configures the mux to serve the "text" service
// "concatbytes" endpoint.
func MountConcatbytesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/concatbytes/{a}/{b}", f)
}

// NewConcatbytesHandler creates a HTTP handler which loads the HTTP request
// and calls the "text" service "concatbytes" endpoint.
func NewConcatbytesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeConcatbytesRequest(mux, dec)
		encodeResponse = EncodeConcatbytesResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "concatbytes")
		ctx = context.WithValue(ctx, goa.ServiceKey, "text")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountConcatstringfieldHandler configures the mux to serve the "text" service
// "concatstringfield" endpoint.
func MountConcatstringfieldHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/concatstringfield/{a}/{b}", f)
}

// NewConcatstringfieldHandler creates a HTTP handler which loads the HTTP
// request and calls the "text" service "concatstringfield" endpoint.
func NewConcatstringfieldHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeConcatstringfieldRequest(mux, dec)
		encodeResponse = EncodeConcatstringfieldResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "concatstringfield")
		ctx = context.WithValue(ctx, goa.ServiceKey, "text")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountConcatbytesfieldHandler configures the mux to serve the "text" service
// "concatbytesfield" endpoint.
func MountConcatbytesfieldHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/concatbytesfield/{a}/{b}", f)
}

// NewConcatbytesfieldHandler creates a HTTP handler which loads the HTTP
// request and calls the "text" service "concatbytesfield" endpoint.
func NewConcatbytesfieldHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeConcatbytesfieldRequest(mux, dec)
		encodeResponse = EncodeConcatbytesfieldResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "concatbytesfield")
		ctx = context.WithValue(ctx, goa.ServiceKey, "text")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/swagger.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/swagger.json", h.ServeHTTP)
}

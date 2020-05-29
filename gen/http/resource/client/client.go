// Code generated by goa v3.1.3, DO NOT EDIT.
//
// resource client HTTP transport
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the resource service endpoint HTTP clients.
type Client struct {
	// All Doer is the HTTP client used to make requests to the All endpoint.
	AllDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the resource service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AllDoer:             doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// All returns an endpoint that makes HTTP requests to the resource service All
// server.
func (c *Client) All() goa.Endpoint {
	var (
		decodeResponse = DecodeAllResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAllRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AllDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("resource", "All", err)
		}
		return decodeResponse(resp)
	}
}

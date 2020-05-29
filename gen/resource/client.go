// Code generated by goa v3.1.3, DO NOT EDIT.
//
// resource client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "resource" service client.
type Client struct {
	AllEndpoint goa.Endpoint
}

// NewClient initializes a "resource" service client given the endpoints.
func NewClient(all goa.Endpoint) *Client {
	return &Client{
		AllEndpoint: all,
	}
}

// All calls the "All" endpoint of the "resource" service.
func (c *Client) All(ctx context.Context) (res []*Resource, err error) {
	var ires interface{}
	ires, err = c.AllEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Resource), nil
}
// Code generated by goa v3.1.3, DO NOT EDIT.
//
// resource HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	resource "github.com/tektoncd/hub/api/gen/resource"
	goa "goa.design/goa/v3/pkg"
)

// AllResponseBody is the type of the "resource" service "All" endpoint HTTP
// response body.
type AllResponseBody []*ResourceResponse

// AllInternalErrorResponseBody is the type of the "resource" service "All"
// endpoint HTTP response body for the "internal-error" error.
type AllInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ResourceResponse is used to define fields on response body types.
type ResourceResponse struct {
	// ID is the unique id of the resource
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the resource
	Name string `form:"name" json:"name" xml:"name"`
	// Display name of the resource
	DisplayName string `form:"displayName" json:"displayName" xml:"displayName"`
	// Type of catalog where resource belongs
	Catalog *CatalogResponse `form:"catalog" json:"catalog" xml:"catalog"`
	// Type of resource
	Type string `form:"type" json:"type" xml:"type"`
	// Description of resource
	Description string `form:"description" json:"description" xml:"description"`
	// Latest version o resource
	LatestVersion string `form:"latest_version" json:"latest_version" xml:"latest_version"`
	// Tags related to resources
	Tags []*Tag `form:"tags" json:"tags" xml:"tags"`
	// Rating of resource
	Rating uint `form:"rating" json:"rating" xml:"rating"`
	// Date when resource was last updated
	LastUpdatedAt string `form:"last_updated_at" json:"last_updated_at" xml:"last_updated_at"`
}

// CatalogResponse is used to define fields on response body types.
type CatalogResponse struct {
	// ID is the unique id of the category
	ID uint `form:"id" json:"id" xml:"id"`
	// Type of support tier
	Type string `form:"type" json:"type" xml:"type"`
}

// Tag is used to define fields on response body types.
type Tag struct {
	// ID is the unique id of the tag
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the tag
	Name string `form:"name" json:"name" xml:"name"`
}

// NewAllResponseBody builds the HTTP response body from the result of the
// "All" endpoint of the "resource" service.
func NewAllResponseBody(res []*resource.Resource) AllResponseBody {
	body := make([]*ResourceResponse, len(res))
	for i, val := range res {
		body[i] = marshalResourceResourceToResourceResponse(val)
	}
	return body
}

// NewAllInternalErrorResponseBody builds the HTTP response body from the
// result of the "All" endpoint of the "resource" service.
func NewAllInternalErrorResponseBody(res *goa.ServiceError) *AllInternalErrorResponseBody {
	body := &AllInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

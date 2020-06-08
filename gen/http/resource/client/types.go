// Code generated by goa v3.1.3, DO NOT EDIT.
//
// resource HTTP client types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

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
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ResourceResponse is used to define fields on response body types.
type ResourceResponse struct {
	// ID is the unique id of the resource
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the resource
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Display name of the resource
	DisplayName *string `form:"displayName,omitempty" json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Type of catalog where resource belongs
	Catalog *CatalogResponse `form:"catalog,omitempty" json:"catalog,omitempty" xml:"catalog,omitempty"`
	// Type of resource
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Description of resource
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Latest version o resource
	LatestVersion *string `form:"latest_version,omitempty" json:"latest_version,omitempty" xml:"latest_version,omitempty"`
	// Tags related to resources
	Tags []*Tag `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	// Rating of resource
	Rating *uint `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	// Date when resource was last updated
	LastUpdatedAt *string `form:"last_updated_at,omitempty" json:"last_updated_at,omitempty" xml:"last_updated_at,omitempty"`
}

// CatalogResponse is used to define fields on response body types.
type CatalogResponse struct {
	// ID is the unique id of the category
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Type of support tier
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Tag is used to define fields on response body types.
type Tag struct {
	// ID is the unique id of the tag
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the tag
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewAllResourceOK builds a "resource" service "All" endpoint result from a
// HTTP "OK" response.
func NewAllResourceOK(body []*ResourceResponse) []*resource.Resource {
	v := make([]*resource.Resource, len(body))
	for i, val := range body {
		v[i] = unmarshalResourceResponseToResourceResource(val)
	}
	return v
}

// NewAllInternalError builds a resource service All endpoint internal-error
// error.
func NewAllInternalError(body *AllInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateAllInternalErrorResponseBody runs the validations defined on
// All_internal-error_Response_Body
func ValidateAllInternalErrorResponseBody(body *AllInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateResourceResponse runs the validations defined on ResourceResponse
func ValidateResourceResponse(body *ResourceResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.DisplayName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("displayName", "body"))
	}
	if body.Catalog == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("catalog", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.LatestVersion == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("latest_version", "body"))
	}
	if body.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rating", "body"))
	}
	if body.LastUpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last_updated_at", "body"))
	}
	if body.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "body"))
	}
	if body.Catalog != nil {
		if err2 := ValidateCatalogResponse(body.Catalog); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range body.Tags {
		if e != nil {
			if err2 := ValidateTag(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateCatalogResponse runs the validations defined on CatalogResponse
func ValidateCatalogResponse(body *CatalogResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	return
}

// ValidateTag runs the validations defined on Tag
func ValidateTag(body *Tag) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

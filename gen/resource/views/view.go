// Code generated by goa v3.1.3, DO NOT EDIT.
//
// resource views
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Detail is the viewed result type that is projected based on a view.
type Detail struct {
	// Type to project
	Projected *DetailView
	// View to render
	View string
}

// DetailView is a type that runs validations on a projected type.
type DetailView struct {
	// ID is the unique id of the resource
	ID *uint
	// Name of the resource
	Name *string
	// Display name of the resource
	DisplayName *string
	// Type of catalog where resource belongs
	Catalog *CatalogView
	// Type of resource
	Type *string
	// Description of resource
	Description *string
	// Latest version o resource
	LatestVersion *string
	// Rating of resource
	Rating *uint
	// Date when resource was last updated
	LastUpdatedAt *string
	// Version of resource
	Versions []*VersionsView
}

// CatalogView is a type that runs validations on a projected type.
type CatalogView struct {
	// ID is the unique id of the category
	ID *uint
	// Type of support tier
	Type *string
}

// VersionsView is a type that runs validations on a projected type.
type VersionsView struct {
	// Version ID of the resource to be fetched
	VersionID *uint
	// Version of the resource to be fetched
	Version *string
}

var (
	// DetailMap is a map of attribute names in result type Detail indexed by view
	// name.
	DetailMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"displayName",
			"catalog",
			"type",
			"description",
			"latest_version",
			"rating",
			"last_updated_at",
			"versions",
		},
	}
)

// ValidateDetail runs the validations defined on the viewed result type Detail.
func ValidateDetail(result *Detail) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateDetailView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateDetailView runs the validations defined on DetailView using the
// "default" view.
func ValidateDetailView(result *DetailView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.DisplayName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("displayName", "result"))
	}
	if result.Catalog == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("catalog", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	if result.LatestVersion == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("latest_version", "result"))
	}
	if result.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rating", "result"))
	}
	if result.LastUpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last_updated_at", "result"))
	}
	if result.Versions == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("versions", "result"))
	}
	if result.Catalog != nil {
		if err2 := ValidateCatalogView(result.Catalog); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range result.Versions {
		if e != nil {
			if err2 := ValidateVersionsView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateCatalogView runs the validations defined on CatalogView.
func ValidateCatalogView(result *CatalogView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	return
}

// ValidateVersionsView runs the validations defined on VersionsView.
func ValidateVersionsView(result *VersionsView) (err error) {
	if result.VersionID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("versionId", "result"))
	}
	if result.Version == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("version", "result"))
	}
	return
}
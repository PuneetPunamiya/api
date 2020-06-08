package hub

import (
	"github.com/jinzhu/gorm"
)

type (

	//Catalog model represents origin repo to which resource belongs to
	Catalog struct {
		gorm.Model
		Name       string
		Type       string
		URL        string
		Owner      string
		ContextDir string
		Resources  []Resource
	}

	//Category model represents categories which associated with group of tags
	Category struct {
		gorm.Model
		Name string `gorm:"size:100;not null;unique"`
		Tags []Tag
	}

	// Tag model represents tags associated with a resource
	Tag struct {
		gorm.Model
		Name       string `gorm:"size:100;not null;unique"`
		Category   Category
		CategoryID int
		Resources  []*Resource `gorm:"many2many:resource_tags;"`
	}

	// Resource represents model which describe resources from a catalog
	Resource struct {
		gorm.Model
		Name        string
		DisplayName string
		Catalog     Catalog
		Type        string
		Rating      float64
		CatalogID   uint
		Versions    []ResourceVersion
		Tags        []*Tag `gorm:"many2many:resource_tags;"`
	}

	// ResourceVersion represents diffrent versions of a Resource
	ResourceVersion struct {
		gorm.Model
		Version     string
		Description string
		URL         string
		Resource    Resource
		ResourceID  uint
	}
)

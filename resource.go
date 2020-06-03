package hub

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
	resource "github.com/tektoncd/hub/api/gen/resource"
)

// resource service example implementation.
// The example methods log the requests and return zero values.
type resourcesrvc struct {
	db     *gorm.DB
	logger *log.Logger
}

// NewResource returns the resource service implementation.
func NewResource(db *gorm.DB, logger *log.Logger) resource.Service {
	return &resourcesrvc{db, logger}
}

// Get all Resources
func (s *resourcesrvc) All(ctx context.Context) (res []*resource.Resource, err error) {

	s.logger.Print("resource.All")

	var all []Resource
	s.db.Find(&all)

	ret := make([]*resource.Resource, len(all))
	for i, r := range all {
		ret[i] = Init(r)
	}
	return
}

// Init Converts
func Init(r Resource) *resource.Resource {
	return &resource.Resource{
		ID:     uint(r.ID),
		Name:   r.Name,
		Type:   r.Type,
		Rating: uint(r.Rating),

		// TODO -->  Pass database data to resource struct data
		// Catalog.ID: r.Catalog.ID,
	}
}

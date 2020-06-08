package hub

import (
	"context"
	"errors"
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

	if err := s.db.Preload("Catalog").
		Preload("Versions", func(db *gorm.DB) *gorm.DB {
			return db.Order("resource_versions.id ASC")
		}).
		Preload("Tags").
		Find(&all).Error; err != nil {
		return []*resource.Resource{}, errors.New("Failed to fetch Resources")
	}

	for _, r := range all {
		tags := []*resource.Tag{}
		for _, t := range r.Tags {
			tags = append(tags, &resource.Tag{
				ID:   t.ID,
				Name: t.Name,
			})
		}
		latestVersion := r.Versions[len(r.Versions)-1]
		res = append(res, &resource.Resource{
			ID:            r.ID,
			Name:          r.Name,
			DisplayName:   r.DisplayName,
			Tags:          tags,
			Catalog:       &resource.Catalog{ID: r.Catalog.ID, Type: r.Catalog.Type},
			Type:          r.Type,
			LatestVersion: latestVersion.Version,
			Description:   latestVersion.Description,
			Rating:        uint(r.Rating),
			LastUpdatedAt: r.UpdatedAt.String(),
		})
	}

	return res, nil
}

package hub

import (
	"context"
	"log"

	resource "github.com/tektoncd/hub/api/gen/resource"
)

// resource service example implementation.
// The example methods log the requests and return zero values.
type resourcesrvc struct {
	logger *log.Logger
}

// NewResource returns the resource service implementation.
func NewResource(logger *log.Logger) resource.Service {
	return &resourcesrvc{logger}
}

// Get all Resources
func (s *resourcesrvc) All(ctx context.Context) (res []*resource.Resource, err error) {
	s.logger.Print("resource.All")
	return
}

// Get one Resource info
func (s *resourcesrvc) Info(ctx context.Context, p *resource.InfoPayload) (res *resource.Resource, err error) {
	
	res = &resource.Resource{
		ID: 1,
		Tags: []*resource.Tag{
			{
				ID:   1,
				Name: "Mail",
			},
		},
	}
	s.logger.Print("resource.Info")
	return
}

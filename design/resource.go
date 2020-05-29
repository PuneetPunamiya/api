package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("resource", func() {
	Description("The resource service provides all resources")

	Error("internal-error", ErrorResult)

	// Method to get all resources
	Method("All", func() {
		Description("Get all Resources")
		Result(ArrayOf(Resource))

		HTTP(func() {
			GET("/resource")
			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
		})
	})
})

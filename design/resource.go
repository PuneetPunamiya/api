package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("resource", func() {
	Description("The resource service provides all resources information")

	Error("internal-error", ErrorResult)

	// Method to get all resources
	Method("All", func() {
		Description("Get all Resources")
		Result(ArrayOf(Resource))

		HTTP(func() {
			GET("/resources")
			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
		})
	})

	// Method to get a resource info
	Method("Info", func() {
		Description("Get one Resource info")
		Payload(func() {
			Attribute("resourceId", UInt, "ID of resource to be shown")
			Required("resourceId")
		})
		Result(Detail)
		HTTP(func() {
			GET("/resource/{resourceId}")
			Response(StatusOK)
			Response("internal-error", StatusInternalServerError)
		})
	})

	// // Method to get a resource info
	// Method("Infotkn", func() {
	// 	Description("Get one Resource info")
	// 	Payload(func() {
	// 		Attribute("type", String, "Type of Resource")
	// 		Attribute("name", String, "Name of Resource")
	// 		Required("type", "name")
	// 	})
	// 	Result(Detail)
	// 	HTTP(func() {
	// 		GET("/resource/{type}/{name}")
	// 		Response(StatusOK)
	// 		Response("internal-error", StatusInternalServerError)
	// 	})
	// })
})

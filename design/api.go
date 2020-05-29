package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("hub", func() {
	Title("Tekton Hub")
	Description("HTTP servie for maintaing Tetkon hub")
	Server("hub", func() {
		Services("resource")
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})

	cors.Origin("/.*localhost/", func() {
		cors.Headers("Content-Type")
		cors.Methods("GET", "POST", "PUT")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
	})
})

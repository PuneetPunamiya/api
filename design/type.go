package design

import (
	. "goa.design/goa/v3/dsl"
)

var Resource = Type("Resource", func() {

	Attribute("id", UInt, "ID is the unique id of the category", func() {
		Example("id", 1)
	})

	Attribute("name", String, "Name of the Category", func() {
		Example("name", "golang")
	})

	Attribute("catalog", Catalog, "Type of catalog where resource belongs")

	Attribute("type", String, "Type of resource", func() {
		Example("type", "task")
	})

	Attribute("description", String, "Description of resource", func() {
		Example("descripiton", "This Task is Golang task to build Go projects")
	})

	Attribute("latest_version", String, "Latest version o resource", func() {
		Example("latest_version", "0.1")
	})

	Attribute("tags", ArrayOf(ResourceTag), "Tags related to resources")

	Attribute("rating", UInt, "Rating of resource", func() {
		Example("rating", 5)
	})

	Attribute("last_updated_at", String, "Date when resource was last updated", func() {
		Example("last_updated_at", "2020-03-26T14:09:08.19599+05:30")
	})

	Required("id", "name", "catalog", "type", "description", "latest_version", "rating", "last_updated_at", "tags")
})

var Catalog = Type("Catalog", func() {
	Attribute("id", UInt, "ID is the unique id of the category", func() {
		Example("id", 1)
	})
	Attribute("type", String, "Type of support tier", func() {
		Example("type", "official")
	})
	Required("id", "type")
})

var ResourceTag = Type("ResourceTag", func() {
	TypeName("Tag")
	Attribute("id", UInt, "ID is the unique id of the tag", func() {
		Example("id", 1)
	})
	Attribute("name", String, "Name of the tag", func() {
		Example("name", "notification")
	})
	Required("id", "name")
})

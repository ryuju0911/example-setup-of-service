package design

import (
	"example/design/types"

	. "goa.design/goa/v3/dsl"
)

var _ = API("example", func() {
	Title("Example Service")
	Description("Example service for setting up a new service")
	Server("example", func() {
		Host("localhost", func() {
			URI("http://0.0.0.0:80/api/example")
		})
	})
})

var _ = Service("example", func() {
	Description("The example service is documentation for setting up a new service")

	Error("BadRequest")

	Method("CreateSomeData", func() {
		Payload(func() {
			Field(1, "id", String)
			Field(2, "content", String)
			Required("id", "content")
		})
		Result(types.SomeData)

		HTTP(func() {
			POST("/api/example")

			Response(StatusOK, func() {
				Description("Created some data")
			})
			Response("BadRequest", StatusBadRequest, func() {
				Description("Failed to create some data")
			})
		})
	})

	Files("/swagger.json", "./gen/http/openapi.json")
})

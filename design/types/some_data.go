package types

import . "goa.design/goa/v3/dsl"

var SomeData = Type("SomeData", func() {
	Attribute("id", String)
	Attribute("content", String)
	Attribute("created_at", String, func() {
		Format(FormatDateTime)
	})

	Required("id", "content", "created_at")
})

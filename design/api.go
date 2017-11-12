package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() {
	Title("The virtual wine cellar")
	Description("A basic example of a CRUD API implemented with goa")
	Host("localhost:8081")
	Scheme("http")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})

var _ = Resource("account", func() {
	DefaultMedia(Account)
	BasePath("/accounts")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all accounts.")
		Response(OK, AccountCollection)
	})

	Action("show", func() {
		Routing(
			GET("/:accountID"),
		)
		Description("Retrieve account with given id. IDs 1 and 2 pre-exist in the system.")
		Params(func() {
			Param("accountID", Integer, "Account ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new account")
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(Created, "/accounts/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:accountID"),
		)
		Description("Change account name")
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:accountID"),
		)
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("bottle", func() {
	DefaultMedia(Bottle)
	BasePath("bottles")
	Parent("account")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all bottles in account optionally filtering by year")
		Params(func() {
			Param("years", ArrayOf(Integer), "Filter by years")
		})
		Response(OK, func() {
			Media(BottleCollection)
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Routing(
			GET("/:bottleID"),
		)
		Description("Retrieve bottle with given id")
		Params(func() {
			Param("bottleID", Integer)
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Record new bottle")
		Payload(BottlePayload, func() {
			Required("name", "vineyard", "varietal", "vintage", "color")
		})
		Response(Created, "^/accounts/[0-9]+/bottles/[0-9]+$")
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PATCH("/:bottleID"),
		)
		Params(func() {
			Param("bottleID", Integer)
		})
		Payload(BottlePayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("rate", func() {
		Routing(
			PUT("/:bottleID/actions/rate"),
		)
		Params(func() {
			Param("bottleID", Integer)
		})
		Payload(func() {
			Member("rating", Integer)
			Required("rating")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:bottleID"),
		)
		Params(func() {
			Param("bottleID", Integer)
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var AccountCollectionLinks = Type("AccountCollectionLinks", func() {
	Attribute("self", HALLink)
	Attribute("members", ArrayOf(HALLink))

	Required("self", "members")
})

var AccountCollection = MediaType("application/vnd.account-collection+json", func() {
	Attributes(func() {
		Attribute("_links", AccountCollectionLinks)

		Required("_links")
	})

	View("default", func() {
		Attribute("_links")
	})
})

var AccountLinks = Type("AccountLinks", func() {
	Attribute("self", HALLink)
	Attribute("bottles", HALLink)

	Required("self", "bottles")
})

// Account is the account resource media type.
var Account = MediaType("application/vnd.account+json", func() {
	Description("A tenant account")
	Attributes(func() {
		Attribute("_links", AccountLinks)
		Attribute("id", Integer, "ID of account", func() {
			Example(1)
		})
		Attribute("name", String, "Name of account", func() {
			Example("test")
		})
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("created_by", String, "Email of account owner", func() {
			Format("email")
			Example("me@goa.design")
		})

		Required("_links", "id", "name", "created_at", "created_by")
	})

	View("default", func() {
		Attribute("_links")
		Attribute("id")
		Attribute("name")
		Attribute("created_at")
		Attribute("created_by")
	})
})

var BottleCollectionLinks = Type("BottleCollectionLinks", func() {
	Attribute("self", HALLink)
	Attribute("members", ArrayOf(HALLink))

	Required("self", "members")
})

var BottleCollection = MediaType("application/vnd.bottle-collection+json", func() {
	Attributes(func() {
		Attribute("_links", BottleCollectionLinks)

		Required("_links")
	})

	View("default", func() {
		Attribute("_links")
	})
})

var BottleLinks = Type("BottleLinks", func() {
	Attribute("self", HALLink)
	Attribute("account", HALLink)

	Required("self", "account")
})

// Bottle is the bottle resource media type.
var Bottle = MediaType("application/vnd.bottle+json", func() {
	Description("A bottle of wine")
	Reference(BottlePayload)
	Attributes(func() {
		Attribute("_links", BottleLinks)
		Attribute("id", Integer, "ID of bottle", func() {
			Example(1)
		})
		Attribute("rating", Integer, "Rating of bottle between 1 and 5", func() {
			Minimum(1)
			Maximum(5)
		})
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("updated_at", DateTime, "Date of last update")
		// Attributes below inherit from the base type
		Attribute("name")
		Attribute("vineyard")
		Attribute("varietal")
		Attribute("vintage")
		Attribute("color")
		Attribute("sweetness")
		Attribute("country")
		Attribute("region")
		Attribute("review")

		Required("_links", "id", "name", "vineyard", "varietal", "vintage", "color")
		Required("created_at", "updated_at")
	})

	View("default", func() {
		Attribute("_links")
		Attribute("id")
		Attribute("name")
		Attribute("rating")
		Attribute("vineyard")
		Attribute("varietal")
		Attribute("vintage")
	})
})

// BottlePayload defines the data structure used in the create bottle request body.
// It is also the base type for the bottle media type used to render bottles.
var BottlePayload = Type("BottlePayload", func() {
	Attribute("name", func() {
		MinLength(2)
		Example("Number 8")
	})
	Attribute("vineyard", func() {
		MinLength(2)
		Example("Asti")
	})
	Attribute("varietal", func() {
		MinLength(4)
		Example("Merlot")
	})
	Attribute("vintage", Integer, func() {
		Minimum(1900)
		Maximum(2020)
		Example(2012)
	})
	Attribute("color", func() {
		Enum("red", "white", "rose", "yellow", "sparkling")
	})
	Attribute("sweetness", Integer, func() {
		Minimum(1)
		Maximum(5)
	})
	Attribute("country", func() {
		MinLength(2)
		Example("USA")
	})
	Attribute("region", func() {
		Example("Napa Valley")
	})
	Attribute("review", func() {
		MinLength(3)
		MaxLength(300)
		Example("Great and inexpensive")
	})
})

var HALLink = Type("HALLink", func() {
	Attribute("href", String)

	Required("href")
})

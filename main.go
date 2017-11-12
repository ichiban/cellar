//go:generate goagen main -d github.com/ichiban/cellar/design
//go:generate goagen app -d github.com/ichiban/cellar/design
//go:generate goose -dir db/migrations sqlite3 db/cellar.db up
//go:generate sqlite3 -init db/seeds.sql db/cellar.db
//go:generate xo "file:db/cellar.db?loc=auto&_foreign_keys=1" -o models
//go:generate xo "file:db/cellar.db?loc=auto&_foreign_keys=1" -o models --query-mode --query-type AccountID --query "SELECT id FROM accounts ORDER BY id"

package main

import (
	"log"

	"database/sql"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ichiban/cellar/app"
	"github.com/knq/dburl"
)

func main() {
	// Create service
	service := goa.New("cellar")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "account" controller
	c := NewAccountController(service)
	app.MountAccountController(service, c)
	// Mount "bottle" controller
	c2 := NewBottleController(service)
	app.MountBottleController(service, c2)

	log.Printf("starting...")

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}
}

var DB *sql.DB

func init() {
	var err error
	DB, err = dburl.Open("file:db/cellar.db?loc=auto&_foreign_keys=1")
	if err != nil {
		log.Fatal(err)
	}
}

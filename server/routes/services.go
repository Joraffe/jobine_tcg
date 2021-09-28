package routes

import (
	"log"

	"github.com/Joraffe/jobine_tcg/server/services/db"
)


// Services describes what services are
// available to all routes in our application
type Services struct {
	Database  db.DatabaseManager
}


// NewRouterServices initializes all of the services available to our routers
func NewRouterServices() *Services {
	services := new(Services)

	database, err := db.New()
	if err != nil {
		log.Fatalf("Error opening database: %s", err.Error())
	}

	services.Database = database

	return services
}

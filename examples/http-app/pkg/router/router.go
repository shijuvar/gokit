package router

import (
	"github.com/gorilla/mux"

	util "github.com/shijuvar/gokit/examples/http-app/pkg/apputil"
	"github.com/shijuvar/gokit/examples/http-app/pkg/postgres"
)

// InitRoutes registers all routes for the application.
func InitRoutes() *mux.Router {

	// Create config for Postgres
	config := postgres.Config{
		Host:     util.AppConfig.DBHost,
		Port:     util.AppConfig.DBPort,
		User:     util.AppConfig.DBUser,
		Password: util.AppConfig.DBPassword,
		Database: util.AppConfig.Database,
	}
	// Creates a Postgres DB instance
	dataStore, err := postgres.New(config)
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router = SetUserRoutes(router, dataStore)
	router = SetProductRoutes(router, dataStore)
	return router
}

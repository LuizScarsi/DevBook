package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a new router with configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
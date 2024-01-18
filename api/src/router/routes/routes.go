package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API HTTP routes
type Route struct {
	URI string
	Method string
	Handler func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

// Configure configures all routes inside router
func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes
	
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	
	return r
}
package routes

import (
	"api/src/middlewares"
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
	routes = append(routes, loginRoute)
	
	for _, route := range routes {
		if route.RequiresAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Auth(route.Handler))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Handler)).Methods(route.Method)
		}
	}
	
	return r
}
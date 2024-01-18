package routes

import "net/http"

// Route represents all API HTTP routes
type Route struct {
	URI string
	Method string
	Handler func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}
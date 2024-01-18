package router

import "github.com/gorilla/mux"

// Generate returns a new router with configured routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
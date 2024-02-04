package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URI: "/login",
	Method: http.MethodPost,
	Handler: controllers.Login,
	RequiresAuth: false,
}
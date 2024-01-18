package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Handler: controllers.SearchUsers,
		RequiresAuth: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodGet,
		Handler: controllers.SearchUser,
		RequiresAuth: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
		RequiresAuth: false,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
		RequiresAuth: false,
	},
}
package rotas

import (
	"api/src/controllers"
	"net/http"
)

var routesUser = []Routes{
	{
		URI:            "/user",
		Method:         http.MethodPost,
		Func:           controllers.CreateUsers,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Func:           controllers.SearchUsers,
		Authentication: true,
	},
	{
		URI:            "/user/{id}",
		Method:         http.MethodGet,
		Func:           controllers.SearchUser,
		Authentication: true,
	},
	{
		URI:            "/user/{id}",
		Method:         http.MethodPut,
		Func:           controllers.UpdateUser,
		Authentication: true,
	},
	{
		URI:            "/user/{id}",
		Method:         http.MethodDelete,
		Func:           controllers.DeleteUser,
		Authentication: true,
	},
	{
		URI:            "/user/{userID}/follow",
		Method:         http.MethodPost,
		Func:           controllers.FollowUser,
		Authentication: true,
	},
}

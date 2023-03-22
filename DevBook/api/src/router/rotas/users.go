package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Routes{
	{
		URI:            "/usuarios",
		Method:         http.MethodPost,
		Func:           controllers.CreateUsers,
		Authentication: false,
	},
	{
		URI:            "/usuarios",
		Method:         http.MethodGet,
		Func:           controllers.SearchUsers,
		Authentication: false,
	},
	{
		URI:            "/usuario/{id}",
		Method:         http.MethodGet,
		Func:           controllers.SearchUser,
		Authentication: false,
	},
	{
		URI:            "/usuario/{id}",
		Method:         http.MethodPut,
		Func:           controllers.UpdateUser,
		Authentication: false,
	},
	{
		URI:            "/usuario/{id}",
		Method:         http.MethodDelete,
		Func:           controllers.DeleteUser,
		Authentication: false,
	},
}

package rotas

import (
	"api/src/controllers"
	"net/http"
)

var routeLogin = Routes{
	URI:            "/login",
	Method:         http.MethodPost,
	Func:           controllers.Login,
	Authentication: false,
}

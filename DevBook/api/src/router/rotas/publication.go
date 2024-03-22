package rotas

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Routes{
	{
		URI:            "/publications",
		Method:         http.MethodPost,
		Func:           controllers.CreatePublication,
		Authentication: true,
	},
	{
		URI:            "/publications",
		Method:         http.MethodGet,
		Func:           controllers.SearchPublications,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationID}",
		Method:         http.MethodGet,
		Func:           controllers.SearchPublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationID}",
		Method:         http.MethodPut,
		Func:           controllers.UpdatePublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationID}",
		Method:         http.MethodDelete,
		Func:           controllers.DeletePublication,
		Authentication: true,
	},
}

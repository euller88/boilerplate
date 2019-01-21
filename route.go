package boilerplate

import "net/http"

var routes = make(Routes, 0)

//Route define como uma rota será escutada
//
//Route defines how a route will be listened
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

//Routes agrupa as rotas à serem escutadas pela API, apenas um tipo de conveniência
//
//Routes groups the routes that will be listened by the API, just a convenience type
type Routes []*Route

//RegisterRoute insere um nova rota para estar avaliável para a API. Precisa ser chamada antes de NewRouter()
//
//RegisterRoute inserts a new route to be available for the API. Must be called before NewRouter()
func RegisterRoute(r *Route) {
	routes = append(routes, r)
}
